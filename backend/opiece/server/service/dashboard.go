package service

import (
	cpu "github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	sysnet "github.com/shirou/gopsutil/v3/net"
	"opiece/server/global"
	"opiece/server/model"
	"time"
)

func GetCpuPercent() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	p := 0.0
	if len(percent) != 0{
		p = percent[0]
	}
	return p
}

func GetMemory() *mem.VirtualMemoryStat{
	vmem, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	return vmem
}

func GetNetStat() []sysnet.IOCountersStat{
	stat, err :=  sysnet.IOCounters(false)
	if err != nil {
		panic(err)
	}
	return stat
}

func GetAllArticlesCount() (int64, error) {
	var count int64
	err := global.ODB.Model(&model.Article{}).Count(&count).Error
	return count, err
}