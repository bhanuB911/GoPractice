package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is where initialization for my program occurs")
}

func main(){
	for i:=0; i<100; i++{
		fmt.Println(i)
	}

    for i:=0; i<100; i++{
        x := rand.Intn(10)
        y := rand.Intn(10)
        
        fmt.Printf("Iteration %v and Value of x is %v and y is %v\t", i, x, y)
        
        if x < 4 && y < 4{
            fmt.Printf("Both are less than 4\t")
            } else if x > 6 && y > 6{
                fmt.Printf("Both are greater than 6 \t")
            } else if x >=4 && x <=6{
                fmt.Printf("x is between 4 and 6\t")
            } else if y != 5{
                fmt.Printf(" y is not equal to 5\t")
            } else {
                fmt.Printf("None of the cases met\t")
            }
                        
        switch{
            case x<4 && y <4:
                fmt.Printf("Both are less than 4\n")
            case x>6 && y >6:
                fmt.Printf("Both are greater than 6\n")
            case x>=4 && x<=6:
                fmt.Printf("x is between 4 and 6\n")
            case y!=5:
                fmt.Printf(" y is not equal to 5\n")
            default:
                fmt.Printf("None of the cases met\n")
        }
    }
}