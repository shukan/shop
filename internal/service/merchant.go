package service

import (
	"context"

	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IMerchantBase interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.MerchantBaseListInput) (out []*entity.MerchantBase, err error)
		// List 分页读取
		List(ctx context.Context, in *do.MerchantBaseListInput) (out *do.MerchantBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.MerchantBase) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.MerchantBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
)

var (
	localMerchantBase IMerchantBase
)

func MerchantBase() IMerchantBase {
	if localMerchantBase == nil {
		panic("implement not found for interface IMerchantBase, forgot register?")
	}
	return localMerchantBase
}

func RegisterMerchantBase(i IMerchantBase) {
	localMerchantBase = i
}
