package server

import (
	"apis/config"
	"apis/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func RunServer() {
	gin.SetMode(gin.ReleaseMode)

	log.Info().Msg(fmt.Sprintf("Server start at http://0.0.0.0:%d", config.G.Port))
	h2Handler := h2c.NewHandler(router.InitRouter(), &http2.Server{})
	s := &http.Server{Handler: h2Handler, Addr: fmt.Sprintf(":%d", config.G.Port)}
	if err := s.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("ListenAndServe")
	}
}
