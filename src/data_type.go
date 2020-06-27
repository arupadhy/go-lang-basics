package main 

import "fmt"

func main() {
  // unsigned int 8 bit i.e 1 byte 0-255
  // uint16 /uint32/ uint64
  // explicit type definition 
  var a uint8 = 1
  var b byte = 3
  fmt.Println(a)
  fmt.Println(b)

  // signed integers
  // signed 8 bit (-128 to 127)
  var x int8 = -3
  fmt.Println(x)
  var y rune = -32423
  // int16 /int32 or rune /int64
  fmt.Println(x)
  fmt.Println(y) 

  // based on os
  // uint (32 or 64)
  // int (same size of uint)
  // uintptr

  // Floating point
  // float32 ( 32 bit floating point numbers)
  // float64 (64 bit floating point number)

  // Complex types
  // complex64
  // complex128

  // Strings - double quote
  // char - single quote
  // Booleans
  // true/false

  var number uint8 =  16
  fmt.Println("Unsgined integer: ", number + 30, "Number: ", number)

  // implicit type definition
  var implicit = 32
  fmt.Println("Implicit unit8 : ", implicit)
  fmt.Printf("%T ", implicit)

  // ommit var keyword
  number1 := 100
  fmt.Printf("%T", number1)

  boolx := true
  var bool2 bool
  numb := 100
  var v uint64
  fmt.Println(boolx)
  fmt.Println(numb)
  fmt.Println("deafult value for usigned int 64 ", v)
  fmt.Println("Default value for bool ", bool2)
}
