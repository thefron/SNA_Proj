package main

import (
	"container/list"
)

type Saved struct {
	saved  map[string]bool
	queued *list.List
}

func NewSaved() *Saved {
	return &Saved{
		make(map[string]bool),
		list.New(),
	}
}

func (s *Saved) Next() *string {
	elem := s.queued.Front()
	if elem == nil {
		return nil
	}
	result := s.queued.Remove(elem).(string)
	return &result
}

func (s *Saved) Enqueue(candidate string) {
	if !s.saved[candidate] {
		s.saved[candidate] = true
		s.queued.PushBack(candidate)
	}
}
