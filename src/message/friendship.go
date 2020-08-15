package message

import (
	"JuneGoBlog/src/dao"
	"github.com/gin-gonic/gin"
)

type FriendShipListResp struct {
	Header         BaseRespHeader       `json:"header"`         // 响应头
	Total          int                  `json:"total"`          // 友链总数
	FriendShipList []dao.FriendShipLink `json:"friendShipList"` // 友链列表
}

type FriendShipListReq struct{}

// 申请添加友链的请求头格式
type FriendApplicationReq struct {
	SiteName string `form:"siteName" json:"siteName"` // 网站名称（必填）
	SiteLink string `form:"siteLink" json:"siteLink"` // 网站链接（必填）
	ImgLink  string `form:"imgLink" json:"imgLink"`   // 网站图标链接
	Intro    string `form:"intro" json:"intro"`       // 网站简介
}

// 申请添加友链响应格式
type FriendApplicationResp struct {
	Header BaseRespHeader `json:"header"`
}

// 友链申请审批请求
type FriendApprovalReq struct {
	FriendshipID int `json:"id"`
	Result       int `json:"result"`
}

// 友链审批响应
type FriendApprovalResp struct {
	Header BaseRespHeader `json:"header"`
}

// 删除友链请求格式
type FriendDeleteReq struct {
	ID int `form:"id" json:"id"` // 要删除的友链ID
}

// 删除友链响应格式
type FriendDeleteResp struct {
	Header BaseRespHeader `json:"header"`
}

func (far FriendApprovalReq) JSON(ctx *gin.Context, jsonReq *BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

func (flr FriendShipListReq) JSON(ctx *gin.Context, jsonReq *BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

func (far FriendApplicationReq) JSON(ctx *gin.Context, jsonReq *BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}

func (flr FriendDeleteReq) JSON(ctx *gin.Context, jsonReq *BaseReqInter) error {
	return ctx.ShouldBindJSON(&jsonReq)
}
