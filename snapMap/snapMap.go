package snapMap

import (
	"github.com/infinivision/common/bheap"
	"github.com/infinivision/common/spinlock"
)

func New() *hashMap {
	return &hashMap{Locker: spinlock.New(), hp: bheap.New(), mp: make(map[string]string)}
}
