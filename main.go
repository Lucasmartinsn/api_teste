package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type projeto struct {
	ID           string `json:"id"`
	Nome_projeto string `json:"nome_projeto"`
	Equipe_resp  string `json:"equipe_resp"`
}
type pessoa struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Id_Projeto string `json:"Id_Projeto"`
	Id_equipe  string `json:"id_equipe"`
}

// album represents data about a record album.
type equipe struct {
	ID           string `json:"id"`
	NOME_EQUIPE  string `json:"nome_equipe"`
	Scrum_master string `json:"scrum_master"`
	Id_projeto   string `json:"id_projeto"`
}

// albums slice to seed record album data.
var equipes = []equipe{
	{ID: "1", NOME_EQUIPE: "log", Scrum_master: "lucas", Id_projeto: "1"},
}
var projetos = []projeto{
	{ID: "1", Nome_projeto: "beta", Equipe_resp: "1"},
}
var pessoas = []pessoa{
	{ID: "1", Nome: "Bruno", Id_Projeto: "45", Id_equipe: "1"},
	{ID: "2", Nome: "Pedro", Id_Projeto: "12", Id_equipe: "3"},
	{ID: "3", Nome: "Caio", Id_Projeto: "13", Id_equipe: "2"},
	{ID: "4", Nome: "lucas", Id_Projeto: "2", Id_equipe: "3"},
}

func main() {
	router := gin.Default()
	router.GET("/projetos", getProjetos)
	router.GET("/projetos/:id", getProjetoByID)
	router.POST("/projetos", postProjeto)
	router.PUT("/projetos/:id", updateProjetosById)

	router.GET("/equipes", getEquipes)
	router.GET("/equipes/:id", getEquipeByID)
	router.POST("/equipes", postEquipe)

	router.GET("/pessoas", getPessoas)
	router.GET("/pessoas/:id", getpessoaByID)
	router.POST("/pessoas", postpessoas)
	router.DELETE("/pessoas/:id", deletePessoaById)
	router.PUT("/pessoas/:id", updatePessoaById)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getEquipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, equipes)
}
func getProjetos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projetos)
}
func getPessoas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pessoas)
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
func postpessoas(c *gin.Context) {
	var newpessoa pessoa

	// Call BindJSON to bind the received JSON to
	// newpessoa.
	if err := c.BindJSON(&newpessoa); err != nil {
		return
	}

	// Add the new pessoa to the slice.
	pessoas = append(pessoas, newpessoa)
	c.IndentedJSON(http.StatusCreated, newpessoa)
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
func getpessoaByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of pessoas, looking for
	// an pessoa whose ID value matches the parameter.
	for _, a := range pessoas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
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
func updateProjetosById(c *gin.Context) {
	id := c.Param("id")
	for i := range projetos {
		if projetos[i].ID == id {
			c.BindJSON(&projetos[i])
			c.IndentedJSON(http.StatusOK, projetos[i])
			return
		}
	}
}
func deletePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range pessoas {
		if a.ID == id {
			pessoas = append(pessoas[:i], pessoas[i+1:]...)
			return
		}
	}
}

func updatePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i := range pessoas {
		if pessoas[i].ID == id {
			c.BindJSON(&pessoas[i])
			c.IndentedJSON(http.StatusOK, pessoas[i])
			return
		}
	}
}
