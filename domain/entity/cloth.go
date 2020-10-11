package entity

type Cloth struct {
	ID       int
	Name     string
	Price    string
	ImageUrl string
	User     *User
}
