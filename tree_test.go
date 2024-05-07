package main

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tests := []struct {
		name  string
		val   int
		left  baseNode
		right baseNode
	}{
		{
			"root node with no child nodes",
			5,
			nil,
			nil,
		},
		{
			"root node with left and right node",
			5,
			NewTree(3, nil, nil),
			NewTree(7, nil, nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := NewTree(test.val, test.left, test.right)

			if root.val != test.val {
				t.Fatalf("unexpected val, exp=%d, got=%d", test.val, root.val)
			}

			if !reflect.DeepEqual(root.left, test.left) {
				t.Fatalf("btree has invalid left node, exp=%v, got=%v", test.left, root.left)
			}

			if !reflect.DeepEqual(root.right, test.right) {
				t.Fatalf("btree has invalid right node, exp=%v, got=%v", test.left, root.right)
			}
		})
	}
}
