package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FeedResponse models.FeedResponse

const (
	MaxVideoNum = 30
)

func (f *FeedResponse) DoNoToken(c *gin.Context) error {
	lastTime := getLastTime(c)
	err := models.QueryVideoListByLimitAndTime(MaxVideoNum, *lastTime, &f.VideoList)
	if err != nil {
		return err
	}
	f.NextTime = lastTime.Unix()
	return nil
}

func (f *FeedResponse) DoHasToken(token string, c *gin.Context) error {
	if claim, ok := middleware.ParseToken(token); ok {
		//token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		lastTime := getLastTime(c)

		//如果用户为登录状态，则更新该视频是否被该用户点赞的状态
		lastTime, err := fillVideoListFields(claim.UserId, &f.VideoList) //不是致命错误，不返回
		if err != nil {
			return err
		}
		err = models.QueryVideoListByLimitAndTime(MaxVideoNum, *lastTime, &f.VideoList)
		if err != nil {
			return err
		}
		f.NextTime = lastTime.Unix()
		return nil
	}
	return nil
}

// FillVideoListFields 填充每个视频的作者信息（因为作者与视频的一对多关系，数据库中存下的是作者的id
// 当userId>0时，我们判断当前为登录状态，其余情况为未登录状态，则不需要填充IsFavorite字段
func fillVideoListFields(userId int64, videos *[]*models.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}

	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	//添加作者信息，以及is_follow状态
	for i := 0; i < size; i++ {
		//var userInfo models.UserInfo
		//user, ok := models.QueryUserLogin(strconv.FormatInt((*videos)[i].AuthorID,10), "id")
		//if !ok {
		//	continue
		//}
		//userInfo.IsFollow = p.GetUserRelation(userId, userInfo.Id) //根据cache更新是否被点赞
		//(*videos)[i].Author = userInfo
		////填充有登录信息的点赞状态
		//if userId > 0 {
		//	(*videos)[i].IsFavorite = p.GetVideoFavorState(userId, (*videos)[i].Id)
		//}
	}
	return &latestTime, nil
}

func getLastTime(c *gin.Context) (latestTime *time.Time) {
	rawTimestamp, ok := c.GetQuery("latest_time")
	if ok {
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil {
			*latestTime = time.Unix(0, intTime)
		}
	}
	*latestTime = time.Now()
	return latestTime
}
