package main

import (
	"fmt"
	"math/rand"
)


func main(){
	/*for i:=0; i<100; i++{
		fmt.Println(i)
	}
	*/

    for i:=0; i<43; i++{
        x := rand.Intn(5)
        //y := rand.Intn(10)
        
        fmt.Printf("Iteration %v and  \t", i)
                        
        switch  x{
            case 0:
                fmt.Printf("Value of x is %v\n",x)
            case 1:
                fmt.Printf("Value of x is %v\n",x)
            case 2:
                fmt.Printf("Value of x is %v\n",x)
            case 3:
                fmt.Printf("Value of x is %v\n",x)
			case 4:
                fmt.Printf("Value of x is %v\n",x)
            default:
                fmt.Printf("Value of x is %v\n",x)
        }
    }
}