package main

import (
"fmt"
"flag"
"sync"
"time"
"bufio"
"os"
"net/http"
)

func main() {

	var c int

	flag.IntVar(&c, "C", 30, "Set the Concurrency")
	flag.Parse()

	var wg sync.WaitGroup
		for i:=0; i<c; i++ {
			wg.Add(1)
			go func () {
				status()
				wg.Done()
			}()
			wg.Wait()
		}
}

func status() {
	time.Sleep(500 * time.Millisecond)
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		link:=scanner.Text()
		
	req,err := http.Get(link)
       		if err != nil {
			return
        	}
         defer req.Body.Close()
         fmt.Println(link,  "[", req.StatusCode , "]")
}

}
