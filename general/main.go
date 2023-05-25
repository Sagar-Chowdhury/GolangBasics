package main

import "fmt"

//note first name then type.

func addNumbers(a int, b int) int {
	return a + b
}

func takeInput() {

	var name string
	var age int

	fmt.Print("Enter your name and age (space-separated): ")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Println(name)
	fmt.Println(age)

}

func main() {

	var message string = "Hello, World"

	fmt.Println(message)

	var number int = 42
	fmt.Println(number)

	sum := addNumbers(100, 1000) //declare and initialize at same time.
	fmt.Println(sum)

	fmt.Println(addNumbers(70, 30))

	age := 18
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("baccha")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i * 100)
	}

	//takeInput()

	var num int = 42
	var ptr *int = &num
	fmt.Println("Value of num before:", num)

	*ptr = 99 // Modifying the value through the pointer

	fmt.Println("Value of num after:", num)

}
