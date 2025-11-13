package models

// appointment for a pet
type Appointment struct {
	ID          int    `json:"id"`
	PetID       int    `json:"pet_id"`
	Date        string `json:"date"` // 🟢 Add this field
	Description string `json:"description"`
}
