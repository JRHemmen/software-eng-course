package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jrhemmen/software-eng-course/demos/apidemo/internal/models"
)

func PingGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func VMsGet(w http.ResponseWriter, r *http.Request) {}

func VMsNameGet(w http.ResponseWriter, r *http.Request) {}

func VMsNameDelete(w http.ResponseWriter, r *http.Request) {}

func VMsNamePatch(w http.ResponseWriter, r *http.Request) {}

func VMsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Decode the request body into a VM struct
	var newVm models.Vm
	err := json.NewDecoder(r.Body).Decode(&newVm)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Check if VM already exists
	var existingVm models.Vm
	collection := MongoClient.Database("testDB").Collection("vms")
	err = collection.FindOne(ctx, map[string]string{"name": newVm.Name}).Decode(&existingVm)
	if err == nil {
		http.Error(w, "VM already exists", http.StatusConflict)
		return
	}

	// Set the created and updated times
	newVm.Created = time.Now()
	newVm.Updated = time.Now()

	// Insert the VM into the database
	_, err = collection.InsertOne(ctx, newVm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the VM as JSON
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newVm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HostsGet(w http.ResponseWriter, r *http.Request) {}

func HostsNameGet(w http.ResponseWriter, r *http.Request) {}

func HostsNameDelete(w http.ResponseWriter, r *http.Request) {}

func HostsNamePatch(w http.ResponseWriter, r *http.Request) {}

func HostsPost(w http.ResponseWriter, r *http.Request) {}

func ClustersGet(w http.ResponseWriter, r *http.Request) {}

func ClustersNameGet(w http.ResponseWriter, r *http.Request) {}

func ClustersNameDelete(w http.ResponseWriter, r *http.Request) {}

func ClustersNamePatch(w http.ResponseWriter, r *http.Request) {}

func ClustersPost(w http.ResponseWriter, r *http.Request) {}
