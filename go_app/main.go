// Created by: Charlotte Jhu
// Created on: May 2023
//
// This program tells the user what triangle they have

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function tells the user what triangle they have
	// variable
	var lengthA float64
	var lengthB float64
	var lengthC float64

	// input
	fmt.Println("Let's find out what triangle you have!")
	fmt.Print("Enter the first length: ")
	fmt.Scanln(&lengthA)
	fmt.Print("Enter the second length: ")
	fmt.Scanln(&lengthB)
	fmt.Print("Enter the third length: ")
	fmt.Scanln(&lengthC)
	fmt.Println("")

	// process	
	var angleA = math.Acos((math.Pow(lengthB, 2)+math.Pow(lengthC, 2)-math.Pow(lengthA, 2))/(2*lengthB*lengthC)) * (180 / math.Pi)
	var angleB = math.Acos((math.Pow(lengthC, 2)+math.Pow(lengthA, 2)-math.Pow(lengthB, 2))/(2*lengthC*lengthA)) * (180 / math.Pi)
	var angleC = math.Acos((math.Pow(lengthA, 2)+math.Pow(lengthB, 2)-math.Pow(lengthC, 2))/(2*lengthA*lengthB)) * (180 / math.Pi)

	var sumOfAngles = angleA + angleB + angleC
	sumOfAngles = math.Round(sumOfAngles)

	angleA = math.Round(angleA)
	angleB = math.Round(angleB)
	angleC = math.Round(angleC)

	if sumOfAngles == 180 {
		if (angleA == 90 || angleB == 90 || angleC == 90) {
			// ouput
			fmt.Println("You have a right triangle.")
		} else if math.Floor(angleA) == 60 && math.Floor(angleB) == 60 && math.Floor(angleC) == 60 {
			// ouput
			fmt.Println("You have an equilateral triangle.")
		} else if angleA == angleB || angleB == angleC || angleC == angleA {
			// ouput
			fmt.Println("You have an isosceles triangle.")
		} else if angleA != angleB && angleB != angleC && angleC != angleA {
			// ouput
			fmt.Println("You have a scalene triangle.")
		} else {
			// ouput
			fmt.Println("You do not have a triangle.")
		}
	} else {
		// ouput
		fmt.Println("You do not have a triangle.")
	}
	fmt.Println("\nDone.")
}
