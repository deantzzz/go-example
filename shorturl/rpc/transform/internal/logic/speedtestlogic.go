package logic

import (
	"context"
	"time"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpeedTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSpeedTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpeedTestLogic {
	return &SpeedTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SpeedTestLogic) SpeedTest(in *transform.SpeedTestReq) (*transform.SpeedTestResp, error) {
	//unix := time.Unix(in.T, 0)
	//l.Info("speed test receive: " + unix.Format(time.RFC3339Nano))
	return &transform.SpeedTestResp{T: time.Now().UnixNano()}, nil
}
