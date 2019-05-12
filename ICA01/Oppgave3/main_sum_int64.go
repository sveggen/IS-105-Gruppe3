package main

import (
	"fmt"
	"os"
	"strconv"
	"./sum"
)

func main() {

	argsNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error occured")
	}

	argsNumber2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Error occured")
	}


	tempVar := sum.SumInt64(int64(argsNumber), int64(argsNumber2))

	fmt.Println(tempVar)

}