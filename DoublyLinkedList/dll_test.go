package doublylinkedlist

import (
	"testing"
)

var testList itemList

var testData []interface{}

type myItem struct {
	value string
}

var myItems = []myItem{
	{"first"},
	{"two"},
	{"three"},
	{"four"},
	{"five"},
}

func init() {
	//Заполняем тестовыми данными
	for _, val := range myItems {
		testList.pushFront(val)
	}

}

func TestLenght(t *testing.T) {
	if testList.Len() != 5 {
		t.Error("Длина не соответствует количеству добавленных элементов")
	}

}
func TestGetLastItem(t *testing.T) {

	item := testList.Last()

	for _, val := range myItems {
		if item.Value() != val {
			t.Error("Нарушен порядок элементов или отсутствуют значения в Списке")
		}
		item = item.Prev()
	}
}
func TestRemove(t *testing.T) {

	item := testList.Last()
	oldLen := testList.Len()
	item.Remove()
	if oldLen == testList.Len() {
		t.Error("После удаления не изменилась длина Списка")
	}

	item = testList.Last()
	for _, val := range myItems[1:] {
		if item.Value() != val {
			t.Error("После удаления нарушен порядок элементов в Списке")
		}
		item = item.Prev()
	}

}
