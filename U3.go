// package main
// import (
// 	"fmt"
// 	"math"
// )

// type point struct{
// 	x float64
// 	y float64
// }

// func (p point)distance(other point)float64{
// 	dx:=p.x-other.x
// 	dy:=p.y-other.y
// 	return math.Sqrt(dx*dx+dy*dy)
// }

// func (p *point)translate(dx,dy float64){
// 	p.x += dx
// 	p.y += dy
// }

// func (p *point)scale(factor float64){
// 	p.x *= factor
// 	p.y *= factor
// }

// func main(){
// 	var n int
// 	fmt.Print("Enter number of points: ")
// 	fmt.Scan(&n)

// 	points:=make([]point,n)

// 	for i:=0;i<n;i++{
// 		fmt.Printf("Enter X and Y for the points %d: ",i+1)
// 		fmt.Scan(&points[i].x,&points[i].y)
// 	}

// 	if n>2{
// 		d:=points[0].distance(points[1])
// 		fmt.Println("Distance between firt 2 priorities:- ",d)
// 	}

// 	if n>=1{
// 		points[0].translate(2,3)
// 	}

// 	if n>=2{
// 		points[1].scale(2)
// 	}

// 	fmt.Println("Updated Points: ")
// 	for i,p:=range points{
// 		fmt.Printf("Point %d: (%.2f,%.2f)\n",i+1,p.x,p.y)
// 	}
// }

// package main
// import "fmt"

// type meter float64
// type username string

// func (m meter)ToKilometer() float64{
// 	return float64(m)/1000
// }

// func (m meter)ToCentimeter()float64{
// 	return float64(m)*100
// }

// func (u username)IsAdmin()bool{
// 	return string(u)=="admin"
// }

// func (u username)mask()string{
// 	name:=string(u)
// 	if len(name)<=2{
// 		return name
// 	}
// 	masked:=string(name[0])
// 	for i:=1;i<len(name)-1;i++{
// 		masked += "*"
// 	}
// 	masked += string(name[len(name)-1])

// 	return masked
// }

// func main(){
// 	var choice int
// 	fmt.Println("1, Meter Conversion")
// 	fmt.Println("2, Username Check")
// 	fmt.Println("Enter Choice: ")
// 	fmt.Scan(&choice)

// 	switch choice{
// 	case 1:
// 		var m meter
// 		fmt.Println("Enter Values in meters: ")
// 		fmt.Scan(&m)

// 		fmt.Println("Kilometers: ",m.ToKilometer())
// 		fmt.Println("Centemeters: ",m.ToCentimeter())

// 	case 2:
// 		var u username
// 		fmt.Print("Enter username: ")
// 		fmt.Scan(&u)

// 		fmt.Println("Is admin: ",u.IsAdmin())
// 		fmt.Println("Masked Username: ",u.mask())

// 	default:
// 		fmt.Println("Invalid Choice")
// 	}
// }

// package main
// import "fmt"

// type numberList []int
// type scoreBoard map[string]int

// func (n numberList)sum()int{
// 	total:=0
// 	for _,num:=range n{
// 		total += num
// 	}
// 	return total
// }

// func (n numberList)average()float64{
// 	if len(n)==0{
// 		return 0
// 	}
// 	return float64(n.sum())/float64(len(n))
// }

// func (s scoreBoard)totalScore()int{
// 	total:=0
// 	for _,score:=range s{
// 		total += score
// 	}
// 	return total
// }

// func (s scoreBoard)Topper()(string,int){
// 	var topname string
// 	maxScore:=-1

// 	for name, score:= range s{
// 		if score>maxScore{
// 			maxScore=score
// 			topname=name
// 		}
// 	}
// 	return topname,maxScore
// }

// func main(){
// 	var n int
// 	fmt.Print("Enter Number of Elements: ")
// 	fmt.Scan(&n)

// 	nums:=make(numberList,n)

// 	for i:=0;i<n;i++{
// 		fmt.Printf("Enter number %d: ",i+1)
// 		fmt.Scan(&nums[i])
// 	}

// 	fmt.Println("Sum: ",nums.sum())
// 	fmt.Println("Average: ",nums.average())

// 	var m int
// 	fmt.Print("\n Enter numbr of students: ")
// 	fmt.Scan(&m)

// 	scores:= make(scoreBoard)

// 	for i:=0;i<m;i++{
// 		var name string
// 		var score int

// 		fmt.Print("Enter name: ")
// 		fmt.Scan(&name)

// 		fmt.Print("Enter Score: ")
// 		fmt.Scan(&score)

// 		scores[name]=score
// 	}

