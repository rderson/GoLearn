// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo

import (
	"errors"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(string, chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// !+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e != nil {
		memo.mu.Unlock()
		select {
		case <-e.ready:
			return e.res.value, e.res.err
		case <-done:
			return nil, errors.New("request canceled")
		}
	}

	e = &entry{ready: make(chan struct{})}
	memo.mu.Unlock()


	value, err := memo.f(key, done)

	select {
	case <-done:
		close(e.ready)
		return nil, errors.New("request canceled")
	default:
		e.res = result{value, err}
		close(e.ready)

		memo.mu.Lock()
		memo.cache[key] = e
		memo.mu.Unlock()

		return value, err
	}
}

//!-
