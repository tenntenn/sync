package try

import (
	"sync"
	"sync/atomic"
)

// Once is a type which added Try method to sync.Once.
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do is same behavior with sync.Once.Do.
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// Try can call sync.Once's Do method with error.
// If f returns non nil error it can retry.
func (o *Once) Try(f func() error) (err error) {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer func() {
			if err == nil {
				atomic.StoreUint32(&o.done, 1)
			}
		}()
		err = f()
	}
	return
}
