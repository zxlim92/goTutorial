package main

import (
	"fmt"
	"math/rand"
)

func getTestMethod(str string) string {
	return "HELLO , " + str

}

func returnTwoVariables() (int, string) {
	return 100, "John"
}

func main() {
	var intArr [3]int32

	fmt.Println(intArr[0])
	fmt.Println(&intArr[0])

	intArr2 := [...]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intArr2)
	var intNum int64 = rand.Int63n(10000)

	var intSlice []int32 = []int32{4, 5, 6}
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	intSlice = append(intSlice, intArr[:]...)

	fmt.Println("Hello, Worldxd!")
	fmt.Println(intNum)
	fmt.Printf("getTestMethod: %s\n", getTestMethod("John"))

	var intNum2, stringNum2 = returnTwoVariables()
	fmt.Println(intNum2, stringNum2)
}
