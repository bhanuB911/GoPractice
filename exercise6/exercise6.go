package main

import (
	"fmt"
	"math/rand"
)

func main(){
	i := 20
	for i > -5 {
		fmt.Println(i)
		if i==0{
			break
		}
		i--
	}

	for j:=0; j<100; j++{
		if j%2== 1 {
			fmt.Printf("%v is Odd number\n", j)
		}
	}

	// nested loops
	for m:=0; m<5; m++{
		for n:=0; n<5; n++{
			fmt.Printf("Outer Loop value is %v \t Inner Loop Value is %v \n", m, n)
		}
		fmt.Println("")
	}

	// For Loop, Data structure - Slice

	xi := []int{42,43,44,45,46,47}

	for i,v := range xi{
		fmt.Println("Ranging over a slice ",i,v)
	}

	//For Loop, Data Structre - Map

	m := map[string]int{
		"James": 42,
		"MoneyPenny": 32,
	}

	for k,v := range m{
		fmt.Println("Ranging over a Map", k, v)
	}

	//Exercise 32
	age1 := m["James"]
	fmt.Println(age1)

	// COMMA OK IDIOM
	if v, ok := m["James"]; ok{
		fmt.Printf("There is a Bond look up and the age of James Bond is %v \n", v)
	}

	age2 := m["Q"]
	fmt.Println(age2)

	// Comma ok IDIOM
	if v, ok := m["Q"]; !ok{
		fmt.Printf("There is no value of Q present in the map so the zero value of the int will be %v\n", v)
	}
	


	//Statement-Statement IDIOM
    c := 1
	for p:=0; p<100; p++{
		if z := rand.Intn(5); z ==3{
			fmt.Printf("Iteration %v \t total count %v \t z is %v\n", p, c, z)
			c++
		}
	}


	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}