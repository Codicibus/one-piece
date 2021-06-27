package service

import (
	"fmt"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
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
	if len(percent) != 0 {
		p = percent[0]
	}
	return p
}

func GetCpuState() []cpu.InfoStat {
	info, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	return info
}

func GetDiskState() map[string]interface{} {
	partitionStat, err := disk.Partitions(true)
	if err != nil {
		panic(err)
	}
	info := make(map[string]interface{})
	for _, pState := range partitionStat {
		usage, err := disk.Usage(pState.Mountpoint)
		if err != nil {
			fmt.Println(err)
			continue
		}
		info["usage"] = usage
		info[pState.Mountpoint] = pState
	}
	return info
}

func GetMemory() *mem.VirtualMemoryStat {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	return vmem
}

func GetNetStat() []sysnet.IOCountersStat {
	stat, err := sysnet.IOCounters(false)
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
