package graphviz

import "syscall/js"

// GraphViz struct representing wasm callbacks
type GraphViz struct {
	graph           [][]Node
	dijkstra, astar js.Func
	bfs, dfs        js.Func
	shutdown        js.Func

	console js.Value
	done    chan struct{}
}

// New construct new GraphViz struct
func New() *GraphViz {
	return &GraphViz{
		console: js.Global().Get("console"),
		done:    make(chan struct{}),
	}
}

// Start setup all callbacks
func (g *GraphViz) Start() {
	// Setup callbacks
	g.setupBFS()
	js.Global().Set("bfs", g.bfs)

	g.setupDFS()
	js.Global().Set("dfs", g.bfs)

	<-g.done
	g.log("Shutting down app")
	g.bfs.Release()
	g.dfs.Release()
	g.dijkstra.Release()
	g.astar.Release()
	g.shutdown.Relase()
}

// utility function to log a message to the UI from inside a callback
func (g *GraphViz) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}

func (g *GraphViz) setupShutdown() {
	g.shutdown = js.Func(func(this js.Value, args []js.Value) interface{} {
		g.done <- struct{}{}
		return nil
	})
}
