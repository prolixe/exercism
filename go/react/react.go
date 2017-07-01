package react

const testVersion = 5

type reactor struct {
	cells []*cell
}

func New() Reactor {

	return &reactor{}
}

type canceler struct {
	cancel func()
}

func (c *canceler) Cancel() {
	c.cancel()
}

type cell struct {
	id         int
	value      int
	react      *reactor
	parents    []*cell
	children   []*cell
	function   func(...int) int
	callbacks  map[int]func(int)
	callbackId int
	dirty      bool
}

func (r *reactor) CreateInput(i int) InputCell {
	newInput := &cell{
		id:    len(r.cells),
		value: i,
		react: r}
	r.cells = append(r.cells, newInput)
	return newInput
}

func (r *reactor) CreateCompute1(parent Cell, f func(int) int) ComputeCell {
	p := parent.(*cell)
	newComputeCell := &cell{
		id:        len(r.cells),
		value:     f(p.value),
		react:     r,
		parents:   []*cell{p},
		function:  func(args ...int) int { return f(args[0]) },
		callbacks: make(map[int]func(int)),
	}
	p.children = append(p.children, newComputeCell)
	r.cells = append(r.cells, newComputeCell)
	return newComputeCell
}

func (r *reactor) CreateCompute2(parent1, parent2 Cell, f func(int, int) int) ComputeCell {

	p1 := parent1.(*cell)
	p2 := parent2.(*cell)

	newComputeCell := &cell{
		id:        len(r.cells),
		value:     f(p1.value, p2.value),
		react:     r,
		parents:   []*cell{p1, p2},
		function:  func(args ...int) int { return f(args[0], args[1]) },
		callbacks: make(map[int]func(int)),
	}

	p1.children = append(p1.children, newComputeCell)
	p2.children = append(p2.children, newComputeCell)

	r.cells = append(r.cells, newComputeCell)
	return newComputeCell
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) update() {
	if c.dirty == false {
		panic("Should not have called update on a non dirty cell")
	}
	c.dirty = false
	//Re-eval!
	pValues := make([]int, 0)
	for _, p := range c.parents {
		pValues = append(pValues, p.value)
	}
	oldValue := c.value

	c.value = c.function(pValues...)

	if c.value != oldValue {
		// If the value changed, set all children to dirty
		for _, child := range c.children {
			child.dirty = true
		}
		// and call all the callbacks
		for _, callback := range c.callbacks {
			callback(c.value)
		}
	}

}

func (c *cell) SetValue(v int) {
	c.value = v
	c.react.update(c.id)
}

func (c *cell) AddCallback(f func(int)) Canceler {

	id := c.callbackId
	c.callbackId++
	c.callbacks[id] = f
	return &canceler{
		cancel: func() {
			delete(c.callbacks, id)
		},
	}
}

func (r *reactor) update(id int) {

	// Check all children of cells[id]
	c := r.cells[id]
	var hasDirtyCell bool
	for _, child := range c.children {
		child.dirty = true
		hasDirtyCell = true
	}

	// loop over until no more dirty cell remains.
	// Be careful to go to each cell by creation order for each loop,
	// to avoid setting a cell dirty more than once!
	for hasDirtyCell {
		hasDirtyCell = false
		for _, cell := range r.cells[id:] {
			if cell.dirty {
				cell.update()
				hasDirtyCell = true
			}
		}
	}
}
