package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct { // struct é um objeto. O que precisa ter em um filme?
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

// "movies" é uma variavel
var movies = []Movie{} // "movie" é um array

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil{ // verificando erro // "nil" é null
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // badrequest é um tipo de erro (erro 400)
	}

	newMovie.ID = len(movies) + 1 // variavel (id) pega o tamanho do array e adicione mais 1
	movies = append(movies, newMovie) // uma "apeend" é adicionar e concatenar no array
	c.JSON(http.StatusOK, newMovie) // statusok (sucesso 200)
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}


func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie) //enviar algo (POST) no endpoint createmovie
	r.GET("/movie", GetMovies)
	r.Run(":8080")
}