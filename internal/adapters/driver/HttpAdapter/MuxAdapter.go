package muxadapter

import (
	"fmt"
	"log"
	"net/http"
	"os"

	drivenadapters "github.com/AlejandroWaiz/goodfriendtimmy/internal/adapters/driven"
	"github.com/AlejandroWaiz/goodfriendtimmy/internal/domain"
	"github.com/gorilla/mux"
)

type MuxAdapter struct {
	domainport domain.DomainPort
	drivenport drivenadapters.DrivenPort
}

func CreateMuxAdapter(domain domain.DomainPort, driven drivenadapters.DrivenPort) *MuxAdapter {
	return &MuxAdapter{domainport: domain, drivenport: driven}
}

func (ma *MuxAdapter) ListenAndServe() {

	port, router := createRouter()

	router.HandleFunc("/", ma.HealthCheck).Methods("GET")
	router.HandleFunc("/Add", ma.AddHttpHandler).Methods("POST")
	router.HandleFunc("/Subtract", ma.SubtractHttpHandler).Methods("POST")
	router.HandleFunc("/Divide", ma.DivideHttpHandler).Methods("POST")
	router.HandleFunc("/Multiply", ma.MultiplyHttpHandler).Methods("POST")
	router.HandleFunc("/GetAll", ma.GetAllSavedOperations).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))

}

func createRouter() (string, *mux.Router) {

	port := fmt.Sprintf(":%s", os.Getenv("MuxPort"))
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("Starting router...")

	return port, router

}
