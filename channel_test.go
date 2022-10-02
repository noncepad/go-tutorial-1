package main_test

import "testing"

func TestChannel(t *testing.T) {
	a := 1
	b := 2
	c := addMyNumbers(a, b)
	if c != 3 {
		t.Fatal("c is not 3")
	}

	ansC := make(chan int, 1)
	go addInSeparateGoRoutine(a, b, ansC)
	ans := <-ansC
	if ans != c {
		t.Logf("ans=%d", ans)
		t.Fatal("answers do not match")
	}
}

func addMyNumbers(a int, b int) int {
	return a + b
}

func addInSeparateGoRoutine(a int, b int, ansC chan<- int) {
	ansC <- addMyNumbers(a, b)
}
