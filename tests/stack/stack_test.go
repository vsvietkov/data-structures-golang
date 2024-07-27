package stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/vsvietkov/data-structures-golang/src/stack"
	"testing"
)

var s *stack.Stack

func TestPush(t *testing.T) {
	s, _ = stack.NewStack(5)
	testData := []int{1, 2, 3, 4, 5}

	for _, v := range testData {
		if err := s.Push(v); err != nil {
			t.Error(err)
		}
	}

	stackData := s.GetElements()
	assert.Equal(t, testData, stackData)
}

func TestPop(t *testing.T) {
	stackData := s.GetElements()

	for i := len(s.GetElements()) - 1; i >= 0; i-- {
		element, err := s.Pop()

		assert.NoError(t, err)
		assert.Equal(t, stackData[i], element)
		assert.Equal(t, len(s.GetElements()), i)
	}
}

func TestPeek(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}

	for _, v := range testData {
		if err := s.Push(v); err != nil {
			t.Error(err)
		}

		lastValue, err := s.Peek()
		assert.NoError(t, err)
		assert.Equal(t, v, lastValue)
	}

	for i := len(testData) - 1; i >= 0; i-- {
		element, err := s.Peek()
		s.Pop()

		assert.NoError(t, err)
		assert.Equal(t, testData[i], element)
	}
}
