package entity

type Cloth struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Price    string `db:"price"`
	ImageUrl string `db:"image_url"`
	Type     string `db:"type"`
	User     *User
}
