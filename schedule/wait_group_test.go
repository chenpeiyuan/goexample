package schedule

import (
	"runtime"
	"sync"
	"testing"
)

func TestWaitGroupBadCase(t *testing.T) {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1) // this goroutine may be not scheduled before wg.Wait()
		defer wg.Done()
		println("2. then goroutine execute")
	}()
	wg.Wait()
	println("1. wait return first")
}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)  // we guarantee wg.Add() executed first
	go func() {
		defer wg.Done()
		println("1. goroutine execute first")
	}()
	wg.Wait() // this line will wait for wg.Done()
	println("2. then wait return")
}

func TestWaitGroupExperiment(t *testing.T) {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		println("1. goroutine execute first")
	}()
	runtime.Gosched() // yields the processor, our goroutine will get a chance to run
	wg.Wait()
	println("2. then wait return")
}
