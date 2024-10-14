package main

import "online-with-go/internal/screen"

func main() {
	s := screen.NewScreen(10, 5)

	s.Draw()
}
