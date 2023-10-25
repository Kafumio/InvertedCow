package dao

import (
	"InvertedCow/model/po"
	"errors"
	"gorm.io/gorm"
)

type UserDao interface {
	// InsertUser 创建用户
	InsertUser(db *gorm.DB, user *po.User) error
	// UpdateUser 更新用户，不更新零值
	UpdateUser(db *gorm.DB, user *po.User) error
	// DeleteUserByID 删除用户
	DeleteUserByID(db *gorm.DB, id uint) error
	// GetUserByID 通过用户id获取用户, 找不到会返回异常而不是nil
	GetUserByID(db *gorm.DB, id uint) (*po.User, error)
	// GetUserByEmail 根据用户邮箱获取用户信息
	GetUserByEmail(db *gorm.DB, email string) (*po.User, error)
	// GetUserByLoginName 根据用户登录名称获取用户
	GetUserByLoginName(db *gorm.DB, loginName string) (*po.User, error)
	// CheckEmail 检测邮箱是否已经存在
	CheckEmail(db *gorm.DB, email string) (bool, error)
	// CheckLoginName 检测loginname是否存在
	CheckLoginName(db *gorm.DB, loginname string) (bool, error)
	// AddUserFollow 用户关注，增加用户关注数以及被关注对象粉丝数
	AddUserFollow(db *gorm.DB, userId, userToId int) error
	// CancelUserFollow 取消关注，减少用户关注数以及被关注对象粉丝数
	CancelUserFollow(db *gorm.DB, userId, userToId int) error
	// GetFollowListByUserId 根据用户Id获取用户的关注列表
	GetFollowListByUserId(db *gorm.DB, userId int) (followList []*po.User, err error)
	// GetFollowerListByUserId 根据用户Id获取用户的粉丝列表
	GetFollowerListByUserId(db *gorm.DB, userId int) (followerList []*po.User, err error)
}

type userDao struct {
}

func NewUserDao() UserDao {
	return &userDao{}
}

func (u *userDao) InsertUser(db *gorm.DB, user *po.User) error {
	return db.Create(user).Error
}

func (u *userDao) UpdateUser(db *gorm.DB, user *po.User) error {
	return db.Model(user).Updates(user).Error
}

func (u *userDao) DeleteUserByID(db *gorm.DB, id uint) error {
	return db.Delete(&po.User{}, id).Error
}

func (u *userDao) GetUserByID(db *gorm.DB, id uint) (*po.User, error) {
	user := &po.User{}
	err := db.First(&user, id).Error
	return user, err
}

func (u *userDao) GetUserByEmail(db *gorm.DB, email string) (*po.User, error) {
	user := &po.User{}
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userDao) GetUserByLoginName(db *gorm.DB, loginName string) (*po.User, error) {
	user := &po.User{}
	err := db.Where("login_name = ?", loginName).First(&user).Error
	return user, err
}

func (u *userDao) CheckEmail(db *gorm.DB, email string) (bool, error) {
	user := &po.User{}
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return user.ID != 0, nil
}

func (s *userDao) CheckLoginName(db *gorm.DB, loginname string) (bool, error) {
	user := &po.User{}
	err := db.Where("login_name = ?", loginname).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *userDao) AddUserFollow(db *gorm.DB, userId, userToId int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE users SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE users SET follower_count=follower_count+1 WHERE id = ?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations` (`user_id`,`follow_id`) VALUES (?,?)", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *userDao) CancelUserFollow(db *gorm.DB, userId, userToId int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE users SET follow_count=follow_count-1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE users SET follower_count=follower_count-1 WHERE id = ?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_id = ? AND follow_id = ?", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *userDao) GetFollowListByUserId(db *gorm.DB, userId int) (followList []*po.User, err error) {
	followList = []*po.User{}
	if err = db.Raw("SELECT u.* FROM user_relations r, users u WHERE r.user_id = ? AND r.follow_id = u.id", userId).Scan(&followList).Error; err != nil {
		return followList, err
	}
	return followList, nil
}

func (s *userDao) GetFollowerListByUserId(db *gorm.DB, userId int) (followerList []*po.User, err error) {
	followerList = []*po.User{}
	if err = db.Raw("SELECT u.* FROM user_relations r, users u WHERE r.follow_id = ? AND r.user_id = u.id", userId).Scan(followerList).Error; err != nil {
		return followerList, err
	}
	return followerList, nil
}
