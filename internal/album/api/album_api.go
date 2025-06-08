package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhonalexis-parra/scaffolding-go/internal/album/domain"
)

type AlbumApi struct {
	albumService domain.AlbumService
}

func NewAlbumApi(albumService domain.AlbumService) AlbumApi {
	return AlbumApi{
		albumService: albumService,
	}
}

func (api *AlbumApi) GetAlbums(c *gin.Context) {

	albums, err := api.albumService.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
