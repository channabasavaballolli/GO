package main

import (
	"fmt"
	"log"
	"net/http"

	"petclinic/db"
	"petclinic/handlers"
	"petclinic/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to PostgreSQL
	db.Connect()

	// Main router
	r := mux.NewRouter()

	// Global Logging Middleware
	r.Use(middleware.LoggingMiddleware)

	// PUBLIC ROUTES

	// Login route (Module 5 Authentication)
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// PROTECTED ROUTES (JWT)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware) // Protect everything under /api

	// Owner routes
	api.HandleFunc("/owners", handlers.CreateOwner).Methods("POST")
	api.HandleFunc("/owners", handlers.GetOwners).Methods("GET")
	api.HandleFunc("/owners/{id}", handlers.GetOwnerByID).Methods("GET")
	api.HandleFunc("/owners/{id}", handlers.UpdateOwner).Methods("PUT")
	api.HandleFunc("/owners/{id}", handlers.DeleteOwner).Methods("DELETE")

	// Pet routes
	api.HandleFunc("/pets", handlers.CreatePet).Methods("POST")
	api.HandleFunc("/pets", handlers.GetPets).Methods("GET")
	api.HandleFunc("/pets/{id}", handlers.GetPetByID).Methods("GET")
	api.HandleFunc("/pets/{id}", handlers.UpdatePet).Methods("PUT")
	api.HandleFunc("/pets/{id}", handlers.DeletePet).Methods("DELETE")

	// Appointment routes
	api.HandleFunc("/appointments", handlers.CreateAppointment).Methods("POST")
	api.HandleFunc("/appointments", handlers.GetAppointments).Methods("GET")
	api.HandleFunc("/appointments/{id}", handlers.GetAppointmentByID).Methods("GET")
	api.HandleFunc("/appointments/{id}", handlers.UpdateAppointment).Methods("PUT")
	api.HandleFunc("/appointments/{id}", handlers.DeleteAppointment).Methods("DELETE")

	// Start server
	fmt.Println("✅ Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
