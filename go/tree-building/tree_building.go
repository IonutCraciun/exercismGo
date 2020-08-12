package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Records slice
type Records []Record

// Record struct
type Record struct {
	ID, Parent int
}

// Node struct
type Node struct {
	ID       int
	Children []*Node
}

func toString(n []*Node) {

	for _, val := range n {
		fmt.Printf("ID %d ", val.ID)
		for _, val2 := range val.Children {
			fmt.Printf("Child: %d ", val2.ID)
		}
		fmt.Println()
	}
}

// Build tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	//sorting the records
	sort.Sort(Records(records))
	if records[0].ID != 0 {
		return nil, errors.New("no root")
	}

	if records[0].Parent != 0 {
		return nil, errors.New("root is not its own parent")
	}

	root := &Node{ID: 0}
	results := make([]*Node, len(records))
	results[0] = root

	for i := 1; i < len(records); i++ {
		currentNode := &Node{ID: records[i].ID}
		parentNode := results[records[i].Parent]

		if records[i].ID != i {
			return nil, fmt.Errorf("missing record with ID %d", i)
		}
		if records[i].Parent >= records[i].ID {
			return nil, errors.New("parent id is not less than child id")
		}

		if parentNode.Children == nil {
			parentNode.Children = []*Node{}
		}

		parentNode.Children = append(parentNode.Children, currentNode)
		results[records[i].ID] = currentNode

	}
	return root, nil
}

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i int, j int) bool {
	return r[i].ID < r[j].ID
}
func (r Records) Swap(i int, j int) {
	temp := r[i]
	r[i] = r[j]
	r[j] = temp
}
