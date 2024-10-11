package main

import (
	"fmt"
	"online-with-go/api"
)

func main() {
	keyboard := api.Keyboard{}

	for {
		key, err := keyboard.GetSingleKey()

		if err != nil {
			return
		}

		fmt.Println(key)
	}

}
