package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevindalam/go-cassandra-api/cassandra"
	"github.com/kevindalam/go-cassandra-api/users"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	router.HandleFunc("/users", users.Get)
	router.HandleFunc("/users/new", users.Post)
	router.HandleFunc("/users/{user_uuid}", users.GetOne)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
