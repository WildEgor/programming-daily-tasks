package main

import (
	"fmt"
	"strings"
	"sync"
)

type DuplicateChecker interface {
	words(size int) []string
	check(string)
	print()
}

type SimpleChecker struct {
	sync.Mutex
	input   string
	Counter map[string]int
}

func (d *SimpleChecker) words(size int) []string {
	words := strings.Fields(d.input)
	var segments []string
	for i := 0; i < len(words); i += size {
		end := i + size
		if end > len(words) {
			end = len(words)
		}
		segments = append(segments, strings.Join(words[i:end], " "))
	}
	return segments
}

func (d *SimpleChecker) check(s string) {
	d.Lock()
	defer d.Unlock()
	d.Counter[s]++
}

func (d *SimpleChecker) print() {
	for k, v := range d.Counter {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func main() {
	checker := &SimpleChecker{
		input:   "apple pineapple banana apple orange banana apple apple",
		Counter: make(map[string]int),
	}

	words := checker.words(3)
	wg := new(sync.WaitGroup)
	wg.Add(len(words))

	for _, segment := range words {
		go func(s string) {
			defer wg.Done()

			sw := strings.Fields(s)
			for _, word := range sw {
				checker.check(word)
			}

		}(segment)
	}

	wg.Wait()

	checker.print()
}
