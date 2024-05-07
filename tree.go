package main

import "fmt"

type baseNode interface {
	isLeaf() bool
	weight() int
}

type tree struct {
	val   int
	left  baseNode
	right baseNode
}

func NewTree(val int, left, right baseNode) *tree {
	return &tree{val, left, right}
}

func (t tree) isLeaf() bool {
	return false
}

func (t tree) weight() int {
	return t.val
}

type leaf struct {
	char      rune
	frequency int
}

func (l leaf) String() string {
	return fmt.Sprintf("%c: %d", l.char, l.frequency)
}

func (l leaf) isLeaf() bool {
	return true
}

func (l leaf) weight() int {
	return l.frequency
}

type nodes []baseNode

func (l nodes) Len() int           { return len(l) }
func (l nodes) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l nodes) Less(i, j int) bool { return l[i].weight() < l[j].weight() }
