package handlers

import (
	"encoding/json"
	"net/http"
	"petclinic/db"
	"petclinic/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateAppointment - POST /appointments
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var a models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(
		`INSERT INTO appointments (pet_id, date, description) VALUES ($1, $2, $3) RETURNING id`,
		a.PetID, a.Date, a.Description,
	).Scan(&a.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(a)
}

// GetAppointments - GET /appointments
func GetAppointments(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, pet_id, date, description FROM appointments")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var a models.Appointment
		if err := rows.Scan(&a.ID, &a.PetID, &a.Date, &a.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		appointments = append(appointments, a)
	}

	json.NewEncoder(w).Encode(appointments)
}

// GetAppointmentByID - GET /appointments/{id}
func GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var a models.Appointment
	err = db.DB.QueryRow("SELECT id, pet_id, date, description FROM appointments WHERE id=$1", id).
		Scan(&a.ID, &a.PetID, &a.Date, &a.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(a)
}

// UpdateAppointment - PUT /appointments/{id}
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	var a models.Appointment
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(
		`UPDATE appointments SET pet_id=$1, date=$2, description=$3 WHERE id=$4`,
		a.PetID, a.Date, a.Description, id,
	)
	if err != nil {
		http.Error(w, "Database update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Appointment updated successfully"})
}

// DeleteAppointment - DELETE /appointments/{id}
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.DB.Exec("DELETE FROM appointments WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Database delete failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Appointment deleted successfully"})
}
