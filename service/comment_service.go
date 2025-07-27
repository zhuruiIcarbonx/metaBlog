package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuruiIcarbonx/metaBlog/base"
	"github.com/zhuruiIcarbonx/metaBlog/base/errorcode"
	"github.com/zhuruiIcarbonx/metaBlog/base/token"
	"github.com/zhuruiIcarbonx/metaBlog/dao"
	"github.com/zhuruiIcarbonx/metaBlog/logger"
	"gorm.io/gorm"
)

type CommentDto struct {
	PostId  int    `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required,max=500"`
}

type CommentListDto struct {
	PostId int `json:"post_id" binding:"required"`
}

func CommentCreate(c *gin.Context) {

	result := base.Result{}

	var dto CommentDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(200, result.FailWeb(400, err.Error()))
		return
	}

	userId := token.GetUserId(c)

	comment := &dao.Comment{
		PostId:  uint(dto.PostId),
		Content: dto.Content,
		UserId:  userId,
	}

	db := dao.InitDb()

	err := db.Transaction(func(tx *gorm.DB) error {

		err := dao.CommentCreate(db, comment)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(200, result.FailCommon(errorcode.Operation_error, err.Error()))
		return
	}

	logger.Log.Info("comment---------------------------:%v", comment)
	c.JSON(200, result.Sucess())

}

func CommentList(c *gin.Context) {

	result := base.Result{}

	var dto CommentListDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(200, result.FailWeb(400, err.Error()))
		return
	}

	logger.Log.Info("[CommentList]dto---------------------------%v", dto)

	db := dao.InitDb()
	list := dao.CommentList(db, dto.PostId)

	c.JSON(200, result.SucessData(list))

}
