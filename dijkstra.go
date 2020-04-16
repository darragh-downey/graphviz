package graphviz

import "syscall/js"

func (g *GraphViz) setupDijkstra() {
	g.dijkstra = js.Func(func(arg js.Value, args []js.Value) interface{} {
		// Dijkstra implementation
		return nil
	})
}
