package connect

import (
	"errors"
)

const testVersion = 3

func ResultOf(b []string) (string, error) {

	graphX, err := createGraph(b, XPlayer)
	if err != nil {
		return "", errors.New("Invalid board")
	}
	graphO, err := createGraph(b, OPlayer)
	if err != nil {
		return "", errors.New("Invalid board")
	}

	for y := range b {
		if Player(b[y][0]) == XPlayer {
			XWinner := DFS(&graphX, Position{0, y})
			if XWinner {
				return "X", nil
			}
		}
	}

	for x := range b[0] {
		if Player(b[0][x]) == OPlayer {
			OWinner := DFS(&graphO, Position{x, 0})
			if OWinner {
				return "O", nil
			}
		}
	}

	return "", nil
}

type Direction struct {
	x, y int
}

// Directions allowed for a hexagonal type grid
var directions = []Direction{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{1, -1},
	{-1, 1},
}

type Position struct {
	x, y int
}

type Graph struct {
	nodes map[Position]*Node
}

type Space struct {
	value byte
	goal  bool
	pos   Position
}

type Player byte

const (
	XPlayer = Player('X')
	OPlayer = Player('O')
)

type State int

const (
	Undiscovered State = iota
	Discovered
	Processed
)

type Node struct {
	payload Space
	adjNode []*Node
	state   State
}

func (g *Graph) AddNode(pos Position, value Space) {
	if _, found := g.nodes[pos]; found {
		panic("Node already exist!")
	}
	g.nodes[pos] = &Node{payload: value, adjNode: make([]*Node, 0)}
}

func (g *Graph) AddEdge(start, end Position) {

	startNode, ok := g.nodes[start]
	if !ok {
		panic("Node start not found!")
	}
	endNode, ok := g.nodes[end]
	if !ok {
		panic("Node end not found!")
	}

	if startNode.payload.value != endNode.payload.value {
		panic("Invalid state, should only connect nodes of same value")

	}

	// Not directed
	startNode.adjNode = append(startNode.adjNode, endNode)
	endNode.adjNode = append(endNode.adjNode, startNode)

}

func createGraph(board []string, player Player) (Graph, error) {
	g := Graph{nodes: make(map[Position]*Node)}
	for y, line := range board {
		for x, val := range line {
			var goal bool
			switch player {
			case XPlayer:
				if x == len(line)-1 {
					//Winner space is at the right
					goal = true
				}
			case OPlayer:
				if y == len(board)-1 {
					// Winner space is at the bottom
					goal = true
				}
			}
			g.AddNode(Position{x, y}, Space{byte(val), goal, Position{x, y}})
		}
	}
	for y, line := range board {
		for x := range line {

			startPos := Position{x, y}
			if !isPosOurPlayer(startPos, player, board) {
				// This space is not ours, no edge to add.
				continue
			}
			for _, dir := range directions {
				endPos := Position{x + dir.x, y + dir.y}
				if isPosInBoard(endPos, board) && isPosOurPlayer(endPos, player, board) {
					g.AddEdge(startPos, endPos)
				}
			}
		}
	}
	return g, nil
}

func isPosInBoard(pos Position, board []string) bool {
	// Assuming the board is rectangular
	if len(board) <= 0 || len(board) <= pos.y || pos.y < 0 {
		return false
	}
	if pos.x < 0 || len(board[0]) <= pos.x {
		return false
	}
	return true
}

func isPosOurPlayer(pos Position, player Player, board []string) bool {
	return Player(board[pos.y][pos.x]) == player
}

func DFS(g *Graph, pos Position) bool {

	node := g.nodes[pos]

	node.state = Discovered
	if node.payload.goal {
		return node.payload.goal
	}

	for _, adjNode := range node.adjNode {
		if adjNode.state == Undiscovered {
			found := DFS(g, adjNode.payload.pos)
			if found {
				return found
			}
		}
	}

	node.state = Processed
	return false
}
