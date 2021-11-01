package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:3600")
}

// album representa o dado sobre o album gravado
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//percorrendo albuns
var albums = []album{
	{ID: "1", Title: "22", Artist: "Taylor Swift", Price: 56.99},
	{ID: "2", Title: "Gelo e Gin", Artist: "Tribo da Periferia", Price: 21.00},
	{ID: "3", Title: "Iluminado", Artist: "Xamã", Price: 36.39},
}

//getAlbums responde com a lista de todos os albuns como JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums adiciona um álbum de JSON recebido no corpo da solicitação.
func postAlbums(c *gin.Context) {
	var newAlbum album

	//chama BindJSON para ligar os JSON recebidos no newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add o novo album no pedaço
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID localiza o album pelo valor do ud, que é um parametro enviado pelo cliente
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Verifica a lista de álbuns procurando pelo álbum que dá match com o id do parâmetro
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
