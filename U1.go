package main
import "fmt"
func main(){
	score := 100
	fmt.Println("score:",score)
	ptr:=&score
	*ptr=150
	fmt.Println("score:",score)
}
