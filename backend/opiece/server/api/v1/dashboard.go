package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"opiece/server/service"
	"opiece/server/service/response"
	"time"
)

func GetSysStat(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Fail(c)
		return
	}
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