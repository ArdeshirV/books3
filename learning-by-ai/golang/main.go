package main

import (
	"errors"
	"fmt"

	c "github.com/ArdeshirV/gocolor"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(c.Error(rec))
		}
		fmt.Print(c.NormalText(""))
	}()

	fmt.Println(c.Prompt("Learning Golang by AI "))
	mainDivide()
}

func mainDivide() {
	result, err := divide(10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return a / b, nil
}
