package main

import "fmt"

type Person struct {
	name string
	days int
}

type Result struct {
	workerId int
	elementsProcessed int
}

func main() {
	persons := map[string]int{"Robert" : 30, "John" : 475, "Elly" : 1022, "Bob" : 99}

	// The minimum amount of workers must 1.
	const workerCount int = 1

	jobs := make(chan Person, len(persons))
	results := make(chan Result, len(persons))

	// start worker pool
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}

	for name, days := range persons {
		jobs <- Person{name, days}
	}

	close(jobs)


	total := make(map[int]int, len(persons))

	// count how may tasks completes in each worker
	for a := 1; a <= len(persons); a++ {
		r := <-results
		total[r.workerId] += r.elementsProcessed
	}

	fmt.Println("\nInfo:")
	fmt.Printf("Workers Count: %d\n", workerCount)
	for workerId, elementsProcessed := range total {
		fmt.Printf("Worker#%d -> %d elements processed\n", workerId, elementsProcessed)
	}
}

func worker(id int, jobs <-chan Person, results chan<- Result) {
	for j := range jobs {
		printNameWeeks(j.name, j.days)
		results <- Result{id, 1}
	}

}

func printNameWeeks(name string, d int) {
	var weekStr, dayStr string
	weeks, days := daysToWeeks(d)
	if weeks > 1 {
		weekStr = fmt.Sprintf("%d weeks", weeks)
	} else {
		weekStr = fmt.Sprintf("%d week", weeks)
	}

	if days > 1 {
		dayStr = fmt.Sprintf(" and %d days", days)
	} else if days == 1 {
		dayStr = fmt.Sprintf(" and %d day", days)
	} else {
		dayStr = ""
	}

	fmt.Printf("%s has worked %s%s in the company\n", name, weekStr, dayStr)
}

func daysToWeeks(days int) (int, int) {
	return days / 7, days % 7
}
