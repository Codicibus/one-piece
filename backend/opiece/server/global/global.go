package global

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"opiece/server/config"
	"time"

	"gorm.io/gorm"
)

var (
	OConfig  *config.Config
	OAuthJWT *jwt.GinJWTMiddleware
	ODB      *gorm.DB
)

var (
	MIMEImages = []string{"image/png"}
)

type OModel struct {
	ID        uint           `gorm:"primarykey" json:"-"` // 主键ID
	CreatedAt time.Time      `json:"-"`                   // 创建时间
	UpdatedAt time.Time      `json:"-"`                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`      // 删除时间
}
