package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	// 身份验证
	rawId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "token解析出错")
		return
	}

	id, ok := rawId.(int64) //保证id是int64
	if !ok {
		models.Fail(c, 1, "user_id不是int64类型")
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 存储视频
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", id, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	title := c.PostForm("title")
	// 视频信息存入数据库
	video := &models.Video{
		Title:    title,
		AuthorID: id,
		PlayUrl:  saveFile,
		CoverUrl: "",
	}
	models.AddVideo(video)

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
