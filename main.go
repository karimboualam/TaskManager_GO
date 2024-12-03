//package main

/*import "fmt"

func main() {
    fmt.Println("Bienvenue dans le projet TaskManager!")
}*/

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Définir une route pour tester
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bienvenue sur l'API TaskManager!"})
	})

	// Lancer le serveur
	r.Run(":8080") // Par défaut sur le port 8080
}
