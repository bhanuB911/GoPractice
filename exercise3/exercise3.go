package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is where initialization for my program occurs")
}

func main(){

  x := rand.Intn(10)
  y := rand.Intn(10)

  fmt.Println(x, y)

  if x < 4 && y < 4{
    fmt.Println("Both are less than 4")
  } else if x > 6 && y > 6{
    fmt.Println("Both are greater than 6")
  } else if x >=4 && x <=6{
    fmt.Println("x is between 4 and 6")
  } else if y != 5{
    fmt.Println(" y is not equal to 5")
  } else {
    fmt.Println("None of the cases met")
  }

  switch{
  case x<4 && y <4:
    fmt.Println("Both are less than 4")
  case x>6 && y >6:
    fmt.Println("Both are greater than 6")
  case x>=4 && x<=6:
    fmt.Println("x is between 4 and 6")
  case y!=5:
    fmt.Println(" y is not equal to 5")
  default:
    fmt.Println("None of the cases met")
  }
}