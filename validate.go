package main

import (
	"errors"
	"fmt"
)

func getCityFromCode(codeToLookup string) (City, error){
	for _, item := range cities{
		if item.code == codeToLookup{
			return item, nil
		}
	}
	message := fmt.Sprintf("%s is not a valid city code", codeToLookup)
	return City{"", "", -999, -999}, errors.New(message)
}

func getCabinFromCode(codeToLookup string) (CabinClass, error){
	for _, item := range cabinClasses{
		if item.code == codeToLookup{
			return item, nil
		}
	}
	message := fmt.Sprintf("%s is not a valid cabin class code", codeToLookup)
	return CabinClass{"", "", -999}, errors.New(message)
}