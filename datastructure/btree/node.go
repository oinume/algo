package btree

const maxChildCount = 5
const halfChildCount = (maxChildCount + 1) / 2

func NewItem(key int64, value interface{}) Item {
	return Item{key: key, value: value}
}

type Item struct {
	key   int64
	value interface{}
}

func (i Item) Key() int64 {
	return i.key
}

func (i Item) Value() interface{} {
	return i.value
}

type Node interface {
	Item() Item
	Children() []Node
	Items() []Item
	IsLeaf() bool
}

func NewNode(item Item) Node {
	return &node{
		item:     item,
		children: make([]Node, 0, maxChildCount),
		items:    make([]Item, 0, maxChildCount),
	}
}

type node struct {
	item     Item
	children []Node
	items    []Item
}

func (n *node) Item() Item {
	return n.item
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) Items() []Item {
	return n.items
}

func (n *node) IsLeaf() bool {
	return len(n.Children()) == 0
}
