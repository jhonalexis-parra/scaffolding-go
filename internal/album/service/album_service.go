package service

import (
	"github.com/jhonalexis-parra/scaffolding-go/internal/album/domain"
)

type AlbumService struct {
}

func NewAlbumService() domain.AlbumService {
	return &AlbumService{}
}

func (s *AlbumService) GetAllAlbums() ([]domain.Album, error) {
	// This is a stub implementation. In a real application, this method would
	// interact with a database or another data source to retrieve the albums.
	albums := []domain.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return albums, nil
}
