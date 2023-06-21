package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is where initialization for my program occurs")
}

func main(){
  x := rand.Intn(250)
  fmt.Printf("Value of x is %v and the Type of x is %T\n",x, x)
  
  if x >= 0 && x <= 100{
    fmt.Println("Between 0 and 100")
  }  else if x >=101 && x <= 200{
    fmt.Println("Between 101 and 200")
  }  else{
    fmt.Println("Between 201 and 250")
  }

  //SwitchCase

  switch{
  case x <=100:
	fmt.Println("Between 0 and 100")
  case x >=101 && x <=200:
	fmt.Println("Between 101 and 200")
  case x>=201:
	fmt.Println("Between 201 and 250")
  default:
	fmt.Println("Either greater than 250 or less than 0")
  }

}