package entity

type User struct {
	ID           int
	Name         string
	SubmitID     string
	Sex          int
	Year         int
	PasswordHash string
	Clothes      []Clothe
}
