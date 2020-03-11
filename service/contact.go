package service

import (
	"../args/contact"
	"../args/user"
	"../dao"
	"../model"
	"../util"
	"../vo"
	"github.com/pkg/errors"
	"time"
)

type ContactService struct {
}

var contactUserDao *dao.ContactUserDao = &dao.ContactUserDao{}
var contactGroupDao *dao.ContactGroupDao = &dao.ContactGroupDao{}
var communityDao *dao.CommunityDao = &dao.CommunityDao{}

// 添加好友
func (*ContactService) Addfriend(userid int64, dstmobile string) error {
	// 查询是否存在该用户
	friendUserInfo, err := userDao.FindUserByMobile(dstmobile)
	if friendUserInfo.Id <= 0 {
		return errors.New("添加的用户不存在")
	}

	// 如果加自己为好友
	if err != nil || userid == friendUserInfo.Id {
		return errors.New("不能添加自己为好友")
	}

	// 判断是否已经加了好友
	// 查询 是否 有符合的 条件
	destid := friendUserInfo.Id
	contactUserInfo, _ := contactUserDao.FindFriendsByOwneridAndDestid(userid, destid)

	// 如果存在记录说明已经添加过好友
	if contactUserInfo.Id > 0 {
		return errors.New("已经添加过好友")
	}

	// 添加好友操作
	contactUserDao.Addfriends(userid, destid)
	return nil
}

// 加载 所有 好友信息
func (service *ContactService) LoadFriendAndGroup(arg *args_user.UserTokenArg) ([]vo.UserFriendsVO, error) {
	// 先判断是否登录

	// 2.查询好友列表
	contactUsers, err := contactUserDao.FindFriendsByOwnerid(arg.Userid)
	if nil != err {
		return nil, errors.New("查询好友失败:" + err.Error())
	}

	// 3.将好友的id串起来
	friendIds := make([]int64, 0)
	for _, contactUser := range contactUsers {
		friendIds = append(friendIds, contactUser.Dstobj)
	}

	users, err := userDao.FindUserByIds(friendIds)
	if nil != err {
		return nil, errors.New("获取好友信息失败:" + err.Error())
	}

	return contactUser2UserFriendsVO(arg.Userid, users), nil
}

// 将 用户信息 转成 好友列表信息
func contactUser2UserFriendsVO(Userid int64, users []model.User) []vo.UserFriendsVO {
	vos := make([]vo.UserFriendsVO, 0)

	for _, user := range users {
		friendVO := vo.UserFriendsVO{Userid, user.Id, "Friend", user.Nickname, user.Avatar}
		vos = append(vos, friendVO)
	}

	return vos
}

// 添加群组信息
func (service *ContactService) CreateGroup(arg *args_contact.GroupArg) (*model.Community, error) {
	community := model.Community{
		Id:       util.IntId(),
		Name:     arg.Groupname,
		Ownerid:  arg.Ownerid,
		Icon:     arg.Avatar,
		Memo:     arg.Groupdesc,
		Type:     arg.Grouptype,
		Createat: time.Now(),
	}

	groupinfo, err := contactGroupDao.CreateGroup(&community)
	if nil != err {
		return nil, errors.New("创建群组失败:" + err.Error())
	}

	return &groupinfo, nil
}

// 加载群聊
func (service *ContactService) LoadCommunity(arg *args_user.UserTokenArg) ([]model.Community, error) {
	// 先判断是否登录

	// 2.查询群列表
	contactGroups, err := contactUserDao.FindContactGroupsByOwnerid(arg.Userid)
	if nil != err {
		return nil, errors.New("查询群列表:" + err.Error())
	}

	// 3.将好友的id串起来
	groupIds := make([]int64, 0)
	for _, contactGroup := range contactGroups {
		groupIds = append(groupIds, contactGroup.Groupid)
	}

	communities, err := communityDao.FindCommunityByIds(groupIds)
	if nil != err {
		return nil, errors.New("获取好友信息失败:" + err.Error())
	}

	return communities, nil
}

// 添加群
func (service *ContactService) JoinCommunity(Userid int64, Dstid string) error {
	// 查询是否存在该用户
	contactGroupInfo, _ := communityDao.FindCommunityByCommid(Dstid)
	if contactGroupInfo.Id <= 0 {
		return errors.New("添加的群不存在")
	}

	// 判断是否已经加了群
	// 查询 是否 有符合的 条件
	destid := contactGroupInfo.Id
	contactUserInfo, _ := contactUserDao.FindCommByCommidAndUserid(Userid, destid)

	// 如果存在记录说明已经添加过群
	if contactUserInfo.Id > 0 {
		return errors.New("已经添加过群")
	}

	// 添加好友操作
	contactUserDao.JoinCommunity(Userid, destid)
	return nil
}
