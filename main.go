package main

import (
	"fmt"
	"foodWander/src/routes"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	fmt.Println("Starting server on port 8080...")

	route := routes.Router()
	err := route.Run(":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
