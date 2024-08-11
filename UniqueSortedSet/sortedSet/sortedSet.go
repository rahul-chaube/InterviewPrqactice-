package sortedset

import (
	"container/heap"
	"time"
)

type Data struct {
	TimeStamp time.Time
	Message   string
	index     int
}

type Priorityque []*Data

func (p Priorityque) Len() int { return len(p) }
func (p Priorityque) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = j
	p[j].index = i
}
func (p Priorityque) Less(i, j int) bool {
	return p[i].TimeStamp.After(p[j].TimeStamp)
}

// func (p *Priorityque) Push(i interface{}) {
// 	n := p.Len()
// 	fmt.Println(" 111 ", reflect.TypeOf(i))
// 	item := i.(*Data)
// 	item.index = n
// 	*p = append(*p, item)
// }

func (p *Priorityque) Push(x interface{}) {
	n := len(*p)
	item := x.(Data)
	item.index = n
	*p = append(*p, &item)
}

func (p *Priorityque) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*p = old[:n-1]
	return item
}

// Creating set

type Set struct {
	Item map[string]*Data
	pq   Priorityque
}

func NewQue() *Set {
	return &Set{
		pq:   make(Priorityque, 0),
		Item: make(map[string]*Data),
	}
}

func (s *Set) Add(value string) {
	if _, exists := s.Item[value]; !exists {
		data := Data{
			Message:   value,
			TimeStamp: time.Now(),
		}
		s.Item[value] = &data
		heap.Push(&s.pq, data)
	}
}

func (s *Set) GetOlder() *Data {
	if s.pq.Len() > 0 {
		item := heap.Pop(&s.pq).(*Data)
		delete(s.Item, item.Message)
		return item
	}
	return nil
}
