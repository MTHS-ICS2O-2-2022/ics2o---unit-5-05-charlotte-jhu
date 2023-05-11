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
	var angleA = Math.acos((lengthB**2+lengthC**2-lengthA**2)/(2*lengthB*lengthC)) * (180 / Math.PI)
	var angleB = Math.acos((lengthA**2+lengthC**2-lengthB**2)/(2*lengthA*lengthC)) * (180 / Math.PI)
	var angleC = Math.acos((lengthA**2+lengthB**2-lengthC**2)/(2*lengthA*lengthB)) * (180 / Math.PI)
	var sumOfAngles = Number((angleA)%.2f) + Number((angleB)%.2f) + Number((angleC)%.2f)

	// input
	fmt.Println("Let's find out what triangle you have!")
	fmt.Println("Enter the first length: ")
	fmt.Scanln(&lengthA)
	fmt.Println("Enter the second length: ")
	fmt.Scanln(&lengthB)
	fmt.Println("Enter the third length: ")
	fmt.Scanln(&lengthC)

	// process
	if sumOfAngles == 180 {
		if (angleA == 90 || angleB == 90 || angleC == 90) {
			// ouput
			fmt.Println("You have a right triangle.")
		} else if (angleA%.2f == 60 && angleB%.2f == 60 && angleC%.2f == 60) {
			// ouput
			fmt.Println("You have an equilateral triangle.")
		} else if (angleA == angleB || angleB == angleC || angleC == angleA) {
			// ouput
			fmt.Println("You have an isosceles triangle.")
		} else if (angleA != angleB && angleB != angleC && angleC != angleA) {
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
}
