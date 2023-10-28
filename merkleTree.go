package MidLevelBlockchain

import (
	"crypto/sha256"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	// Create leaf nodes for each data block
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	// While there's more than 1 node, keep hashing till we reach root
	for len(nodes) > 1 {
		level := []MerkleNode{}

		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
				level = append(level, *node)
			} else {
				// Handle odd number of nodes
				node := NewMerkleNode(&nodes[i], &nodes[i], nil)
				level = append(level, *node)
			}
		}

		nodes = level
	}

	tree := MerkleTree{&nodes[0]}
	return &tree
}
