package doublylinkedlist

type item struct {
	ptrNext *item
	ptrPrev *item
	data    interface{}
}

func (i *item) remove() {
	i.ptrNext.ptrPrev = i.ptrPrev
	i.ptrPrev.ptrNext = i.ptrNext
}

func (i *item) Value() interface{} { return i.data }

func (i *item) Next() *item { return i.ptrNext }

func (i *item) Prev() *item { return i.ptrPrev }
