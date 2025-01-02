package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")
	var names []string;
	// names={"hemal"};
	names = append(names, "hemal")
	fmt.Println(names);

	var sayings = [3]string{"World is full of anger","Be kind and spread love","Hell yess"};
	fmt.Println(sayings);

	var acceptings = [...]string{"let go count and we shall relax","Best of go"};
	fmt.Println(acceptings);
	
	for i,v := range sayings{
		fmt.Println("The ",i+1," th Saying is : ",v);
	} 
}