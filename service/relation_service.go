package service

import (
	"InvertedCow/dao"
	e "InvertedCow/error"
	"InvertedCow/model/po"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"strconv"
)

const (
	// FollowListKey 关注列表的前缀key
	FollowListKey = "followlist-"
	// FollowerListKey 粉丝列表的前缀key
	FollowerListKey = "followerlist-"
)

type RelationService interface {
	// AddFollow 关注其他用户
	AddFollow(userId, userToId int) *e.Error
	// CancelFollow 取消关注
	CancelFollow(userId, userToId int) *e.Error
	// GetFollowList 获取关注列表
	GetFollowList(userId int) ([]*po.User, *e.Error)
	// GetFollowerList 获取好友列表
	GetFollowerList(userId int) ([]*po.User, *e.Error)
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

func (r *relationService) AddFollow(userId, userToId int) *e.Error {
	if err := r.userDao.AddUserFollow(r.db, userId, userToId); err != nil {
		return e.ErrAddFollowFailed
	}
	if exists, _ := r.redis.Exists(FollowListKey + strconv.Itoa(userId)).Result(); exists > 0 {
		r.redis.Del(FollowListKey + strconv.Itoa(userId))
	}
	if exists, _ := r.redis.Exists(FollowerListKey + strconv.Itoa(userToId)).Result(); exists > 0 {
		r.redis.Del(FollowerListKey + strconv.Itoa(userToId))
	}
	return nil
}

func (r *relationService) CancelFollow(userId, userToId int) *e.Error {
	if err := r.userDao.CancelUserFollow(r.db, userId, userToId); err != nil {
		return e.ErrCancelFollowFailed
	}
	if exists, _ := r.redis.Exists(FollowListKey + strconv.Itoa(userId)).Result(); exists > 0 {
		r.redis.Del(FollowListKey + strconv.Itoa(userId))
	}
	if exists, _ := r.redis.Exists(FollowerListKey + strconv.Itoa(userToId)).Result(); exists > 0 {
		r.redis.Del(FollowerListKey + strconv.Itoa(userToId))
	}
	return nil
}

func (r *relationService) GetFollowList(userId int) ([]*po.User, *e.Error) {
	var followList []*po.User
	var err error
	if exists, _ := r.redis.Exists(FollowListKey + strconv.Itoa(userId)).Result(); exists > 0 {
		err = r.redis.Get(FollowListKey + strconv.Itoa(userId)).Scan(&followList)
		if err == nil {
			return followList, nil
		}
	}
	followList, err = r.userDao.GetFollowListByUserId(r.db, userId)
	if err != nil {
		return nil, e.ErrUserUnknownError
	}
	for i, _ := range followList {
		followList[i].IsFollow = true
	}
	r.redis.Set(FollowListKey+strconv.Itoa(userId), followList, 0)
	return followList, nil
}

func (r *relationService) GetFollowerList(userId int) ([]*po.User, *e.Error) {
	var followerList []*po.User
	var err error
	if exists, _ := r.redis.Exists(FollowerListKey + strconv.Itoa(userId)).Result(); exists > 0 {
		err = r.redis.Get(FollowerListKey + strconv.Itoa(userId)).Scan(&followerList)
		if err == nil {
			return followerList, nil
		}
	}
	followerList, err = r.userDao.GetFollowerListByUserId(r.db, userId)
	if err != nil {
		return nil, e.ErrUserUnknownError
	}
	for i, _ := range followerList {
		followerList[i].IsFollow = true
	}
	r.redis.Set(FollowerListKey+strconv.Itoa(userId), followerList, 0)
	return followerList, nil
}
