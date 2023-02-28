type node struct {
    key      []byte
    value    []byte
    version  int64
    height   int8
    size     int
    hash     []byte
    left     *node
    right    *node
    bloom    *bloom.BloomFilter // Add Bloom filter field
}

func (n *node) Has(key []byte) bool {
    if n.bloom != nil && !n.bloom.Test(key) {
        // If the key is not in the Bloom filter, it does not exist in the tree
        return false
    }
    // Perform an expensive search to confirm whether the key exists in the tree
    _, found := n.get(key)
    return found
}

func (n *node) Set(key, value []byte, version int64) *node {
    ...
    // Create a new node and add the key to the Bloom filter
    newNode := &node{key: key, value: value, version: version, height: 1, size: 1, bloom: bloom.New(1000000, 5)}
    newNode.bloom.Add(key)
    ...
    return newNode
}

func (n *node) Delete(key []byte) *node {
    ...
    // Remove the key from the Bloom filter
    if n.bloom != nil {
        n.bloom.Remove(key)
    }
    ...
    return n.rebalance()
}
