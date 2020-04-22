package core

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext struct {
	MainRouter *mux.Router
	APIRouter  *mux.Router
	DB         *mongo.Database
}
