package stack_array

import (
	"errors"
	"golang-first-step/about_leetcode/model"
)

type Pet struct {
	petType string
	// 时间戳
	count int
}

func (p *Pet) SetCount(count int) {
	p.count = count
}

func (p *Pet) GetCount() int {
	return p.count
}

func (p *Pet) GetType() string {
	return p.petType
}

// CatDogQueue 猫狗队列
type CatDogQueue struct {
	dogQueue *model.MyQueue[Pet]
	catQueue *model.MyQueue[Pet]
	// 队列时间戳
	count int
}

// PollAll 按照顺序出队列
func (cd CatDogQueue) PollAll() (Pet, error) {
	if !cd.dogQueue.IsEmpty() && !cd.catQueue.IsEmpty() {
		dog, _ := cd.dogQueue.Peek()
		cat, _ := cd.catQueue.Peek()

		if dog.GetCount() > cat.GetCount() {
			c, _ := cd.catQueue.Dequeue()
			return c, nil
		} else {
			d, _ := cd.dogQueue.Dequeue()
			return d, nil
		}
	} else if !cd.dogQueue.IsEmpty() {
		d, _ := cd.dogQueue.Dequeue()
		return d, nil
	} else if !cd.catQueue.IsEmpty() {
		c, _ := cd.catQueue.Dequeue()
		return c, nil
	} else {
		return Pet{}, errors.New("empty queue")
	}
}

// Add 猫狗队列入队列
func (cd *CatDogQueue) Add(p Pet) bool {
	if "dog" == p.petType {
		cd.dogQueue.Enqueue(p)
		cd.count++
		p.SetCount(cd.count)
		return true
	} else if "cat" == p.petType {
		cd.catQueue.Enqueue(p)
		cd.count++
		p.SetCount(cd.count)
		return true
	}

	return false
}

func NewCatDogQueue() *CatDogQueue {
	return &CatDogQueue{
		dogQueue: model.NewMyQueue[Pet](),
		catQueue: model.NewMyQueue[Pet](),
		count:    0,
	}
}
