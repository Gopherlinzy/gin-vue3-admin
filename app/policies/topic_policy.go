// Package policies 用户授权
package policies

import (
	"github.com/Gopherlinzy/gohub/app/models/topic"
	"github.com/Gopherlinzy/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
