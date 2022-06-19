package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Benchmark1(b *testing.B) {
	var c counter
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&c.n,1)
	}
}

func Benchmark2(b *testing.B) {
	var mu sync.Mutex
	var c counter
	for i := 0; i < b.N; i++ {
		mu.Lock()
		c.n++
		mu.Unlock()
	}
}
