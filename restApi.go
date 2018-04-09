package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Objeto deporte
type Sport struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"firstname,omitempty"`
}

var deportes []Sport //Lista con los deportes
/*
* Muestra todos los datos referentes a deportes en formato JSON
 */
func GetDeportesHandler(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(deportes)
}

/*
* Muestra los datos referentes al deporte con el id del parametro
 */
func GetSportHandler(writer http.ResponseWriter, request *http.Request) {
	ids := mux.Vars(request)
	for _, i := range deportes {
		if i.ID == ids["id"] {
			json.NewEncoder(writer).Encode(&i)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Sport{})
}

/*
* Crea un nuevo deporte con el id del parametro
* Se deberia no pasar ningun parametro de id y generarlo automaticamente
 */
func CreateDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	ids := mux.Vars(request)
	var dep Sport
	_ = json.NewDecoder(request.Body).Decode(&dep)
	dep.ID = ids["id"]
	dep.Name = "Tenis"
	deportes = append(deportes, dep)
	json.NewEncoder(writer).Encode(deportes)
}

/*
* Elimina el deporte con el id del parametro
 */
func DeleteDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	ids := mux.Vars(request)
	for i, it := range deportes {
		if it.ID == ids["id"] {
			deportes = append(deportes[:i], deportes[i+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(deportes)
}

/*
* Actualiza los datos del id del parametro
 */
func UpdateDeporteHandler(writer http.ResponseWriter, request *http.Request) {
	ids := mux.Vars(request)
	for i, it := range deportes {
		if it.ID == ids["id"] {
			deportes = append(deportes[:i], deportes[i+1:]...)
			var dep Sport
			_ = json.NewDecoder(request.Body).Decode(&dep)
			dep.ID = ids["id"]
			dep.Name = "Hockey"
			deportes = append(deportes, dep)
			json.NewEncoder(writer).Encode(dep)
			return
		}
	}
}
func main() {
	//Iniciamos el array de prueba
	deportes = append(deportes, Sport{ID: "1", Name: "Futbol"})
	deportes = append(deportes, Sport{ID: "2", Name: "Basket"})

	router := mux.NewRouter() //Creamos un nuevo router
	router.HandleFunc("/deporte", GetDeportesHandler).Methods("GET")
	router.HandleFunc("/deporte/{id}", GetSportHandler).Methods("GET")
	router.HandleFunc("/deporte/{id}", CreateDeporteHandler).Methods("POST")
	router.HandleFunc("/deporte/{id}", DeleteDeporteHandler).Methods("DELETE")
	router.HandleFunc("/deporte/{id}", UpdateDeporteHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router)) //Iniciamos el server y esperamos a la escucha en el puerto 3000

}
