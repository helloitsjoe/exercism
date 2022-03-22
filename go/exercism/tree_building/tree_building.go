// Package tree_building builds trees
package tree_building

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Seen = map[int]bool

func insert(parent *Node, toInsert Record) *Node {
	if parent.ID == toInsert.Parent {
		parent.Children = append(parent.Children, &Node{toInsert.ID, nil})
		return parent
	}

	for _, child := range parent.Children {
		if child.ID == toInsert.Parent {
			child.Children = append(child.Children, &Node{toInsert.ID, nil})
		} else if child.Children != nil {
			insert(child, toInsert)
		}
	}

	return parent
}

func validateRecords(records []Record) error {
	if records[0].Parent > 0 {
		return errors.New("root node has parent")
	}

	if records[0].ID > 0 {
		return errors.New("no root node")
	}

	seen := Seen{}

	for i := 1; i < len(records); i++ {
		record := records[i]

		if seen[record.ID] == true {
			return errors.New("duplicate node")
		}

		if record.ID == 0 {
			return errors.New("duplicate root node")
		}

		if record.ID != records[i-1].ID+1 {
			return errors.New("non-continuous")
		}

		if record.ID == record.Parent {
			return errors.New("non-continuous")
		}

		if record.ID < record.Parent {
			return errors.New("parent is higher than record")
		}

		seen[record.ID] = true
	}

	return nil
}

// Build the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if err := validateRecords(records); err != nil {
		return nil, err
	}

	root := &Node{records[0].ID, nil}

	for _, record := range records[1:] {
		root = insert(root, record)
	}

	return root, nil
}
