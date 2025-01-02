package main 

import "fmt"

func main (){

		var a =10;
		fmt.Println(a);

		var  b int;
		b=1;
		fmt.Println(b);

		var c string;
		c="Shreyash";
		fmt.Println(c);
		
		var e="Ghanekar"
		var d=c+e;
		fmt.Println(d);

		var c1 byte;
		c1='c';
		fmt.Println(c1);

		
		fmt.Println(len(d));

		var t1 = []int{1,23,234,345};
		fmt.Println(t1);
		fmt.Println(len(t1));

		fmt.Println(c[0:2]);
		fmt.Println(c[:2]);
		fmt.Println(c[0:]);
		fmt.Println(c[:]);

		//create a copy of string
		var t2=c[:];
		fmt.Println(c[0:2]);
		fmt.Printf(string(t2[1]));








	}