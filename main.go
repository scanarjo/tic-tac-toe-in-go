package main

import "os"

func main() {
	game := NewCLIGameWithIO(os.Stdin, os.Stdout)

	game.Start()
}
