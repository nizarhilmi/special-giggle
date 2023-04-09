package controller

import (
	"api-go-basic/data"
	"api-go-basic/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type movieInput struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var store = data.ListMovies{}

// CreateMovie godoc
// @Summary Create New Movie.
// @Description Creating a new Movie.
// @Tags Movie
// @Param Body movieInput true "the body to create a new movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var input movieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	movie := models.Movie{Title: input.Title, Year: input.Year}
	if err := store.Create(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"Data": movie})
}

// GetAllMovies godoc
// @Summary Get all movies.
// @Description Get a list of Movies.
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovies(c *gin.Context) {
	movies, err := store.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": movies})
}
