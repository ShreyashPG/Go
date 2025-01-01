package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type Don struct {
	Person
	saying []string
}

func driveVehicle(first string , age int) bool {

		fmt.Println("Lest see if ",first,"is eligible to drive or not ?")
		if(age >=18 ){
			return true;
		}
		return false;
}


func changeAge(p Person ){
		p.age=p.age+100;
}


func main() {
	fmt.Println("Hello Everyone on Board !!")

	p1 := Person{
		first: "Baby",
		last:  "John",
		age:   28,
	}

	d1 := Don{
		Person: p1,
		saying: []string{"Call me Father !!", "Yes this is it!!"},
	}

    driveFlag:=driveVehicle(d1.first,d1.age);

	fmt.Println(driveFlag);
	// fmt.Println("")
	
	//Age of person p1 does not changes and he can live
	changeAge(p1);
	fmt.Println(p1.age);

	fmt.Println(d1)


	fmt.Printf("%t\n",p1.age);
	fmt.Printf("%t\n",p1);



	}
