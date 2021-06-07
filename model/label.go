package model

import (
	"blogServer/database"
)

var LabelManager = LabelDao{}

type Label struct {
	LabelId   uint   `gorm:"label_id;primaryKey;not null"`
	LabelName string `gorm:"label_name;not null"`
}

func (article Label) TableName() string {
	return "label"
}

type LabelDao struct {
}

func (labelDao LabelDao) GetLabelsByArticleId(articleId uint) ([]Label, error) {
	var labels []Label
	result := database.MysqlDb.Joins("inner join article_label al on label.label_id = al.label_id and al.article_id = ?", articleId).Find(&labels)
	if result.Error != nil {
		return nil, result.Error
	}
	return labels, nil
}

// HasLabels 是否包含
func (labelDao LabelDao) HasLabels(labelIds []uint) bool {
	for _, value := range labelIds {
		if !labelDao.HasLabel(value) {
			return false
		}
	}
	return true
}

// HasLabel 是否包含
func (labelDao LabelDao) HasLabel(labelId uint) bool {
	var count int64
	result := database.MysqlDb.Model(&Label{}).Where("label_id = ?", labelId).Count(&count)
	if result.Error != nil {
		return false
	}
	return count > 0
}
