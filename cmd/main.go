package main

import (
	"clean-architecture/internal/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
