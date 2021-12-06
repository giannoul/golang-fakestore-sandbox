package product

type Rating struct {
	Rate  float64
	Count int
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}
