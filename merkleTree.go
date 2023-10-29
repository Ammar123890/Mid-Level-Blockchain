package MidLevelBlockchain

import (
	"crypto/sha256"
	"fmt"
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
		fmt.Printf("Leaf Node created with data: %x\n", node.Data) // Display leaf node creation
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
		fmt.Printf("Parent Node created with left child: %x and right child: %x resulting in hash: %x\n", left.Data, right.Data, node.Data) // Display parent node creation
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

	fmt.Println("Leaf nodes created...") // Initial statement after all leaf nodes are created

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

		fmt.Println("Next level of parent nodes created...") // Statement after a new level of parent nodes is created

		nodes = level
	}

	tree := MerkleTree{&nodes[0]}
	return &tree
}
