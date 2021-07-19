package curd

import (
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/service/curd"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

type PageData struct {
}

func (this PageData) PulickPage(context *gin.Context) {
	uid := int(context.GetFloat64(constant.ValidatorPrefix + "uid"))
	partnerId := int(context.GetFloat64(constant.ValidatorPrefix + "partner_id"))
	table := strings.Replace(context.GetString(constant.ValidatorPrefix+"model"), ".", "_", -1)

	domain := context.GetString(constant.ValidatorPrefix + "domain")

	search := context.GetString(constant.ValidatorPrefix + "search")
	searchfields := context.GetString(constant.ValidatorPrefix + "search_fields")

	fields := context.GetString(constant.ValidatorPrefix + "fields")
	offset := int(context.GetFloat64(constant.ValidatorPrefix + "page"))
	limit := int(context.GetFloat64(constant.ValidatorPrefix + "limit"))
	order := context.GetString(constant.ValidatorPrefix + "order")

	fmt.Print(uid, partnerId, search, searchfields, fields, offset, limit, order, table)

	//type Domain struct{
	//	filed1 string
	//	sign string
	//	value interface{}
	//}
	//a1 := []string{}
	//json.Unmarshal([]byte(domain),a1)
	//
	//a := "[('id','=',1),('id','=',1),('id','=',1),('id','=',1)]"
	//
	//b := []int{1,2,3,4,5}
	//fmt.Println(a)
	//fmt.Println(a[0])
	//fmt.Println(a[1])
	//fmt.Println(b)
	////a11 := &Domain{}
	////fmt.Println(domain)
	////json.Unmarshal([]byte(domain),a11)
	//FindAllSubmatch
	//k:= regexp.MustCompile(`/\\((.+?)\\)/g`).FindAllStringSubmatch(domain, -1)
	//
	//yhCount := 0
	//yCount = strings.Count(domain,"|")
	//hCount = strings.Count(domain,"&")

	//domain = `['|','&',('id','=',1),('active','=',False),('users_id.state','=','未认领')]`
	//var orIndex, andIndex, fIndex, kIndex int

	datalist := domainList(domain)

	fmt.Println(datalist)

	fmt.Println("=========Sfenxi")
	var treeNode *TreeNode
	//var newNode *TreeNode
	for i := 0; i < len(datalist); i++ {
		if treeNode == nil {
			bl := strings.Contains(datalist[0], "(")
			if bl {
				newNode := new(TreeNode)
				newNode.value = "&"
				newNode.left = new(TreeNode)
				newNode.left.value = datalist[0]
				treeNode = newNode
			} else {
				treeNode = new(TreeNode)
				treeNode.value = datalist[i]
			}
			continue
		}
		ret := InsertNode(treeNode, datalist[i])
		if ret == nil {
			newNode := new(TreeNode)
			newNode.value = "&"
			newNode.left = treeNode
			newNode.right = new(TreeNode)
			newNode.right.value = datalist[i]
			treeNode = newNode
		}
		fmt.Println(treeNode)
		//newNode := findEmptyNode2(treeNode,datalist[i])

		//switch(newNode.value){
		//case "|","&":
		//	if newNode.left == nil{
		//		t := new(TreeNode)
		//		t.value = datalist[i]
		//		newNode.left = t
		//		continue
		//	}
		//
		//	if newNode.right == nil{
		//		t := new(TreeNode)
		//		t.value = datalist[i]
		//		newNode.right = t
		//		continue
		//	}
		//case "!":
		//	if newNode.left == nil{
		//		t := new(TreeNode)
		//		t.value = datalist[i]
		//		newNode.left = t
		//		continue
		//	}
		//}
		//
		////var newNode *TreeNode
		//if strings.Contains(" | ! &",datalist[i]){
		//	newNode = new(TreeNode)
		//	newNode.value = datalist[i]
		//}
		//t:= findEmptyNode2(treeNode)
		//
		//if t== nil{
		//	treeNode = newNode
		//}

	}

	//sql:=fenxi2(treeNode,0,datalist)
	fmt.Println("=========Efenxi")
	//fmt.Println(sql)

	fmt.Println("=========", datalist)

	//t := time.Now()
	//fmt.Println(t)
	//ret := fn(50)
	//fmt.Println(ret)
	//fmt.Println(t.Sub(time.Now()).Seconds())
	//for _, match := range regexp.MustCompile("\\((.+?)\\)").FindAllString(domain, -1) {
	//	fmt.Printf("%v\n", match)
	//}

	//for k,v := range domain{
	//	fmt.Println(k,v)
	//	a1 := v.(Domain)
	//	//strings.Split(domain)
	//}

	result := curd.CreateUserCurdFactory().PublicPage(uid, partnerId, domain, search, searchfields, fields, offset, limit, order, table)
	if result != nil {
		response.Success(context, constant.CurdStatusOkMsg, result)
		return
	}
	response.Fail(context, constant.CurdSelectFailCode, constant.CurdSelectFailMsg, "")
	return

}

func fn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fn(n-1) + fn(n-2)
	}
}

type TreeNode struct {
	Value       string
	Left, Right *TreeNode
}

