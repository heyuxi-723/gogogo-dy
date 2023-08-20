package service

import "github.com/RaymondCode/simple-demo/models"

// 点赞列表
func FavoriteList(userIdInt64 int64) (videoList []models.Video, err error) {
	//var res []models.Video
	favorites, err := models.QueryFavoriteByUserID(userIdInt64)
	if err != nil {
		return nil, err
	}
	videoIds := make([]int64, len(favorites))
	for i, favorites := range favorites {
		videoIds[i] = favorites.VideoId
		video := models.Video{}
		result := models.DB.Model(&models.Video{}).Where("videoId = ?", videoIds[i]).Find(&video)
		if result == nil {
			return nil, result.Error
		}
		videoList = append(videoList, video)
	}
	return videoList, err
}
