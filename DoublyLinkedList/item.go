package doublylinkedlist

type item struct {
	ptrList *itemList
	ptrNext *item
	ptrPrev *item
	data    interface{}
}

func (i *item) Remove() {
	if i.ptrNext != nil {
		i.ptrNext.ptrPrev = i.ptrPrev
	} else {
		i.ptrList.lastItem = i.ptrPrev
	}

	if i.ptrPrev != nil {
		i.ptrPrev.ptrNext = i.ptrNext
	} else {
		i.ptrList.firstItem = i.ptrNext
	}

	i.ptrNext = nil
	i.ptrPrev = nil

	i.ptrList.lenght = i.ptrList.lenght - 1

}

func (i *item) Value() interface{} { return i.data }

func (i *item) Next() *item { return i.ptrNext }

func (i *item) Prev() *item { return i.ptrPrev }
