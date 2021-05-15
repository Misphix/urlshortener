package httpserver

import (
	"urlshorterner/internal/shorturl"

	"github.com/gin-gonic/gin"
)

func NewHttpServer() (*gin.Engine, error) {
	engine := gin.Default()
	handler, err := shorturl.NewShortURLHandler()
	if err != nil {
		return nil, err
	}

	v1 := engine.Group("api/v1")
	{
		v1.POST("/urls", handler.UploadURL)
		v1.DELETE("/urls/:urlID", handler.DeleteURL)
	}
	engine.GET("/:urlID", handler.RedirectURL)

	return engine, nil
}
