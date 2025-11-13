package handlers

import (
	"encoding/json"
	"net/http"
	"petclinic/db"
	"petclinic/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateOwner - POST /owners
func CreateOwner(w http.ResponseWriter, r *http.Request) {
	var o models.Owner
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(
		`INSERT INTO owners (name, contact, email) VALUES ($1, $2, $3) RETURNING id`,
		o.Name, o.Contact, o.Email,
	).Scan(&o.ID)
	if err != nil {
		http.Error(w, "Database insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(o)
}

// GetOwners - GET /owners
func GetOwners(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, contact, email FROM owners")
	if err != nil {
		http.Error(w, "Database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var owners []models.Owner
	for rows.Next() {
		var o models.Owner
		if err := rows.Scan(&o.ID, &o.Name, &o.Contact, &o.Email); err != nil {
			http.Error(w, "Error scanning data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		owners = append(owners, o)
	}

	json.NewEncoder(w).Encode(owners)
}

// GetOwnerByID - GET /owners/{id}
func GetOwnerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var o models.Owner
	err = db.DB.QueryRow("SELECT id, name, contact, email FROM owners WHERE id=$1", id).
		Scan(&o.ID, &o.Name, &o.Contact, &o.Email)
	if err != nil {
		http.Error(w, "Owner not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(o)
}

// UpdateOwner - PUT /owners/{id}
func UpdateOwner(w http.ResponseWriter, r *http.Request) {
	var o models.Owner
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(
		`UPDATE owners SET name=$1, contact=$2, email=$3 WHERE id=$4`,
		o.Name, o.Contact, o.Email, id,
	)
	if err != nil {
		http.Error(w, "Database update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Owner not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Owner updated successfully"})
}

// DeleteOwner - DELETE /owners/{id}
func DeleteOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("DELETE FROM owners WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Database delete failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Owner not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Owner deleted successfully"})
}
