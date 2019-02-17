#!/usr/bin/env python
# -*- coding: utf-8 -*-
# http://d.hatena.ne.jp/naoya/20090412/btree
# t: min degree

class BTree:
    def __init__(self, t=2):
        self.t = t
        self.root = BTree.Node(t)
        self.root.is_leaf = True

    def insert(self, k):
        r = self.root
        if len(r) == 2 * self.t - 1:
            s = BTree.Node(self.t)
            s.children.append(r)
            s.split_child(0, r)
            s.insert_nonfull(k)
            self.root = s
        else:
            r.insert_nonfull(k)

    def delete(self, k):
        r = self.root
        if r.search(k) is None:
            return
        r.delete(k)
        if len(r) == 0:
            self.root = r.children[0]

    def search(self, k):
        return self.root.search(k)

    def show(self):
        self.root.show(1)

    class Node:
        def __init__(self, t):
            self.t = t
            self.keys = []
            self.children = []
            self.is_leaf = False

        def __len__(self):
            return len(self.keys)

        def search(self, k):
            i = 0
            while (i < len(self) and self.keys[i] < k):
                i += 1

            if i < len(self) and self.keys[i] == k:
                return (self, i)

            if self.is_leaf:
                return
            else:
                return self.children[i].search(k)

        def split_child(self, i, y):
            t = self.t
            z = BTree.Node(t)

            z.is_leaf = y.is_leaf
            z.keys = y.keys[t:]
            if not y.is_leaf:
                z.children = y.children[t:]

            self.children.insert(i + 1, z)
            self.keys.insert(i, y.keys[t - 1])

            y.keys = y.keys[0:t - 1]
            y.children = y.children[0:t]

        def locate_subtree(self, k):
            i = 0
            while (i < len(self)):
                if k < self.keys[i]:
                    return i
                i += 1
            return i

        def insert_nonfull(self, k):
            if self.is_leaf:
                i = 0
                for i in xrange(len(self)):
                    if k < self.keys[i]:
                        self.keys.insert(i, k)
                        return self
                self.keys.append(k)
            else:
                i = self.locate_subtree(k)
                c = self.children[i]
                if (len(c) == 2 * self.t - 1):
                    self.split_child(i, c)
                    if k > self.keys[i]:
                        c = self.children[i + 1]
                c.insert_nonfull(k)

        def show(self, pad):
            print "%s%s:[children=%d,leaf=%s]" % ('-' * pad, self.keys, len(self.children), self.is_leaf)
            if self.is_leaf:
                return
            else:
                for c in self.children:
                    c.show(pad + 1)

        ## 要検証
        def delete(self, k):
            t = self.t
            flag = False

            for i, x in enumerate(self.keys):
                if k == x:
                    flag = True  # 現在着目中の節に k が見つかった
                    if self.is_leaf:
                        ## 1. キーが x に存在し、x が葉
                        self.keys.remove(k)
                    else:
                        ## 2. キーが x に存在し、x が内部節点
                        if i > 0 and len(self.children[i]) > t - 1:
                            ## 2a
                            self.keys[i] = self.children[i].keys.pop()
                        elif len(self.children[i + 1]) > t - 1:
                            ## 2b
                            self.keys[i] = self.children[i + 1].keys.pop(0)
                        else:
                            ## 2c
                            self.children[i].keys += [self.keys.pop(0)] + self.children[i + 1].keys
                            del (self.children[i + 1])
                            self.children[i].delete(k)

            if not flag:
                ## 3. 現在着目中の節 x に k がなかった => k がある部分木の根cを決める
                i = self.locate_subtree(k)
                c = self.children[i]
                if len(c) > t - 1:
                    c.delete(k)
                else:
                    ## 3a. c自身はt-1個しかキーを持たないが、兄弟がt以上持ってる
                    if i > 0 and len(self.children[i - 1]) > t - 1:
                        flag = True
                        c.keys.insert(0, self.keys[i - 1])  ## cと左兄弟を分離するxのキーをcに
                        self.keys[i - 1] = self.children[i - 1].keys.pop()  ## 左兄弟の最後のキーをxに
                    elif len(self) > i and len(self.children[i + 1]) > t - 1:
                        flag = True
                        c.keys.append(self.keys[i])  ## c と右兄弟を分離するxのキーをcに
                        self.keys[i] = self.children[i + 1].keys.pop(0)  ## 右兄弟の最初のキーをxに

                    if flag:
                        c.delete(k)
                    else:
                        ## 3b. 左右兄弟もt - 1しか持ってなかった
                        if i > 0:
                            l = self.children[i - 1]
                            c.keys = l.keys + [self.keys.pop(i - 1)] + c.keys
                            del (self.children[i - 1])
                        else:
                            r = self.children[i + 1]
                            c.keys += [self.keys.pop(i)] + r.keys
                            del (self.children[i + 1])
                        c.delete(k)


import sys
import random

## order 5 の BTree の例
tree = BTree(2)
list = range(4)
#random.shuffle(list)

for x in list:
    tree.insert(x)
tree.show()

if len(sys.argv) > 1:
    result = tree.search(int(sys.argv[1]))
    if result:
        print "HIT: %d" % result[0].keys[result[1]]
