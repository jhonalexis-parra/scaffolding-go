package domain

type AlbumService interface {
	GetAllAlbums() ([]Album, error)
}