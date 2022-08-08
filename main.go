package main

import (
	"fmt"
	"strconv"
	"time"
)

func YearBorn(year string) int {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
	}

	thisYear := time.Now().Year()
	age := thisYear - yearInt
	fmt.Print("You have ", age, " years old. 😱")
	return age
}

func main() {
	var year string

	fmt.Print("Write in which year you were born: ")
	fmt.Scan(&year)

	YearBorn(year)
}
