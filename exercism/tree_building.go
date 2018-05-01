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
	} else if root := newBuild(records); root == nil {
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
func newBuild(records []Record) *Node {
	nodeArray := make([]Node, len(records))

	for ii, record := range records {
		nodeArray[ii].ID = ii

		if ii != 0 {
			node := &nodeArray[record.Parent]

			node.Children = append(node.Children, &nodeArray[ii])
		}
	}

	return &nodeArray[0]
}

//
// This iteration uses a array of nodes to cut out long look up times.
//
// The result for the first benchmark is something like 700 times faster.
//
// This is average length solution but still 20 lines longer than the shortest.
//
