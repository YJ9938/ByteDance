package controller

import (
	"net/http"
	"strconv"

	"github.com/YJ9938/DouYin/model"
	"github.com/YJ9938/DouYin/service"
	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	Response
}

type FavoriteListResponse struct {
	Response
	VideoList []model.VideoDisplay `json:"video_list,omitempty"`
}

// 因为response 和 Response一样, 在错误处理时使用 Error
func FavoriteAction(c *gin.Context) {
	//rawId := c.Query("user_id")
	token := c.Query("token")
	rawVideoId := c.Query("video_id")
	rawActionType := c.Query("action_type")
	if token == "" || rawVideoId == "" || rawActionType == "" {
		Error(c, 1, "参数获取失败")
		return
	}

	//user_id, _ := strconv.ParseInt(rawId, 10, 64)
	video_id, _ := strconv.ParseInt(rawVideoId, 10, 64)
	actiontype, _ := strconv.ParseInt(rawActionType, 10, 64)
	if actiontype != 1 && actiontype != 2 {
		Error(c, 1, "操作类型不符")
		return
	}

	claims := parseToken(token)
	if claims == nil {
		Error(c, 1, "身份鉴权失败")
		return
	}
	user_id, _ := strconv.ParseInt(claims.Id, 10, 64)

	// write to databse
	favoriteService := service.FavoriteService{
		User_id:     user_id,
		Video_id:    video_id,
		Action_type: actiontype,
	}
	if err := favoriteService.FavoriteAction(); err != nil {
		Error(c, 1, "点赞操作信息写入数据库出错")
		return
	} else {
		msg := ""
		if actiontype == 1 {
			msg = "点赞成功"
		} else {
			msg = "取消点赞成功"
		}
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  msg,
		})
	}
}

func FavoriteList(c *gin.Context) {
	rawId := c.Query("user_id")
	token := c.Query("token")

	claims := parseToken(token)
	if claims == nil || claims.Id != rawId {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "token鉴权失败"},
			VideoList: nil,
		})
	}

	user_id, _ := strconv.ParseInt(rawId, 10, 64)
	favoroteService := service.FavoriteService{
		User_id: user_id,
	}
	videoList, err := favoroteService.FavoriteList()
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "获取点赞列表失败"},
			VideoList: nil,
		})
	} else {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "获取点赞列表成功"},
			VideoList: videoList,
		})
	}
}
