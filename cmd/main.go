package main

import (
	"github.com/gin-gonic/gin"
	albumRegistry "github.com/jhonalexis-parra/scaffolding-go/internal/album/registry"
)

func main() {
	albumHandlers := albumRegistry.CreateAlbumApi()

	router := gin.Default()
	router.GET("/albums", albumHandlers.GetAlbums)

	router.Run("localhost:8080")
}
