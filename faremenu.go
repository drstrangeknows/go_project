package main

import "fmt"

func main() {
	var origin string

	fmt.Println("**** Flying High Fare Calculator ****")

	validOriginCode := false
	var originCity City
	var originError error

	for !validOriginCode {
		fmt.Print("Enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originError = getCityFromCode(origin)

		if originError == nil {
			fmt.Println("You have entered city " + originCity.cityName)
			validOriginCode = true
		} else {
			fmt.Println(originError)
		}
	}

	fmt.Println("You've entered origin " + origin)

	fmt.Println(cities[0].cityName)
	fmt.Println(cities[1].latitude)
}
