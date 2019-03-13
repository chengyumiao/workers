package workers

import "sync"

type WorkersPool struct {
	workers chan bool
	WG *sync.WaitGroup
	Mu *sync.Mutex
}

func (this *WorkersPool) GetOneWorker() {
	this.workers <- true
	if this.WG != nil {
		this.Mu.Lock()
		defer this.Mu.Unlock()
		this.WG.Add(1)
	}
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
	}
}

func NewWorkersPoolWithWG(size int) *WorkersPool {
	return &WorkersPool{
		workers: make(chan bool, size),
		WG: &sync.WaitGroup{},
		Mu: &sync.Mutex{},
	}
}
