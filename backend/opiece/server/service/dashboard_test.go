package service

import (
	"fmt"
	"testing"
)

func TestGetCpuPercent(t *testing.T) {
	fmt.Println(GetCpuPercent())
}

func TestGetMemory(t *testing.T) {
	fmt.Println(GetMemory())
}

func TestGetNetStat(t *testing.T) {
	GetNetStat()
}