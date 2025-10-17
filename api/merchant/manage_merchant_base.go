package merchant

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// start fo front

// start fo manage
type MerchantBaseAdd struct {
	Id                      uint        `json:"id"         `                    // 自增
	MerchantIsSelfsupport   bool        `json:"merchant_is_selfsupport" `       // 是否自营(BOOL):1-默认;0-非默认
	MerchantName            string      `json:"merchant_name"     `             // 商户名称
	MerchantOpeningHours    string      `json:"merchant_opening_hours"     `    // 商户营业时间
	MerchantCloseHours      string      `json:"merchant_close_hours"     `      // 商户关闭时间
	MerchantIntl            string      `json:"merchant_intl"     `             // 商户手机号码国家代码
	MerchantTel             string      `json:"merchant_tel"     `              // 商户手机号码
	MerchantLogo            string      `json:"merchant_logo"     `             // 商户logo
	MerchantLongitude       string      `json:"merchant_longitude"     `        // 商户经度
	MerchantLatitude        string      `json:"merchant_latitude"     `         // 商户维度
	ProvinceId              int         `json:"province_id"     `               // 商户地址省id
	Province                string      `json:"province"     `                  // 商户地址省
	CityId                  int         `json:"city_id"     `                   // 商户地址市id
	City                    string      `json:"city"     `                      // 商户地址市
	CountyId                int         `json:"county_id"     `                 // 商户地址区id
	County                  string      `json:"county"     `                    // 商户地址区
	MerchantDistrictId      string      `json:"merchant_district_id"     `      // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea            string      `json:"merchant_area"     `             // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress string      `json:"merchant_detailed_address"     ` // 商户详细地址
	MerchantIsEnable        bool        `json:"merchant_is_enable"  `           // 是否启用(BOOL):1-启用;0-禁用
	CreateTime              *gtime.Time `json:"create_time"          `          // 创建时间
}

type MerchantBaseListReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/list" tags:"商户管理" method:"get" summary:"商户管理列表接口"`
	ml.BaseList

	MerchantName     string `json:"merchant_name"  type:"LIKE"     ` // 商户名称
	MerchantIsEnable bool   `json:"merchant_is_enable"  `            // 是否启用(BOOL):1-启用;0-禁用
}

type MerchantBaseAddReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/add" tags:"商户管理" method:"post" summary:"商户管理添加接口"`

	MerchantBaseAdd
}

type MerchantBaseEditReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/edit" tags:"商户管理" method:"post" summary:"商户管理编辑接口"`

	Id uint `json:"id"   ` // 商户管理编号`
	MerchantBaseAdd
}
type MerchantBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
type MerchantBaseRemoveReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/remove" tags:"商户管理" method:"post" summary:"商户管理删除接口"`
	Id     uint `json:"id"   ` // 商户管理编号
}

type MerchantBaseRemoveRes struct {
}
type MerchantBaseEditRes struct {
	MerchantId interface{} `json:"merchant_id"   dc:"商户信息"`
}
type MerchantBaseEditStateReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/editState" tags:"商户管理" method:"post" summary:"商户管理修改状态接口"`

	Id               uint `json:"id"   `                // 商户管理编号`
	MerchantIsEnable bool `json:"merchant_is_enable"  ` // 是否启用(BOOL):1-启用;0-禁用
}

type MerchantBaseEditStateRes struct {
	Id interface{} `json:"id"   dc:"商户管理信息"`
}
type MerchantBaseDetailReq struct {
	g.Meta `path:"/manage/merchant/merchantBase/detail" tags:"商户管理" method:"get" summary:"商户详情接口"`
	Id     uint `json:"id" v:"required#请输入商户ID"   dc:"商户ID"`
}

type MerchantBaseDetailRes struct {
	model.MerchantInfoVo
}
