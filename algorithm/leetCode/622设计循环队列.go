package leetCode

/*
   MyCircularQueue(k): 构造器，设置队列长度为 k 。
   Front: 从队首获取元素。如果队列为空，返回 -1 。
   Rear: 获取队尾元素。如果队列为空，返回 -1 。
   enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
   deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
   isEmpty(): 检查循环队列是否为空。
   isFull(): 检查循环队列是否已满。
*/

type MyCircularQueue struct {
	data    []int //用于存放队列数据的数组
	front   int   //指向头指针
	rear    int   //指向尾指针
	maxSize int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func ConstructorQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		data:    make([]int, k),
		front:   -1,
		rear:    -1,
		maxSize: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.front = 0
	}
	this.rear = (this.rear + 1) % this.maxSize
	this.data[this.rear] = value

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.rear {
		this.front = -1
		this.rear = -1
		return true
	}
	this.front = (this.front + 1) % this.maxSize
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.rear]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.rear+1)%this.maxSize == this.front {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
