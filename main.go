package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite" // Importa o driver SQLite que não requer CGO
)

var db *sql.DB

func main() {
	// Abre (ou cria) o banco de dados SQLite
	var err error
	db, err = sql.Open("sqlite", "./contador.db")
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	// Cria a tabela se não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS contador (id INTEGER PRIMARY KEY, visitas INTEGER)`)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Insere a primeira linha caso a tabela esteja vazia
	_, err = db.Exec(`INSERT INTO contador (id, visitas) VALUES (1, 0) ON CONFLICT(id) DO NOTHING`)
	if err != nil {
		log.Fatalf("Erro ao inserir dados iniciais: %v", err)
	}

	http.HandleFunc("/api/contador-visitas", contadorHandler)

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func contadorHandler(w http.ResponseWriter, r *http.Request) {
	// Incrementa o contador no banco de dados
	_, err := db.Exec(`UPDATE contador SET visitas = visitas + 1 WHERE id = 1`)
	if err != nil {
		http.Error(w, "Erro ao atualizar contador", http.StatusInternalServerError)
		return
	}

	// Lê o valor atualizado
	row := db.QueryRow(`SELECT visitas FROM contador WHERE id = 1`)
	var visitas int
	err = row.Scan(&visitas)
	if err != nil {
		http.Error(w, "Erro ao ler contador", http.StatusInternalServerError)
		return
	}

	// Retorna o valor em JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"visitas": visitas})
}
