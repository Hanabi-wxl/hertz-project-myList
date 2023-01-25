package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"hertz-mylist/base/result"
	"hertz-mylist/biz/model/gorm"
	"hertz-mylist/biz/model/user"
	"log"
	"net/http"
	"time"
)

var (
	HzJwtMw     *jwt.HertzJWTMiddleware
	IdentityKey = "claim"
)

type Claim struct {
	ID       int
	Username string
}

func JwtMwInit() {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		// 置所属领域名称
		Realm: "hertz jwt",
		// 用于设置签名密钥
		Key: []byte("hertz-mylist&&sinrewxljntm"),
		// 设置 token 过期时间
		Timeout: time.Hour * 8,
		// 设置最大 token 刷新时间
		MaxRefresh: time.Hour * 4,
		// 设置 token 的获取源
		TokenLookup: "header: Authorization",
		// 设置从 header 中获取 token 时的前缀
		TokenHeadName: "Bearer",
		// 用于设置检索身份的键
		IdentityKey: IdentityKey,
		// 登陆成功后为向 token 中添加自定义负载信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Claim); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 从 token 提取用户信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			claim := claims[IdentityKey].(map[string]interface{})
			return &Claim{
				ID:       int(claim["ID"].(float64)),
				Username: claim["Username"].(string),
			}
		},
		// 认证
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginUser user.User
			if err := c.BindAndValidate(&loginUser); err != nil {
				return "", result.NewIError(10000, "数据绑定错误")
			}
			password := loginUser.Password
			if err := gorm.DB.Where("username = ?", loginUser.Username).First(&loginUser).Error; err != nil {
				return nil, result.NewIError(10001, "用户名不存在")
			}
			if check := loginUser.CheckPassword(password); check {
				return Claim{
					ID:       loginUser.ID,
					Username: loginUser.Username,
				}, nil
			} else {
				return nil, result.NewIError(10002, "密码错误")
			}
		},
		// 鉴权
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			return true
		},
		// 登录响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":   code,
				"token":  token,
				"detail": "登陆成功",
			})
		},
		// 设置 jwt 校验流程发生错误时响应所包含的错误信息
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		// 无权限
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":   10011,
				"detail": "无权限或用户认证已过期",
				"data":   nil,
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	HzJwtMw = authMiddleware
}
