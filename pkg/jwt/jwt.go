package jwt

import (
	"errors"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {
	SignKey    []byte        // 秘钥
	MaxRefresh time.Duration // 刷新 token 的最大刷新时间
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(configYaml.Gohub_Config.App.Key),
		MaxRefresh: time.Duration(configYaml.Gohub_Config.JWT.MaxRefreshTime) * time.Minute,
	}
}

// ParserToken 解析 Token, 中间件调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 2. 解析错误
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors == jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}
	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JWTCustomClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	t := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	// 首次签名时间 > (当前时间 - 最大允许刷新时间)
	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// IssueToken 生成 Token, 在登录成功时调用
func (jwt *JWT) IssueToken(userID string, userName string) string {
	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime.Unix(),
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 签名生效时间
			IssuedAt:  jwtpkg.NewNumericDate(app.TimeNowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(expireAtTime),            // 签名过期时间
			Issuer:    configYaml.Gohub_Config.App.Name,               // 签名颁发者
		},
	}

	// 2. 根据 claims 生成 token 对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() time.Time {
	timezone := app.TimeNowInTimezone()

	var expireTime int64
	if configYaml.Gohub_Config.App.Debug {
		expireTime = configYaml.Gohub_Config.JWT.DebugExpireTime
	} else {
		expireTime = configYaml.Gohub_Config.JWT.ExpireTime
	}

	expire := time.Duration(expireTime) * time.Minute
	return timezone.Add(expire)
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwtpkg.Token) (interface{}, error) {
			return jwt.SignKey, nil
		})
}

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
