package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	output := os.Args[len(os.Args)-1]
	matches := os.Args[1 : len(os.Args)-1]

	fmt.Println(output)
	fmt.Printf("%v\n", matches)

	out, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln("failed to open output file:", err)
	}
	defer out.Close()

	for _, file := range matches {
		fmt.Println(file)

		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		out.Write(content)
		out.WriteString("\n")
	}
}
