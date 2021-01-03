package database

//DBer Operaciones basicas para la base de datos
type DBer interface {
	InsertOne(interface{}) (string, error)
	FindAll([]interface{}) error
	FindByID(string, interface{}) error
	DeleteOne(string) error
	Udate(string, interface{}) (int64, error)
	Purge() error
}
