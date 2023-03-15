package web

import (
	"errors"
	"log"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

type Products struct {
	Products []Product
}

// construct
func NewProduct(id int, name string, price int) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

/* func NewProducts() *Products{
	products := Products{
		Products: []Product{
			{1, "Mouse", 150},
			{2, "Keyboard", 400},
			{3, "Headset", 400},
		},
	}
	return &products
} */

// method products
func (products *Products) getProduct() []Product {
	if products == nil {
		return nil
	}
	return products.Products
}

func (products *Products) findProductById(id int) (*Product, error) {
	for _, val := range products.Products {
		if val.ID == id {
			return &Product{
				ID: val.ID,
				Name: val.Name,
				Price: val.Price,
			}, nil
		}
	}
	return &Product{}, errors.New("Product tidak ditemukan")
}

func (products *Products) addProduct(product *Product) {
	products.Products = append(products.Products, *product)
	log.Printf("%s Berhasil ditambahkan!", product.Name)
}
