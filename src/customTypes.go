package main

import "fmt"
// go land does not have classes and relies on structs

type car struct {
  class string
  speed float64
  color string
  cost float64
}

// two types of method in go
// value receiver and pointer receiver

const usixteenBitMax float64 = 65535
const maxSpeed float64 = 1.60

// This is a method which is value reciever and can be called on car
// this is associated with car struct
func (c car) kmh() float64 {
  return c.speed * (c.speed/maxSpeed)
}

func (c car) mph() float64 {
  return c.kmh() / maxSpeed
}

// pointer receiver method will modify the value
func (c *car) updateColor(newColor string) {
  c.color = newColor
}

/**
 why would you ever use value receiver instead of pointer receiver

Value receiver makes a copy (efficient for small struct)
pointer receiver does not make a copy (efficient for large structs)

look at go code for when to use what

**/


func main() {
  myCar := car{
    class: "Sedan",
    speed: 123.3,
    color: "Dark Blue",
    cost: 342432.333}
/**
  could also be used as 
  myCar := car{"Sedan", 123.3, "Dark Blue", 342432.333}
  can be hard to read though
**/

  fmt.Println(myCar.class)
  fmt.Println("Calling method on car now ", myCar.kmh())
  fmt.Println("Calling method on car now with more ", myCar.mph())
  fmt.Println("Color of the color is: ", myCar.color)
  // this method call will modify myCar since the method accepts pointer 
  myCar.updateColor("Green")
  fmt.Println("Color of the car after updating is ", myCar.color)
}
