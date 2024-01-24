package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"strconv"
	"uptrace-example/biz"
)

type UserController struct {
	service biz.UserService
}

func NewUserController(store biz.UserService) *UserController {
	return &UserController{service: store}
}

func (u *UserController) GetUserById(c *gin.Context) {
	// 需要通过该 url 看，才能看到完整的链路信息
	// https://app.uptrace.dev/traces/3043/17acab0d42eb0050f6409c9bfee157c1?time_gte=20240122T113200&time_dur=3600&span=244382601524
	otelplay.PrintTraceID(c.Request.Context())

	idParam := c.Param("id")
	//转换为int64
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		// 附加错误信息，让 uptrace 收集
		_ = c.Error(err)
		Fatal(c, 400, err.Error())
		return
	}

	// 需要从请求中拿到 context，传播给下一层，才能形成完整的链路信息
	user, err := u.service.FindByID(c.Request.Context(), id)
	if err != nil {
		// 附加错误信息，让 uptrace 收集
		_ = c.Error(err)
		Fatal(c, 500, err.Error())
		return
	}

	// 可以通过 traceparent 和 tracestate 两个 header 来传递链路信息
	Ok(c, user)
}
