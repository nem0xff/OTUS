package doublylinkedlist

type itemList struct {
	firstItem *item
	lastItem  *item
	lenght    int
}

func (l *itemList) pushFront(val interface{}) {

	var newItem item

	if l.lenght == 0 && l.firstItem == nil && l.lastItem == nil {
		//Обрабатываем самый первый элемент
		l.lastItem = &newItem
	} else {
		// Все последующие элементы
		newItem.ptrNext = l.firstItem
		newItem.ptrNext.ptrPrev = &newItem
	}

	l.firstItem = &newItem
	newItem.data = val
	newItem.ptrList = l
	l.lenght = l.lenght + 1

}

func (l *itemList) pushBack(val interface{}) {

	var newItem item

	if l.lenght == 0 && l.firstItem == nil && l.lastItem == nil { // Это первый элемент можно использовать pushFront
		l.pushFront(val)
	}

	newItem.ptrPrev = l.lastItem
	l.lastItem = &newItem

	newItem.ptrPrev.ptrNext = &newItem

	l.lenght = l.lenght + 1
	newItem.data = val
	newItem.ptrList = l
}

func (l *itemList) First() *item { return l.firstItem }

func (l *itemList) Last() *item { return l.lastItem }

func (l *itemList) Len() int { return l.lenght }
