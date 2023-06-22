
package main

import (
	"fmt"
)

type Person struct{
	firstname string
	lastname string
	flavour []string
}

type Engine struct{
	electric bool
}


func main() {

	//Exercise-53

	p1 := Person{
		firstname: "James",
		lastname: "Bond",
		flavour: []string{"Chocolate", "banana", "guaua with mango"},
	}

	p2 := Person{
		firstname: "MoneyPenny",
		lastname: "Jenny",
		flavour: []string{"Mango", "strawberry"},
	} 

	fmt.Printf("%T - %#v\n", p1,p1)
	fmt.Printf("%T - %#v\n", p2,p2)

	fmt.Println("Favourite Ice Cream flavours \t", p1.flavour)
	fmt.Println("Favourite Ice Cream flavours \t", p2.flavour)


	for _, v := range p1.flavour {
		fmt.Println(p1.firstname, "Favourite is", v)
	}

	for _, v := range p2.flavour {
		fmt.Println(p2.firstname, "Favourite is", v)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	//Exercise-54

	am := map[string]Person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for _, v := range am{

		fmt.Println(v)
		for _, v2 := range v.flavour {
			fmt.Println(v.firstname, v.lastname, v2)
		}
	
	}

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	//Exercise-55
	type engine struct{
		electric bool
	}

	type vehicle struct{
		engine
		make string
		model string
		doors string
		color string
	}

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make: "German",
		model: "Y",
		doors: "Automatic",
		color: "Blue",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make: "korean",
		model: "polo",
		doors: "Manual",
		color: "Grey",
	}

	fmt.Println(v1.electric,v1.make,v1.model,v1.doors,v1.color)
	fmt.Println(v1.electric,v2.make,v2.model,v2.doors,v2.color)

	fmt.Println(v1.engine.electric, v1.color)
	fmt.Println(v1.engine.electric,v2.color)

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	//Exercise-56

	as1:= struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "James Bond",
		friends: map[string]int{
			"MoneyPenny": 27,
			"Dr_No": 44,
			"Q": 28,
			"M": 38,
		},
		favDrinks: []string{"Vodka", "Martini"},
	}

	for k,v := range as1.friends{
		fmt.Println(as1.first,"-friends-", k,v)
	}

	for _,v := range as1.favDrinks{
		fmt.Println(as1.first,"-drinks-",v)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")
}

