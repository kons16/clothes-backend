package entity

type Cordinate struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	TopClothID  string `db:"top_cloth_id"`
	PantClothID string `db:"pant_cloth_id"`
	UserID      int    `db:"user_id"`
}
