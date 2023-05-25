package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')

	fmt.Print("Enter a rating")
	rat, _ := reader.ReadString('\n')
	rat = strings.TrimSuffix(rat, "\n")
	rat = strings.TrimSuffix(rat, "\r")
	val, err := strconv.ParseFloat(rat, 32)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(val)

	fmt.Printf("Hello, %s! You are %s years old.\n", name, age)

}
