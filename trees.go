package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values from
// the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
}
