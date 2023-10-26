package service

import (
	"InvertedCow/dao"
	e "InvertedCow/error"
	"InvertedCow/model/po"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const (
	// FollowListKey 关注列表的前缀key
	FollowListKey = "followlist"
	// FollowerListKey 粉丝列表的前缀key
	FollowerListKey = "followerlist"
)

type RelationService interface {
	// AddFollow 关注其他用户
	AddFollow(userId, userToId uint) *e.Error
	// CancelFollow 取消关注
	CancelFollow(userId, userToId uint) *e.Error
	// GetFollowList 获取关注列表
	GetFollowList(userId uint) ([]*po.User, *e.Error)
	// GetFollowerList 获取好友列表
	GetFollowerList(userId uint) ([]*po.User, *e.Error)
}

type relationService struct {
	db      *gorm.DB
	redis   *redis.Client
	userDao dao.UserDao
}

func NewRelationService(db *gorm.DB, redis *redis.Client, userDao dao.UserDao) RelationService {
	return &relationService{
		db:      db,
		redis:   redis,
		userDao: userDao,
	}
}

func (r *relationService) AddFollow(userId, userToId uint) *e.Error {
	if err := r.userDao.AddUserFollow(r.db, userId, userToId); err != nil {
		return e.ErrAddFollowFailed
	}
	r.redis.SAdd(fmt.Sprintf("%s-%d", FollowListKey, userId), userToId)
	r.redis.SAdd(fmt.Sprintf("%s-%d", FollowerListKey, userToId), userId)
	return nil
}

func (r *relationService) CancelFollow(userId, userToId uint) *e.Error {
	if err := r.userDao.CancelUserFollow(r.db, userId, userToId); err != nil {
		return e.ErrCancelFollowFailed
	}
	r.redis.SRem(fmt.Sprintf("%s-%d", FollowListKey, userId), userToId)
	r.redis.SRem(fmt.Sprintf("%s-%d", FollowerListKey, userToId), userId)
	return nil
}

func (r *relationService) GetFollowList(userId uint) ([]*po.User, *e.Error) {
	followList, err := r.userDao.GetFollowListByUserId(r.db, userId)
	if err != nil {
		return nil, e.ErrUserUnknownError
	}
	for i, _ := range followList {
		followList[i].IsFollow = true
		userToId := followList[i].ID
		r.redis.SAdd(fmt.Sprintf("%s-%d", FollowListKey, userId), userToId)
		r.redis.SAdd(fmt.Sprintf("%s-%d", FollowerListKey, userToId), userId)
	}
	return followList, nil
}

func (r *relationService) GetFollowerList(userId uint) ([]*po.User, *e.Error) {
	followerList, err := r.userDao.GetFollowerListByUserId(r.db, userId)
	if err != nil {
		return nil, e.ErrUserUnknownError
	}
	for i, _ := range followerList {
		followerList[i].IsFollower = true
		userToId := followerList[i].ID
		r.redis.SAdd(fmt.Sprintf("%s-%d", FollowerListKey, userId), userToId)
		r.redis.SAdd(fmt.Sprintf("%s-%d", FollowListKey, userToId), userId)
	}
	return followerList, nil
}
