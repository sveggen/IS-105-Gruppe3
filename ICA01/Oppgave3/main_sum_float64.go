package main

import (
	"fmt"
	"os"
	"strconv"
	"./sum"
)

func main() {

	argsNumber, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error occured")
	}

	argsNumber2, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		fmt.Println("Error occured")
	}


	tempVar := sum.SumFloat64(float64(argsNumber), float64(argsNumber2))

	fmt.Println(tempVar)

}