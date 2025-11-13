package tes

import (
	"fmt"
	"strconv"
)

func SumNumber(data int) int {
	dataString := strconv.Itoa(data)
	result := data
	for i := 0; i < len(dataString); i++ {
		convertedVal, _ := strconv.Atoi(string(dataString[i]))
		result += convertedVal

	}
	return result
}

func main() {
	var generatorNumber = make(map[int]int)
	result := 0
	for i := 1; i <= 5000; i++ {
		sumNumber := SumNumber(i)
		if generatorNumber[sumNumber] != 0 {
			result -= generatorNumber[sumNumber]
		}
		if generatorNumber[sumNumber] == 0 {
			generatorNumber[sumNumber] = i
			result += i
		}
	}
	fmt.Println("Result", result)
}
