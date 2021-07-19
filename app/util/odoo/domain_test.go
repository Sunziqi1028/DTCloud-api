package odoo

import (
	"fmt"
	"testing"
)

func TestDomainTree(t *testing.T) {
	buf := []string{
		`[('id','=',1),('id','=',1),('id','=',1)]`,
		`['|','!','&',('id','=',1),('active','=',False),('users_id.state','=','未认领'),('id','=',1),('users_id.state','=','未认领')]`,
		`['|','&',('id','=',1),('active','=',False),('users_id.state','=','未认领'),'|','&',('id','=',1),('active','=',False),('users_id.state','=','未认领')]`,
		`[('id','=',1),('id','=',1),'!','|','&',('id','=',1),('active','=',False),('users_id.state','=','未认领'),'|','&',('id','=',1),('active','=',False),('users_id.state','=','未认领')]`,
	}
	dlist := DomainFormStringToList(buf[1])
	var treeNode *TreeNode
	treeNode = DomainTree(treeNode, dlist)
	treeNode.Traverse()

	fmt.Print(TreeBuff)

	//fmt.Println(treeNode)
}
