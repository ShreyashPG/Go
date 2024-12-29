package main

import "fmt"

type Person struct {
	first string ;
	last string ;
	age int ;
	sayings []string;
}

func main() {
	fmt.Println("This is how strcut works");

	Person1 := Person{
		first: "James",	
		last: "Bond",
		age : 38,
		sayings : []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	Person2 := Person{
		first: "Jammy",	
		last: "Bond",
		age : 38,
		sayings : []string{"Shaken", "Wishes?", "Say never"},
	}
	Person3 := Person{
		first: "Jefer",	
		last: "Bond",
		age : 38,
		sayings : []string{"Not stirred", "Any wishes?", "Never say"},
	}
	fmt.Println(Person1);
	fmt.Println(Person2);
	fmt.Println(Person3);

	map1 :=map[string]Person{"James":Person1, "Jammy":Person2, "Jefer":Person3};

	for i, v := range map1 {
		fmt.Println(i, v);
	}

	persons:=[]Person{Person1, Person2, Person3};

	for i, v := range persons {
		fmt.Println(i, v);
	}

}