package task8

import "sync/atomic"

type Sema struct {
	count int64
}

func NewSemaphore() *Sema {
	return &Sema{count: 0}
}

func (s *Sema) Add(n int) {
	for i := 0; i < n; i++ {
		atomic.AddInt64(&s.count, 1)
	}
}

func (s *Sema) Wait() {
	for atomic.LoadInt64(&s.count) > 0 {
	}
}

func (s *Sema) Done() {
	atomic.AddInt64(&s.count, -1)
}
