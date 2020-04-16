package graphviz

import "syscall/js"

func (g *GraphViz) setupDFS() {
	g.dfs = js.Func(func(arg js.Value, args []js.Value) interface{} {
		// DFS implementation
		return nil
	})
}
