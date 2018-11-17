package btree

import "fmt"

var ErrAlreadyExists = fmt.Errorf("the key already exists")

type Tree struct {
	root Node
}

func NewTree(root Node) *Tree {
	return &Tree{
		root: root,
	}
}

func (tree *Tree) Insert(item Item) error {
	if tree.root == nil {
		tree.root = NewNode(item)
		return nil
	}

	tree.insert(tree.root, item)
	return nil
}

/*
// 木 t にエントリー e を挿入する
    private Node insert(Node t, Pair e) {
        if (t == null) return new NodeA(e, null, null);
        int i;
        for (i = 0; i < t.es.size(); i++) {
            final int cmp = e.key.compareTo(t.es.get(i).key);
            // e.key < t.es.get(i).key
            if (cmp < 0) {
                t.ns.set(i, insert(t.ns.get(i), e));
                return balance(t, i);
            }
            else if (cmp == 0) {
                t.es.set(i, e);
                return t;
            }
        }
        t.ns.set(i, insert(t.ns.get(i), e));
        return balance(t, i);
    }
*/

func (tree *Tree) insert(parent Node, newItem Item) (Node, error) {
	i := 0
	for i = 0; i < len(parent.Items()); i++ {
		item := parent.Items()[i]
		if newItem.Key() == item.Key() {
			parent.Items()[i] = newItem
			return parent, nil // TODO: return already exists?
		}
		if newItem.Key() < item.Key() {
			inserted, _ := tree.insert(parent.Children()[i], newItem)
			parent.Children()[i] = inserted
			return tree.balance(parent, i), nil
		}
	}
	inserted, _ := tree.insert(parent.Children()[i], newItem)
	parent.Children()[i] = inserted
	return tree.balance(parent, i), nil
}

/*
   // 挿入時のアクティブなノードと反応して木を変形する
    // アクティブでなければ何もしない
    private Node balance(Node t, int i) {
        Node ni = t.ns.get(i);
        if (!activeI(ni)) return t;
        // 以下、ni はアクティブ。つまり２分岐
        t.es.add(i, ni.es.get(0));
        t.ns.set(i, ni.ns.get(1));
        t.ns.add(i, ni.ns.get(0));
        return t.es.size() < m ? t : split(t);
    }
*/
func (tree *Tree) balance(node Node, position int) Node {
	return node
}

/*
 private abstract class Node { // ノードの型(抽象型)
        List<Pair> es = new ArrayList<Pair>(m);     // 要素のリスト
        List<Node> ns = new ArrayList<Node>(m + 1); // 枝のリスト

        // 挿入時のアクティブなノードを通常のノードへ変換する
        abstract Node deactivate();

        // 枝が１本の余分なノードを切り詰める
        Node trim() { return ns.size() == 1 ? ns.get(0) : this; }
    }
*/
