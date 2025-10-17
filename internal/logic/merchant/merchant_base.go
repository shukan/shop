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

// Find 读取信息
func (s *sMerchantBase) Find(ctx context.Context, in *do.MerchantBaseListInput) (out []*entity.MerchantBase, err error) {
	out, err = dao.MerchantBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sMerchantBase) List(ctx context.Context, in *do.MerchantBaseListInput) (out *do.MerchantBaseListOutput, err error) {
	out, err = dao.MerchantBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sMerchantBase) Add(ctx context.Context, in *do.MerchantBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.MerchantBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sMerchantBase) Edit(ctx context.Context, in *do.MerchantBase) (affected int64, err error) {
	_, err = dao.MerchantBase.Edit(ctx, in.Id, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sMerchantBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.MerchantBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// getById 获取商户详情
func (s *sMerchantBase) GetById(ctx context.Context, merchantId uint) (*entity.MerchantBase, error) {
	merchantInfo, err := dao.MerchantBase.Get(ctx, merchantId)
	if err != nil {
		return nil, err
	}

	if merchantInfo == nil {
		return nil, nil
	}

	return merchantInfo, nil
}
