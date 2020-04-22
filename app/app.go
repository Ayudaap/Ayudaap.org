//https://github.com/antonybudianto/go-rest-starter/blob/master/starter/app/app.go

import (

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"../core"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"fmt"
	"log"
	"net/http"
	"os"

	
)

// App level struct containing its dependencies
type App struct {
	AppCtx *core.AppContext
}

const dbDriver = "mongodb"
const dbName = "ayudaap_db"

//TODO: Cambiar en productivo
//var dbUsername = os.Getenv("DB_USERNAME") //Remplazar en productivo
//var dbPassword = os.Getenv("DB_PASSWORD") //Remplazar en productivo

// Initialize App dependencies
func (a *App) Initialize() {

	a.AppCtx = &core.AppContext{}

	//TODO: Cambiar cadena de conexion
	//connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUsername, dbPassword, dbHost,dbPort)
	connectionString := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(connectionString)
	
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
        log.Fatal(err)
	}
	
	defer client.Disconnect(ctx)
	a.AppCtx.DB = client
	
	a.AppCtx.MainRouter = mux.NewRouter()
	a.AppCtx.APIRouter = a.AppCtx.MainRouter.PathPrefix("/api").Subrouter()
	a.initializeRoutes()
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.AppCtx.MainRouter))
}

func (a *App) initializeRoutes() {
	user.New(a.AppCtx.DB, a.AppCtx.APIRouter)
}