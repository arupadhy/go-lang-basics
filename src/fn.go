package main

import "fmt"

func add(a, b float64) float64 {
  return a + b
}

// go function can return tuple kind of stuff
func multiple(a, b string) (string, string) {
  return a,b
}

const PI float64 = 3.14
func main() {
  var num1, num2 float64 = 3.5, 5.2
  /**
  same thing as above
  num1, num2 := 3.5, 5.2
  **/
  fmt.Println("const pi value: ", PI)
  fmt.Printf("%T \n", num1)
  fmt.Println(add(num1,num2))

  w1, w2 := "Hey", "there"
  fmt.Println(multiple(w1,w2))
  
  var int1 int = 32
  var b float64 = float64(int1)
  x := b;
 
  fmt.Println(" b as float64 ", b)
  fmt.Printf("%T \n", x)
}
