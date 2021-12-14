package util

import (
)

type GraphNode struct {
	Id string
	Connected map[string]GraphNode
	Data interface{}
}

func MakeNode(id string, data interface{}) GraphNode {
	return GraphNode{
		Id: id,
		Connected: map[string]GraphNode{},
		Data: data,
	}
}


func (g GraphNode) IsConnected(node GraphNode) bool {
	if _,exists := g.Connected[node.Id]; exists {
		return true
	}

	return false
}

type Graph map[string]GraphNode

func (g Graph) AddNode(node GraphNode) {
	g[node.Id] = node
}

func (g Graph) GetNode(id string) *GraphNode {

	// Make node
	if node,exists := g[id]; exists {
		return &node
	}

	return nil
}

func (g Graph) NodeExists(id string) bool {

	return !(g.GetNode(id) == nil)
}

func (g Graph) ConnectNodes(aId string, bId string) {

	if g.NodeExists(aId) && g.NodeExists(bId) {
		a := *g.GetNode(aId)
		b := *g.GetNode(bId)

		// Connect a to b
		if !a.IsConnected(b) {
			a.Connected[b.Id] = b
		}

		// Connect b to a
		if !b.IsConnected(a) {
			b.Connected[a.Id] = a
		}
	}
}

func (g Graph) Traverse(
	startAt string,
	actionFunc func(node GraphNode, path []string) bool,
	visitFunc func(node GraphNode, path []string) bool,
) {
	g.visit(startAt, actionFunc, visitFunc, []string{startAt})
}

func (g Graph) visit(
	at string,
	actionFunc func(node GraphNode, path []string) bool,
	visitFunc func(node GraphNode, path []string) bool,
	path []string,
) {

	if g.NodeExists(at) {

		node := *g.GetNode(at)

        if !actionFunc(node, path) {
			return
        }

		if len(node.Connected) > 0 {
			for connectedId,childNode := range node.Connected {

				if visitFunc(childNode, path) {
					g.visit(connectedId, actionFunc, visitFunc, append(path, connectedId))
				}

			}
		}
	}
}
