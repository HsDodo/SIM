package handler

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"os"
	"path"
	"server/common/response"
	"server/file/api/internal/logic"
	"server/file/api/internal/svc"
	"server/file/api/internal/types"
	models "server/models/file"
	"server/utils"
	"strings"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}
		// 获取 图片类型：avatar, group_avatar, chat
		imageType := r.FormValue("imageType")
		switch imageType {
		case "avatar", "group_avatar", "chat":
		default:
			response.Response(w, nil, errors.New("imageType只能为 avatar,group_avatar,chat"))
			return
		}

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			response.Response(w, nil, err)
			return
		}
		mSize := float64(fileHeader.Size) / float64(1024) / float64(1024) // 获取文件大小 单位MB

		if mSize > svcCtx.Config.FileSize {
			response.Response(w, nil, fmt.Errorf("图片大小超过限制，最大只能上传%.2fMB大小的图片", svcCtx.Config.FileSize))
			return
		}
		// 获取文件后缀，判断是否在白名单中
		nameList := strings.Split(fileHeader.Filename, ".")
		var suffix string
		if len(nameList) > 1 {
			suffix = nameList[len(nameList)-1]
		}
		if !utils.InList(svcCtx.Config.WhiteList, suffix) {
			response.Response(w, nil, errors.New("图片非法"))
			return
		}

		imageData, err := io.ReadAll(file)
		imageHash := utils.MD5(imageData)

		l := logic.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage(&req)

		var fileModel models.FileModel
		err = svcCtx.DB.Where("hash = ?", imageHash).First(&fileModel).Error
		if err == nil {
			resp.URL = fileModel.WebPath()
			logx.Infof("文件 %s hash重复", fileHeader.Filename)
			response.Response(w, resp, err)
			return
		}
		// 拼接路径
		dirPath := path.Join(svcCtx.Config.UploadDir, imageType)
		_, err = os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, os.ModePerm)
		}
		fileName := fileHeader.Filename
		newFileModel := models.FileModel{
			UserID:   req.UserID,
			FileName: fileName,
			Size:     fileHeader.Size,
			Hash:     imageHash,
			Uid:      uuid.New(),
		}
		newFileModel.Path = path.Join(dirPath, newFileModel.Uid.String()+"."+suffix)
		err = os.WriteFile(newFileModel.Path, imageData, os.ModePerm)
		if err != nil {
			response.Response(w, nil, err)
			return
		}
		// 文件信息入库
		err = svcCtx.DB.Create(&newFileModel).Error
		if err != nil {
			logx.Error(err)
			response.Response(w, nil, err)
			return
		}
		resp.URL = newFileModel.WebPath()
		response.Response(w, resp, err)
	}
}
