package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Ici on vérifie si "mt" correspond bien à 2, ce qui indique que le message est en binaire ([]byte).
// 1 indique un message Texte.
// Point à vigiler, je ne suis pas sûr si on récupère automatiquement un MT de 1 ou 2, suivant le serveur ou le navigateur.
func IsMessageTypeValid(mt int, conn *websocket.Conn) bool {
	if mt != websocket.BinaryMessage {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Attention, vous n'avez pas mis que des caractères autorisés."))
		if err != nil {
			log.Printf("Error %s when sending message to client", err)
			return false
		}
		return false
	}
	return true
}

// Vérification que le dernier message de l'utilisateur à plus de deux secondes.
// Pour éviter les spam.
func SafetyMessageRateSending2Seconds(db *sql.DB, senderID string, conn *websocket.Conn) bool {
	var createdAt time.Time
	err := db.QueryRow("SELECT CreatedAt FROM messages WHERE SenderID = ? ORDER BY CreatedAt DESC LIMIT 1", senderID).Scan(&createdAt)
	if err == sql.ErrNoRows { // Si il n'y a pas d'historique d'échanges, ceci ignore la règle. Sinon il y a blocage.
		return true
	}
	if err != nil { // Gestion des erreurs.
		log.Printf("Error during verification of ruleSendingMessageTimeRateLimit2Seconds : %v", err)
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Attention, nous rencontrons une erreur du côté serveur."))
		return false
	}
	if time.Since(createdAt) < 2*time.Second { //Affichage message d'erreur auprès de l'utilisateur.
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Attention, vous ne pouvez pas spammer l'envoi de messages."))
		return false
	}
	return true
}

// Sécurité pour les messages vides.
func IsMessageNotEmpty(message []byte, conn *websocket.Conn) bool {
	if len(message) == 0 {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Attention, vous ne pouvez pas envoyer de messages vides."))
		return false
	}
	return true
}

// Sécurité pour les messages trop grands.
func IsMessageTooTall(message []byte, conn *websocket.Conn) bool {
	if len(message) > 800 { // 800 bytes ou caractères.
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Attention, votre message excède la taille limite."))
		return false
	}
	return true
}

// Ecriture du message envoyé dans la BDD
func CreateMessageInBDD(db *sql.DB, message []byte, senderID string, receiverID string) error {
	query := `INSERT INTO messages (SenderID, ReceiverID, Content) VALUES (?, ?, ?)`
	_, err := db.Exec(query, senderID, receiverID, message)
	return err
}

// Ecriture du message chez les utilisateurs.
func WriteMessagesFromBddToUserScreen(db *sql.DB, conn *websocket.Conn, receiverID string, senderID string, lastChecked *time.Time) {
	// 1. On cherche les messages qui sont destinés à l'utilisateur actuel (receiverID)
	// et qui ont été envoyés par son contact (senderID) après la dernière vérification.
	rows, err := db.Query("SELECT Content, CreatedAt FROM messages WHERE ReceiverID  = ? AND SenderID = ? AND CreatedAt > ? ORDER BY CreatedAt ASC", receiverID, senderID, *lastChecked)
	if err != nil {
		log.Printf("Error querying messages: %v", err)
		return
	}
	defer rows.Close() //Fermeture de la connexion une fois la vérification finie.

	for rows.Next() {
		var content string
		var createdAt time.Time

		// 2. On scanne les données de la ligne en cours.
		err := rows.Scan(&content, &createdAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		// 3. On envoie le contenu trouvé au client via la websocket.
		err = conn.WriteMessage(websocket.TextMessage, []byte(content))
		if err != nil {
			log.Printf("Error sending message to websocket: %v", err)
			return
		}

		// 4. On met à jour le curseur de temps avec la date du message qu'on vient d'envoyer.
		*lastChecked = createdAt
	}
}
