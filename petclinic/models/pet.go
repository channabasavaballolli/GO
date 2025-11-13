package models

// pet belonging to an owner
type Pet struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
	OwnerID int    `json:"owner_id"`
}
