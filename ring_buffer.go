package stats

// RingBuffer that
type RingBuffer struct {
	values []int64
	size   int
	head   int
	tail   int
}

// NewRingBuffer creates a new ring buffer of int64's
func NewRingBuffer(samples int) *RingBuffer {
	return &RingBuffer{
		values: make([]int64, samples),
	}
}

// Queue queues a value on the buffer
func (b *RingBuffer) Queue(v int64) {
	next := b.head + 1

	if next >= len(b.values) {
		next = 0
	}

	if b.size < len(b.values) {
		b.size++
	}

	b.values[b.head] = v
	b.head = next
}

// Dequeue dequeues a value off the buffer
func (b *RingBuffer) Dequeue() int64 {
	v := b.values[b.tail]

	next := b.tail + 1

	if next >= len(b.values) {
		next = 0
	}

	if b.size > 0 {
		b.size--
	}

	b.tail = next

	return v
}

// Iterate iterates through the ring buffer
func (b *RingBuffer) Iterate(fn func(v int64)) {
	for b.size > 0 {
		fn(b.Dequeue())
	}
}
