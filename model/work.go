package model

import (
	"time"

	"zhuxuyang/spider/resource"
)

type Work struct {
	ID          int64
	SourceID    int64
	BindISBN    string
	CreatedAt   *time.Time
	DeletedTime int64
}

func (m *Work) TableName() string {
	return "work"
}

func SaveWork(work *Work) {
	dbWork := FindWork(work.SourceID)
	if dbWork.ID == 0 && !BookExisted(work.SourceID) {
		err := resource.GetDB().Save(work).Error
		if err != nil {
			resource.Logger.Error("SaveWork err" + err.Error())
		}
	}
}

func FindWork(SourceID int64) *Work {
	w := &Work{}
	err := resource.GetDB().Model(&Work{}).
		Where("source_id=?", SourceID).
		First(&w).Error
	if err != nil {
		resource.Logger.Error("FindWork err" + err.Error())
	}
	return w
}

func DeleteWork(work *Work) {
	if work.ID > 0 {
		err := resource.GetDB().Model(&Work{}).Where("id=?", work.ID).
			Update(map[string]interface{}{
				"deleted_time": time.Now().Unix(),
			}).Error
		if err != nil {
			resource.Logger.Error("DeleteWork err" + err.Error())
		}
	}
}

func FindWorkList(lastID int64) []*Work {
	list := make([]*Work, 0)
	err := resource.GetDB().Model(&Work{}).
		Where("id>?", lastID).
		Where("deleted_time=0").
		Order("id asc").
		Limit(1000).
		Scan(&list).Error
	if err != nil {
		resource.Logger.Error("FindWorkList err " + err.Error())
	}
	return list
}
