package main

import "fmt"

func main(){

	a := 10
	b := 20
	c := a + b
	fmt.Println(c)

	for i:=0 ; i<20 ;i++ {
		fmt.Println(i);
	}


	//slice of int
	a1 := []int {10, 20 , 20 , 40};

	//slice of string
	a4 := []string {"James", "Bond", "Miss"};
	fmt.Println(a4);

	fmt.Println(a1);

	for i, v := range a1 {
		fmt.Println(i, v);
	}

	a2:=map [string] int{"James": 38, "Bond": 40, "Miss": 30};

	for i , v := range a2 {
		fmt.Println(i,v);
	}

}