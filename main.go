package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		total int
		nb    int
		err   error
	)

	for i, arg := range os.Args[1:] {
		nb, err = strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("param %d must be an int", i)
			os.Exit(1)
		}

		total += nb
	}

	fmt.Println("result:", total)
}
