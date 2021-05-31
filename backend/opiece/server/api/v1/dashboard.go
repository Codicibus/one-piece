package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"opiece/server/service"
	"time"
)

func GetArticleStat(c *gin.Context) {
	conn := service.NewWSUpgrader(c)
	defer conn.Close()
	for {
		count, err := service.GetAllArticlesCount()
		if err != nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			continue
		}
		data, err := json.Marshal(map[string]int64{
			"article_count": count,
		})
		if err != nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("数据序列化失败"))
			continue
		}
		_ = conn.WriteMessage(websocket.TextMessage, data)
		time.Sleep(time.Second*10)
	}
}

func GetSysStat(c *gin.Context) {
 	conn := service.NewWSUpgrader(c)
	defer conn.Close()
	for {
		cpuPercent := service.GetCpuPercent()
		memStat := service.GetMemory()
		netStat := service.GetNetStat()
		data, err := json.Marshal(map[string]interface{}{
			"cpu_percent": cpuPercent,
			"mem_stat": memStat,
			"net_stat": netStat,
		})
		if err != nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("数据序列化失败"))
			continue
		}
		_ = conn.WriteMessage(websocket.TextMessage, data)
		time.Sleep(time.Second * 5)
	}
}