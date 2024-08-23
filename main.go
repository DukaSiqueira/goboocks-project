package main

import (
	"fmt"
	"gobooks/internal/service"
)

func main() {
	book := service.Book{
		ID:     1,
		Title:  "Selva",
		Author: "Eduardo",
		Genre:  "Ação",
	}

	fmt.Println(book)
}
