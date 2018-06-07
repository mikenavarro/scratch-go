package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Cluster struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var clusters []Cluster

func main() {
	clusters = append(clusters, Cluster{"1", "cluster1"})
	clusters = append(clusters, Cluster{"2", "cluster2"})

	router := mux.NewRouter()
	router.HandleFunc("/cluster", GetClusters).Methods("GET")
	router.HandleFunc("/clusters", GetClusters).Methods("GET")
	router.HandleFunc("/cluster/{id}", GetCluster).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func GetClusters(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clusters)
	fmt.Println(clusters)
}

func GetCluster(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, c := range clusters {
		if c.ID == params["id"] {
			json.NewEncoder(w).Encode(c)
		} else {
			fmt.Println("shoot")
		}
	}

}
