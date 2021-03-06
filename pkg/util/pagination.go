package util

import (
	"gin-blog/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// 获取分页
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result

}
