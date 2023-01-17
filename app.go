package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

type postBody struct {
	UserId  string `json:"userid"`
	DocId   string `json:"docid"`
	Content string `json:"content"`
}

func sendRes(w http.ResponseWriter, jsonData *simplejson.Json) {

	// JSON encode jsonData
	payload, err := jsonData.MarshalJSON()
	if err != nil {
		log.Println(err, "\tstatus_code: 992")
		http.Error(w, "Internal Error", http.StatusMethodNotAllowed)
		return
	}

	// Return response JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func ProcessPostRequest(w http.ResponseWriter, r *http.Request) {

	// Accept only POST
	if r.Method != "POST" {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create empty return JSON
	jsonData := simplejson.New()

	var postBody postBody
	// Decode post body
	err := json.NewDecoder(r.Body).Decode(&postBody)
	if err != nil {
		log.Println("Error in UnMarsal")
		jsonData.Set("status", "Error")
		jsonData.Set("status_code", "990")
		jsonData.Set("message", "Cannot Unmarshal Post Body")
		sendRes(w, jsonData)
		return
	}

	log.Println("PostBody", postBody.Content)

	err = Mongo_Update_Content(postBody.UserId, postBody.DocId, postBody.Content)
	if err != nil {
		log.Println("Error in Updating Content in Mongo")
		jsonData.Set("status", "Error")
		jsonData.Set("status_code", 400)
		jsonData.Set("message", "Failed to update content in Mongo DB")
		sendRes(w, jsonData)
		return
	}

	jsonData.Set("status", "Success")
	jsonData.Set("status_code", 200)
	jsonData.Set("message", "Content Updated")
	sendRes(w, jsonData)
}

func ProcessGetRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Create empty return JSON
	jsonData := simplejson.New()

	var userid string = r.FormValue("userid")
	var docid string = r.FormValue("docid")

	content, err := Mongo_Get_Content(userid, docid)
	if err != nil {
		log.Println("Error in Updating Content in Mongo")
		jsonData.Set("status", "Error")
		jsonData.Set("status_code", 400)
		jsonData.Set("message", "Failed to update content in Mongo DB")
		sendRes(w, jsonData)
		return
	}

	jsonData.Set("status", "Success")
	jsonData.Set("status_code", 200)
	jsonData.Set("content", content)
	sendRes(w, jsonData)

	return
}

func genUUID() string {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())

	return id.String()
}

func main() {

	log.Println("Devina Started")
	router := NewRouter()
	if err := http.ListenAndServe(":2007", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//Connection Related
	router.HandleFunc("/getcontent", ProcessGetRequest)
	router.HandleFunc("/postcontents", ProcessPostRequest)

	return router
}
