// Package tree provides a refactored Go solution for the Exercism tree-building slug.
package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build takes a list of array of records and build a tree of records using their parent and record ids
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if len(records) == 0 {
		return nil, nil
	} else if err := newChk(records); err != nil {
		return nil, err
	} else if root := newBuild(records, &Node{}); root == nil {
		return nil, fmt.Errorf("Bad tree, naughty tree")
	} else {
		return root, nil
	}
}

// newChk sanity checks the record ids
func newChk(records []Record) (err error) {
	id := 0

	if records[0].Parent != 0 || records[0].ID != 0 {
		return fmt.Errorf("Bad root record")
	}

	for ii, record := range records {
		if ii != 0 {
			if record.ID >= len(records) {
				return fmt.Errorf("Record id out of range")
			}

			if record.ID <= record.Parent {
				return fmt.Errorf("Bad parent id")
			}

			if id == record.ID {
				return fmt.Errorf("Bad record id")
			}
		}

		id = records[ii].ID
	}

	return nil
}

// newBuild builds the record tree their parent and record ids
func newBuild(records []Record, root *Node) *Node {
	node := root

	for ii, record := range records {
		if ii != 0 {
			if record.Parent != node.ID {
				node = findNode(node, record.Parent)

				if node == nil || node.ID != record.Parent {
					node = findNode(root, record.Parent)
				}

				if node == nil || node.ID != record.Parent {
					return nil
				}
			}

			node.Children = append(node.Children, &Node{ID: record.ID})
		}
	}

	return root
}

// findNode returns the subnode with the given id or nil if not found
func findNode(node *Node, id int) *Node {
	for _, child := range node.Children {
		if child.ID == id {
			return child
		} else if child.ID > id {
			return nil
		} else if node := findNode(child, id); node != nil {
			return node
		}
	}

	return nil
}

//
// A disappointing lot - no one gave any before/after benchmark figures,
// only one had comments and one looked like the orignal.
//
// All flattened the loop structure, most used sort - something like
// two on the children and four on all records.
//
// Several used a map or just an array to find a node given an ID where
// I used a search routine.  I should have done that too.
//
// Several preallocated an array of nodes.  Good idea.
//
// Mine was about the longest solution - twice a long as the shortest.
// Several did not look neat or tidy.
//
