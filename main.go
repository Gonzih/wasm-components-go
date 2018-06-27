package main

import (
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %s\r\n", err)
		panic(err)
	}
}

func main() {
	// c := make(chan struct{}, 0)

	// store := NewStore()

	cmp, err := NewComponent("helloTemplate", "root", func(cmp *GenericComponent) error {
		cmp.props = struct {
			Label string
		}{
			Label: "markup from props",
		}

		return nil
	})

	checkErr(err)
	checkErr(cmp.Render())

	// <-c
}
