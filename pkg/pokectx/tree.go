package pokectx

type node struct {
	children map[string]*node
	value    string
}

func (n *node) set(path ...string) {
	if len(path) == 0 {
		return
	}

	if len(path) == 1 {
		n.value = path[0]
		return
	}

	child, ok := n.children[path[0]]
	if !ok {
		child = &node{
			children: make(map[string]*node),
		}
		n.children[path[0]] = child
	}

	child.set(path[1:]...)
}

func (n *node) get(path ...string) (string, bool) {
	if len(path) == 0 {
		return n.value, true
	}

	child, ok := n.children[path[0]]
	if !ok {
		return "", false
	}

	return child.get(path[1:]...)
}
