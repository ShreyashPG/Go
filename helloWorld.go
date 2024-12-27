package main 

import "fmt";

func main(){
	fmt.Println("Hello World from Go")

	s:= fmt.Sprintln("Hello  World 2nd time from Go")

	fmt.Println(s)

	// foo();
	foo()

	// a,b:=bar("James Bond", 38)
	a, b:=bar("James Bond", 38)
	fmt.Println(a);
	fmt.Println(b, "is the age of James bond after 10 years")

	fmt.Println(`Hello This 
	can 
	be  
	written 
	this way 
	also `)
}

func foo (){
	// fmt.Print("The foo is here")
	fmt.Println("The foo is here")
}

func bar(x string , y int) (string , int ) {
	// return 	fmt.Sprintln(x,"is here ! ");
	return 	fmt.Sprint(x," is here ! "), y+10;


}