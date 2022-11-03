package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func GetResponse() *Response {
	return &Response{
		httpCode: http.StatusOK,
		result: &result{
			Code:    0,
			Message: "",
			Data:    nil,
			Cost:    "",
		},
	}
}
func BadRequest(ctx *gin.Context, data ...any) {
	GetResponse().WithDataFailure(http.StatusBadRequest, ctx, data)
}

// Success 业务成功响应
func Success(ctx *gin.Context, data ...any) {
	if data != nil {
		GetResponse().WithDataSuccess(ctx, data[0])
		return
	}
	GetResponse().Success(ctx)
}

// Fail 业务失败响应
func Fail(ctx *gin.Context, code int, message *string, data ...any) {
	var msg string
	if message == nil {
		msg = ""
	} else {
		msg = *message
	}
	if data != nil {
		GetResponse().WithData(data[0]).FailCode(ctx, code, msg)
		return
	}
	GetResponse().FailCode(ctx, code, msg)
}

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

type Response struct {
	httpCode int
	result   *result
}

// Fail 错误返回
func (r *Response) Fail(ctx *gin.Context) {
	r.SetCode(http.StatusInternalServerError)
	r.json(ctx)
}

// FailCode 自定义错误码返回
func (r *Response) FailCode(ctx *gin.Context, code int, msg ...string) {
	r.SetCode(code)
	if msg != nil {
		r.WithMessage(msg[0])
	}
	r.json(ctx)
}

// Success 正确返回
func (r *Response) Success(ctx *gin.Context) {
	r.SetCode(http.StatusOK)
	r.json(ctx)
}

// WithDataSuccess 成功后需要返回值
func (r *Response) WithDataSuccess(ctx *gin.Context, data interface{}) {
	r.SetCode(http.StatusOK)
	r.WithData(data)
	r.json(ctx)
}

func (r *Response) WithDataFailure(code int, ctx *gin.Context, data interface{}) {
	r.SetHttpCode(code)
	if data != nil {
		r.WithData(data)
	}
	r.json(ctx)
}

// SetCode 设置返回code码
func (r *Response) SetCode(code int) *Response {
	r.result.Code = code
	return r
}

// SetHttpCode 设置http状态码
func (r *Response) SetHttpCode(code int) *Response {
	r.httpCode = code
	return r
}

type defaultRes struct {
	Result any `json:"result"`
}

// WithData 设置返回data数据
func (r *Response) WithData(data interface{}) *Response {
	switch data.(type) {
	case string, int, bool:
		r.result.Data = &defaultRes{Result: data}
	default:
		r.result.Data = data
	}
	return r
}

// WithMessage 设置返回自定义错误消息
func (r *Response) WithMessage(message string) *Response {
	r.result.Message = message
	return r
}

// json 返回 gin 框架的 HandlerFunc
func (r *Response) json(ctx *gin.Context) {
	if len(strings.Trim(r.result.Message, " ")) > 0 {
		if r.result.Code >= http.StatusBadRequest {
			r.result.Message = fmt.Sprintf("msg: %s\r\nplease visit https://http.cat/%d for reason", r.result.Message, r.result.Code)
		}
	}

	r.result.Cost = time.Since(ctx.GetTime("requestStartTime")).String()
	ctx.AbortWithStatusJSON(r.httpCode, r.result)
}
