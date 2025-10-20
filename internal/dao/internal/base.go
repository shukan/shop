// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

// BaseDao is the data access object for the table merchant_base.
type BaseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BaseColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BaseColumns defines and stores column names for the table merchant_base.
type BaseColumns struct {
	Id                      string // 自增主键
	MerchantIsSelfsupport   string // 是否自营(BOOL):1-默认;0-非默认
	MerchantName            string // 商户名称
	MerchantOpeningHours    string // 商户营业时间
	MerchantCloseHours      string // 商户关闭时间
	MerchantIntl            string // 商户手机号码国家代码
	MerchantTel             string // 商户手机号码
	MerchantLogo            string // 商户logo
	MerchantLongitude       string // 商户经度
	MerchantLatitude        string // 商户维度
	ProvinceId              string // 商户地址省id
	Province                string // 商户地址省
	CityId                  string // 商户地址市id
	City                    string // 商户地址市
	CountyId                string // 商户地址区id
	County                  string // 商户地址区
	MerchantDistrictId      string // 商户地址省市区id拼接字符串(英文逗号拼接)
	MerchantArea            string // 商户地址省市区拼接字符串(英文逗号拼接)
	MerchantDetailedAddress string // 商户详细地址
	MerchantIsEnable        string // 是否启用(BOOL):1-启用;0-禁用
	CreateTime              string // 创建时间
	UpdateTime              string // 更新时间
}

// baseColumns holds the columns for the table merchant_base.
var baseColumns = BaseColumns{
	Id:                      "id",
	MerchantIsSelfsupport:   "merchant_is_selfsupport",
	MerchantName:            "merchant_name",
	MerchantOpeningHours:    "merchant_opening_hours",
	MerchantCloseHours:      "merchant_close_hours",
	MerchantIntl:            "merchant_intl",
	MerchantTel:             "merchant_tel",
	MerchantLogo:            "merchant_logo",
	MerchantLongitude:       "merchant_longitude",
	MerchantLatitude:        "merchant_latitude",
	ProvinceId:              "province_id",
	Province:                "province",
	CityId:                  "city_id",
	City:                    "city",
	CountyId:                "county_id",
	County:                  "county",
	MerchantDistrictId:      "merchant_district_id",
	MerchantArea:            "merchant_area",
	MerchantDetailedAddress: "merchant_detailed_address",
	MerchantIsEnable:        "merchant_is_enable",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
}

// NewBaseDao creates and returns a new DAO object for table data access.
func NewBaseDao(handlers ...gdb.ModelHandler) *BaseDao {
	return &BaseDao{
		group:    "merchant",
		table:    "merchant_base",
		columns:  baseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BaseDao) Columns() BaseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BaseDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Find 查询数据
func (dao *BaseDao) Find(ctx context.Context, in *do.MerchantBaseListInput) (out []*entity.MerchantBase, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 对象转换
	if err := query.Scan(&out); err != nil {
		return out, err
	}

	return out, nil
}

// List 分页读取
func (dao *BaseDao) List(ctx context.Context, in *do.MerchantBaseListInput) (out *do.MerchantBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).Where("deleted_time IS NULL").OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.MerchantBaseListOutput{}
	out.Page = in.Page
	out.Size = in.Size

	// 查询记录总数
	count, err1 := query.Count()
	if err1 != nil {
		return nil, err1
	}

	out.Records = count
	out.Total = int(math.Ceil(float64(count) / float64(out.Size)))

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 分页
	query = query.Page(in.Page, in.Size)

	// 对象转换
	if err := query.Scan(&out.Items); err != nil {
		return out, err
	}

	return out, nil
}

// Add 新增
func (dao *BaseDao) Add(ctx context.Context, in *do.MerchantBase) (lastInsertId int64, err error) {
	data := do.MerchantBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *BaseDao) Edit(ctx context.Context, id any, in *do.MerchantBase) (int64, error) {
	data := do.MerchantBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// Remove 根据主键删除
func (dao *BaseDao) Remove(ctx context.Context, id any) (int64, error) {
	//res, err := dao.Ctx(ctx).WherePri(id).Delete()
	res, err := dao.Ctx(ctx).WherePri(id).Update(
		g.Map{
			"deleted_time": gdb.Raw("NOW()"),
		})
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Gets 读取多条记录
func (dao *BaseDao) Gets(ctx context.Context, id any) (entitys []*entity.MerchantBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Get 读取一条记录
func (dao *BaseDao) Get(ctx context.Context, id any) (one *entity.MerchantBase, err error) {
	var entitys []*entity.MerchantBase
	entitys, err = dao.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(entitys) > 0 {
		one = entitys[0]
	}

	return one, err
}