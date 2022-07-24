package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	k := strings.Split(a, "")
	for j := 0; j < len(k); j++ {
		r := k[j]
		m, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Ошикбка")
		}
		z := float64(m)
		s := math.Pow(z, 2)
		fmt.Println(s)
	}
}
