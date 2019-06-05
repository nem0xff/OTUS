package shortener

import (
	"errors"
	"testing"
)

var rep *Repository

var testData Data
var updateData Data
var testID int

func init() {
	var config = RepositoryConfig{
		deduplication: true,
	}
	rep, _ = NewRepository(config)

	//Заполнение тестовых структур
	newUrl := makePseudoURL()
	testData = Data{
		url:        newUrl,
		dateOfSave: "",
	}

	updateUrl := makePseudoURL()
	updateData = Data{
		url:        updateUrl,
		dateOfSave: "",
	}

}

func TestRepositoryCreate(t *testing.T) {
	var err error
	testID, err = rep.Create(testData)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Добавлена запись c ID = %v", testID)
	}
}

func TestRepositoryUpdate(t *testing.T) {
	var err error

	if compareEntryId(testID, updateData) {
		t.Error(errors.New("Добавляемая запись уже существует"))
		t.FailNow()
	}

	err = rep.Update(testID, updateData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Logf("Обновлена запись c ID = %v", testID)
	}

	if !compareEntryId(testID, updateData) {
		t.Error(errors.New("Не совпадает тестовая и текущая запись в хранилище"))
		t.FailNow()
	}

}

func compareEntryId(id int, entryData Data) bool {
	curretEntry := rep.storage[id]
	if curretEntry.url == entryData.url {
		return true
	} else {
		return false
	}

}
