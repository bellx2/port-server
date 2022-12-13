package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func hello(w http.ResponseWriter, r *http.Request) {
	rAddr := r.Host
	method := r.Method
	path := r.URL.Path
	log.Printf("%s [%s] %s\n", rAddr, method, path)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	var(
		st = flag.Int("p", 58000, "start port")
		num = flag.Int("n", 1, "port nums")
	)	
	flag.Parse()
	
	if *num < 1 {
		fmt.Println("num must > 0")
		return
	}

	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", hello)

	var wg sync.WaitGroup

	for i := *st; i < *st + *num; i++ {
		p := i
		wg.Add(1)
		go func() {
			port := fmt.Sprintf(":%d", p)
			http.ListenAndServe(port, serverMuxA)
		}()
	}
	fmt.Printf("Open Port From %d to %d\n", *st, *st + *num -1)
	wg.Wait()
}