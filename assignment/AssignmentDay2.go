//Create two functions where function 1 accepts 2 numbers and function 2 accepts 3 numbers and returns a number as it's sum.Print these values using fmt package

package main

import (
	"fmt"
)

//Function which returns sum of 2 numbers
func sumOf2numbers(num1 int, num2 int) int{
return num1 + num2
}

//Function which returns sum of 3 numbers
func sumOf3numbers(num1 int, num2 int, num3 int) int{
return num1 + num2 + num3
}

//main function
func main(){
	var num1, num2, num3 = 10, 20, 30
	fmt.Println("sum of 2 numbers(10,20) = ", sumOf2numbers(num1,num2))
	fmt.Println("sum of 3 numbers(10,20,30) = ", sumOf3numbers(num1,num2,num3))
}

