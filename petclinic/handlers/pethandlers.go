package handlers

import (
	"encoding/json"
	"net/http"
	"petclinic/db"
	"petclinic/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePet - POST /pets
func CreatePet(w http.ResponseWriter, r *http.Request) {
	var p models.Pet
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(
		`INSERT INTO pets (name, species, age, owner_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		p.Name, p.Species, p.Age, p.OwnerID,
	).Scan(&p.ID)
	if err != nil {
		http.Error(w, "Database insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

// GetPets - GET /pets
func GetPets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, species, age, owner_id FROM pets")
	if err != nil {
		http.Error(w, "Database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &p.Age, &p.OwnerID); err != nil {
			http.Error(w, "Error scanning data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		pets = append(pets, p)
	}

	json.NewEncoder(w).Encode(pets)
}

// GetPetByID - GET /pets/{id}
func GetPetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p models.Pet
	err = db.DB.QueryRow("SELECT id, name, species, age, owner_id FROM pets WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Species, &p.Age, &p.OwnerID)
	if err != nil {
		http.Error(w, "Pet not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(p)
}

// UpdatePet - PUT /pets/{id}
func UpdatePet(w http.ResponseWriter, r *http.Request) {
	var p models.Pet
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(
		`UPDATE pets SET name=$1, species=$2, age=$3, owner_id=$4 WHERE id=$5`,
		p.Name, p.Species, p.Age, p.OwnerID, id,
	)
	if err != nil {
		http.Error(w, "Database update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Pet updated successfully"})
}

// DeletePet - DELETE /pets/{id}
func DeletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.DB.Exec("DELETE FROM pets WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Database delete failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Pet deleted successfully"})
}
