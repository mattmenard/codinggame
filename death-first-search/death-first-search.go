package main

import (
	"fmt"
	"os"
	"sync"
	"strconv"
)


func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int

	var rmIndex1, rmIndex2 int

	graph := NewGraph()
	var exitGateways []int
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)

		if(!(graph.ContainsNode(strconv.Itoa(N1)))){
			node := &Node{strconv.Itoa(N1)}
			graph.AddNode(node)
		}
		
		if(!(graph.ContainsNode(strconv.Itoa(N2)))){
			node := &Node{strconv.Itoa(N2)}
			graph.AddNode(node)
		}

		graph.AddEdge(graph.GetNode(strconv.Itoa(N1)), graph.GetNode(strconv.Itoa(N2)))
    }

    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)

		exitGateways = append(exitGateways, EI)
    }

    for {
        // SI: The index of the node on which the Bobnet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)

		fmt.Fprint(os.Stderr,"  ||  Agent Node: ",SI)

		nodeSelected := false

		for egIndex := range exitGateways {
			if(graph.IsConnected(strconv.Itoa(SI), strconv.Itoa(exitGateways[egIndex]))) {
				rmIndex1 = SI
				rmIndex2 = exitGateways[egIndex]
				nodeSelected = true
				break
			}
		}

		if(!nodeSelected) {
			for egIndex := range exitGateways {
				eggw := exitGateways[egIndex]
				nodes := graph.GetNodes()
				
				for i := 0; i < len(nodes); i++ {
					if(graph.IsConnected(strconv.Itoa(eggw), nodes[i].name)) {
						rmIndex1, _ = strconv.Atoi(nodes[i].name)
						rmIndex2 = eggw
						break
					}
				}
			}
		}

		graph.RemoveEdge(graph.GetNode(strconv.Itoa(rmIndex1)), graph.GetNode(strconv.Itoa(rmIndex2)))

        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // Example: 0 1 are the indices of the nodes you wish to sever the link between
        fmt.Println(strconv.Itoa(rmIndex1),strconv.Itoa(rmIndex2))
    }
}


// ----- Undirected Graph -----
type Node struct {
	name  string
}

type Edge struct {
	node *Node
}

type UndirectedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

// --------------------------

// ------- Node Pair --------
type NodePair struct {
	node1 string
	node2 string
}

// --------------------------

// ++++++++++++++++++++++++++ Undirected Graph Methods ++++++++++++++++++++++++++

/******************************************************************************
 * Function: NewGraph
 * Description: Constructor function to create a new undirected graph instance
 *              and initialize the map of edges
 * Arguments: None
 * Returns:
 *    - Undirected Graph
******************************************************************************/
func NewGraph() *UndirectedGraph {
	return &UndirectedGraph{
		Edges: make(map[string][]*Edge),
	}
}

/******************************************************************************
 * Function: GetNode
 * Description: From the undirected graph instance, find and return the node
 *              with the supplied node.
 * Arguments:
 *     - name - string: The name of the node to find
 * Returns:
 *     - node - The node with the name matching the argument
******************************************************************************/
func (g *UndirectedGraph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

/******************************************************************************
 * Function: GetNodes
 * Description: From the undirected graph instance, get the list of all the
 *              nodes in the graph.
 * Arguments: Node
 * Returns:
 *     - nodes - A list of all the nodes in the graph
******************************************************************************/
func (g *UndirectedGraph) GetNodes() (nodes []*Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	nodes = g.Nodes
	return nodes
}

/******************************************************************************
 * Function: GetNodeNames
 * Description: From the undirected graph instance, get the list of the names
 *              of all the nodes in the graph.
 * Arguments: Node
 * Returns:
 *     - nodeNames - A list of the names of all the nodes in the graph
******************************************************************************/
func (g *UndirectedGraph) GetNodeNames() (nodeNames []string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for _, n := range g.Nodes {
		nodeNames = append(nodeNames, n.name)
	}

	return nodeNames
}

/******************************************************************************
 * Function: ContainsNode
 * Description: Returns true is the node is in the graph and false otherwise.
 * Arguments: None
 * Returns:
 *     - bool: True is the node was found and false otherwise.
******************************************************************************/
func (g *UndirectedGraph) ContainsNode(nodeName string) bool {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	found := false

	if(len(g.Nodes) > 0) {
		for _, n := range g.Nodes {
			if(n.String() == nodeName){
				found = true
				break;
			}
	    }
    }

	return found
}

/******************************************************************************
 * Function: AddNode
 * Description: Add a node to the undirected graph.
 * Arguments:
 *     - n - Node: The node to add to the graph
 * Returns: None
******************************************************************************/
func (g *UndirectedGraph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	g.Nodes = append(g.Nodes, n)
}

/******************************************************************************
 * Function: RemoveNode
 * Description: Remove a node to the undirected graph.
 * Arguments:
 *     - n - Node: The node to remove from the graph
 * Returns: None
******************************************************************************/
func (g *UndirectedGraph) RemoveNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if(g.ContainsNode(n.name)) {
		tempArr := make([]*Node, len(g.Nodes))

		for index := range g.Nodes {

			if(g.Nodes[index].name != n.name){
			    tempArr = append(tempArr, g.Nodes[index])
			}
		}

		g.Nodes = tempArr
	}
}

