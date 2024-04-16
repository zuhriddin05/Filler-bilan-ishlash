package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zuhriddin05/Filler-bilan-ishlash/str"
)

func main() {

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		u1 := str.User{}
		st := strings.SplitAfterN(scanner.Text(), " ", 2)
		if st[0] == "" {
			continue
		}
		if st[0] == "Name: " {
			u1.Name = st[1]
		}
		if st[0] == "Age: " {
			u1.Age = st[1]
		}
		if st[0] == "Occupation: " {
			u1.Occupation = st[1]
		}
		fmt.Printf(u1.Name)
		fmt.Printf(u1.Age)
		fmt.Printf(u1.Occupation)
		fmt.Println("")
	}
}
