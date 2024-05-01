package main

import (
	"fmt"
	"sync"
)

func worker(i int,wg *sync.WaitGroup){
	defer wg.Done()//signal that this goroutine is done
	fmt.Printf("worker %d started\n",i)
	//some task is happening
	fmt.Printf("worker %d end\n",i)
	
}

//worker(1)
//worker(2)
//worker(3)

func main() {
	// fmt.Println("Exploring gorouting started")
	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i:=1;i<=3;i++{
		wg.Add(1)//increment the waitgroup counter
		go worker(i,&wg);
	}
	wg.Wait()
	fmt.Println("Workers task complete")
}