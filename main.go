package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("error reading file: ", err)

		os.Exit(1)
	}

	table := getCharFrequency([]rune(string(data)))
	sortedNodes := getNodes(table)

	getBTree(sortedNodes)
}

func getCharFrequency(data []rune) map[rune]int {
	table := make(map[rune]int)

	for _, c := range data {
		table[c] = table[c] + 1
	}

	return table
}

func getNodes(charFrequency map[rune]int) nodes {
	var l nodes

	for k, v := range charFrequency {
		l = append(l, leaf{char: k, frequency: v})
	}

	return l
}

func getBTree(l nodes) baseNode {
	if len(l) == 0 {
		return nil
	}

	n := make(nodes, len(l))
	copy(n, l)
	sort.Sort(n)

	for len(n) > 1 {
		left := n[0]
		right := n[1]

		weight := left.weight() + right.weight()

		root := NewTree(weight, left, right)
		n = append(n, root)

		n = n[2:]
		sort.Sort(n)
	}

	return n[0]
}
