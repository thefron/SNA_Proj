package main

import (
	"container/list"
	"sync"
)

type Saved struct {
	saved  map[string]bool
	queued *list.List
	lock   *sync.Mutex
}

func NewSaved() *Saved {
	return &Saved{
		make(map[string]bool),
		list.New(),
		&sync.Mutex{},
	}
}

func (s *Saved) Next() *string {
	s.lock.Lock()
	defer s.lock.Unlock()
	elem := s.queued.Front()
	if elem == nil {
		return nil
	}
	result := s.queued.Remove(elem).(string)
	return &result
}

func (s *Saved) Enqueue(candidate string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.saved[candidate] {
		s.saved[candidate] = true
		s.queued.PushBack(candidate)
	}
}
