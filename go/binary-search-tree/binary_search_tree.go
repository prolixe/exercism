package binarysearchtree

const testVersion = 1

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(d int) *SearchTreeData {
	return &SearchTreeData{data: d}
}
func (std *SearchTreeData) Insert(d int) {
	if d <= std.data {
		if std.left != nil {
			std.left.Insert(d)
		} else {
			std.left = Bst(d)
		}
	} else {
		if std.right != nil {
			std.right.Insert(d)
		} else {
			std.right = Bst(d)
		}
	}
}

func (std *SearchTreeData) walkTreeSlice(f func(int) string, s *[]string) {
	if std.left != nil {
		std.left.walkTreeSlice(f, s)
	}
	*s = append(*s, f(std.data))
	if std.right != nil {
		std.right.walkTreeSlice(f, s)
	}
}
func (std *SearchTreeData) walkTreeInt(f func(int) int, s *[]int) {
	if std.left != nil {
		std.left.walkTreeInt(f, s)
	}
	*s = append(*s, f(std.data))
	if std.right != nil {
		std.right.walkTreeInt(f, s)
	}

}
func (std *SearchTreeData) MapString(f func(int) string) []string {
	s := make([]string, 0)
	std.walkTreeSlice(f, &s)
	return s

}
func (std *SearchTreeData) MapInt(f func(int) int) []int {
	i := make([]int, 0)
	std.walkTreeInt(f, &i)
	return i
}
