package global

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"opiece/server/config"
	"time"

	"gorm.io/gorm"
)

var (
	OConfig *config.Config
	OAuthJWT *jwt.GinJWTMiddleware
)


type OModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
