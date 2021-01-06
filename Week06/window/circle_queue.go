package window

import "fmt"

type CircleQueue struct {
	data    []interface{} // 数据空间
	front   int           // 头
	rear    int           // 尾
	maxSize int           // 容量
}

func NewCircleQueue(maxSize int) *CircleQueue {
	return &CircleQueue{
		data:    make([]interface{}, maxSize),
		front:   0,
		rear:    0,
		maxSize: maxSize,
	}
}

// 入队
func (q *CircleQueue) AddLast(data interface{}) {
	if q.isFull() {
		q.front = q.rear + 1
	}
	q.data[q.rear] = data
	q.rear = (q.rear + 1) % q.maxSize
}

// 出队
func (q *CircleQueue) Pop() interface{} {
	if q.isEmpty() {
		fmt.Println("队列为空...")
		return nil
	}
	data := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.maxSize
	return data
}

// 获取队列内元素的数量
func (q *CircleQueue) GetQueueNum() int {
	return (q.rear - q.front + q.maxSize) % q.maxSize
}

// 获取头元素
func (q *CircleQueue) GetFront() interface{} {
	if !q.isEmpty() {
		return q.data[q.front]
	} else {
		return nil
	}
}

// 获取尾元素
func (q *CircleQueue) GetRear() interface{} {
	if !q.isEmpty() {
		return q.data[(q.rear-1)/q.maxSize]
	} else {
		return nil
	}
}

// 清空队列
func (q *CircleQueue) Clear() {
	if q.isEmpty() {
		return
	}
	for i := 0; i < q.GetQueueNum(); i++ {
		q.Pop()
	}
}

// 判断队列是否已满
func (q *CircleQueue) isFull() bool {
	return q.front == (q.rear+1)%q.maxSize
}

// 判断队列是否为空
func (q *CircleQueue) isEmpty() bool {
	return q.front == q.rear
}
