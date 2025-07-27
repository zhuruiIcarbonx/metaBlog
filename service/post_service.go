package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuruiIcarbonx/metaBlog/base"
	"github.com/zhuruiIcarbonx/metaBlog/base/errorcode"
	"github.com/zhuruiIcarbonx/metaBlog/base/token"
	"github.com/zhuruiIcarbonx/metaBlog/dao"
	"gorm.io/gorm"
)

type PostDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostListDto struct {
	Title string `json:"title"`
}

type PostUpateDto struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func PostCreate(c *gin.Context) {

	result := base.Result{}

	var dto PostDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(200, result.FailWeb(400, err.Error()))
		return
	}

	userId := token.GetUserId(c)

	post := &dao.Post{
		Title:   dto.Title,
		Content: dto.Content,
		UserId:  userId,
	}

	db := dao.InitDb()
	err := db.Transaction(func(tx *gorm.DB) error {

		err := dao.PostCreate(db, post)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(200, result.FailCommon(errorcode.Operation_error, err.Error()))
		return
	}

	fmt.Printf("---------------------------post:%v", post)
	c.JSON(200, result.Sucess())

}

func PostList(c *gin.Context) {

	result := base.Result{}

	var dto PostListDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(200, result.FailWeb(400, err.Error()))
		return
	}

	fmt.Printf("[PostList]dto---------------------------%v", dto)

	db := dao.InitDb()
	post := &dao.Post{
		Title: dto.Title,
	}
	list := dao.PostList(db, post)

	c.JSON(200, result.SucessData(list))

}

func PostOne(c *gin.Context) {

	result := &base.Result{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	db := dao.InitDb()
	post := dao.PostOne(db, userId)

	c.JSON(200, result.SucessData(post))

}

func PostUpdate(c *gin.Context) {

	result := base.Result{}

	var dto PostUpateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(200, result.FailWeb(400, err.Error()))
		return
	}
	fmt.Printf("[PostList]dto---------------------------%v", dto)

	userId := token.GetUserId(c)

	db := dao.InitDb()
	storedOne := dao.PostOne(db, int(dto.ID))
	if storedOne.ID == 0 {
		c.JSON(200, result.SucessData(errorcode.No_data))
		return

	}
	if storedOne.UserId != userId {
		c.JSON(200, result.SucessData(errorcode.No_data_permission))
		return

	}

	post := &dao.Post{
		Title:   dto.Title,
		Content: dto.Content,
	}
	post.ID = dto.ID //放里面报错，拿出来

	err := db.Transaction(func(tx *gorm.DB) error {

		err := dao.PostUpate(db, post)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(200, result.FailCommon(errorcode.Operation_error, err.Error()))
		return
	}

	c.JSON(200, result.SucessData(post))

}

func PostDelete(c *gin.Context) {

	result := base.Result{}

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, result.Fail(errorcode.Param_error))
		return
	}
	fmt.Printf("[PostList]idStr---------------------------%v", idStr)

	userId := token.GetUserId(c)

	db := dao.InitDb()
	storedOne := dao.PostOne(db, id)
	if storedOne.UserId != userId {
		c.JSON(200, result.SucessData(errorcode.No_data_permission))
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {

		err := dao.CommentDelete(db, id)
		if err != nil {
			return err
		}
		err = dao.PostDelete(db, id)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(200, result.FailCommon(errorcode.Operation_error, err.Error()))
		return
	}

	c.JSON(200, result.Sucess())

}

// userId := token.GetUserId(c)
// 	log.Printf("cureent userId is %d", userId)
// 	username := token.GetUsername(c)
// 	log.Printf("cureent userId is %s", username)
