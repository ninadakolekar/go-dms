package models

// User ... Represents a user
type User struct {
	AvailableApp  bool
	AvailableAuth bool
	AvailableCr   bool
	AvailableInit bool
	AvailableQA   bool
	AvailableRw   bool
	Username      string
	Name          string
}
