package main

import (
	"fmt"
)

func main() {
	// 整数型の変数iを宣言し、値を20にする。続いて、iを浮動小数点数型の変数fに代入し、iとfの値を出力せよ
	var i int = 20
	var f float64 = float64(i)
	fmt.Println("The value of i is:", i)
	fmt.Println("The value of f is:", f)

	// 整数変数と浮動小数点数変数の両方に代入できるvalueという定数を宣言するプログラムを書き、iという整数とfという浮動小数点変数に代入し、iとfを出力せよ
	const value = 42
	i = value
	f = float64(value)
	fmt.Println("The value of i is:", i)
	fmt.Println("The value of f is:", f)

	// byte型の変数b、int32型の変数smallI、uint64型の変数bigIのそれぞれに、その型の最大有効値を代入し、次に各変数に1を加え、その値を出力せよ
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b++
	smallI++
	bigI++
	fmt.Println("The value of b is:", b)
	fmt.Println("The value of smallI is:", smallI)
	fmt.Println("The value of bigI is:", bigI)
}
