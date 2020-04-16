package graphviz

import "syscall/js"

func (g *GraphViz) setupAStar() {
	g.astar = js.Func(func(arg js.Value, args []js.Value) interface{} {
		// A* implementation
		return nil
	})
}
