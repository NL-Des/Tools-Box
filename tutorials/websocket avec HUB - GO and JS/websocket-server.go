package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// La connexion HTTP est "Upgradée" en une connexion WebSocket (WS).
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Déclarations temporaires.
// C'est pour imiter l'importation/réception des données des utilisateurs qui écrivent dans la websocket.
var pathDB string

// On part du principe que la BDD est initialisée avant dans le main ou le server.
var db *sql.DB

// Chaque nouvelle discussion va générer une nouvelle connexion.
// La fonction sera rutilisée et recréée pour chaque discussion.
func wsHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade de la connexion.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading: %v \n", err)
		return
	}
	defer conn.Close()

	// Obtention des ID à partir des informations de la page.
	senderID := r.URL.Query().Get("ID de l'envoyeur")
	receiverID := r.URL.Query().Get("ID du receveur")
	// Variable créée pour vérifier périodiquement si il y a de nouveaux messages à afficher.
	var lastChecked = time.Now()

	// Si il manque l'ID de l'envoyeur ou du destinataire.
	if senderID == "" {
		log.Printf("Missing senderID")
		return
	}
	if receiverID == "" {
		log.Printf("Missing receiverID")
		return
	}

	// Goroutine de réception des messages.
	go func() {
		for {
			// Réception du message de l'utilisateur en []byte.
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error %s when reading message from client", err)
				return
			}
			if IsMessageTypeValid(mt, conn) != true {
				return
			}
			if SafetyMessageRateSending2Seconds(db, senderID, conn) != true {
				return
			}
			if IsMessageNotEmpty(message, conn) != true {
				return
			}
			if IsMessageTooTall(message, conn) != true {
				return
			}
			err = CreateMessageInBDD(db, message, senderID, receiverID)
			if err != nil {
				log.Printf("Error %s when writing the new message in the BDD", err)
			}
		}
	}()

	// Goroutine d'écriture des messages.
	go func() {
		for {
			WriteMessagesFromBddToUserScreen(db, conn, receiverID, senderID, &lastChecked)
			time.Sleep(2 * time.Second)
		}
	}()
}
