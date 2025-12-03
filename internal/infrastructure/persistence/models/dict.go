package models

import (
	"time"
)

const TableNameDictType = "dict_types"
const TableNameDictDatum = "dict_data"

// DictType mapped from table <dict_types>
type DictType struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code      string    `gorm:"column:code;not null;comment:字典类型编码" json:"code"`                       // 字典类型编码
	Name      string    `gorm:"column:name;not null;comment:字典key" json:"name"`                        // 字典key
	Status    int32     `gorm:"column:status;not null;default:1;comment:字典状态:1-正常,2-禁用" json:"status"` // 字典状态:1-正常,2-禁用
	Sort      int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                           // 排序
	Remark    string    `gorm:"column:remark;not null;comment:备注" json:"remark"`                       // 备注
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int32     `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName DictType's table name
func (*DictType) TableName() string {
	return TableNameDictType
}

// DictDatum mapped from table <dict_data>
type DictDatum struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TypeCode  string    `gorm:"column:type_code;not null;comment:字典类型编码" json:"type_code"`             // 字典类型编码
	Label     string    `gorm:"column:label;not null;comment:字典key" json:"label"`                      // 字典key
	Value     string    `gorm:"column:value;not null;comment:字典值" json:"value"`                        // 字典值
	Status    int32     `gorm:"column:status;not null;default:1;comment:字典状态:1-正常,2-禁用" json:"status"` // 字典状态:1-正常,2-禁用
	Sort      int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                           // 排序
	Remark    string    `gorm:"column:remark;not null;comment:备注" json:"remark"`                       // 备注
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int32     `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName DictDatum's table name
func (*DictDatum) TableName() string {
	return TableNameDictDatum
}
