package main

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

type Node struct {
	Key   []byte
	Value []byte
	Left  *Node
	Right *Node
}

type MerklePatriciaTree struct {
	Root *Node
}

func (mpt *MerklePatriciaTree) Put(key, value []byte) {
	mpt.Root = mpt.putHelper(mpt.Root, key, value, 0)
}

func (mpt *MerklePatriciaTree) putHelper(current *Node, key, value []byte, index int) *Node {
	if current == nil {
		// If the node doesn't exist, create a new node and add the key-value pair
		return &Node{Key: key, Value: value}
	}
	if len(current.Key) == 0 {
		// If the node has an empty key, it's a leaf node and should be updated with the new value
		current.Value = value
		return current
	}

	// Compare the current node's key with the key we're trying to insert
	if bytes.HasPrefix(key[index:], current.Key) {
		// If the current node's key is a prefix of the key we're trying to insert,
		// recursively update the left subtree with the remaining part of the key
		current.Left = mpt.putHelper(current.Left, key, value, index+len(current.Key))
	} else {
		// If the current node's key is not a prefix of the key we're trying to insert,
		// recursively update the right subtree with the remaining part of the key
		current.Right = mpt.putHelper(current.Right, key, value, index)
	}

	// Update the node's key by taking the longest common prefix of the left and right subtrees
	prefix := longestCommonPrefix(current.Left, current.Right)
	current.Key = prefix
	return current
}

func (mpt *MerklePatriciaTree) Get(key []byte) ([]byte, error) {
	node := mpt.getHelper(mpt.Root, key, 0)
	if node == nil {
		return nil, fmt.Errorf("key not found")
	}
	return node.Value, nil
}

func (mpt *MerklePatriciaTree) getHelper(current *Node, key []byte, index int) *Node {
	if current == nil {
		return nil
	}
	if len(current.Key) == 0 {
		if bytes.Equal(current.Value, key) {
			return current
		}
		return nil
	}
	if bytes.HasPrefix(key[index:], current.Key) {
		return mpt.getHelper(current.Left, key, index+len(current.Key))
	}
	return mpt.getHelper(current.Right, key, index)
}

func longestCommonPrefix(left, right *Node) []byte {
	minLength := min(len(left.Key), len(right.Key))
	for i := 0; i < minLength; i++ {
		if left.Key[i] != right.Key[i] {
			return left.Key[:i]
		}
	}
	return left.Key[:minLength]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (n *Node) Hash() []byte {
	enc, _ := rlp.Encode([]interface{}{n.Key, n.Value, n.Left, n.Right})
	return enc
}

func (mpt *MerklePatriciaTree) Hash() []byte {
	return mpt.Root.Hash()
}
