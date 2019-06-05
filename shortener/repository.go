package shortener

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

type Data struct {
	url        string
	dateOfSave string
}

type RepositoryConfig struct {
	deduplication bool
}

type Repository struct {
	RepositoryConfig
	idCount int
	storage map[int]Data
	hashMap map[string]int
}

func NewRepository(newConfig RepositoryConfig) (*Repository, error) {
	var err error
	newRepository := Repository{}
	newRepository.storage = make(map[int]Data)
	newRepository.hashMap = make(map[string]int)
	newRepository.idCount = 0
	newRepository.RepositoryConfig = newConfig

	return &newRepository, err
}

func (r *Repository) Create(data Data) (int, error) {
	var id int
	var err error
	id = r.idCount

	data.dateOfSave = fmt.Sprint(time.Now().Format("15:04:05 02.01.2006"))
	r.storage[id] = data

	if r.deduplication {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(data.url)))
		r.hashMap[hash] = id
	}

	r.idCount = r.idCount + 1
	return id, err

}

func (r *Repository) Update(id int, data Data) error {
	var err error
	if val, ok := r.storage[id]; ok {
		if val.url == data.url {
			r.storage[id] = data
		} else {
			newMd5Sum := fmt.Sprintf("%x", md5.Sum([]byte(data.url)))
			oldMd5Sum := fmt.Sprintf("%x", md5.Sum([]byte(val.url)))
			delete(r.hashMap, oldMd5Sum)
			r.hashMap[newMd5Sum] = id
			r.storage[id] = data
		}
	} else {
		err = errors.New("Нет такой записи")
	}

	return err
}

func (r *Repository) Delete(ID int) error {
	var err error
	entry, err := r.GetByID(ID)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(entry.url)))
	delete(r.hashMap, hash)
	delete(r.storage, ID)
	return err
}

func (r *Repository) GetByID(id int) (Data, error) {
	var err error
	result := r.storage[id]
	return result, err
}

func (r *Repository) GetByHash(hash string) (Data, error) {
	var err error
	id := r.hashMap[hash]
	result := r.storage[id]
	return result, err
}

func (r *Repository) SetDeduplication(deduplicationStatus bool) error {
	var err error
	r.RepositoryConfig.deduplication = deduplicationStatus
	return err
}

func (r *Repository) GetDeduplication() (bool, error) {
	return r.RepositoryConfig.deduplication, nil
}
