package merchant

import (
	"context"

	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sMerchantBase struct{}

func init() {
	service.RegisterMerchantBase(NewMerchantBase())
}

func NewMerchantBase() *sMerchantBase {
	return &sMerchantBase{}
}

func (s *sMerchantBase) Find(ctx context.Context, in *do.MerchantBaseListInput) (out []*entity.MerchantBase, err error) {
	/*out, err = dao.MerchantBase.Find(ctx, in)

	return out, err*/
	return
}

func (s *sMerchantBase) List(ctx context.Context, in *do.MerchantBaseListInput) (out *do.MerchantBaseListOutput, err error) {
	out, err = dao.MerchantBase.List(ctx, in)

	return out, err
}

func (s *sMerchantBase) Add(ctx context.Context, in *do.MerchantBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.MerchantBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

func (s *sMerchantBase) Edit(ctx context.Context, in *do.MerchantBase) (affected int64, err error) {
	_, err = dao.MerchantBase.Edit(ctx, in.Id, in)

	if err != nil {
		return 0, err
	}
	return
}

func (s sMerchantBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	/*affected, err = dao.MerchantBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err*/
	return
}
