// Queue ma'lumot tuzilmasi

package main

import "fmt"

func main() {
	testQueue()
}

func testQueue() {
	queue := NewQueue()
	queue.Enqueue(100)
	fmt.Println("queue: ", queue)
	queue.Enqueue(25).Enqueue(77)
	fmt.Println("queue: ", queue)
	fmt.Println("is queue empty? ", queue.IsEmpty())

	var result, _ = queue.Peek()
	fmt.Println("the next item in the queue: ", result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
}

// Queue turini e'lon qilamiz
type Queue struct {
	data []int
}

// NewQueue : yangi Queue obyektini qaytarib beradi
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty : queue'ning bo'shligini tekshirad
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : queue'dagi keyingi elementni qaytarib beradi
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Enqueue : queue'ga yangi element qo'shadi va queue'ga ko'rsatkich qaytarib beradi
func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : queue'dan keyingi elementni o'chirib tashlab, uning qiymatini qaytarib beradi
func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:] // birinchi elementdan boshlab barcha elementlarni qaytaradi
	return element, nil
}
