package pokectx

type te4nickNode struct {
	children map[string]*te4nickNode
	value    string
}

func (n *te4nickNode) set(path ...string) {
	if len(path) == 0 {
		return
	}

	if len(path) == 1 {
		n.value = path[0]
		return
	}

	child, ok := n.children[path[0]]
	if !ok {
		child = &te4nickNode{
			children: make(map[string]*te4nickNode),
		}
		n.children[path[0]] = child
	}

	child.set(path[1:]...)
}

func (n *te4nickNode) get(path ...string) (value string, found bool) {
	if len(path) == 0 {
		return n.value, true
	}

	child, ok := n.children[path[0]]
	if !ok {
		return "", false
	}

	return child.get(path[1:]...)
}

func (n *te4nickNode) getChildrenNames(path ...string) (childrenNames []string, found bool) {
	if len(path) == 0 {
		for name := range n.children {
			childrenNames = append(childrenNames, name)
		}
		return
	}

	child, ok := n.children[path[0]]
	if !ok {
		return nil, false
	}

	return child.getChildrenNames(path[1:]...)
}
