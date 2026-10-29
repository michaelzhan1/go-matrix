package main

import (
	"fmt"

	"github.com/michaelzhan1/go-matrix/matrix/vector"
)

func main() {
	x := vector.Vector([]float32{1, 2, 3})
	y := vector.Vector([]float32{4, 5, 6})

	fmt.Println(x.Add(y))
}
