package api

import (
	"github.com/gin-gonic/gin"
	"user/app/business"
	"user/app/http/params"
	"user/app/models"
	"user/lib"
)

func Info(ctx *gin.Context) {
	lib.Ok(ctx, models.AuthUser)
}

// UpdateInfo 改个人信息
func UpdateInfo(ctx *gin.Context) {
	var query params.UserInfo
	err := ctx.Bind(&query)
	if err != nil {
		panic(err)
	}
	lib.Ok(ctx, business.UpdateUserInfo(query))
}

// UploadAvatar 上传头像
func UploadAvatar(ctx *gin.Context) {
	lib.Ok(ctx, business.UploadAvatar(ctx))
}
