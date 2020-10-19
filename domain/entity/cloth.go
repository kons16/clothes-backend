package entity

type Cloth struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Price    string `db:"price"`
	ImageUrl string `db:"image_url"`
	// 服の追加する際に初期は base64 で受け取る
	ImageBase64 string
	Type        string `db:"type"`
}
