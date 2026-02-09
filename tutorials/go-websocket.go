// MARK: Instructions

// A télécharger pour avoir GO : "sudo snap install go --classic"
// A créer pour que GO fonctionne avec l'import nécessaire pour le websocket : "go mod init tutorial"
// A télécharger pour que le go puisse gérer le websocket : "go get github.com/gorilla/websocket"

/*
Instructions à suivre pour pouvoir tester le code dans un environnement de simulation d'un websocket.
# Télécharger
"wget https://github.com/vi/websocat/releases/download/v1.13.0/websocat.x86_64-unknown-linux-musl"

# Rendre le fichier exécutable
"chmod +x websocat.x86_64-unknown-linux-musl"

# Le déplacer vers /usr/local/bin
"sudo mv websocat.x86_64-unknown-linux-musl /usr/local/bin/websocat"

# Vérifier l'installation (Juste pour vérifier qu'il n'y ais pas de retours bizzares)
"websocat --version"
*/

// MARK: Comment l'utiliser :
// Ouvrez un premier terminal pour lancer le serveur websocket : "go run go-websocket.go"
// Ouvrez un second terminal pour pouvoir intérargir avec : "websocat ws://localhost:8080/ws"
// Dans le second terminal il y aura deux comportements :
//
//	-si vous ne tapez pas "start" : une phrase apparaît.
//	-si vous tapez start : un compteur commence.
//
// Vous pouvez tenter de vous connecter avec un navigateur à partir de http://localhost:8080/
// Vous pourrez voir des erreurs dans le premier terminal, celui du serveur.

// MARK: Le code
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Configuration de la connexion WebSocket.
// Transforme les requêtes HTTP entrantes en connexions WebSocket.
type webSocketHandler struct {
	upgrader websocket.Upgrader
	// "Upgrader" vérifie que la requête est une demande de passage en WebSocket
	// Il gère le Handshake entre le client et le serveur.
	// Il tansforme (upgrade) la connexion HTTP en une connexion WebSocket persistante.
}

// ServeHTTP rend compatible l'objet avec le serveur web GO (On peut l'utiliser avec http.ListenAndServe).
func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// L'upgrader vérifie les en-têtes de la requête HTTP (précisément : "Connection: Upgrade")
	// Si positif, alors "c" représente la connexion WebSocket active.
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}
	// Rôle du defer : Avant que la fonction ServeHTTP ne se termine, lance les instructions suivantes.
	// Pourquoi ? En cas de coupure ou de plantage
	defer func() {
		log.Println("closing connection")
		c.Close()
	}()

	// Création d'un boucle infinie pour réceptionner les messages.
	for {
		mt, message, err := c.ReadMessage()
		// "c" contient la connection de la websocket que nous lisons.
		// ReadMessage va lire les messages en entré, mais va s'arrêter ici dans l'exécution du code si il n'y a rien.
		// ReadMessage renvois trois variables.
		// "mt" : Message Type. C'est un entier. Si valeur 1 : Texte | Si valeur 2 : Binaire
		// "message" : c'est un tableau d'octets ([]byte). Appelé charge utile ou payload, envoyépar le client.
		// "err" : pour l'objet erreur.
		if err != nil {
			log.Printf("Error %s when reading message from client", err)
			return
		}
		// Ici on vérifie si mt correspond bien à 2, ce qui indique que le message est en binaire.
		if mt == websocket.BinaryMessage {
			err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}
			return
		}
		log.Printf("Receive message %s", string(message))

		// Dans le cadre du tuto, si l'utilisateur n'entre pas "start", rien ne se passe
		if strings.Trim(string(message), "\n") != "start" { // Enlève le "\n" qui pourrait être inséré si l'utilisateur appuie sur la touche "Entrée".
			//Ecrit un message d'erreur avec le message envoyé par l'utilisateur.
			err = c.WriteMessage(websocket.TextMessage, []byte("Le mot pour débuter, c'est start :P"))
			if err != nil {
				log.Printf("Error %s whend sending message to client", err)
				return
			}
			continue
		}
		log.Println("start responding to client...")

		// initialisation d'un compteur pour indiquer l'ordre des messages.
		i := 1
		// Initialisation d'une boucle infinie pour le fil des messages envoyés.
		for {
			response := fmt.Sprintf("Notification %d", i)
			err = c.WriteMessage(websocket.TextMessage, []byte(response))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}
			i++
			// Pour éviter des envois de messages permanents, il y a un temps d'attente de 2 secondes entre chaque envoi.
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}
	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
