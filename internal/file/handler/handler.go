package handler

import (
	"github.com/HaHadaxigua/melancholy/internal/basic/middleware"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
	"github.com/HaHadaxigua/melancholy/internal/file/service"
	"github.com/HaHadaxigua/melancholy/internal/global/consts"
	"github.com/HaHadaxigua/melancholy/internal/global/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupFileRouters(r gin.IRouter) {
	// open

	// secured
	secured := r.Group("/file", middleware.JWT)
	// 文件夹
	secured.POST("/create", createFolder)
	secured.GET("/space", fileSpace)
	secured.PATCH("/folder", modifyFolder)
	secured.DELETE("/folder/:id", deleteFolder)
}

func createFolder(c *gin.Context) {
	req := &msg.ReqFolderCreate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	} else {
		uid := c.GetInt(consts.UserID)
		req.UserID = uid
	}
	if err := service.File.CreateFolder(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}

func fileSpace(c *gin.Context) {
	uid := c.GetInt(consts.UserID)
	if rsp, err := service.File.ListFileSpace(uid); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(rsp))
	}
}

func modifyFolder(c *gin.Context) {
	req := &msg.ReqFolderUpdate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	} else {
		req.UserID = c.GetInt(consts.UserID)
	}
	if err := service.File.UpdateFolder(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Ok(nil))
	}
}
func deleteFolder(c *gin.Context) {
	folderID := c.Param("id")
	if folderID == "" {
		c.JSON(http.StatusBadRequest, response.NewErr(nil))
		return
	}
	if err := service.File.DeleteFolder(folderID, c.GetInt(consts.UserID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErr(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(nil))
}
