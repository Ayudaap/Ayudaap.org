package repository

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN objeto de conexión de la BD */
var MongoCN *mongo.Client
var once sync.Once

// Nombre de la base de datos
const DataBase string = "AyudaapDb"

// Inicializa una nueva instancia
func init() {
	once.Do(conectarBD)
	//return MongoCN
}

// ConectarDB inicia una conexión de hacia la BD
func conectarBD() {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASSWORD")

	//var cadenaConexion = fmt.Printf("mongodb+srv://%s:%s@%h:%p",user,pass,host,port)

	//cadenaConexion := fmt.Sprintf("mongodb://%s:%s", host, port)
	cadenaConexion := fmt.Sprintf("mongodb://localhost:27017")
	clienteOpts := options.Client().ApplyURI(cadenaConexion)
	cliente, err := mongo.Connect(context.TODO(), clienteOpts)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = cliente.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Println("Conexión Exitosa con la BD")
	MongoCN = cliente
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

// Obtienene la colecion y el contexto de trabaj
// `dataBase` Nombre de la base de datos
// `collection` Nombre de la conexion
func GetCollection(dataBase string, collection string) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	db := MongoCN.Database(dataBase)
	col := db.Collection(collection)

	return col, ctx, cancel
}
