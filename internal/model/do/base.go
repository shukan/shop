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
	g.Meta                     `orm:"table:merchant_base, do:true"`
	Id                         any         // 自增主键
	MerchantIsSelfsupport      any         // 是否自营(BOOL):1-默认;0-非默认
	MerchantName               any         // 商户名称
	MerchantLogo               any         // 商户logo
	CreateTime                 *gtime.Time // 创建时间
	UpdateTime                 *gtime.Time // 更新时间
	DeletedTime                *gtime.Time // 删除时间
	LegalRepresentativeRame    any         // 法人姓名
	CorporateIDnumber          any         // 法人证件号
	CorporateIDphotoFront      any         // 法人证件照正面
	CorporateIDphotoBack       any         // 法人证件照反面
	MerchantIntl               any         // 商户手机号码国家代码
	MerchantTel                any         // 商户手机号码
	MerchantLongitude          any         // 商户经度
	MerchantLatitude           any         // 商户维度
	ProvinceId                 any         // 商户地址省id
	Province                   any         // 商户地址省
	CityId                     any         // 商户地址市id
	City                       any         // 商户地址市
	CountyId                   any         // 商户地址区id
	County                     any         // 商户地址区
	MerchantDistrictId         any         // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea               any         // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress    any         // 商户详细地址
	BusinessLicense            any         // 商户营业执照
	UnifiedSocialCreditCode    any         // 商户统一社会信用代码
	OtherInformation           any         // 商户其他资料图片
	MerchantReviewStatus       any         // 商户审核状态(0=待审核;1=审核通过;2=审核未通过)
	MerchantSettleinStatus     any         // 商户入驻状态(0=未入驻;1=已入住;2=入驻失败;3=退出入驻)
	MerchantReviewFailReason   any         // 商户审核失败原因
	MerchantReviewFailImage    any         // 商户审核失败图片
	MerchantExitSettleinReason any         // 商户退出入驻原因
	MerchantBanStatus          any         // 商户封禁状态(0=未封禁;1=已封禁;2=解除封禁审核中;3=解除封禁审核失败)
	MerchantBanReason          any         // 商户封禁原因
	MerchantBanImage           any         // 商户封禁图片
	MerchantApplyUnbanReason   any         // 商户申请解封原因
	MerchantApplyUnbanImage    any         // 商户申请解封图片
	MerchantUnbanFailReason    any         // 商户申请解封失败原因
	MerchantUnbanFailImage     any         // 商户申请解封失败图片
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