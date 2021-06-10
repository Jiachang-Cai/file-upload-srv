package middlewares

import (
	"errors"
	"fmt"

	"fmtsmsapi/utils/app"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
)

func ValidateToken(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("jwt-token")
		if jwtToken == "" {
			app.ErrorNoResponse(c, "401", "请重新登录")
			c.Abort()
			return
		}
		type MyCustomClaims struct {
			Id int64 `json:"id"`
			jwt.StandardClaims
		}
		token, err := jwt.ParseWithClaims(jwtToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			log.Error(app.LogError(c, err, errors.New(jwtToken)))
			app.ErrorNoResponse(c, "401", "服务错误")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Println(claims)
			//user, exist := cache.Memory.UserInfo.Get(claims.Id, models.DB.Self)
			//if !exist || user.Status == 0 {
			//	app.ErrorNoResponse(c, "401", "你已经被停用,没有权限访问该接口")
			//	c.Abort()
			//	return
			//}
			//app.CurrentUser = user
			//c.Set("user_id", fmt.Sprintf("%v", user.Id))
			c.Next()
		} else {
			app.ErrorNoResponse(c, "401", "没有权限访问该接口")
			c.Abort()
			return
		}
	}
}
