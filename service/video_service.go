package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context, authorId int64, data *multipart.FileHeader, title string) {
	// 存储视频
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", authorId, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 视频信息存入数据库
	video := &models.Video{
		Title:    title,
		AuthorID: authorId,
		PlayUrl:  config.Config.Url + "/static/videos/" + finalName,
		CoverUrl: "",
	}
	models.AddVideo(video)
}
