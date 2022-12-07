// Package policies 用户授权
package policies

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/topic"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
