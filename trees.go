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

// Same determines whether the trees t1 and t2
// contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}
