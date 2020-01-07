package types

import (
	"time"

	"github.com/mkzilla/koala/pkg/types/config"
)

type Status struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	GroupID    int64      `json:"groupID" xorm:"groupID"`
	UserID     int64      `json:"userID" xorm:"userID"`
	Desc       string     `json:"desc" xorm:"'desc' varchar(5000)"`
	CreateTime *time.Time `json:"createTime" xorm:"'createTime' created"`
	UpdateTime *time.Time `json:"updateTime" xorm:"'updateTime' updated"`
}

func InsertStatus(s Status) (Status, error) {
	_, err := config.DBEngine.InsertOne(&s)
	return s, err
}

func GetRecentStatusByUser(id interface{}) (Status, error) {
	var s Status
	_, err := config.DBEngine.Where("userID = ?", id).Desc("id").Get(&s)
	return s, err
}
