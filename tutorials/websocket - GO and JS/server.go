// MARK: Instructions

// Etapes si vous n'avez pas du langage GO sur votre espace de travail.
// A télécharger pour avoir GO : "sudo snap install go --classic"
// A télécharger pour que le GO puisse gérer le websocket : "go get github.com/gorilla/websocket"
// A créer pour que GO fonctionne avec l'import nécessaire pour le websocket : "go mod init tutorial"

// Etapes si vous avez du GO sur votre espace de travail.
// 1- A télécharger pour que le GO puisse gérer le websocket : "go get github.com/gorilla/websocket"
// 2- Lancer le serveur qui va gérer le websocket : go run server.go
// 3- Utiliser Live Server, qui va ouvrir une nouvelle page web pour gérer la discussion.
// -----------------------, Si vous n'avez pas Live Server, télécharger le ici comme extention de VSCode : https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// MARK: Du HTTP au WS
// La simple connection HTTP est upgradée vers le protocole binaire WebSocket.
var upgrader = websocket.Upgrader{ // ".Upgrader" contient les paramètres de configuration pour cette transition.
	// CheckOrigin vérifie la requête HTTP entrante.
	CheckOrigin: func(r *http.Request) bool {
		// Dans le cadre de l'exemple, toute les requêtes de connexions sont acceptées.
		return true
	},
}

// MARK: Websocket
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// L'upgrader vérifie les en-têtes de la requête HTTP (précisément : "Connection: Upgrade")
	// Si positif, alors "conn" représente la connexion WebSocket active.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	// Rôle du defer : Avant que la fonction wsHandler ne se termine, lance les instructions suivantes.
	// Pourquoi ? En cas de coupure ou de plantage.
	defer conn.Close()

	// Boucle for infinie pour écouter les messages entrants.
	for {
		_, message, err := conn.ReadMessage()
		// "conn" contient la connection de la websocket que nous lisons.
		// ReadMessage va lire les messages en entrée, mais va s'arrêter ici dans l'exécution du code si il n'y a rien.
		// ReadMessage renvois trois variables.
		// "_" : Message Type. Dans notre cas on l'ignore. C'est un entier. Si valeur 1 : Texte | Si valeur 2 : Binaire
		// ---------- Ici nous nous attendons un format précis, donc pas de vérifications sur le Type de Message.
		// "message" : c'est un tableau d'octets ([]byte). Appelé charge utile ou payload, envoyé par le client.
		// "err" : pour l'objet erreur.
		if err != nil {
			fmt.Printf("Error %s when reading message from client \n", err)
			break
		}
		// Affiche dans le terminal le message envoyé.
		// Le format "%s" dans "Printf" permet de convertir le []byte en String lisible pour nous humains dans le terminal.
		fmt.Printf("Received : %s\n", (strings.Trim(string(message), "\n"))) // "Trim" est ici pour éliminer les sauts de lignes dans le message.

		// "conn.WriteMessage" va écrire le message dans le client websocket contenu dans "conn".
		// "websocket.TextMessage" ceci force la lecture des données binaires en []byte, à devenir du texte (UTF-8).
		// "message" c'est ce qui va être écrit.
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message : %s", err)
			break
		}
	}
}

// MARK: Serveur
func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Websocket server started on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server : %s", err)
	}
}
