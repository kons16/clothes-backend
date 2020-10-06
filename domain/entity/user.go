package entity

type User struct {
	ID           int64
	Name         string
	SubmitID     string
	Sex          int64
	Year         int64
	PasswordHash string
	Clothes      []Clothe
}
