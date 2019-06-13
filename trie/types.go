package trie

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io/ioutil"
	"sync"
)

const (
	K_LIMIT  = 0x0F
	CHILDREN = 0x100
)

type FoldFunc (func([]byte, []byte, interface{}) interface{})

type Trie interface {
	Get([]byte) []byte
	Set([]byte, []byte) error
	Fold(FoldFunc, interface{}) interface{}
}

type tree struct {
	sync.Mutex
	tree *TreeNode
}

type TreeNode struct {
	Value    []byte
	Children map[uint8]*TreeNode
}

func (t *tree) Fold(f FoldFunc, b interface{}) interface{} {
	t.Lock()
	defer t.Unlock()
	return t.tree.Fold([]byte{}, f, b)
}

func (t *tree) Get(k []byte) []byte {
	t.Lock()
	defer t.Unlock()
	return t.tree.Get(k)
}

func (t *tree) Set(k, v []byte) error {
	t.Lock()
	defer t.Unlock()
	return t.tree.Set(k, v)
}

func (a *tree) Show() []byte {
	var buf, data bytes.Buffer

	zw := gzip.NewWriter(&buf)
	if err := gob.NewEncoder(&data).Encode(a.tree); err != nil {
		return []byte{}
	}
	zw.Write(data.Bytes())
	zw.Close()
	return buf.Bytes()
}

func (a *tree) Read(buf []byte) ([]byte, error) {
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
	if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&a.tree); err != nil {
		return []byte{}, err
	}
	return []byte{}, nil
}
