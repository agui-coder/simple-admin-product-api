package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuPageLogic {
	return &GetSpuPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSpuPageLogic) GetSpuPage(req *types.SpuPageReq) (resp *types.SpuPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
