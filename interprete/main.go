package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CodigoEnviado struct {
	Contenido string `json:"contenido"`
}

func manejarEnviarcodigo(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo JSON en una estructura CodigoEnviado
	var codigo CodigoEnviado
	err := json.NewDecoder(r.Body).Decode(&codigo)
	if err != nil {
		http.Error(w, "Error al leer el contenido JSON", http.StatusBadRequest)
		return
	}

	// Aquí puedes hacer lo que necesites con el contenido del textarea.
	// Por ejemplo, imprimirlo en la consola.
	fmt.Println("codigo recibido:", codigo.Contenido)

	// Puedes enviar una respuesta al cliente si lo deseas.
	respuesta := map[string]string{"mensaje": "Codigo recibido correctamente"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func main() {
	http.HandleFunc("/enviar-texto", manejarEnviarcodigo)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("localhost:3000", nil)
}