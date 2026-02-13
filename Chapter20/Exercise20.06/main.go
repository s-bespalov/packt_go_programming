package main

import "fmt"

func main() {
	finished := make(chan bool)
	names := []string{"Packt"}
	go func() {
		names = append(names, "Electric")
		names = append(names, "Boogalo")
		finished <- true
	}()
	<-finished
	for _, name := range names {
		fmt.Println(name)
	}
}
