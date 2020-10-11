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

// ログイン後に得られる構造体
type LoginGetUser struct {
	ID           int    `db:"id"`
	PasswordHash string `db:"password_hash"`
}
