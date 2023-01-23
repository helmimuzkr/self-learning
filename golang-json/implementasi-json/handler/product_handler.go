package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/helmimuzkr/model"
)

func PostFormToJSON(w http.ResponseWriter, r *http.Request) {
	productId := r.PostFormValue("product_id")
	name := r.PostFormValue("name")
	price, _ := strconv.Atoi(r.PostFormValue("price"))
	galleries := []string{r.PostFormValue("gallery1"), r.PostFormValue("gallery2"), r.PostFormValue("gallery3")}

	product := &model.Product{
		ProductId: productId,
		Name:      name,
		Price:     price,
		Galleries: galleries,
	}

	fmt.Println(product)

	path := "json/product.json"

	// Deteksi apakah file sudah ada
	var _, err = os.Stat(path)
	// If file not exist, then create new file
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Created File!")
		file.Close()
	}

	// Open file fo JSON-encoding
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Marshaling
	encoder := json.NewEncoder(file)
	encoder.Encode(product)

	fmt.Println("Encoding completed.")
	w.Write([]byte("Submit data successfuly"))
}
