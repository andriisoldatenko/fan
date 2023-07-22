package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println()
	f, err := os.Open("./test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		for key, value := range record {
			fmt.Printf("%v  %v\n", key, value)
		}
	}
}
