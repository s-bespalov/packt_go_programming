package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		processStdIn()
	} else {
		processFileOrInput()
	}
}

func rot13(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = ((r - 'A' + 13) % 26) + 'A'
		} else if r >= 'a' && r <= 'z' {
			runes[i] = ((r - 'a' + 13) % 26) + 'a'
		}
	}
	return string(runes)
}

func processStdIn() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading stdin:", err)
			return
		}
		encoded := rot13(input)
		fmt.Print(encoded)
	}
}

func processFileOrInput() {
	var inputReader io.Reader
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("error opening file:", err)
			return
		}
		defer file.Close()
		inputReader = file
	} else {
		fmt.Print("Enter text:")
		inputReader = os.Stdin
	}
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		encoded := rot13(scanner.Text())
		fmt.Println(encoded)
	}
}
