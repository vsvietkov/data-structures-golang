package stack

import (
	"github.com/vsvietkov/data-structures-golang/src/stack"
	"testing"
)

func BenchmarkPush(b *testing.B) {
	s, _ := stack.NewStack(b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := s.Push(i); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPop(b *testing.B) {
	s, _ := stack.NewStack(b.N)

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.Pop()

		if err != nil {
			b.Error(err)
		}
	}
}
