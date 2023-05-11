// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: May 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  // This function tells the user what kind of triangle they have
  // input
  const lengthA = parseFloat(document.getElementById("length-a").value)
  const lengthB = parseFloat(document.getElementById("length-b").value)
  const lengthC = parseFloat(document.getElementById("length-c").value)

  // process
  // using cosine law
  const angleA = Math.acos((lengthB**2 + lengthC**2 - lengthA**2) / (2 * lengthB * lengthC)) * (180/Math.PI)
  const angleB = Math.acos((lengthC**2 + lengthA**2 - lengthB**2) / (2 * lengthC * lengthA)) * (180/Math.PI)
  const angleC = Math.acos((lengthA**2 + lengthB**2 - lengthC**2) / (2 * lengthA * lengthB)) * (180/Math.PI)
  const sumOfAngles = Number((angleA).toFixed(2)) + Number((angleB).toFixed(2)) + Number((angleC).toFixed(2))
  
  if (sumOfAngles == 180) {
    if (angleA==90 || angleB==90 || angleC==90) {
      // ouput
      document.getElementById("answer").innerHTML = "You have a right triangle."
    } else if (angleA==60 && angleB==60 && angleC==60) {
      // ouput
      document.getElementById("answer").innerHTML = "You have an equilateral triangle."
    } else if (angleA==angleB || angleB==angleC || angleC==angleA) {
      // ouput
      document.getElementById("answer").innerHTML = "You have an isosceles triangle."
    } else (angleA != angleB && angleB != angleC && angleC != angleA) {
      // ouput
      document.getElementById("answer").innerHTML = "You have a scalene triangle."
    }
  } else { 
    // ouput
    document.getElementById("answer").innerHTML = "You do not have a triangle."
  }
}
