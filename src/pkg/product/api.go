package product

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getHttpRequestToApi(url string) *[]byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //func ReadAll(r io.Reader) ([]byte, error)
	if err != nil {
		log.Fatal(err)
	}
	return &body
}

func (p *Product) GetById(id int) Product {
	var b *[]byte
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	b = getHttpRequestToApi(url)
	json.Unmarshal([]byte(*b), &p)
	return *p
}

func (p *Product) GetAll() []Product {
	var b *[]byte
	var prods *[]Product
	url := "https://fakestoreapi.com/products"
	b = getHttpRequestToApi(url)
	json.Unmarshal([]byte(*b), &prods)
	return *prods
}
