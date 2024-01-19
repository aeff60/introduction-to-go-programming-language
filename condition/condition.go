package main

import "fmt"

var score int

func main() {
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println(" score = ", score)
	if score >= 80 {
		fmt.Println("you got A grade")
	} else if score >= 70 {
		fmt.Println("you got B grade")
	} else if score >= 60 {
		fmt.Println("you got C grade")
	} else if score >= 50 {
		fmt.Println("you got D grade")
	} else {
		fmt.Println("you got F grade")
	}
}
