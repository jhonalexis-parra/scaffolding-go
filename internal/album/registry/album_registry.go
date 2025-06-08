package registry

import (
	"github.com/jhonalexis-parra/scaffolding-go/internal/album/api"
	"github.com/jhonalexis-parra/scaffolding-go/internal/album/domain"
	"github.com/jhonalexis-parra/scaffolding-go/internal/album/service"
)

func CreateAlbumApi() api.AlbumApi {
	albumService := service.NewAlbumService()
	return api.NewAlbumApi(albumService)
}

func CreateAlbumService() domain.AlbumService {
	return service.NewAlbumService()
}
