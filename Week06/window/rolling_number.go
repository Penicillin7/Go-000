package window

import (
	"sort"
	"sync"
	"time"
)

var currentTime = time.Now().Unix()

type RollingNumber struct {
	timeInMilliseconds       int
	numberOfBuckets          int
	bucketSizeInMilliseconds int
	buckets                  *CircleQueue
	mux                      sync.RWMutex
}

func NewRollingNumber(timeInMilliseconds int, numberOfBuckets int) *RollingNumber {
	if timeInMilliseconds%numberOfBuckets != 0 {
		panic("The timeInMilliseconds must divide equally into numberOfBuckets")
	}
	return &RollingNumber{
		timeInMilliseconds:       timeInMilliseconds,
		numberOfBuckets:          numberOfBuckets,
		bucketSizeInMilliseconds: timeInMilliseconds / numberOfBuckets,
		buckets:                  NewCircleQueue(numberOfBuckets),
	}
}

// 获取当前bucket
func (r *RollingNumber) getCurrentBucket() *Bucket {
	currentBucket := r.buckets.GetRear().(*Bucket)
	bucketSizeInMilliseconds := int64(r.bucketSizeInMilliseconds)
	if currentBucket != nil && currentTime < currentBucket.windowStart+bucketSizeInMilliseconds {
		return currentBucket
	}
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.buckets.GetRear() == nil {
		newBucket := NewBucket(currentTime)
		r.buckets.AddLast(newBucket)
		return newBucket
	} else {
		for i := 0; i < r.numberOfBuckets; i++ {
			lastBucket := r.buckets.GetRear().(*Bucket)
			if currentTime < lastBucket.windowStart+bucketSizeInMilliseconds {
				return lastBucket
			} else if currentTime-lastBucket.windowStart+bucketSizeInMilliseconds > int64(r.timeInMilliseconds) {
				r.reset()
				r.getCurrentBucket()
			} else {
				r.buckets.AddLast(NewBucket(lastBucket.windowStart + bucketSizeInMilliseconds))
			}
		}
	}
	return r.buckets.GetRear().(*Bucket)
}

// 获取所有bucket
func (r *RollingNumber) getValues(event Event) []int32 {
	lastBucket := r.getCurrentBucket()
	if lastBucket == nil {
		return []int32{}
	}
	result := make([]int32, r.buckets.GetQueueNum())
	for idx, b := range r.buckets.data {
		bucket := b.(*Bucket)
		result[idx] += bucket.GetAdder(event)
	}
	return result
}

// 根据event type 获取所有Bucket 某index 总和
func (r *RollingNumber) getRollingSum(event Event) int32 {
	if r.getCurrentBucket() == nil {
		return 0
	}
	var sum int32 = 0
	for _, i := range r.buckets.data {
		bucket := i.(*Bucket)
		sum += bucket.GetAdder(event)
	}
	return sum
}

// 根据event type 获取最后一个bucket
func (r *RollingNumber) getValueOfLatestBucket(event Event) int32 {
	if r.getCurrentBucket() == nil {
		return 0
	}
	return r.buckets.GetRear().(*Bucket).GetAdder(event)
}

// 获取getValues结果的最大值
func (r *RollingNumber) getRollingMaxValue(event Event) int32 {
	result := r.getValues(event)
	if len(result) == 0 {
		return 0
	} else {
		r.quickSort(result)
		return result[len(result)-1]
	}
}

// 快速排序
func (r *RollingNumber) quickSort(a []int32) {
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

// 自增
func (r *RollingNumber) Increment(event Event) {
	r.getCurrentBucket().Increment(event)
}

// 增加指定值
func (r *RollingNumber) Add(event Event, value int32) {
	r.getCurrentBucket().Add(event, value)
}

// 保留最大值
func (r *RollingNumber) UpdateRollingMax(event Event, value int32) {
	r.getCurrentBucket().UpdateMaxUpdater(event, value)
}

// 清空bucket桶
func (r *RollingNumber) reset() {
	r.buckets.Clear()
}
