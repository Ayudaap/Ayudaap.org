package database

//DBer Operaciones basicas para la base de datos
type DBer interface {
	InsertOne(registro interface{}) (string, error)
	FindAll() ([]interface{}, error)
	FindByID(ID string) (interface{}, error)
	DeleteOne(ID string) error
	Udate(ID string, registro interface{}) (int64, error)
	Purge() error
}
