package logic

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/hewenyu/gozero-ent/rpc/internal/svc"
	"github.com/hewenyu/gozero-ent/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *pb.Empty) (*pb.EmptyReply, error) {
	// initialize table structure
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.WithContext(l.ctx).Errorw("database error", logx.Field("detail", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.EmptyReply{}, nil
}