// 	fmt.Println("Total Score: ",scores.totalScore())

// 	topName,topScore:=scores.Topper()
// 	fmt.Println("Topper: ",topName,"With Score",topScore)

// }

// package main
// import (
// 	"fmt"
// 	"math"
// )

// type point struct{
// 	x float64
// 	y float64
// }

// type path []point

// //distance for point
// func (p point)distance(other point)float64{
// 	dx:=p.x-other.x
// 	dy:=p.y-other.y
// 	return math.Sqrt(dx*dx + dy*dy)
// }

// func (p path)distance() float64{
// 	total:=0.0
// 	for i:=0;i<len(p)-1;i++{
// 		total += p[i].distance(p[i+1])
// 	}
// 	return total
// }

// func main(){
// 	var n int
// 	fmt.Print("Enter number of points")
// 	fmt.Scan(&n)

// 	path:=make(path, n)

// 	for i:=0;i<n;i++{
// 		fmt.Printf("Enter X and Y for point %d: ",i+1)
// 		fmt.Scan(&path[i].x,&path[i].y)
// 	}

// 	for i:=0;i<len(path)-1;i++{
// 		d:=path[i].distance(path[i+1])
// 		fmt.Printf("Distance btw point %d and %d: %.2f\n",i+1,i+2,d)
// 	}

// 	fmt.Println("Total path Distance: ",path.distance())
// }

// package main
// import "fmt"

// type IntSet struct{
// 	words []uint64
// }

// func (s *IntSet)add(x int){
// 	word:=x/64
// 	bit:=uint(x%64)
// 	for word>=len(s.words){
// 		s.words=append(s.words, 0)
// 	}
// 	s.words[word]|=1<<bit
// }

// func (s *IntSet)has(x int)bool{
// 	word:=x/64
// 	bit:=uint(x%64)
// 	if word>=len(s.words){
// 		return false
// 	}
// 	return s.words[word]&(1<<bit)!=0
// }

// func (s *IntSet)UnionWith(t *IntSet){
// 	for i,tword:=range t.words{
// 		if i<len(s.words){
// 			s.words[i] |= tword
// 		}else{
// 			s.words=append(s.words,tword)
// 		}
// 	}
// }

// func (s *IntSet)string() string{
// 	result:="{"
// 	first:=true

// 	for i, word:= range s.words{
// 		for j:=0;j<64;j++{
// 			if word&(1<<uint(j))!=0{
// 				if !first{
// 					result +=" "
// 				}
// 				result += fmt.Sprintf("%d",64*i+j)
// 				first=false
// 			}
// 		}
// 	}

// 	result += "}"
// 	return result
// }

// func main(){

// 	var s1,s2 IntSet
// 	s1.add(1)
// 	s1.add(2)
// 	s2.add(3)
// 	s2.add(4)

// 	fmt.Println("Set 1: ",s1.string())
// 	fmt.Println("Set 2: ",s2.string())

// 	fmt.Println("Is 3 in set 1? ",s1.has(3))

// 	s1.UnionWith(&s2)
// 	fmt.Println("Union: ",s1.string())
// }

// package main
// import "fmt"

// type notifier interface{
// 	notify(message string)
// }

// type email struct{}
// type sms struct{}
// type pushNotification struct{}

// func (e email)notify(message string){
// 	fmt.Println("Email sent: ",message)
// }

// func (s sms)notify(message string){
// 	fmt.Println("SMS sent: ",message)
// }

// func (p pushNotification)notify(message string){
// 	fmt.Println("Push Notification: ",message)
// }

// func sendNotification(n notifier,message string){
// 	n.notify(message)
// }

// func main(){
// 	var choice int
// 	var msg string

// 	fmt.Println("1. Email")
// 	fmt.Println("2. SMS")
// 	fmt.Println("3. Push Notification")
// 	fmt.Println("Enter Choice: ")
// 	fmt.Scan(&choice)

// 	fmt.Print("Enter Message: ")
// 	fmt.Scan(&msg)

