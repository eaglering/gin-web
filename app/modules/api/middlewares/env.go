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
		m.OS = c.Request.Header.Get("OS")
		m.Version = c.Request.Header.Get("Version")
		timestamp := c.Request.Header.Get("Timestamp")
		token := c.Request.Header.Get("Token")
		sign := c.Request.Header.Get("Signature")

		if m.OS == "" || m.Version == "" || timestamp == "" || token == "" || sign == "" {
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
		m.User, _ = m.getUserInfo(token)
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

func (m *Env) getUserInfo (token string) (map[string]string, error) {
	r := redis.Instance()
	conn := r.Get()
	defer conn.Close()
	return redis2.StringMap(conn.Do("GET", token))
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