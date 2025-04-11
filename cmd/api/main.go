package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Incident struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Reporter    string             `bson:"reporter"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
}

var (
	client             *mongo.Client
	incidentCollection *mongo.Collection
	tmpl               *template.Template
)

func toLower(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}

func main() {
	// Conexión a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://JFMG:contraseña123@jfmg.xzezjjy.mongodb.net/?retryWrites=true&w=majority&appName=JFMG")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	incidentCollection = client.Database("incidentDB").Collection("incidents")

	// Configurar templates con funciones personalizadas
	tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"toLower": toLower,
	}).ParseGlob("templates/*.html"))

	// Configurar rutas
	router := mux.NewRouter()
	router.HandleFunc("/", listIncidents).Methods("GET")
	router.HandleFunc("/create", createIncident).Methods("POST")
	router.HandleFunc("/update/{id}", updateIncident).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteIncident).Methods("POST")

	// Archivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Servidor iniciado en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func listIncidents(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := incidentCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var incidents []Incident
	for cursor.Next(ctx) {
		var incident Incident
		cursor.Decode(&incident)
		incidents = append(incidents, incident)
	}

	data := struct {
		Title     string
		Incidents []Incident
	}{
		Title:     "Gestor de Incidentes",
		Incidents: incidents,
	}

	tmpl.ExecuteTemplate(w, "index.html", data)
}

func createIncident(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	incident := Incident{
		Reporter:    r.FormValue("reporter"),
		Description: r.FormValue("description"),
		Status:      "Pendiente",
		CreatedAt:   time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := incidentCollection.InsertOne(ctx, incident)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateIncident(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	newStatus := r.FormValue("status")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = incidentCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": newStatus}},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteIncident(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = incidentCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
