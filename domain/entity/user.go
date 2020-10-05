package entity

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Clothes      []Clothe
}
