package tree

import (
	"errors"
	"sort"
)

const testVersion = 4

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type NodeSlice []*Node

func (n NodeSlice) Len() int {
	return len(n)
}

func (n NodeSlice) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NodeSlice) Less(i, j int) bool { return n[i].ID < n[j].ID }

type Stack struct {
	nodes []*Node

	count int
}

func NewStack(size int) *Stack {
	s := Stack{}
	s.nodes = make([]*Node, size)
	return &s
}

func (s *Stack) Push(n *Node) {
	if len(s.nodes) > s.count {
		s.nodes[s.count] = n
		s.count++
		return
	}
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Validate the records
	for _, r := range records {
		if r.ID <= r.Parent && r.ID != r.Parent {
			return nil, errors.New("ID equal or greater than Parent, while not being root!")
		}
	}

	// Find root
	var root Node
	var rootFound bool
	processedRecord := make(map[int]bool)
	for _, r := range records {
		if r.ID == r.Parent {
			root = Node{ID: r.ID}
			rootFound = true
			break
		}
	}
	if !rootFound {
		return nil, errors.New("root not found, invalid records")
	}
	nodesStack := *NewStack(100)
	nodesStack.Push(&root)
	processedRecord[root.ID] = true
	nodeSize := 1
	for nodeSize < len(records) {
		currentNode := nodesStack.Pop()
		if currentNode == nil {
			break
		}
		nodeSize += findChildren(records, currentNode)
		for _, c := range currentNode.Children {
			if processedRecord[c.ID] {
				return nil, errors.New("Already processed this record id. Either cycle or duplicates")
			}
			nodesStack.Push(c)
			processedRecord[c.ID] = true
		}
	}

	// Testing for continuous ids
	for i := range records {
		if !processedRecord[i] {
			return nil, errors.New("Non-countinous record id!")
		}
	}

	return &root, nil

}

func findChildren(records []Record, node *Node) (foundCount int) {

	// Linear Search for all children of node in the record
	for _, r := range records {
		if r.Parent == node.ID && r.ID != r.Parent {
			node.Children = append(node.Children, &Node{ID: r.ID})
			foundCount++
		}
	}
	sort.Sort(NodeSlice(node.Children))
	return

}
