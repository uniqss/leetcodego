package main

import "fmt"

func main() {
	var s string
	var k int
	//s = "leet2code3"
	//k = 10
	//s = "ha22"
	//k = 5
	s = "a2345678999999999999999"
	k = 1
	ret := decodeAtIndex(s, k)
	fmt.Println(ret)
}

type strFragmentNode struct {
	parent *strFragmentNode
	str    string
	repeat int
	maxLen int
}

func newFragmentNode(parent *strFragmentNode, appendStr string) *strFragmentNode {
	ret := strFragmentNode{
		parent: parent,
		str:    appendStr,
		repeat: 1,
	}
	if parent != nil {
		ret.maxLen = parent.maxLen + len(appendStr)
	} else {
		ret.maxLen = len(appendStr)
	}
	return &ret
}

func appendTail(node *strFragmentNode, tail string) {
	node.str += tail
	node.maxLen += len(tail)
}

func appendRepeat(node *strFragmentNode, repeat int) {
	node.repeat *= repeat
	node.maxLen = node.maxLen * repeat
}

func decodeAtIndex(s string, k int) string {
	k--
	root := newFragmentNode(nil, "")
	var currNode *strFragmentNode = nil
	parsingStr := false
	i:= 0
	for ;i< len(s) && s[i] >= '0' || s[i] <= '9';i++ {}
	for ;i < len(s);i++ {
		si := s[i]
		if si >= '0' && si <= '9' {
			parsingStr = false
			appendRepeat(currNode, int(si - '0'))
		} else {
			if parsingStr {
				appendTail(currNode, s[i:i + 1])
			} else {
				if currNode == nil {
					currNode = newFragmentNode(root, s[i:i + 1])
				} else {
					currNode = newFragmentNode(currNode, s[i:i + 1])
				}
			}
			parsingStr = true
		}
		if currNode != nil && currNode.maxLen > k {
			node := currNode
			for node != nil {
				mod := k % (node.maxLen / node.repeat)
				if node.parent != nil && mod < node.parent.maxLen {
					node = node.parent
					k = mod
					continue
				} else {
					return node.str[mod - node.parent.maxLen:mod - node.parent.maxLen + 1]
				}
			}
		}
	}
	return ""
}