/*
*****************************************************************************
  - Function: AddNodes
  - Description: Add a collection of nodes to the undirected graph.
  - Arguments:
  - - graph - UndirectedGraph: The undirected graph instance to add the
    nodes to
  - - names - string: The names of the nodes to add to the undirected graph
  - Returns:
  - - nodes - Map: The map of nodes that were added to the undirected graph

*****************************************************************************
*/
func AddNodes(graph *UndirectedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

/******************************************************************************
 * Function: AddEdge
 * Description: Adds an edge connecting 2 nodes to the graph.
 * Arguments:
 *     - n1 - Node: The first of 2 nodes connected by the edge
 *     - n2 - Node: The second of 2 nodes connected by the edge
*      - weight - int: The weight of the edge between the nodes
 * Returns: None
******************************************************************************/
func (g *UndirectedGraph) AddEdge(n1, n2 *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1})
}

/******************************************************************************
 * Function: IsConnected
 * Description: Determines if 2 nodes in the graph are connected to each other
 * Arguments:
 *     - node1 - string: The name of the first node
 *     - node2 - string: The name of the second node
 * Returns: bool - True if the nodes are connected and false otherwise.
******************************************************************************/
func (g *UndirectedGraph) IsConnected(node1, node2 string) bool {
	isconnected := false

	edges := g.Edges[node1]

	for _, edge := range edges {
		if (edge.node.name) == node2 {
			isconnected = true
			break
		}
	}

	return isconnected
}

/******************************************************************************
 * Function: RemoveEdge
 * Description: Removed an edge connecting 2 nodes in the graph.
 * Arguments:
 *     - n1 - Node: The first of 2 nodes connected by the edge
 *     - n2 - Node: The second of 2 nodes connected by the edge
 * Returns: None
******************************************************************************/
func (g *UndirectedGraph) RemoveEdge(n1, n2 *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if(g.IsConnected(n1.name, n2.name)) {
		
		delete(g.Edges, n1.name)
		delete(g.Edges, n2.name)
	}
}

/******************************************************************************
 * Function: String
 * Description: Returns the name of the node.
 * Arguments: None
 * Returns:
 *     - string: The name of the node.
******************************************************************************/
func (n *Node) String() string {
	return n.name
}

/******************************************************************************
 * Function: String
 * Description: Returns the name of the node and value of the edge.
 * Arguments: None
 * Returns:
 *     - string: The name of the node and the value of the edge.
******************************************************************************/
func (e *Edge) String() string {
	return e.node.String()
}

/******************************************************************************
 * Function: String
 * Description: Returns the names of all the nodes and the values of all the
 *              edges contained in the graph.
 * Arguments: None
 * Returns:
 *     - string: The names of all the nodes and the values of all the edges.
******************************************************************************/
func (g *UndirectedGraph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		s = s + n.String() + " ->"
		for _, c := range g.Edges[n.name] {
			s = s + " " + c.node.String()
		}
		s = s + "\n"
	}
	return
}
