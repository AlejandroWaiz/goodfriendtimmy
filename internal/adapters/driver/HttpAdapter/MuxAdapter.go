package muxadapter

import (
	"fmt"
	"github.com/AlejandroWaiz/goodfriendtimmy/internal/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type MuxAdapter struct {
	domainport domain.DomainPort
}

func CreateMuxAdapter(domain domain.DomainPort) *MuxAdapter {
	return &MuxAdapter{domainport: domain}
}

func (ma *MuxAdapter) ListenAndServe() {

	port, router := createRouter()

	router.HandleFunc("/", ma.HealthCheck).Methods("GET")
	router.HandleFunc("/Add", ma.AddHttpHandler).Methods("POST")
	router.HandleFunc("/Subtract", ma.SubtractHttpHandler).Methods("POST")
	router.HandleFunc("/Divide", ma.DivideHttpHandler).Methods("POST")
	router.HandleFunc("/Multiply", ma.MultiplyHttpHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))

}

func createRouter() (string, *mux.Router) {

	port := fmt.Sprintf(":%s", "3000")
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("Starting router...")

	return port, router

}
