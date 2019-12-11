package gopy

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func test2() {
	l := NewLock()
	defer l.Unlock()
}

func test(wg *sync.WaitGroup) {
	t := time.Now()
	for time.Since(t) < time.Second*2 {
		func() {
			l := NewLock()
			defer l.Unlock()
			test2()
			time.Sleep(time.Duration(float64(time.Millisecond) * (1 + rand.Float64())))
		}()
	}
	wg.Done()
}

func TestLock(t *testing.T) {
	l := InitAndLock()
	l.Unlock()
	defer func() {
		l.Lock()
		Finalize()
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(4)
	go test(&wg)
	go test(&wg)
	go test(&wg)
	test(&wg)
	wg.Wait()
}
