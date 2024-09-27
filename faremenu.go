package main

import (
	"fmt"
)

func main() {
	var origin, destination, cabinClass string

	fmt.Println("**** Airfare Calculator ****")

	var validOrginEntered = false
	var originCity City
	var originErr error

	for !validOrginEntered {

		fmt.Print("enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originErr = getCityFromCode(origin)

		if originErr == nil {
			fmt.Println("You've entered " + originCity.cityName)
			validOrginEntered = true
		} else {
			fmt.Println(originErr)
		}
	}
	var validDestinatioEntered = false
	var destinationCity City
	var destinatioErr error

	for !validDestinatioEntered {

		fmt.Print("enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinatioErr = getCityFromCode(destination)

		if originErr == nil {
			fmt.Println("You've entered " + destinationCity.cityName)
			validDestinatioEntered = true
		} else {
			fmt.Println(destinatioErr)
		}
	}
	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {
		for i := range cabinClasses{
			fmt.Println(cabinClasses[i].code + "=" + cabinClasses[i].className)
		}
		fmt.Print("Enter cabin class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println("You've entered " + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}
	distance := CalculateDistance(float64(destinationCity.longtitude)/1000,
		float64(destinationCity.latitude)/1000,
		float64(originCity.longtitude)/1000,
		float64(originCity.latitude)/1000)
	fmt.Printf("\nDistance = %.1f km\n", distance)
	rateFloat64 := float64(enteredCabinClass.rate) / 100

	fmt.Printf("$ per km = %.2f\n", rateFloat64)
	fmt.Printf("Total fare = $%.2f\n", rateFloat64*float64(distance))
}
