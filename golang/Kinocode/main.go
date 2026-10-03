package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name   string
}

func (u User) bmi(weight, height float64) (result float64) {
	result = (weight / height / height) * 10000
	return
}

type Student struct {
	name string
	math, english float64
}

func (s Student) avg() float64 {
	return (s.math + s.english) / 2
}

func main() {
	// 正常な変数
	num := 42
	num01 := 3.14
	num_02 := 100
	fmt.Println("The value of num is:", num)
	fmt.Println("The value of num01 is:", num01)
	fmt.Println("The value of num_02 is:", num_02)

	// 無効な変数
	// 42num := 42 // このような変数名は使用できません
	// num-01 := 3.14 // このような変数名は使用できません

	//数値型
	var intNum int = 42
	var floatNum float64 = 3.14
	var uintNum uint = 100
	fmt.Println(
		"intNum:", intNum,
		"floatNum:", floatNum,
		"uintNum:", uintNum,
		"types:", reflect.TypeOf(intNum), reflect.TypeOf(floatNum), reflect.TypeOf(uintNum),
	)

	// 論理演算子
	x := 8
	y := 5
	fmt.Println("x > y:", x > y)   // true
	fmt.Println("x < y:", x < y)   // false
	fmt.Println("x == y:", x == y) // false
	fmt.Println("x != y:", x != y) // true

	// 文字列型
	var str string = "Hello, Kinocode!"
	fmt.Println("str:", str, "type:", reflect.TypeOf(str))

	// ブール型
	a := 10
	b := 20
	result := a < b // true
	fmt.Println("boolVal:", result, "type:", reflect.TypeOf(result))

    // 配列型
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[0], arr[1], arr[2])
	fmt.Println("arr:", arr, "type:", reflect.TypeOf(arr))

	// スライス型
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice, "type:", reflect.TypeOf(slice))

	// マップ型
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("m:", m, "type:", reflect.TypeOf(m))

	// 構造体型
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("p:", p, "type:", reflect.TypeOf(p))

	// 条件分岐
	if age := 20; age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// ループ処理
	for i := range 5 {
		if i == 3 {
			break // ループを終了
		}
		fmt.Println("i:", i)
	}

	arr2 := [...]int{1, 2, 3, 4, 12}
	sum := 0
	for _, V := range arr2 {
		fmt.Println("V:", V)
		sum += V
	}
	fmt.Println("Sum:", sum) // 15になる

	// 関数
	sayHello("Kinocode")
	result1 := cal(10, 5)
	fmt.Println("Result:", result1)

	// 無名関数
	func() {
		fmt.Println("This is an anonymous function.")
	}() // 即時実行

	// クロージャ
	closure := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}
	add5 := closure(5)
	fmt.Println("add5(3):", add5(3)) // 8になる

	//構造体
	var u User
	u.name = "Alice"
	fmt.Println("User:", u)
	ubmi := u.bmi(90, 172) //60kg, 1.7mのBMIを計算
	fmt.Println("Alice's BMI:", ubmi)


	student1 := Student{name: "Alice", math: 90, english: 85}
	fmt.Println("Student1:", student1)
	fmt.Println("Average score:", student1.avg())
}

func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func cal(a, b int) (r int) {
	r = a + b
	return 
}