package shortener

//Shortener - interface of shortener
type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(url string) (string, error)
}

type IRepository interface {
	Create(data Data) (int, error)
	Update(id int, data Data) error
	Delete(id int) error
	GetByID(id int) (Data, error)
	GetByHash(hash string) (Data, error)
}

type IKeyGenerator interface {
	GenerateKey(data int) (string, error)
	ResolvKey(key string) (int, error)
}
