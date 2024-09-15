package logic

import (
	"context"

	"server/file/api/internal/svc"
	"server/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileLogic {
	return &FileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileLogic) File(req *types.UploadFileRequest) (resp *types.UploadFileResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.UploadFileResponse{}, nil
}
