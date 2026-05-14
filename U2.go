// package main
// import "fmt"

// type Student struct{
// 	RollNo int
// 	Name string
// 	CGPA float64
// }

// func main(){
// 	var s Student
// 	fmt.Println("Before Assignment : ")
// 	fmt.Println("Rollno: ",s.RollNo)
// 	fmt.Println("Name: ",s.Name)
// 	fmt.Println("CGPA: ",s.CGPA)

// 	s.RollNo=101
// 	s.Name="Vinay"
// 	s.CGPA=8.5

// 	fmt.Println("\nAfter Assignment : ")
// 	fmt.Println("Rollno: ",s.RollNo)
// 	fmt.Println("Name: ",s.Name)
// 	fmt.Println("CGPA: ",s.CGPA)
// }


//P-2
// package main
// import "fmt"

// type Employee struct{
// 	ID int
// 	Name string
// 	salary float64
// }

// func main(){
// 	e1:=Employee{
// 		ID:1,
// 		Name:"Vinay",
// 		salary:5000,
// 	}

// 	e2:=e1

// 	e2.salary=7000

// 	fmt.Println("e1: ",e1)
// 	fmt.Println("e2: ",e2)
// }

// package main
// import "fmt"

// type student struct{
// 	name string
// 	CGPA float64
// }

// func updateCGPA(s *student, value float64){
// 	s.CGPA=value
// }

// func main(){
// 	s:=student{
// 		name: "Vinay",
// 		CGPA: 7.5,
// 	}

// 	updateCGPA(&s,9.2)

// 	fmt.Println("Updated Student: ",s)
// }

// package main
// import "fmt"

// type stu struct{
// 	RollNo int
// 	name string
// 	cgpa float64
// }

// func main(){
// 	s:=stu{
// 		RollNo: 01,
// 		name:"vinay",
// 	}
// 	fmt.Println(s)
// }


// package main
// import "fmt"

// type point struct{
// 	x int
// 	y int
// }

// func main(){
// 	p1:=point{1,2}
// 	p2:=point{1,2}
// 	result:=p1==p2
// 	fmt.Println("are p1 and p2 equal? ",result)
// // }

// package main 
// import "fmt"

// func add(x int,y int)int{
// 	return x+y
// }

// func main(){
// 	result:=add(5,3)
// 	fmt.Println(result)
// }


// package main 
// import "fmt"

// func increment(x int){
// 	x=x+1
// 	fmt.Println("Inside Function: ",x)
// }

// func main(){
// 	num:=10
// 	fmt.Println("Before Function: ",num)
// 	increment(num)
// 	fmt.Println("After Function: ",num)
// }


// package main 
// import "fmt"

// func update(x *int){
// 	*x=*x+1
// }

// func main(){
// 	num:=10
// 	fmt.Println("Before: ",num)
// 	update(&num)
// 	fmt.Println("After: ",num)
// }

// package main 
// import "fmt"

// func divide(a,b int)(int,error){
// 	if b==0{
// 		return 0,fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a/b, nil
// }

// func main(){
// 	result,err:=divide(10,0)

// 	if err!=nil{
// 		fmt.Println("Error: ",err)
// 	}else{
// 		fmt.Println("Result: ",result)
// 	}
// }

// package main
// import "fmt"

// func minMax(a,b int)(min int ,max int){
// 	if a<b{
// 		min=a
// 		max=b
// 	}else {
// 		min=b
// 		max=a
// 	}
// 	return
// }

// func main(){
// 	minval,maxval:=minMax(10,5)
// 	fmt.Println("Min: ",minval)
// 	fmt.Println("Max: ",maxval)
// }

// package main
// import "fmt"

// func factorial(n int)int{
// 	if n==0{
// 		return 1
// 	}
// 	return n*factorial(n-1)
// }

// func main(){
// 	result:=factorial(5)
// 	fmt.Println("factorial: ",result)
// }


//varadic behaves like slice
// package main
// import "fmt"

// func sum(nums ...int)int{
// 	total:=0
// 	for _,num:=range nums{
// 		total+=num
// 	}
// 	return total
// }

// func main(){
// 	// fmt.Println(sum(1,2))
// 	// fmt.Println(sum(1,2,3,4,5))
// 	// fmt.Println(sum())

// 	numbers:=[]int{10,20,30}
// 	fmt.Println(sum(numbers...))
// }

// package main
// import "fmt"

// func main(){
// 	square:=func(x int)int{
// 		return x*x
// 	}

// 	result:=square(5)
// 	fmt.Println("Square: ",result)
// }



// package main
// import "fmt"

// func demo(){
// 	defer fmt.Println("End")
// 	fmt.Println("Start")
	
// }

// func main(){
// 	demo()
// }



