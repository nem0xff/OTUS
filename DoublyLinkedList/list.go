package doublylinkedlist

type itemList struct {
	firstItem *item
	lastItem  *item
	lenght    int
}

func (l *itemList) pushFront(val interface{}) {

	var newItem item

	if l.lenght == 0 && l.firstItem == nil && l.lastItem == nil { //Обрабатываем самый первый элемент
		newItem.data = val
		l.firstItem = &newItem
		l.lastItem = &newItem
		l.lenght = l.lenght + 1
		return
	}

	newItem.ptrNext = l.firstItem
	l.firstItem = &newItem

	newItem.ptrNext.ptrPrev = &newItem

	l.lenght = l.lenght + 1
	newItem.data = val

}

func (l *itemList) pushBack(val interface{}) {

	var newItem item

	if l.lenght == 0 && l.firstItem == nil && l.lastItem == nil { // Это первый элемент
		l.pushFront(val)
	}

	newItem.ptrPrev = l.lastItem
	l.lastItem = &newItem

	newItem.ptrPrev.ptrNext = &newItem

	l.lenght = l.lenght + 1
	newItem.data = val

}

func (l *itemList) Remove(i *item) {
	i.remove()
	l.lenght = l.lenght - 1
}

func (l *itemList) First() *item { return l.firstItem }

func (l *itemList) Last() *item { return l.lastItem }

func (l *itemList) Len() int { return l.lenght }
