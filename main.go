package main

import "fmt"

func main() {
	a := createServer()
	a.routes()
	a.serve()
	fmt.Println("Running on 4000")
}
