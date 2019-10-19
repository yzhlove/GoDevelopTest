package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("%d %v \n", i, int(i*10/100))
	}

}
