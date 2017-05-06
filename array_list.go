// list project array_list.go
package list

import (
	"bytes"
	"fmt"
)

type ArrayList struct {
	a []interface{}
}

// 创建和初始化ArrayList的方法
func NewArrayList() *ArrayList {
	return &ArrayList{a: make([]interface{}, 0)}
}

// 向ArrayList中添加元素的方法
func (list *ArrayList) Add(e interface{}) bool {
	list.a = append(list.a, e)
	return true
}

// 删除ArrayList中指定的元素
func (list *ArrayList) RemoveElement(e interface{}) {
	j := -1
	for i, v := range list.a {
		j = i
		if v == e {
			break
		}
	}
	list.a = append(list.a[:j], list.a[j+1:])
}

// 删除ArrayList中指定的元素
func (list *ArrayList) RemoveIndex(index int) {
	list.a = append(list.a[:index], list.a[index+1:])
}

// 清除ArrayList中的所有元素
func (list *ArrayList) Clear() {
	list.a = make([]interface{}, 0)
}

// 判断ArrayList是否包含指定元素
func (list *ArrayList) Contains(e interface{}) bool {
	for _, v := range list.a {
		if v == e {
			return true
		}
	}
	return false
}

// 获取ArrayList中元素值数量
func (list *ArrayList) Len() int {
	return len(list.a)
}

// 生成ArrayList的一个快照
func (list *ArrayList) Elements() []interface{} {
	initialLen := len(list.a)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, v := range list.a {
		if actualLen < initialLen {
			snapshot[actualLen] = v
		} else {
			snapshot = append(snapshot, v)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// 获取ArrayList自身字符串表示形式
func (list *ArrayList) String() string {
	var buf bytes.Buffer
	buf.WriteString("List{")
	first := true
	for _, v := range list.a {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", v))
	}
	buf.WriteString("}")
	return buf.String()
}
