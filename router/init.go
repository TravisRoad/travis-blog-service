package router

import (
	"github.com/TravisRoad/travis-blog-service/api"
	"github.com/TravisRoad/travis-blog-service/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (rt *Router) Register(r *gin.Engine) {
	rr := r.Group("/api").Use(middleware.Auth())

	likeApi := &api.LikeApi{}
	rr.POST("/like", likeApi.AddLike)
	rr.GET("/like", likeApi.GetLike)
}
