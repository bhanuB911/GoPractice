package main

import (
	"fmt"
)

func main(){

	// Exercise-42

	a := [5]int{0,1,2,3,4}
	fmt.Printf("%v \t %T\n", a, a)
	fmt.Println(len(a))

	for _, v := range a {
		fmt.Printf("%v - %T\n", v, v)
	}

	fmt.Println("")
	
    // Exercise-43

	b := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Printf("%v \t %T\n", b, b)
	fmt.Println(len(b))

	for _, v := range b {
		fmt.Printf("%v - %T\n", v, v)
	}
    fmt.Printf("%#v \n", b)
	fmt.Println("")
	fmt.Println("--------------------")


	// Exercise-44

	fmt.Println(b[0:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])
	fmt.Println("")
	fmt.Println("--------------------")

	// Exercise-45

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("")
	fmt.Println("--------------------")

	// Exercise-46

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(z)
    z = append(z[:3],z[6:]...)
	fmt.Println(z)
	fmt.Println("--------------------")


	// Exercise-47

	states := make([]string,0,50)
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,)
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i:=0; i<len(states); i++{
		fmt.Printf("Index - %v \t Value -%v \n", i, states[i])
	}
	fmt.Println("--------------------")

	// Exercise-48

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "MoneyPenny", "I'm 008."}
	fmt.Println(jb)
	fmt.Println(mp)

	xxs := [][]string{jb, mp}
	fmt.Println(xxs)

	for _, v1 := range xxs{
		//fmt.Printf("Outer Record - %v \t Type - %T\n", v1, v1)
		fmt.Println(v1)
		for _, v2 := range v1{
			//fmt.Println("Inner Record - %v \n", v2)
			fmt.Println(v2)
		} 
	}

	fmt.Println("--------------------")


}