package window

import "sync/atomic"

type Bucket struct {
	windowStart int64
	adder       []int32
	maxUpdater  []int32
}

func NewBucket(windowStart int64) *Bucket {
	return &Bucket{
		windowStart: windowStart,
		adder:       make([]int32, EventCount),
		maxUpdater:  make([]int32, EventCount),
	}
}

func (b *Bucket) GetAdder(event Event) int32 {
	return b.adder[event]
}

func (b *Bucket) GetMaxUpdater(event Event) int32 {
	return b.maxUpdater[event]
}

func (b *Bucket) Increment(event Event) {
	atomic.AddInt32(&b.adder[event], 1)
}

func (b *Bucket) Add(event Event, value int32) {
	atomic.AddInt32(&b.adder[event], value)
}

func (b *Bucket) UpdateMaxUpdater(event Event, value int32) {
	if b.maxUpdater[event] < value {
		atomic.AddInt32(&b.maxUpdater[event], value)
	}
}
