package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcomee , Here we are exploring strings package")
	
	fmt.Println(strings.HasPrefix("test", "e"))
	fmt.Println(strings.HasSuffix("test", "st"))

	//lists of methods in strings package
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.LastIndex("test", "e"))
	fmt.Println(strings.Contains("test", "e"))
	fmt.Println(strings.Replace("teset", "e", "E",1))
	fmt.Println(strings.Replace("tesete", "e", "E",2))

	fmt.Println(strings.Replace("teset", "e", "E", -1))
	fmt.Println(strings.Split("test", "e"))
	fmt.Println(strings.SplitAfter("test", "e"))
	fmt.Println(strings.SplitAfterN("test", "e", 2))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.ToTitle("test"))
	

}
