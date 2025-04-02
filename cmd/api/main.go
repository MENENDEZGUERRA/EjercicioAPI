package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Incident struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Reporter    string             `bson:"reporter" json:"reporter"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

var client *mongo.Client
var incidentCollection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://JFMG:contraseña123@jfmg.xzezjjy.mongodb.net/?retryWrites=true&w=majority&appName=JFMG")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB")

	incidentCollection = client.Database("incidentDB").Collection("incidents")

	// Configurar el router y las rutas
	router := mux.NewRouter()
	router.HandleFunc("/incidents", createIncident).Methods("POST")
	router.HandleFunc("/incidents", getIncidents).Methods("GET")
	router.HandleFunc("/incidents/{id}", getIncident).Methods("GET")
	router.HandleFunc("/incidents/{id}", updateIncident).Methods("PUT")
	router.HandleFunc("/incidents/{id}", deleteIncident).Methods("DELETE")

	// Iniciar el servidor en el puerto 8000
	fmt.Println("Servidor iniciado en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// createIncident crea un nuevo incidente.
func createIncident(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var incident Incident
	err := json.NewDecoder(r.Body).Decode(&incident)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	if incident.Reporter == "" {
		http.Error(w, "El campo 'reporter' es obligatorio", http.StatusBadRequest)
		return
	}
	if len(incident.Description) < 10 {
		http.Error(w, "La descripción debe tener al menos 10 caracteres", http.StatusBadRequest)
		return
	}

	incident.Status = "pendiente"
	incident.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := incidentCollection.InsertOne(ctx, incident)
	if err != nil {
		http.Error(w, "Error al crear el incidente", http.StatusInternalServerError)
		return
	}
	incident.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(incident)
}

// getIncidents obtiene la lista de todos los incidentes. -- GET
func getIncidents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var incidents []Incident
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := incidentCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error al obtener los incidentes", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var incident Incident
		cursor.Decode(&incident)
		incidents = append(incidents, incident)
	}
	json.NewEncoder(w).Encode(incidents)
}

// getIncident obtiene un incidente específico por ID. -- GET/{id}
func getIncident(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idParam := params["id"]

	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var incident Incident
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = incidentCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&incident)
	if err != nil {
		http.Error(w, "Incidente no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(incident)
}

// updateIncident actualiza únicamente el campo "status" de un incidente.
func updateIncident(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idParam := params["id"]

	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var updateData struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	if updateData.Status == "" {
		http.Error(w, "El campo 'status' es obligatorio", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"status": updateData.Status}}

	result, err := incidentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error al actualizar el incidente", http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "Incidente no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bson.M{"message": "Incidente actualizado exitosamente"})
}

// deleteIncident elimina un incidente por ID. -- DELETE
func deleteIncident(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idParam := params["id"]

	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := incidentCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Error al eliminar el incidente", http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Incidente no encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bson.M{"message": "Incidente eliminado exitosamente"})
}
