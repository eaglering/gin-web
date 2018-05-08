package middlewares

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"net/http"
	"gin-web/app/common/pools/redis"
	"strconv"
	"time"
	redis2 "github.com/garyburd/redigo/redis"
	"gin-web/app/common/libraries/feature"
	"encoding/json"
)

const (
	DefaultConfig = "app/modules/api/config/features.yaml"
)

type Env struct {
	OS 			string
	Version 	string
	User    	map[string]string
	Features	*feature.Features
}

func (m *Env) New() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.OS = c.Request.Header.Get("os")
		m.Version = c.Request.Header.Get("version")
		timestamp := c.Request.Header.Get("timestamp")
		token := c.Request.Header.Get("token")
		sign := c.Request.Header.Get("signature")

		if m.OS == "" || m.Version == "" || timestamp == "" || sign == "" {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		t, err := strconv.ParseInt(timestamp, 10, 64)
		now := time.Now().Unix()
		if err != nil || now - t < 10 * 60 {
			c.AbortWithStatus(http.StatusRequestTimeout)
			return
		}
		// Open Signature check
		//result := m.checkSign(fmt.Sprintf("os=%v&version=%v&timestamp=%v%v",
		//	m.OS, m.Version, timestamp, Config.Server.ApiSecret), sign)
		//if !result {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		if token != "" {
			m.User, _ = m.GetUserInfo(token)
		}
		m.Features = &feature.Features{}
		m.Features.Init(DefaultConfig, feature.Condition{
			OS: m.OS,
			Version: feature.Expression{Origin: m.Version},
		})
		c.Set("Env", m)
		c.Next()
	}
}

func (m *Env) checkSign(signString string, sign string) bool {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signString))
	cipher := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipher) == strings.ToLower(sign)
}

func (m *Env) GetUserInfo (token string) (result map[string]string, err error) {
	r := redis.Instance()
	conn := r.Get()
	defer conn.Close()
	t, err := redis2.Bytes(conn.Do("GET", token))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(t, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *Env) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m.User == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}