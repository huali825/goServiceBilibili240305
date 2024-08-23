package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

// MiddlewareBuilder 是一个结构体，用于构建中间件
type MiddlewareBuilder struct {
	Namespace  string // 命名空间
	Subsystem  string // 子系统
	Name       string // 名称
	Help       string // 帮助信息
	InstanceID string // 实例ID
}

// Build 是 MiddlewareBuilder 的一个方法，用于构建中间件
func (m *MiddlewareBuilder) Build() gin.HandlerFunc {
	// method：HTTP请求的方法，例如GET、POST、PUT、DELETE等。这个标签用于区分不同类型的HTTP请求。
	// pattern：请求的路径模式。例如，对于一个URL为/api/v1/users的请求，pattern可能被设置为/api/v1/users。这个标签用于区分不同的API端点。
	// status：HTTP响应的状态码。例如，200表示请求成功，404表示资源未找到，500表示服务器内部错误。这个标签用于区分不同类型的HTTP响应。
	labels := []string{"method", "pattern", "status"}
	// 创建一个 SummaryVec，用于记录响应时间
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_resp_time",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.9:   0.01,
			0.99:  0.005,
			0.999: 0.0001,
		},
	}, labels)
	// 注册 SummaryVec
	prometheus.MustRegister(summary)
	// 创建一个 Gauge，用于记录活跃请求
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_active_req",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
	})
	// 注册 Gauge
	prometheus.MustRegister(gauge)
	// 返回一个中间件函数
	return func(ctx *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		// 增加活跃请求计数
		gauge.Inc()
		// 在请求结束时执行
		defer func() {
			// 计算请求持续时间
			duration := time.Since(start)
			// 减少活跃请求计数
			gauge.Dec()
			// 404????
			// 获取请求路径
			pattern := ctx.FullPath()
			// 如果路径为空，则设置为 unknown
			if pattern == "" {
				pattern = "unknown"
			}
			// 记录响应时间
			summary.WithLabelValues(
				ctx.Request.Method,
				pattern,
				strconv.Itoa(ctx.Writer.Status()),
			).Observe(float64(duration.Milliseconds()))
		}()
		// 你最终就会执行到业务里面
		ctx.Next()
	}
}
