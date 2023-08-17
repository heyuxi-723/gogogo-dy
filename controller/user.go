package controller

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]models.User{
	"zhangleidouyin": {
		Id:   1,
		Name: "zhanglei",
	},
}

type UserResponse struct {
	models.Response
	User models.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	res := &service.UserLoginResponse{}
	err := res.Register(username, password)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	res := &service.UserLoginResponse{}
	err := res.Login(username, password)
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}

func UserInfo(c *gin.Context) {
	//自己的userId
	myUserId, ok := c.Get("user_id")
	if !ok {
		models.Fail(c, 1, "tokne解析出错")
		return
	}

	if myUserId, ok = myUserId.(int64); !ok {
		models.Fail(c, 1, "用户名ID解析出错")
		return
	}

	//要查询的userid
	userId := c.Query("user_id")

	res := &service.UserInfoResponse{}
	err := res.GetUserInfo(userId, strconv.FormatInt(myUserId.(int64), 10))
	if err != nil {
		models.Fail(c, 1, err.Error())
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "ok"
	c.JSON(http.StatusOK, res)
}