// 	switch choice{
// 	case 1:
// 		sendNotification(email{},msg)
// 	case 2:
// 		sendNotification(sms{},msg)
// 	case 3:
// 		sendNotification(pushNotification{},msg)
// 	default:
// 		fmt.Println("Invalid choice")
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type shape interface{
// 	Area() float64
// }

// type circle struct{
// 	radius float64
// }

// func (c circle)Area() float64{
// 	return  math.Pi*c.radius*c.radius
// }

// type rectangle struct{
// 	width float64
// 	height float64
// }

// func (r rectangle)Area()float64{
// 	return r.width*r.height
// }

// type triangle struct{
// 	base float64
// 	height float64
// }

// func (t triangle)Area()float64{
// 	return 0.5*t.base*t.height
// }

// func printArea(s shape)float64{
// 	area:=s.Area()
// 	fmt.Println("Area: ",area)
// 	return area
// }

// func main(){
// 	var shapes []shape

// 	var n int
// 	fmt.Println("Enter number of shapes: ")
// 	fmt.Scan(&n)

// 	for i:=0;i<n;i++{
// 		var choice int
// 		fmt.Println("\n1.Circle\n2.Rectangle\n3.Triangle")
// 		fmt.Print("Enter type: ")
// 		fmt.Scan(&choice)

// 		switch choice{
// 		case 1:
// 			var r float64
// 			fmt.Print("Enter Radius: ")
// 			fmt.Scan(&r)
// 			shapes=append(shapes,circle{r})

// 		case 2:
// 			var w, h float64
// 			fmt.Print("Enter base and height: ")
// 			fmt.Scan(&w, &h)
// 			shapes=append(shapes,triangle{w,h})
		
// 		case 3:
// 			var b,h float64
// 			fmt.Print("Enter base and height: ")
// 			fmt.Scan(&b,&h)
// 			shapes=append(shapes, triangle{b,h})
// 		}
// 	}

// 	total:=0
// 	for _,s:=range shapes{
// 		total+=int(printArea(s))
// 	}

// }





// package main
// import (
// 	"fmt"
// )

// type paymentMethod interface{
// 	pay(amount float64)
// }

// //implementingmethods

// type creditCard struct{}
// func (c creditCard)pay(amount float64){
// 	fmt.Println("Paid",amount,"using credit card")
// }

// type upi struct{}
// func (u upi)pay(amount float64){
// 	fmt.Println("Paid",amount,"using UPI")
// }

// type netBanking struct{}
// func (n netBanking)pay(amount float64){
// 	fmt.Println("Paid",amount,"using NetBanking")
// }

// func processPayment(p paymentMethod,amount float64){
// 	p.pay(amount)
// }

// func main(){
// 	var choice int
// 	var amount float64

// 	fmt.Println("1. Credit card")
// 	fmt.Println("2. UPI")
// 	fmt.Println("3. Net Banking")
// 	fmt.Println("Enter Choice: ")
// 	fmt.Scan(&choice)

// 	fmt.Print("Enter Amount: ")
// 	fmt.Scan(&amount)

// 	switch choice{
// 	case 1:
// 		processPayment(creditCard{},amount)
// 	case 2:
// 		processPayment(upi{},amount)
// 	case 3:
// 		processPayment(netBanking{},amount)
// 	default:
// 		fmt.Println("Invalid Choice")
// 	}
// }

// package main
// import (
// 	"fmt"
// )

// type logger interface{
// 	log(message string)
// }

// type consoleLogger struct{}
// func (c consoleLogger)log(message string){
// 	fmt.Println("Console: ",message)
// }

// type fileLogger struct{}
// func(f fileLogger)log(message string){
// 	fmt.Println("Writting to file: ",message)
// }

// type mocklogger struct{
// 	logs []string
// }

// func (m *mocklogger)log(message string){
// 	m.logs=append(m.logs, message)
// }

// func processData(l logger){
// 	l.log("processing started")
// 	l.log("processing completed")
// }

// func main(){
// 	mock :=&mocklogger{}
// 	processData(mock)
// 	fmt.Println("Stored logs: ",mock.logs)

// 	console:=consoleLogger{}
// 	processData(console)
// }


package main
import (
	"fmt"
	"math"
	"strings"
)

func main(){
	var n int
	fmt.Print("Enter number of inputs: ")
	fmt.Scan(&n)

	data:=make([]interface{},n)

	for i:=0;i<n;i++{
		var input string
		fmt.Print("Enter Value: ")
		fmt.Scan(&input)

		var intVal int
		_, err1:=fmt.Sscanf(input, "%d",&intVal)

		var floatVal float64
		_, err2:= fmt.Sscanf(input,"%f",&floatVal)

		if err1==nil{
			data[i]=intVal
		}else if err2==nil{
			data[i]=floatVal
		}else{
			data[i]=input
		}
	}

	fmt.Println("\n Processed Output: ")

	for _,value:=range data{
		switch v:=value.(type){
		case int:
			fmt.Println("Integer: ",v,"-> Double: ",v*2)
		case float64:
			fmt.Println("Float: ",v,"-> Rounded: ",math.Round(v))
		case string:
			fmt.Println("String: ",v,"-> Upercase: ",strings.ToUpper(v))
		}
	}
}


