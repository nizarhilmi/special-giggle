package data

import (
	"api-go-basic/models"
	"fmt"
)

type ListMovies struct {
	Movies []models.Movie
}

func (store *ListMovies) Create(movie *models.Movie) error {
	movie.ID = len(store.Movies) + 1
	store.Movies = append(store.Movies, *movie)
	return nil
}

func (store *ListMovies) ReadAll() ([]models.Movie, error) {
	return store.Movies, nil
}

func (store *ListMovies) Read(id int) (*models.Movie, error) {
	for _, movie := range store.Movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("movie with ID %d not found", id)
}

func (store *ListMovies) Update(id int, movie *models.Movie) error {
	for i, m := range store.Movies {
		if m.ID == id {
			movie.ID = id
			store.Movies[i] = *movie
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", id)
}

func (store *ListMovies) Delete(id int) error {
	for i, movie := range store.Movies {
		if movie.ID == id {
			store.Movies = append(store.Movies[:i], store.Movies[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("movie with ID %d not found", id)
}
