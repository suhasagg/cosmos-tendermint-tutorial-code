import (
    "github.com/tendermint/iavl"
    "github.com/willf/bloom"
)

type IavlPlusTree struct {
    iavl.Tree
    filter *bloom.BloomFilter
}

func NewIavlPlusTree() *IavlPlusTree {
    tree := iavl.NewMutableTree(nil, 100)
    filter := bloom.NewWithEstimates(1000000, 0.01)
    return &IavlPlusTree{Tree: tree, filter: filter}
}

func (t *IavlPlusTree) Set(key, value []byte) error {
    err := t.Tree.Set(key, value)
    if err == nil {
        t.filter.Add(key)
    }
    return err
}

func (t *IavlPlusTree) Has(key []byte) bool {
    if t.filter.Test(key) {
        return true
    }
    return t.Tree.Has(key)
}
