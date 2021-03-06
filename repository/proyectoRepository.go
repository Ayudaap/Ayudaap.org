package repository

import (
	"log"
	"time"

	"Ayudaap.org/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProyectosRepository Repositorio de base de datos
type ProyectosRepository struct {
	DbRepo MongoRepository
}

//ProyectosCollection Nombre de la tabla de proyectos
const ProyectosCollection string = "proyectos"

//InsertProyecto Inserta una nueva instancia de Proyecto
func (o *ProyectosRepository) InsertProyecto(proyecto models.Proyecto) string {
	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	proyecto.Auditoria = models.Auditoria{
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	resultado, err := col.InsertOne(ctx, proyecto)
	if err != nil {
		log.Fatal(err)
	}

	ObjectID, _ := resultado.InsertedID.(primitive.ObjectID)

	var result string = ObjectID.Hex()
	return result
}

//GetAllProyectos Obtiene todas los proyectos
func (o *ProyectosRepository) GetAllProyectos() []models.Proyecto {
	var proyectos []models.Proyecto

	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	datos, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next(ctx) {
		var Proyecto models.Proyecto
		err := datos.Decode(&Proyecto)
		if err != nil {
			log.Fatal(err)
		}
		proyectos = append(proyectos, Proyecto)
	}
	return proyectos
}

//GetProyectoByID Obtiene una Proyecto por Id
func (o *ProyectosRepository) GetProyectoByID(id string) *models.Proyecto {
	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	Oid, _ := primitive.ObjectIDFromHex(id)

	var Proyecto *models.Proyecto
	err := col.FindOne(ctx, bson.M{"_Id": Oid}).Decode(&Proyecto)
	if err != nil {
		return nil
	}

	return Proyecto
}

//DeleteProyecto Elimina una Proyecto
func (o *ProyectosRepository) DeleteProyecto(id string) (int, error) {
	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": oID}

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount <= 0 {
		return 0, nil
	}

	return int(result.DeletedCount), nil
}

//UpdateProyecto Actualiza una Proyecto retornando el total de elementos que se modificaron
func (o *ProyectosRepository) UpdateProyecto(proyecto *models.Proyecto) (int64, error) {

	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	filter := bson.M{"_id": proyecto.ID}
	update := bson.M{"$set": proyecto}

	proyecto.Auditoria.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//PurgarProyectos Purgarproyectos borra toda la collecion
func (o *ProyectosRepository) PurgarProyectos() error {

	col, ctx, cancel := o.DbRepo.GetCollection(ProyectosCollection)
	defer cancel()

	err := col.Drop(ctx)

	return err
}
