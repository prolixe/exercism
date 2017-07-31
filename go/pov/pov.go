package pov

import "fmt"

const testVersion = 2

type Graph struct {
	edges map[string]*edgeNode
}

type edgeNode struct {
	adjacent string
	next     *edgeNode
	previous *edgeNode
}

func New() *Graph {

	return &Graph{edges: make(map[string]*edgeNode)}
}

// AddNode add a Node (also called a vertice)
func (g *Graph) AddNode(nodeLabel string) {
	pEdge := &edgeNode{}
	pEdge.next = g.edges[nodeLabel]
	if pEdge.next != nil {
		pEdge.next.previous = pEdge
	}
	g.edges[nodeLabel] = pEdge
}

// AddArc add an arc (or an edges) betwen 2 nodes
func (g *Graph) AddArc(from, to string) {
	g.AddNode(from)
	g.edges[from].adjacent = to
}

func (g *Graph) ArcList() []string {
	arc := make([]string, 0)
	for label, node := range g.edges {
		pEdge := node
		for pEdge != nil && len(pEdge.adjacent) > 0 {
			arc = append(arc, fmt.Sprintf("%s -> %s", label, pEdge.adjacent))
			pEdge = pEdge.next
		}
	}
	return arc
}

func (g *Graph) removeArc(from, to string) {
	if g.edges[from] != nil {
		pEdge := g.edges[from]
		for pEdge != nil && pEdge.adjacent != to {
			pEdge = pEdge.next
			if g.edges[from] == pEdge {
				fmt.Println("looping!!!")
				break
			}
		}
		if pEdge == nil {
			return
		}
		if pEdge.previous != nil {
			pEdge.previous.next = pEdge.next
		}
		if pEdge.next != nil {
			pEdge.next.previous = pEdge.previous
		}
		if pEdge == g.edges[from] {
			g.edges[from] = pEdge.next
		}
	}
}

func (g *Graph) findParent(child string) (*edgeNode, string) {
	for label, node := range g.edges {
		pEdge := node
		for pEdge != nil {
			if pEdge.adjacent == child {
				return pEdge, label
			}
			pEdge = pEdge.next
		}
	}
	return nil, ""
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// Change the direction of every edges between the old and new root.
	var parentName string
	_, parentName = g.findParent(newRoot)
	childName := newRoot
	for childName != oldRoot {
		//fmt.Printf("child '%s' is becoming the parent of '%s'\n", childName, parentName)
		_, parentParentName := g.findParent(parentName)
		//fmt.Printf("the parent of '%s' is '%s'\n", parentName, parentParentName)
		g.removeArc(parentName, childName)
		g.AddArc(childName, parentName)

		childName = parentName
		parentName = parentParentName
		if parentParentName == "" {
			break
		}

	}
	return g
}
