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
	CreatedAt time.Time      `json:"created_at"`          // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`      // 删除时间
}
