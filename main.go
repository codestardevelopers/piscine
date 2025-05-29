package main

import (
	"fmt"
	"piscine/piscine"
)

func main() {
	s := "Hello 78 World!  a 4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
