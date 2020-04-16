package graphviz

import "syscall/js"

func (g *GraphViz) setupBFS() {
	g.bfs = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// BFS implementation
		return nil
	})
}

func bfs() {

}
