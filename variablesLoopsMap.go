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

	a1 := []int {10, 20 , 20 , 40};

	fmt.Println(a1);

	for i, v := range a1 {
		fmt.Println(i, v);
	}

	a2:=map [string] int{"James": 38, "Bond": 40, "Miss": 30};

	for i , v := range a2 {
		fmt.Println(i,v);
	}

}