package snapMap

import (
	"github.com/infinivision/common/bheap"
)

func New() *hashMap {
	return &hashMap{hp: bheap.New(), mp: make(map[string]string)}
}
