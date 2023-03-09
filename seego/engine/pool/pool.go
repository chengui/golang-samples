package pool

import (
	"fmt"
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
)

type EnginePool struct {
	size    int
	lock    sync.RWMutex
	killed  *simplelru.LRU
	working *simplelru.LRU
	factory func() *Engine
}

type EngineWrapper struct {
	*Engine
	Lock *sync.Mutex
}

func NewEnginePool(size int, factory func() *Engine) *EnginePool {
	p := &EnginePool{
		size:    size,
		factory: factory,
	}
	p.killed, _ = simplelru.NewLRU(size, nil)
	p.working, _ = simplelru.NewLRU(size, nil)
	return p
}

func (p *EnginePool) Init() error {
	for i := 0; i < p.size; i++ {
		k := fmt.Sprintf("__%d", i)
		p.killed.Add(k, &EngineWrapper{p.factory(), new(sync.Mutex)})
	}
	return nil
}

func (p *EnginePool) Add(sid string) *EngineWrapper {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.working.Len() < p.size {
		if _, v, ok := p.killed.RemoveOldest(); ok {
			p.working.Add(sid, v)
			return v.(*EngineWrapper)
		} else {
			return nil
		}
	} else {
		if _, v, ok := p.working.RemoveOldest(); ok {
			p.working.Add(sid, v)
			return v.(*EngineWrapper)
		} else {
			return nil
		}
	}
}

func (p *EnginePool) Get(sid string) (*EngineWrapper, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if val, ok := p.working.Get(sid); ok {
		return val.(*EngineWrapper), true
	} else {
		return nil, false
	}
}

func (p *EnginePool) Del(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if val, ok := p.working.Get(sid); ok {
		p.working.Remove(sid)
		p.killed.Add(sid, val)
	}
	return nil
}
