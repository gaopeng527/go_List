// list project list.go
package list

type List interface {
	Add(e interface{}) bool
	RemoveElement(e interface{})
	RemoveIndex(index int)
	Clear()
	Contains(e interface{}) bool
	Len() int
	Elements() []interface{}
	String() string
}

// 将列表other添加到列表one中
func AddList(one List, other List) {
	if one == nil || other == nil || other.Len() == 0 {
		return
	}
	for _, v := range other.Elements() {
		one.Add(v)
	}
}

// 判断列表 one 是否是列表 other 的超集
func IsSuperlist(one List, other List) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen <= otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}

// 生成列表 one 对列表 other 的差集
func Difference(one List, other List) List {
	if one == nil {
		return nil
	}
	differencedList := NewArrayList()
	if other == nil || other.Len() == 0 {
		AddList(differencedList, one)
		return differencedList
	}
	for _, v := range one.Elements() {
		if !other.Contains(v) {
			differencedList.Add(v)
		}
	}
	return differencedList
}

// 判断给定value是否为列表
func IsList(value interface{}) bool {
	if _, ok := value.(List); ok {
		return true
	}
	return false
}
