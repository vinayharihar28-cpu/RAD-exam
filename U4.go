// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:9090")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server running on port: 9090")

// 	sem := make(chan struct{}, 3)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}

// 		sem <- struct{}{}
// 		go handleClient(conn, sem)
// 	}
// }

// func handleClient(conn net.Conn, sem chan struct{}) {
// 	defer conn.Close()
// 	defer func() { <-sem }()

// 	fmt.Println("Client Connected:", conn.RemoteAddr())

// 	for i := 0; i < 10; i++ {
// 	_, err := conn.Write([]byte(time.Now().Format("15:04:05\n")))
// 	if err != nil {
// 		fmt.Println("Client disconnected:", conn.RemoteAddr())
// 		return
// 	}
// 	time.Sleep(time.Second)
// }
// }




// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// type Result struct {
// 	id  int
// 	msg string
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	resultChan := make(chan Result)
// 	var wg sync.WaitGroup

// 	// Ordering goroutine
// 	go func() {
// 		expected := 0
// 		buffer := make(map[int]string)

// 		for res := range resultChan {
// 			buffer[res.id] = res.msg

// 			for {
// 				if msg, ok := buffer[expected]; ok {
// 					fmt.Println("Output:", msg)
// 					delete(buffer, expected)
// 					expected++
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}()

// 	id := 0
// 	fmt.Println("Enter message:")

// 	for scanner.Scan() {
// 		text := scanner.Text()

// 		wg.Add(1)
// 		go func(id int, text string) {
// 			defer wg.Done()
// 			time.Sleep(time.Millisecond * 500)
// 			resultChan <- Result{id, text}
// 		}(id, text)

// 		id++
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	// Now safe to close channel
// 	close(resultChan)

// 	// Give time for final prints (optional but helpful)
// 	time.Sleep(time.Second)
// }



// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start")
// 	panic("Something went wrong!")
// 	fmt.Println("End") // ❌ never executes
// }

// package main
// import (
// 	"fmt"
// 	"math"
// 	"sync"
// )

// func producer(nums []int,out chan<-int){
// 	defer close(out)
// 	for _,n:=range nums{
// 		out <-n
// 	}
// }


// func transformer(in <-chan int,out chan<-float64,errChan chan<-error,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for num:=range in{
// 		if num<0{
// 			errChan<-fmt.Errorf("Invalid input (negative number): %d",num)
// 			continue
// 		}
// 		result:=math.Sqrt(float64(num))
// 		out<-result
// 	}
// }

// func consumer(out<-chan float64,done chan<-bool){
// 	for res:=range out{
// 		fmt.Println("Result: ",res)
// 	}
// 	done<-true
// }

// func main(){
// 	nums:=[]int{1,4,9,-2,16,-5,25}
// 	inputChan:=make(chan int)
// 	outputChan:=make(chan float64)
// 	errchan:=make(chan error)
// 	done:=make(chan bool)

// 	var wg sync.WaitGroup

// 	go producer(nums,inputChan)
// 	wg.Add(1)
// 	go transformer(inputChan,outputChan,errchan,&wg)

// 	go func(){
// 		wg.Wait()
// 		close(outputChan)
// 		close(errchan)
// 	}()

// 	go func(){
// 		for err:=range errchan{
// 			fmt.Println("Error: ",err)
// 		}
// 	}()

// 	go consumer(outputChan,done)
// 	<-done
// 	fmt.Println("Pipeline finished Gracefully!!")
// }

// package main
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int,tasks <-chan int, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	for task:=range tasks{
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		time.Sleep(10*time.Second)
// 		fmt.Printf("Worker %d finished task %d\n",id,task)
// 	}
// }


// func main(){
// 	numWorkers:=5
// 	numTasks:=15

// 	tasks:=make (chan int,5)

// 	var wg sync.WaitGroup

// 	for i:=1;i<=numWorkers;i++{
// 		wg.Add(1)
// 		go worker(i,tasks,&wg)
// 	}

// 	for i:=1;i<=numTasks;i++{
// 		fmt.Println("Sending tasks: ",i)
// 		tasks<-i
// 	}

// 	close(tasks)

// 	wg.Wait()
// 	fmt.Println("All tasks comleted")
// }


// package main
// import (
// 	"fmt"
// 	"time"
// )

// func worker(ch chan int,id int){
// 	time.Sleep(time.Second)
// 	ch<-id
// }

// func main(){
// 	ch:=make (chan int)
// 	for i:=0;i<=5;i++{
// 		go worker(ch,i)
// 	}

// 	fmt.Println("Recived: ",<-ch)
// 	time.Sleep(2*time.Second)
// }

// package main
// import "fmt"

// func main() {
// 	ch := make(chan int,6)

// 	ch <- 1
// 	ch <- 2

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
	
// }

// package main
// import "fmt"

// func main(){
// 	message:=make(chan string)
// 	go func(){
// 		message<-"Hello from gorotuine"
// 	}()
// 	msg:=<-message
// }


// package main
// import "fmt"

// func main(){
// 	numbers:=make(chan int)
// 	go func(){
// 		numbers<-15
// 	}()
// 	nu:=<-numbers
// 	fmt.Println(nu)
// }


package main
import (
	"fmt"
	"sync"

)

func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
			fmt.Println(counter)

		}()

	}
	wg.Wait()
}