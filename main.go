package main

import "fmt"

func update(x *string) {
	*x = "Athulya"
}
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

	// //Sort
	// var ar = []int{34, 56, 8, 2, fmt
	// sort.Ints(ar)
	// fmt.Println(ar)
	// index:=sort.SearchInts(ar,8)
	// fmt.Println(index)

	//loops
	// for x := 0; x < 5; x++ {
	// 	fmt.Println(x)
	// }
	// name:=[]string{"India","Anertica","Sweden","Istanbul"}
	// for index,values:=range name{
	// 	fmt.Println(index,values)
	// }

	//functions
	// func wishMe(name string) {
	// 	fmt.Printf("Hi Mr %v",name)

	// }
	// func cylceName(n []string, f func(string)) {
	// 	for _, v := range n {
	// 		f(v)
	// 	}
	// }
	// var names=[]string{"Athulya","nIYA","Jesuit"}
	// cylceName(names,wishMe)

	//maps
	// menu := map[string]float64{
	// 	"india":  23.4,
	// 	"usa":    263.3,
	// 	"france": 34.8,
	// }
	// for k,v :=range menu{
	// 	fmt.Println(k,"-",v)
	// }
	// menu["indian"]=23.44
	// fmt.Println(menu)

	//pointers
	name := "athulya"
	m := &name
	fmt.Println(name,m)
	update(m)
	fmt.Println(*m)


}
