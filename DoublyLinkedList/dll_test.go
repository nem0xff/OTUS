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

func TestPushFrontAndGetLastItem(t *testing.T) {

	for _, val := range myItems {
		testList.PushFront(val)
	}

	item := testList.Last()

	for _, val := range myItems {
		if item.Value() != val {
			t.Error("Нарушен порядок элементов или отсутствуют значения в Списке")
		}
		item = item.Prev()
	}
}

func TestLenght(t *testing.T) {

	if testList.Len() != 5 {
		t.Error("Длина не соответствует количеству добавленных элементов")
	}
}

func TestRemoveLastItem(t *testing.T) {

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

func TestRemoveFirstItem(t *testing.T) {

	item := testList.First()
	oldLen := testList.Len()
	item.Remove()
	if oldLen == testList.Len() {
		t.Error("После удаления не изменилась длина Списка")
	}

	item = testList.Last()
	for _, val := range myItems[1:4] {
		if item.Value() != val {
			t.Error("После удаления нарушен порядок элементов в Списке")
		}
		item = item.Prev()
	}
}

func TestRemoveMiddleItem(t *testing.T) {
	item := testList.Last().Prev()
	oldLen := testList.Len()
	item.Remove()
	if oldLen == testList.Len() {
		t.Error("После удаления не изменилась длина Списка")
	}

	item = testList.Last()
	testItems := append(myItems[1:2], myItems[3]) // делаем слайс из второго и четвертого элементов.
	for _, val := range testItems {
		if item.Value() != val {
			t.Error("После удаления нарушен порядок элементов в Списке")
		}
		item = item.Prev()
	}

}
