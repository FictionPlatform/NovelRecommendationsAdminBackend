package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Trace 链路追踪中间件 (OpenTelemetry 版本)
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 获取 Tracer
		// "gin-middleware" 是 instrumentation name，用于标识追踪来源
		tracer := otel.Tracer("gin-middleware")

		// 2. 从请求头中提取追踪上下文
		// 这相当于 OpenTracing 的 Extract 操作
		// 如果请求头中没有追踪信息，ctx.Request.Context() 将返回一个空的 Context
		reqCtx := otel.GetTextMapPropagator().Extract(
			ctx.Request.Context(),
			propagation.HeaderCarrier(ctx.Request.Header),
		)

		// 3. 开启一个新的 Span
		// tracer.Start 会自动判断 reqCtx 中是否有父 Span：
		//   - 如果有，则创建子 Span (ChildOf)
		//   - 如果没有，则创建根 Span
		opName := ctx.Request.URL.Path
		newCtx, span := tracer.Start(reqCtx, opName)

		// 4. 确保 Span 结束
		defer span.End()

		// 5. 将 Span 存储到 gin.Context 中 (保持原代码逻辑，方便业务代码手动获取)
		// 注意：类型由 opentracing.Span 变为 trace.Span
		ctx.Set("traceSpan", span)

		// 6. 【关键步骤】将包含 Span 的新 Context 注入到 Request 中
		// 这样后续的 Controller、Service、DB 调用才能拿到链路上下文
		ctx.Request = ctx.Request.WithContext(newCtx)

		ctx.Next()
	}
}
