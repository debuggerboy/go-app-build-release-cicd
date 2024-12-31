package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/debuggerboy/go-app-build-release-cicd/src/chaffCutter"
	"github.com/debuggerboy/go-app-build-release-cicd/src/errorHandler"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your Name: ")
	inputName, err := reader.ReadString('\n')
	errorHandler.CheckError(err, "Checking inputName")

	fmt.Print("Enter Your Age: ")
	inputAge, err := reader.ReadString('\n')
	errorHandler.CheckError(err, "checking inputAge")

	// The below code snippet, is now wrapped as chaffCutter.SnipStringToInt64 function
	//age, err := strconv.ParseInt(strings.TrimSpace(inputAge), 0, 32)
	age := chaffCutter.SnipStringToInt64(inputAge)
	name := chaffCutter.SnipSpaceFromString(inputName)

	fmt.Println("============= Debug =============")
	fmt.Printf("The type of inputAge is %T.\n", inputAge)
	fmt.Printf("The type of age is %T.\n", age)
	fmt.Println("============= Debug =============")

	fmt.Println("Hello", name, "! you are", age, "years old !!")

}
