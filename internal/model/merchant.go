package model

import "golershop.cn/internal/model/entity"

type MerchantInfoVo struct {
	entity.MerchantBase
	Id                      int64   `json:"id"                        orm:"id"                        ` // 自增主键
	MerchantIsSelfsupport   int     `json:"merchant_is_selfsupport"   orm:"merchant_is_selfsupport"   ` // 是否自营(BOOL):1-默认;0-非默认
	MerchantName            string  `json:"merchant_name"             orm:"merchant_name"             ` // 商户名称
	MerchantOpeningHours    string  `json:"merchant_opening_hours"    orm:"merchant_opening_hours"    ` // 商户营业时间
	MerchantCloseHours      string  `json:"merchant_close_hours"      orm:"merchant_close_hours"      ` // 商户关闭时间
	MerchantIntl            string  `json:"merchant_intl"             orm:"merchant_intl"             ` // 商户手机号码国家代码
	MerchantTel             string  `json:"merchant_tel"              orm:"merchant_tel"              ` // 商户手机号码
	MerchantLogo            string  `json:"merchant_logo"             orm:"merchant_logo"             ` // 商户logo
	MerchantLongitude       float64 `json:"merchant_longitude"        orm:"merchant_longitude"        ` // 商户经度
	MerchantLatitude        float64 `json:"merchant_latitude"         orm:"merchant_latitude"         ` // 商户维度
	ProvinceId              int     `json:"province_id"               orm:"province_id"               ` // 商户地址省id
	Province                string  `json:"province"                  orm:"province"                  ` // 商户地址省
	CityId                  int     `json:"city_id"                   orm:"city_id"                   ` // 商户地址市id
	City                    string  `json:"city"                      orm:"city"                      ` // 商户地址市
	CountyId                int     `json:"county_id"                 orm:"county_id"                 ` // 商户地址区id
	County                  string  `json:"county"                    orm:"county"                    ` // 商户地址区
	MerchantDistrictId      string  `json:"merchant_district_id"      orm:"merchant_district_id"      ` // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea            string  `json:"merchant_area"             orm:"merchant_area"             ` // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress string  `json:"merchant_detailed_address" orm:"merchant_detailed_address" ` // 商户详细地址
	MerchantIsEnable        int     `json:"merchant_is_enable"        orm:"merchant_is_enable"        ` // 是否启用(BOOL):1-启用;0-禁用
}
