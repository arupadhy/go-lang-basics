package main

import "fmt"

func main() {
  x := 10
  a := &x // memory address of x is stored in a

  fmt.Println("Value of a is : ", a)
  fmt.Println("Value at a is: ", *a)
  
  *a = 100
  fmt.Println(x)
  
  fmt.Println("*a * *a : ", *a * *a)
}
