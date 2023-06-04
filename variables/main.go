package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/packageone"
	"os"
	"time"
)

func main() {
	newString := packageone.PublicVar
	fmt.Println(newString)
}

const prompt = "and then press ENTER."

func mainOld2() {

	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2  // [2, 10]
	var secondNumber = rand.Intn(8) + 2 // [2, 10]
	var subtraction = rand.Intn(8) + 2  // [2, 10]
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guest the Number Game")
	fmt.Println("_____________________")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally though of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}

func mainOLD() {

	// first way
	var firstNumber int
	firstNumber = 2

	// second way
	var secondNumber = 5

	// third way
	thirdNumber := 7

	fmt.Println(firstNumber)
	fmt.Println(secondNumber)
	fmt.Println(thirdNumber)
}
