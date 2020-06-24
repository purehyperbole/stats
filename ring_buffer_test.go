package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestRingBufferQueue(t *testing.T) {
    b := NewRingBuffer(10)

    for i := 0; i < 10; i++ {
        b.Queue(int64(i))
		assert.Equal(t, int64(i), b.values[i])
    }

	assert.Equal(t, 10, b.size)
	assert.Equal(t, 0, b.head)
	assert.Equal(t, 0, b.tail)

	b.Queue(10)

	assert.Equal(t, int64(10), b.values[0])
	assert.Equal(t, 1, b.head)
	assert.Equal(t, 0, b.tail)
}

func TestRingBufferDequeue(t *testing.T) {
	b := NewRingBuffer(10)

	for i := 0; i < 10; i++ {
		b.Queue(int64(i))
	}

	for i := 0; i < 20; i++ {
		assert.GreaterOrEqual(t, b.Dequeue(), int64(0))
	}
}

func TestRingBufferIterate(t *testing.T) {
	b := NewRingBuffer(10)

	for i := 0; i < 100; i++ {
		b.Queue(int64(i))
	}

	i := int64(90)

	b.Iterate(func(v int64) {
		assert.Equal(t, i, v)
		i++
	})
}
