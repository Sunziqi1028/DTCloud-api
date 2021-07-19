package odoo

import (
	"regexp"
	"strings"
)

type TreeNode struct {
	Value       string
	Left, Right *TreeNode
}

func DomainTree(t *TreeNode, arg map[int]string) *TreeNode {
	var treeNode *TreeNode

	for i := 0; i < len(arg); i++ {
		if treeNode == nil {
			bl := strings.Contains(arg[0], "(")
			if bl {
				newNode := new(TreeNode)
				newNode.Value = "&"
				newNode.Left = new(TreeNode)
				newNode.Left.Value = arg[0]
				treeNode = newNode
			} else {
				treeNode = new(TreeNode)
				treeNode.Value = arg[i]
			}
			continue
		}
		ret := InsertNode(treeNode, arg[i])
		if ret == nil {
			newNode := new(TreeNode)
			newNode.Value = "&"
			newNode.Left = treeNode
			newNode.Right = new(TreeNode)
			newNode.Right.Value = arg[i]
			treeNode = newNode
		}
	}

	return treeNode
}

func InsertNode(t *TreeNode, parm string) *TreeNode {
	if t.Value == "" {
		t.Value = parm
		return t
	}

	switch t.Value {
	case "|", "&":
		if t.Left == nil {
			newNode := new(TreeNode)
			newNode.Value = parm
			t.Left = newNode
			return t
		} else {
			if !strings.Contains(t.Left.Value, "(") {
				ret := InsertNode(t.Left, parm)
				if ret != nil {
					return ret
				}
			}
		}

		if t.Right == nil {
			newNode := new(TreeNode)
			newNode.Value = parm
			t.Right = newNode
			return t
		} else {
			if !strings.Contains(t.Right.Value, "(") {
				ret := InsertNode(t.Right, parm)
				if ret != nil {
					return ret
				}
			}
		}
	case "!":
		if t.Left == nil {
			newNode := new(TreeNode)
			newNode.Value = parm
			t.Left = newNode
			return t
		} else {
			if !strings.Contains(t.Left.Value, "(") {
				ret := InsertNode(t.Left, parm)
				if ret != nil {
					return ret
				}
			}
		}
	}

	//todo
	return nil
}

func DomainFormStringToList(domain string) map[int]string {
	list := make(map[int]string)

	for i := 0; i < len(domain); {
		var (
			min      int = 0
			s        int = 0
			orIndex  int = 0
			andIndex int = 0
			fIndex   int = 0
			kIndex   int = 0
		)

		listlen := len(list)
		if listlen > 0 {
			s = i + len(list[listlen-1])
		}

		orIndex = strings.Index(domain[s:], "|")
		if orIndex > 0 {
			if min > orIndex || min == 0 {
				min = orIndex

				list[listlen] = "|"
			}
		}
		andIndex = strings.Index(domain[s:], "&")
		if andIndex > 0 {
			if min >= andIndex || min == 0 {
				min = andIndex
				list[listlen] = "&"
			}
		}
		fIndex = strings.Index(domain[s:], "!")
		if fIndex > 0 {
			if min >= fIndex || min == 0 {
				min = fIndex
				list[listlen] = "!"
			}
		}
		kIndex = strings.Index(domain[s:], "(")
		if kIndex > 0 {
			if min >= kIndex || min == 0 {
				min = kIndex
				list[listlen] = regexp.MustCompile("\\((.+?)\\)").FindString(domain[s:])
			}
		}

		if min == 0 {
			break
		}

		i = s + min

	}
	return list
}
