package interceptor

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	result2 "InvertedCow/model/vo"
	"InvertedCow/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

var releasePath = []string{"/account/signIn", "/account/signUp", "/account/code/send"}

// TokenAuthorize
//
//	@Description: token拦截器
//	@return gin.HandlerFunc
func TokenAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检验是否在放行名单
		path := c.Request.URL.Path
		for _, releaseStartPath := range releasePath {
			if strings.HasPrefix(path, releaseStartPath) {
				c.Next()
				return
			}
		}
		// 检验是否携带token
		r := result2.NewResult(c)
		token := c.Request.Header.Get("token")
		claims, err := utils.ParseToken(token)
		userInfo := &dto.UserInfo{
			ID:            claims.ID,
			Avatar:        claims.Avatar,
			LoginName:     claims.LoginName,
			Username:      claims.Username,
			Email:         claims.Email,
			FollowCount:   claims.FollowCount,
			FollowerCount: claims.FollowerCount,
		}
		if err != nil || userInfo == nil {
			r.Error(e.ErrSessionInvalid)
			c.Abort()
			return
		}
		if c.Keys == nil {
			c.Keys = make(map[string]interface{}, 1)
		}
		c.Keys["user"] = userInfo
	}
}
