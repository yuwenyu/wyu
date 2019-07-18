package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func M() gin.HandlerFunc {
	return func(c *gin.Context) {
		H := c.Request.Host
		c.Set("M",
			gin.H{
				"H":H,
				"Ln":mLn(H),
			},
		)
		c.Next()
	}
}

func S() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func mLn(domain string) string {
	strPreDomain := strings.Split(domain, ".")
	switch strPreDomain[0] {
	case "cn":
		return "cn"
	case "en":
		return "en"
	case "fr":
		return "fr"
	default:
		return "cn"
	}
}

