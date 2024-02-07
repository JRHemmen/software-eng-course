package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient *mongo.Client
)

func Start() {

	// Connect to MongoDB
	ConnectMongo("mongodb://mongouser1:password123@localhost:27017")

	// Start the API routes
	StartRoutes()

	// Disconnect Mongo on close
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	defer func() {
		if err = MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ConnectMongo(mongoURI string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		panic(err)
	}

	// Test the connection
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
}

func StartRoutes() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)                    // Log API requests
	r.Use(middleware.Recoverer)                 // Recover from panics without crashing the server
	r.Use(middleware.Timeout(60 * time.Second)) // Set a timeout for requests

	// Test routes
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/ping", PingGet)
	})

	// VM routes
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/vms", VMsGet)
		r.Get("/api/v1/vms/{name}", VMsNameGet)
		r.Delete("/api/v1/vms/{name}", VMsNameDelete)
		r.Patch("/api/v1/vms/{name}", VMsNamePatch)
		r.Post("/api/v1/vms", VMsPost)
	})

	// Hypervisor (Host) routes
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/hosts", HostsGet)
		r.Get("/api/v1/hosts/{name}", HostsNameGet)
		r.Delete("/api/v1/hosts/{name}", HostsNameDelete)
		r.Patch("/api/v1/hosts/{name}", HostsNamePatch)
		r.Post("/api/v1/hosts", HostsPost)
	})

	// Cluster routes
	r.Group(func(r chi.Router) {
		r.Get("/api/v1/clusters", ClustersGet)
		r.Get("/api/v1/clusters/{name}", ClustersNameGet)
		r.Delete("/api/v1/clusters/{name}", ClustersNameDelete)
		r.Patch("/api/v1/clusters/{name}", ClustersNamePatch)
		r.Post("/api/v1/clusters", ClustersPost)
	})

	http.ListenAndServe("0.0.0.0:8080", r)
}
