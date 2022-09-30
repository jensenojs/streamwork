package engine

type Node = string

type Edge = map[Node]Node

type WebServer struct {
	jobName   string
	sources   []Node
	operators []Node
	edges     Edge
}

func NewWebServer(jobName string, connectionList []Connection) *WebServer {
	var w = &WebServer{
		jobName: jobName,
	}
	var incomingCount = map[Node]int{}

	for _, c := range connectionList {
		from := c.from.GetName()
		to := c.to.GetName()

		count := incomingCount[to] // default value is 0
		incomingCount[from] = count
		incomingCount[to] = count + 1

		w.edges[from] = to
	}

	for k, v := range incomingCount {
		if v == 0 {
			w.sources = append(w.sources, k)
		} else {
			w.operators = append(w.operators, k)
		}
	}
	return w
}

func (w *WebServer) Start() {

}