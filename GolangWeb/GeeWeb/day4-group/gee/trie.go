package gee

import (
	"fmt"
	"strings"
)

// 前缀树节点
type node struct {
	Pattern  string
	Part     string
	Children []*node
	IsWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{Pattern=%s, Part=%s, IsWild=%t}", n.Pattern, n.Part, n.IsWild)
}

// 插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	Part := parts[height]
	child := n.matchChild(Part)
	if child == nil {
		child = &node{Part: Part, IsWild: Part[0] == ':' || Part[0] == '*'}
		n.Children = append(n.Children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	Part := parts[height]
	Children := n.matchChildren(Part)

	for _, child := range Children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// 把所有节点加入 list
func (n *node) travel(list *([]*node)) {
	if n.Pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.Children {
		child.travel(list)
	}
}

// 匹配单个子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

// 匹配所有符合条件的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
