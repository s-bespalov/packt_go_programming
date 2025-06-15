package main

import (
	"log"
	"strings"

	"github.com/s-bespalov/packt_go_programming/Chapter11/Activity11.01/ssn"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate | log.Llongfile | log.Lmsgprefix)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	log.Printf("Checking data %#v\n", validateSSN)
	for i, v := range validateSSN {
		log.Printf("Validate data %#v %d of %d\n", v, i, len(validateSSN))
		v = strings.ReplaceAll(v, "-", "")
		_, err := ssn.NewSSN(v)
		if err != nil {
			log.Printf("the value of %v caused an error: %v\n", v, err)
		}
	}
}
