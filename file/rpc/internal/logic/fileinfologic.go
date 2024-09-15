package logic

import (
	"context"
	models "server/models/file"
	"strings"

	"server/file/rpc/internal/svc"
	"server/file/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileInfoLogic) FileInfo(in *proto.FileInfoRequest) (*proto.FileInfoResponse, error) {
	// 根据uuid获取文件信息
	file := models.FileModel{}
	err := l.svcCtx.DB.Take(&file, "uid = ?", in.FileUuid).Error
	if err != nil {
		return nil, err
	}
	strList := strings.Split(file.FileName, ".") // 		./xxx/xxx/xxx/xxx.txt
	fileInfo := proto.FileInfoResponse{
		FileName: file.FileName,
		FileHash: file.Hash,
		FilePath: file.Path,
		FileSize: file.Size,
	}
	if len(strList) > 0 {
		fileInfo.FileType = strList[len(strList)-1]
	}
	return &fileInfo, nil
}
