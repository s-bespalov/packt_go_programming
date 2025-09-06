package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Sam", "Name of the person to say hallo")
	quietFlag = flag.Bool("quiet", false, "Toggle to be quit when saying hello")
)

func main() {
	flag.Parse()
	if !*quietFlag {
		qreeting := fmt.Sprintf("Hello %s! Welcome to the command line", *nameFlag)
		fmt.Println(qreeting)
	}
}
