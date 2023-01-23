package model

type Product struct {
	ProductId string   `json:"product_id"`
	Name      string   `json:"name"`
	Price     int      `json:"price"`
	Galleries []string `json:"galleies"`
}
