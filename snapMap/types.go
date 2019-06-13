package snapMap

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"

	"github.com/infinivision/common/bheap"
	"github.com/infinivision/common/miscellaneous"
	"github.com/infinivision/common/typeclass"
)

type FoldFunc (func([]byte, []byte, interface{}) interface{})

type Map interface {
	Get([]byte) []byte
	Set([]byte, []byte) error
	Fold(FoldFunc, interface{}) interface{}
}

type hashMap struct {
	sync.Locker
	hp bheap.BHeap
	mp map[string]string
}

func (m *hashMap) Fold(f FoldFunc, b interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	for k, v := range m.mp {
		b = f([]byte(k), []byte(v), b)
	}
	return b
}

func (m *hashMap) Get(k []byte) []byte {
	m.Lock()
	defer m.Unlock()
	if v, ok := m.mp[string(k)]; !ok {
		return []byte{}
	} else {
		return []byte(v)
	}
}

func (m *hashMap) Set(k, v []byte) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.mp[string(k)]; !ok {
		m.hp.Insert(typeclass.NewString(string(k)))
	}
	m.mp[string(k)] = string(v)
	return nil
}

func (a *hashMap) show() []byte {
	data := []byte{}
	for x := a.hp.Extract(); x != nil; x = a.hp.Extract() {
		k := x.(typeclass.OrdString).String()
		v := a.mp[k]
		data = append(data, miscellaneous.Eslice([]byte(k))...)
		data = append(data, miscellaneous.Eslice([]byte(v))...)
	}
	return data
}

func (a *hashMap) Show() []byte {
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	zw.Write(a.show())
	zw.Close()
	return buf.Bytes()
}

func (a *hashMap) Read(buf []byte) ([]byte, error) {
	var err error
	var k, v []byte

	zr, err := gzip.NewReader(bytes.NewReader(buf))
	if err != nil {
		return []byte{}, err
	}
	if err = zr.Close(); err != nil {
		return []byte{}, err
	}
	data, err := ioutil.ReadAll(zr)
	if err != nil {
		return []byte{}, err
	}
	for len(data) > 0 {
		if k, data, err = miscellaneous.Dslice(data); err != nil {
			return []byte{}, err
		}
		if v, data, err = miscellaneous.Dslice(data); err != nil {
			return []byte{}, err
		}
		if err = a.Set(k, v); err != nil {
			return []byte{}, err
		}
	}
	return data, nil
}
