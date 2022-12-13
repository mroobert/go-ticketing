package main

import (
	"fmt"

	"github.com/go-ticketing/pkgs/validator"
)

func main() {
	vld := validator.New()
	fmt.Println(vld)
}
