package main

import (
	"fmt"

	"github.com/giannoul/golang-fakestore/pkg/product"
)

func main() {
	p := product.Product{}
	fmt.Printf("%#v", p.GetById(1))
	fmt.Println()
	fmt.Printf("%#v", p.GetAll())

}
