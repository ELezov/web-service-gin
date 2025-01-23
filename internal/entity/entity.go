package entity

type Wish struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	IsReserved  bool    `json:"is_reserved"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
	ProductUrl  string  `json:"product_url"`
}
