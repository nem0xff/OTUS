package doublylinkedlist

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	var testList itemList
	for i := 0; i < 10; i++ {

		testList.pushFront(i)
	}
}
