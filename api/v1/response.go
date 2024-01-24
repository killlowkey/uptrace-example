package v1

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func Ok(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"traceId": trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID(),
		"code":    200,
		"msg":     "success",
		"data":    data,
	})
}

func Fatal(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"traceId": trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID(),
		"code":    code,
		"msg":     msg,
	})
}
