package merchant

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/merchant"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	MerchantBase = cMerchantBase{}
)

type cMerchantBase struct{}

// =================== 管理端使用 =========================

// 列表
func (c *cMerchantBase) List(ctx context.Context, req *merchant.MerchantBaseListReq) (res *merchant.MerchantBaseListRes, err error) {
	input := do.MerchantBaseListInput{}
	gconv.Scan(req, &input)
	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.MerchantBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// 添加
func (c *cMerchantBase) Add(ctx context.Context, req *merchant.MerchantBaseAddReq) (res *merchant.MerchantBaseEditRes, err error) {
	input := do.MerchantBase{}
	gconv.Scan(req, &input)
	var result, error = service.MerchantBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &merchant.MerchantBaseEditRes{
		MerchantId: result,
	}

	return
}

// 编辑
func (c *cMerchantBase) Edit(ctx context.Context, req *merchant.MerchantBaseEditReq) (res *merchant.MerchantBaseEditRes, err error) {

	input := do.MerchantBase{}
	gconv.Scan(req, &input)

	var result, error = service.MerchantBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &merchant.MerchantBaseEditRes{
		MerchantId: result,
	}

	return
}
