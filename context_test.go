package main_test

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	a := 1
	b := 2
	c := addMyNumbers(a, b)
	if c != 3 {
		t.Fatal("c is not 3")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ansC := make(chan int, 1)
	go addInSeparateGoRoutineWithSleep(a, b, ansC)
	var ans int
	var err error
	select {
	case <-ctx.Done():
		err = errors.New("timed out")
	case ans = <-ansC:
	}
	if err != nil {
		t.Fatal(err)
	}

	if ans != c {
		t.Logf("ans=%d", ans)
		t.Fatal("answers do not match")
	}
}

func addInSeparateGoRoutineWithSleep(a int, b int, ansC chan<- int) {
	time.Sleep(3 * time.Second)
	ansC <- addMyNumbers(a, b)
}
