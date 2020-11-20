package tree

import (
	"fmt"
	"sort"
)

/*
	BenchmarkTwoTree BenchmarkTwoTree-8
		162
		7425929 ns/op
		256 B/op
		10 allocs/op
	BenchmarkTenTree BenchmarkTenTree-8
		1052
		1106465 ns/op
		5777 B/op
		153 allocs/op
	BenchmarkShallowTree BenchmarkShallowTree-8
		612
		1934966 ns/op
		706396 B/op
		10023 allocs/op
*/

// Record contains ID and Parent ID fields
type Record struct {
	ID     int
	Parent int
}

// Node contains ID and Children fields
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a Tree from a list of records (id, parentID)
func Build(record []Record) (*Node, error) {
	if len(record) == 0 {
		return nil, nil
	}
	// Sort record by record.Parent && record.ID
	sort.Slice(record, func(i, j int) bool {
		if record[i].Parent < record[j].Parent {
			return true
		}
		if record[i].Parent > record[j].Parent {
			return false
		}
		return record[i].ID < record[j].ID
	})

	var t *Node
	for idx, r := range record {
		if isInvalid, errMsg := invalidRecord(&record, idx); isInvalid {
			return &Node{0, nil}, errMsg
		}
		t = insert(t, r)
	}
	return t, nil
}

func insert(t *Node, r Record) *Node {
	if t == nil {
		return &Node{r.ID, nil}
	}
	if r.Parent == t.ID {
		t.Children = append(t.Children, &Node{r.ID, nil})
		return t
	}
	for _, child := range t.Children {
		if r.Parent == child.ID {
			child.Children = append(child.Children, &Node{r.ID, nil})
		}
	}
	return t
}

func invalidRecord(record *[]Record, idx int) (bool, error) {
	// Assert 	0 <= record.ID < len(record)
	// 			record.Parent < record.ID
	// 			no duplicates
	// 			for all non-roots:
	// 				parent != id (some could be combined)
	// 				parent < id  (but I prefer obvious error messages)
	if (*record)[idx].ID >= len(*record) {
		return true, fmt.Errorf("record.ID (%v) >= len(record)", (*record)[idx].ID)
	} else if (*record)[idx].ID == 0 && (*record)[idx].Parent != 0 {
		return true, fmt.Errorf("root node has parent")
	} else if idx > 0 {
		if (*record)[idx] == (*record)[idx-1] {
			return true, fmt.Errorf("duplciate node")
		} else if (*record)[idx].Parent == (*record)[idx].ID {
			return true, fmt.Errorf("record parent == record id")
		} else if (*record)[idx].Parent > (*record)[idx].ID {
			return true, fmt.Errorf("record.Parent must be < record.ID")
		}
	}
	return false, nil
}
