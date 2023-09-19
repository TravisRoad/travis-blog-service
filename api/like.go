package api

import (
	"net/http"

	"github.com/TravisRoad/travis-blog-service/service"
	"github.com/gin-gonic/gin"
)

type LikeApi struct{}

type AddLikeReq struct {
	Field     string `json:"field"`
	Operation string `json:"op"`
}

func (r *LikeApi) AddLike(c *gin.Context) {
	req := AddLikeReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	if len(req.Field) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "field is required",
		})
		return
	}

	ls := new(service.LikeService)
	if err := ls.AddLike(req.Field, req.Operation, c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})
}

func (r *LikeApi) GetLike(c *gin.Context) {
	field := c.Query("field")
	ls := new(service.LikeService)
	ans, err := ls.GetLike(field)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": ans,
		"msg":  "",
	})
}
