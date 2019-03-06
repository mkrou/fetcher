package fetcher

import (
	"sync"
	"time"
)

var (
	singleton *mock
	once      sync.Once
)

type Fetcher interface {
	Get() (string, error)
	List() ([]string, error)
}

func NewFetcher(inParallel int) Fetcher {
	once.Do(func() {
		singleton = &mock{
			sem: make(chan struct{}, inParallel),
		}
	})

	return singleton
}

type mock struct {
	sem chan struct{}
}

func (m *mock) Get() (string, error) {
	m.lock()
	defer m.unlock()

	<-time.After(time.Second)

	return "some string", nil
}

func (m *mock) List() ([]string, error) {
	m.lock()
	defer m.unlock()

	<-time.After(time.Second)

	return []string{"some string"}, nil
}

func (m *mock) lock() {
	m.sem <- struct{}{}
}

func (m *mock) unlock() {
	<-m.sem
}
