package dao

import (
	"../model"
)

type CommunityDao struct {
}

func (dao *CommunityDao) FindCommunityByIds(groupIds []int64) ([]model.Community, error) {
	communities := make([]model.Community, 0)
	err := DbEngine.In("id", groupIds).Find(&communities)

	return communities, err
}

func (dao *CommunityDao) FindCommunityByCommid(groupId string) (model.Community, error) {
	comm := model.Community{}
	_, err := DbEngine.Where("id=?", groupId).Get(&comm)

	return comm, err
}
