package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	contador int
	mutex    sync.Mutex
)

func main() {
	http.HandleFunc("/api/contador-visitas", contadorHandler)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}

func contadorHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	contador++
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"visitas": contador})
}
