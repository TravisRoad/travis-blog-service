package service

import (
	"github.com/TravisRoad/travis-blog-service/global"
	"github.com/TravisRoad/travis-blog-service/helper"
	"github.com/TravisRoad/travis-blog-service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LikeService struct{}

func GetRealIp(c *gin.Context) string {
	if ip := c.GetHeader("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
		return ip
	}
	return c.ClientIP()
}

func (ls *LikeService) AddLike(field, op string, c *gin.Context) error {
	ip := GetRealIp(c)
	userAgent := c.Request.UserAgent()

	likeOp := model.LikeOp{
		Operation: "add",
		Field:     field,
		UserAgent: userAgent,
		IPAddress: uint32(helper.IPv4ToUInt32(ip)),
	}
	like := model.Like{
		Field: field,
	}

	tx := global.DB.Begin()
	if err := tx.Debug().Model(&model.Like{}).Where("field = ?", field).FirstOrCreate(&like).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&likeOp).Error; err != nil {
		tx.Rollback()
		return err
	}

	expr := gorm.Expr("number + 1")
	if op == "minus" {
		expr = gorm.Expr("number - 1")
	}

	if err := tx.Model(&model.Like{}).Where("id = ?", like.ID).Update("number", expr).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ls *LikeService) GetLike(field string) (int, error) {
	like := model.Like{}
	if err := global.DB.Where("field = ?", field).First(&like).Error; err != nil {
		return 0, err
	}
	return like.Number, nil
}
