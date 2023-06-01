package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("ds")

	// //string
	// var num string="Robert"
	// num1:="Masala"
	// fmt.Println(num,num1)
	// //Numbers
	// var sum int8=6
	// fmt.Println(sum)
	// //float
	// var num3 float32=34.5
	// fmt.Println(num3)

	//string formatting
	// a:=23
	// b:="India is"
	// fmt.Printf("India is my country %v and %v",a,b)
	// fmt.Printf("India is my country %T and %T",a,b)
	// var c=fmt.Sprintf("India is my country %T and %T",a,b)
	// fmt.Println(c)

	// //Arrays
	// var num = [4]int{2, 3, 4}
	// fmt.Println(num)
	// //slice
	// var num2 = []int{2, 3, 4}
	// num2=append(num2, 23)
	// fmt.Println(num2)
	// num3:=num[1:2]
	// fmt.Println(num3)

	// //Strings
	// var str string="India is my country"
	// fmt.Println(strings.Contains(str,"is"))
	// fmt.Println((strings.Index(str,"is")))
	// fmt.Println(strings.Split(str," "))

	var ar = []int{34, 56, 8, 2, 47}
	sort.Ints(ar)
	fmt.Println(ar)
	index:=sort.SearchInts(ar,8)	
	fmt.Println(index)
}