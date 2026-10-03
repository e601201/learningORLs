package main

import (
	"fmt"
)

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	fmt.Println(x, y, z) // 4 2 2
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x) // x: [a b i j y]
	fmt.Println("y:", y) // y: [a b i j y]
	fmt.Println("z:", z) // z: [i j y]

	x2 := []int{1, 2, 3, 4}
	num := copy(x2[:3], x2[1:]) // [1 2 3]  ← [2 3 4]
	fmt.Println(x2, num) // [2 3 4 4] 3

	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:] // 配列からスライスを作成
	fmt.Println(xArray, xSlice)

	xSlice2 := []int{1, 2, 3, 4}
	xArray2 := [4]int(xSlice2) // スライスから配列を作成
	fmt.Println(xArray2, xSlice2)
	smallArray := [2]int(xSlice2)
	xSlice2[0] = 10
	fmt.Println(xSlice2) // [10 2 3 4]
	fmt.Println(xArray2) // [1 2 3 4]
	fmt.Println(smallArray) // [1 2]
}
