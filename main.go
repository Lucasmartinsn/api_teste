package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
/*type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}*/
type projeto struct {
	ID           string `json:"id"`
	Nome_projeto string `json:"nome_projeto"`
	Equipe_resp  string `json:"equipe_resp"`
}

// album represents data about a record album.
type equipe struct {
	ID           string `json:"id"`
	NOME_EQUIPE  string `json:"nome_equipe"`
	Scrum_master string `json:"scrum_master"`
	Id_projeto   string `json:"id_projeto"`
}

// albums slice to seed record album data.
/*var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}*/

// albums slice to seed record album data.
var equipes = []equipe{
	{ID: "1", NOME_EQUIPE: "log", Scrum_master: "lucas", Id_projeto: "1"},
}
var projetos = []projeto{
	{ID: "1", Nome_projeto: "beta", Equipe_resp: "1"},
}

func main() {
	router := gin.Default()
	router.GET("/projetos", getProjetos)
	router.GET("/albums/:id", getProjetoByID)
	router.POST("/albums", postProjeto)

	router.GET("/equipes", getEquipes)
	router.GET("/equipes/:id", getEquipeByID)
	router.POST("/equipes", postEquipe)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getEquipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, equipes)
}
func getProjetos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projetos)
}

// postAlbums adds an album from JSON received in the request body.
func postEquipe(c *gin.Context) {
	var newEquipe equipe

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newEquipe); err != nil {
		return
	}

	// Add the new album to the slice.
	equipes = append(equipes, newEquipe)
	c.IndentedJSON(http.StatusCreated, newEquipe)
}
func postProjeto(c *gin.Context) {
	var newProjeto projeto

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProjeto); err != nil {
		return
	}

	// Add the new album to the slice.
	projetos = append(projetos, newProjeto)
	c.IndentedJSON(http.StatusCreated, newProjeto)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getEquipeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range equipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipes not found"})
}
func getProjetoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range projetos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
