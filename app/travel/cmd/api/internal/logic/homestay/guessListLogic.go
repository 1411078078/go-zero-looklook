package homestay

import (
	"context"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
	"looklook/common/tool"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
)

type GuessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGuessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GuessListLogic {
	return GuessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GuessListLogic) GuessList(req types.GuessListReq) (*types.GuessListResp, error) {

	//@todo 先返回所有列表数据 ， 后续根据当前标签关联相同属性的民宿推荐给用户.

	var resp []types.Homestay

	list, err := l.svcCtx.HomestayModel.FindPageList(0, 5)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "req : %+v , err : %v", req, err)
	}

	if len(list) > 0 {
		for _, homestay := range list {
			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)

			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

			resp = append(resp, typeHomestay)
		}
	}

	return &types.GuessListResp{
		List: resp,
	}, nil
}
