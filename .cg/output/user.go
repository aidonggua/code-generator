package domain

import "time"

type User struct {
	Id        int64     `json:"id"`         // 主键
	Name      string    `json:"name"`       // 姓名
	Age       int       `json:"age"`        // 年龄
	Birth     time.Time `json:"birth"`      // 生日
	GmtCreate time.Time `json:"gmt_create"` // 创建时间
	GmtModify time.Time `json:"gmt_modify"` // 修改时间
	Deleted   int64     `json:"deleted"`    // 删除标记
}

func (*User) TableName() string {
	return "user"
}
