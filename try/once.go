package try

import "sync"

// Once is a type which added Try method to sync.Once.
type Once struct {
	sync.Once
}

// Try can call sync.Once's Do method with error.
// If f returns non nil error it can retry.
func (o *Once) Try(f func() error) (err error) {
	o.Do(func() {
		err = f()
	})
	if err != nil {
		o.Once = sync.Once{}
	}
	return err
}
