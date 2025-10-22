// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Base is the golang structure for table base.
type MerchantBase struct {
	Id                         int64       `json:"id"                            orm:"id"                            ` // 自增主键
	MerchantIsSelfsupport      int         `json:"merchant_is_selfsupport"       orm:"merchant_is_selfsupport"       ` // 是否自营(BOOL):1-默认;0-非默认
	MerchantName               string      `json:"merchant_name"                 orm:"merchant_name"                 ` // 商户名称
	MerchantLogo               string      `json:"merchant_logo"                 orm:"merchant_logo"                 ` // 商户logo
	CreateTime                 *gtime.Time `json:"create_time"                   orm:"create_time"                   ` // 创建时间
	UpdateTime                 *gtime.Time `json:"update_time"                   orm:"update_time"                   ` // 更新时间
	DeletedTime                *gtime.Time `json:"deleted_time"                  orm:"deleted_time"                  ` // 删除时间
	LegalRepresentativeRame    string      `json:"legal_representative_rame"     orm:"legal_representative_rame"     ` // 法人姓名
	CorporateIDnumber          string      `json:"corporate_i_dnumber"           orm:"corporate_IDnumber"            ` // 法人证件号
	CorporateIDphotoFront      string      `json:"corporate_i_dphoto_front"      orm:"corporate_IDphoto_front"       ` // 法人证件照正面
	CorporateIDphotoBack       string      `json:"corporate_i_dphoto_back"       orm:"corporate_IDphoto_back"        ` // 法人证件照反面
	MerchantIntl               string      `json:"merchant_intl"                 orm:"merchant_intl"                 ` // 商户手机号码国家代码
	MerchantTel                string      `json:"merchant_tel"                  orm:"merchant_tel"                  ` // 商户手机号码
	MerchantLongitude          float64     `json:"merchant_longitude"            orm:"merchant_longitude"            ` // 商户经度
	MerchantLatitude           float64     `json:"merchant_latitude"             orm:"merchant_latitude"             ` // 商户维度
	ProvinceId                 int         `json:"province_id"                   orm:"province_id"                   ` // 商户地址省id
	Province                   string      `json:"province"                      orm:"province"                      ` // 商户地址省
	CityId                     int         `json:"city_id"                       orm:"city_id"                       ` // 商户地址市id
	City                       string      `json:"city"                          orm:"city"                          ` // 商户地址市
	CountyId                   int         `json:"county_id"                     orm:"county_id"                     ` // 商户地址区id
	County                     string      `json:"county"                        orm:"county"                        ` // 商户地址区
	MerchantDistrictId         string      `json:"merchant_district_id"          orm:"merchant_district_id"          ` // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea               string      `json:"merchant_area"                 orm:"merchant_area"                 ` // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress    string      `json:"merchant_detailed_address"     orm:"merchant_detailed_address"     ` // 商户详细地址
	BusinessLicense            string      `json:"business_license"              orm:"business_license"              ` // 商户营业执照
	UnifiedSocialCreditCode    string      `json:"unified_social_credit_code"    orm:"unified_social_credit_code"    ` // 商户统一社会信用代码
	OtherInformation           string      `json:"other_information"             orm:"other_information"             ` // 商户其他资料图片
	MerchantReviewStatus       int         `json:"merchant_review_status"        orm:"merchant_review_status"        ` // 商户审核状态(0=待审核;1=审核通过;2=审核未通过)
	MerchantSettleinStatus     int         `json:"merchant_settlein_status"      orm:"merchant_settlein_status"      ` // 商户入驻状态(0=未入驻;1=已入住;2=入驻失败;3=退出入驻)
	MerchantReviewFailReason   string      `json:"merchant_review_fail_reason"   orm:"merchant_review_fail_reason"   ` // 商户审核失败原因
	MerchantReviewFailImage    string      `json:"merchant_review_fail_image"    orm:"merchant_review_fail_image"    ` // 商户审核失败图片
	MerchantExitSettleinReason string      `json:"merchant_exit_settlein_reason" orm:"merchant_exit_settlein_reason" ` // 商户退出入驻原因
	MerchantBanStatus          int         `json:"merchant_ban_status"           orm:"merchant_ban_status"           ` // 商户封禁状态(0=未封禁;1=已封禁;2=解除封禁审核中;3=解除封禁审核失败)
	MerchantBanReason          string      `json:"merchant_ban_reason"           orm:"merchant_ban_reason"           ` // 商户封禁原因
	MerchantBanImage           string      `json:"merchant_ban_image"            orm:"merchant_ban_image"            ` // 商户封禁图片
	MerchantApplyUnbanReason   string      `json:"merchant_apply_unban_reason"   orm:"merchant_apply_unban_reason"   ` // 商户申请解封原因
	MerchantApplyUnbanImage    string      `json:"merchant_apply_unban_image"    orm:"merchant_apply_unban_image"    ` // 商户申请解封图片
	MerchantUnbanFailReason    string      `json:"merchant_unban_fail_reason"    orm:"merchant_unban_fail_reason"    ` // 商户申请解封失败原因
	MerchantUnbanFailImage     string      `json:"merchant_unban_fail_image"     orm:"merchant_unban_fail_image"     ` // 商户申请解封失败图片
}
