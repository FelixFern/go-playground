package counter

import "sync"

type Counter struct {
	Mu    sync.Mutex
	Value int
}
