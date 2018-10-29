package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := newArr()
	fmt.Println(a)

	fmt.Println(a.shuffle())
}

type arr []int8

func newArr() arr {
	return arr{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func (a arr) shuffle() arr {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range a {
		np := r.Intn(len(a) - 1)
		a[i], a[np] = a[np], a[i]
	}
	return a
}