func fenxi(index int, datalist map[int]string) string {
	var (
		str   string
		parm1 string
		parm2 string
	)
	if index >= len(datalist) {
		return ""
	}

	if strings.Contains(" | ! &", datalist[index+1]) {
		parm1 = fenxi(index+1, datalist)
	} else {
		parm1 = datalist[index+1]
	}

	if strings.Contains(" | &", datalist[index+2]) {
		parm2 = fenxi(index+2, datalist)
	} else {
		parm2 = datalist[index+2]
	}

	switch datalist[index] {
	case "|":
		str = fmt.Sprintf(" %s or %s ", parm1, parm2)
	case "!":
		str = fmt.Sprintf(" not %s ", parm1)
	case "&":
		str = fmt.Sprintf(" %s and %s ", parm1, parm2)
	}

	fmt.Printf("%d :%s", index, str)
	return str
}

//func add(t *tree, value int) *tree {
//if t == nil { //第一层（中间层）或者树的末端
//	t = new(tree) //new(tree) 返回的是一个*tree 是一个地址
//	t.value = value
//	return t
//}
func add(t *TreeNode, data string) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		if strings.Contains(" | ! &", data) {
			t.value = data
		} else {
			t.value = "&"
		}
	}
	//
	//
	//switch(datalist[i]){
	//case "|":
	//	treeNode.left = datalist[i+1]
	//	str = fmt.Sprintf(" %s or %s ",parm1,parm2 )
	//case "!":
	//	str = fmt.Sprintf(" not %s ",parm1)
	//case "&":
	//	str = fmt.Sprintf(" %s and %s ",parm1 ,parm2 )
	//}
	//
	//if value < t.value {
	//	t.left = add(t.left, value) //比树根小，往左边下移
	//} else {
	//	t.right = add(t.right, value) //比树根大，往右边下移
	//}
	return t
}
func fenxi2(treeNode TreeNode, index int, datalist map[int]string) string {

	for i := 0; i < len(datalist); i++ {
		//if strings.Contains(" | ! &",datalist[i]){
		//	treeNode.value = datalist[i]
		//}else{
		//	treeNode.left = datalist[i+1]
		//	switch(datalist[i]){
		//	case "|":
		//		treeNode.left = datalist[i+1]
		//		str = fmt.Sprintf(" %s or %s ",parm1,parm2 )
		//	case "!":
		//		str = fmt.Sprintf(" not %s ",parm1)
		//	case "&":
		//		str = fmt.Sprintf(" %s and %s ",parm1 ,parm2 )
		//	}
		//
		//}
	}
	var (
		str   string
		parm1 string
		parm2 string
	)

	if index >= len(datalist) {
		return ""
	}

	if strings.Contains(" | ! &", datalist[index+1]) {
		parm1 = fenxi(index+1, datalist)
	} else {
		parm1 = datalist[index+1]
	}

	if strings.Contains(" | &", datalist[index+2]) {
		parm2 = fenxi(index+2, datalist)
	} else {
		parm2 = datalist[index+2]
	}

	switch datalist[index] {
	case "|":
		str = fmt.Sprintf(" %s or %s ", parm1, parm2)
	case "!":
		str = fmt.Sprintf(" not %s ", parm1)
	case "&":
		str = fmt.Sprintf(" %s and %s ", parm1, parm2)
	}

	fmt.Printf("%d :%s", index, str)
	return str
}

func domainList(domain string) map[int]string {
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

func findEmptyNode(t *TreeNode) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		return t
	}
	switch t.value {
	case "|", "&":
		if t.left == nil || t.right == nil {
			return t
		} else {
			return findEmptyNode(t)
		}
	case "!":
		if t.left == nil {
			return t
		} else {
			return findEmptyNode(t)
		}
	}
	return nil
}
func findEmptyNode2(t *TreeNode, parm string) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		t.value = parm
		return t
	}

	switch t.value {
	case "|", "&":
		if t.left == nil {
			return t
		} else {
			if !strings.Contains("(", t.left.value) {
				ret := findEmptyNode2(t.left, parm)
				if ret != nil {
					return ret
				}
			}
		}

		if t.right == nil {
			return t
		} else {
			if !strings.Contains("(", t.right.value) {
				ret := findEmptyNode2(t.right, parm)
				if ret != nil {
					return ret
				}
			}
		}
	case "!":
		if t.left == nil {
			return t
		} else {
			if !strings.Contains("(", t.left.value) {
				ret := findEmptyNode2(t.left, parm)
				if ret != nil {
					return ret
				}
			}
		}
	}

	//todo
	return nil
}

func InsertNode(t *TreeNode, parm string) *TreeNode {
	if t.value == "" {
		t.value = parm
		return t
	}

	switch t.value {
	case "|", "&":
		if t.left == nil {
			newNode := new(TreeNode)
			newNode.value = parm
			t.left = newNode
			return t
		} else {
			if !strings.Contains(t.left.value, "(") {
				ret := InsertNode(t.left, parm)
				if ret != nil {
					return ret
				}
			}
		}

		if t.right == nil {
			newNode := new(TreeNode)
			newNode.value = parm
			t.right = newNode
			return t
		} else {
			if !strings.Contains(t.right.value, "(") {
				ret := InsertNode(t.right, parm)
				if ret != nil {
					return ret
				}
			}
		}
	case "!":
		if t.left == nil {
			newNode := new(TreeNode)
			newNode.value = parm
			t.left = newNode
			return t
		} else {
			if !strings.Contains(t.left.value, "(") {
				ret := InsertNode(t.left, parm)
				if ret != nil {
					return ret
				}
			}
		}
	}

	//todo
	return nil
}
