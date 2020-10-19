package entity

type Cordinate struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	TopClothID  int    `db:"top_cloth_id"`
	PantClothID int    `db:"pant_cloth_id"`
	UserID      int    `db:"user_id"`
}
