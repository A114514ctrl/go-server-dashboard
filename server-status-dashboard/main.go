package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// 缓存 CPU 使用率（后台协程定时采集，避免每次 API 请求阻塞 1 秒）
var (
	cachedCPUPercent float64
	cpuMu            sync.RWMutex
)

func init() {
	go func() {
		for {
			percents, err := cpu.Percent(time.Second, false)
			if err == nil && len(percents) > 0 {
				cpuMu.Lock()
				cachedCPUPercent = math.Round(percents[0]*10) / 10
				cpuMu.Unlock()
			}
		}
	}()
}

// SystemInfo 系统硬件指标
type SystemInfo struct {
	CPUPercent    float64 `json:"cpu_percent"`
	Cores         int     `json:"cores"`
	MemoryTotalGB float64 `json:"memory_total_gb"`
	MemoryUsedGB  float64 `json:"memory_used_gb"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskTotalGB   float64 `json:"disk_total_gb"`
	DiskUsedGB    float64 `json:"disk_used_gb"`
	DiskPercent   float64 `json:"disk_percent"`
	UptimeHours   uint64  `json:"uptime_hours"`
}

// ServiceStatus 外部服务状态
type ServiceStatus struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Status    string `json:"status"` // "up" 或 "down"
	LatencyMs int64  `json:"latency_ms"`
}

// 要监控的外部服务列表
var services = []struct{ Name, URL string }{
	{"Baidu (百度)", "https://www.baidu.com"},
	{"Taobao (淘宝)", "https://www.taobao.com"},
	{"Google (谷歌)", "https://www.google.com"},
}

func main() {
	r := gin.Default()

	// 首页 - 提供大屏 HTML
	r.GET("/", func(c *gin.Context) {
		c.File("dashboard.html")
	})

	// API
	api := r.Group("/api")
	{
		api.GET("/system", handleSystemInfo)
		api.GET("/services", handleServices)
	}

	fmt.Println("=== Server Status Dashboard ===")
	fmt.Println("  Local:  http://localhost:8686")
	fmt.Println("  API:    http://localhost:8686/api/system")
	fmt.Println("  API:    http://localhost:8686/api/services")
	r.Run(":8686")
}

func handleSystemInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": collectSystemInfo(),
	})
}

func collectSystemInfo() SystemInfo {
	info := SystemInfo{}

	// CPU（使用后台采集的缓存值）
	info.Cores = runtime.NumCPU()
	cpuMu.RLock()
	info.CPUPercent = cachedCPUPercent
	cpuMu.RUnlock()

	// 内存
	if vmem, err := mem.VirtualMemory(); err == nil {
		info.MemoryTotalGB = math.Round(float64(vmem.Total)/1073741824*10) / 10
		info.MemoryUsedGB = math.Round(float64(vmem.Used)/1073741824*10) / 10
		info.MemoryPercent = math.Round(vmem.UsedPercent*10) / 10
	}

	// 磁盘（优先 / ，Windows 回退 C:）
	d, err := disk.Usage("/")
	if err != nil {
		d, err = disk.Usage("C:")
	}
	if err == nil {
		info.DiskTotalGB = math.Round(float64(d.Total)/1073741824*10) / 10
		info.DiskUsedGB = math.Round(float64(d.Used)/1073741824*10) / 10
		info.DiskPercent = math.Round(d.UsedPercent*10) / 10
	}

	// 运行时间
	if uptime, err := host.Uptime(); err == nil {
		info.UptimeHours = uptime / 3600
	}

	return info
}

func handleServices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": checkAllServices(),
	})
}

func checkAllServices() []ServiceStatus {
	var wg sync.WaitGroup
	results := make([]ServiceStatus, len(services))

	for i, svc := range services {
		wg.Add(1)
		go func(idx int, s struct{ Name, URL string }) {
			defer wg.Done()
			results[idx] = checkService(s)
		}(i, svc)
	}

	wg.Wait()
	return results
}

func checkService(svc struct{ Name, URL string }) ServiceStatus {
	ss := ServiceStatus{
		Name:   svc.Name,
		URL:    svc.URL,
		Status: "down",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, svc.URL, nil)
	if err != nil {
		return ss
	}

	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		return ss
	}
	defer resp.Body.Close()

	ss.LatencyMs = elapsed
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		ss.Status = "up"
	}

	return ss
}
