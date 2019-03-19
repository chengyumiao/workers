package workers

import (
	"sync"
	"errors"
)

type WorkersPool struct {
	Size int
	workers chan bool
	WG *sync.WaitGroup
	Mu *sync.Mutex
}

func (this *WorkersPool) GetOneWorker() error{
	if len(this.workers) >= this.Size {
		return errors.New("workers exceed the maximum quantity!")
	}

	this.workers <- true
	if this.WG != nil {
		this.Mu.Lock()
		defer this.Mu.Unlock()
		this.WG.Add(1)
	}
	return nil
}

func (this *WorkersPool) ReleaseOneWorker() {
	<- this.workers
	if this.WG != nil {
		this.Mu.Lock()
		defer this.Mu.Unlock()
		this.WG.Done()
	}
}

func (this *WorkersPool) Wait() {
	if this.WG != nil {
		this.WG.Wait()
		return
	}
	return
}

func NewWorkersPool(size int) *WorkersPool{
	return &WorkersPool{
		workers: make(chan bool, size),
		Size: size,
	}
}

func NewWorkersPoolWithWG(size int) *WorkersPool {
	return &WorkersPool{
		workers: make(chan bool, size),
		WG: &sync.WaitGroup{},
		Mu: &sync.Mutex{},
		Size:size,
	}
}
