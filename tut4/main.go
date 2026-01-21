package main
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(id int, ch chan int){
	defer wg.Done()
	for v := range ch{
		fmt.Println("id ",id , "int ", v)
	}
	
}

func main(){
	
	ch := make(chan int)

	for i := 0 ; i < 5 ; i++{
		wg.Add(1)
		go worker(i,ch)
	}

	for ii := 0 ; ii <= 10 ; ii ++ {
		ch <- ii
	}
	close(ch)
	wg.Wait()
	}



