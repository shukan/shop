// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// Base is the golang structure of table merchant_base for DAO operations like Where/Data.
type MerchantBase struct {
	g.Meta                  `orm:"table:merchant_base, do:true"`
	Id                      any         // 自增主键
	MerchantIsSelfsupport   any         // 是否自营(BOOL):1-默认;0-非默认
	MerchantName            any         // 商户名称
	MerchantOpeningHours    any         // 商户营业时间
	MerchantCloseHours      any         // 商户关闭时间
	MerchantIntl            any         // 商户手机号码国家代码
	MerchantTel             any         // 商户手机号码
	MerchantLogo            any         // 商户logo
	MerchantLongitude       any         // 商户经度
	MerchantLatitude        any         // 商户维度
	ProvinceId              any         // 商户地址省id
	Province                any         // 商户地址省
	CityId                  any         // 商户地址市id
	City                    any         // 商户地址市
	CountyId                any         // 商户地址区id
	County                  any         // 商户地址区
	MerchantDistrictId      any         // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea            any         // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress any         // 商户详细地址
	MerchantIsEnable        any         // 是否启用(BOOL):1-启用;0-禁用
	CreateTime              *gtime.Time // 创建时间
	UpdateTime              *gtime.Time // 更新时间
}
type MerchantBaseListInput struct {
	ml.BaseList
	Where MerchantBase // 查询条件
}

type MerchantBaseListOutput struct {
	Items   []*entity.MerchantBase // 列表
	Page    int                    // 分页号码
	Total   int                    // 总页数
	Records int                    // 数据总数
	Size    int                    // 单页数量
}

type MerchantBaseListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}