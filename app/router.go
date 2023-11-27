package app

import "github.com/gin-gonic/gin"

// https://go.dev/doc/tutorial/web-service-gin

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", getHello)
	router.GET("/ping", getPing)

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	return router
}
