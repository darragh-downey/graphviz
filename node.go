package graphviz

// Node represents element in grid
type Node struct {
	row, col                 int
	north, south, east, west int
}
