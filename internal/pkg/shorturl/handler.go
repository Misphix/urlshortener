package shorturl

import (
	"fmt"
	"net/http"
	"path"
	"time"
	"urlshorterner/internal/configmanager"
	"urlshorterner/internal/pkg/shorturl/database"
	"urlshorterner/internal/pkg/shorturl/service"

	"github.com/gin-gonic/gin"
)

type ShortURLHandler struct {
	s        *service.ShortURL
	basePath string
}

type UploadRequest struct {
	URL      string    `json:"url"`
	ExpireAt time.Time `json:"expireAt"`
}

type ShortUrlResponse struct {
	ID       string `json:"id"`
	ShortURL string `json:"shortUrl"`
}

func NewShortURLHandler() (*ShortURLHandler, error) {
	config, err := configmanager.Get()
	if err != nil {
		return nil, err
	}

	db, err := database.NewMySQL(config.Database.DSN)
	if err != nil {
		return nil, err
	}

	return &ShortURLHandler{
		s:        service.NewShortURL(db),
		basePath: fmt.Sprintf("%s:%d", config.HTTPServer.Domain, config.HTTPServer.Port),
	}, nil
}

func (h *ShortURLHandler) UploadURL(ctx *gin.Context) {
	var request UploadRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	urlID, err := h.s.Shorter(request.URL, request.ExpireAt)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	response := ShortUrlResponse{
		ID:       urlID,
		ShortURL: path.Join(h.basePath, urlID),
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ShortURLHandler) DeleteURL(ctx *gin.Context) {
	urlID, ok := ctx.Params.Get("urlID")
	if !ok {
		ctx.Status(http.StatusBadRequest)
	}

	if err := h.s.DeleteURL(urlID); err != nil {
		_, ok := err.(*database.DatabaseError)
		if ok {
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (h *ShortURLHandler) RedirectURL(ctx *gin.Context) {
	urlID, ok := ctx.Params.Get("urlID")
	if !ok {
		ctx.Status(http.StatusBadRequest)
	}

	url, err := h.s.GetURL(urlID)
	if err != nil {
		_, ok := err.(*database.DatabaseError)
		if ok {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusNotFound)
		}
		return
	}

	ctx.Redirect(http.StatusFound, url)
}
