package httpapis

import (
	"github.com/Golang-Tools/optparams"
)

type FeastHttpClientConfig struct {
	Query_Address string `json:"query_address" jsonschema:"required,description=连接服务的主机地址"`
	Query_Timeout int    `json:"query_timeout,omitempty" jsonschema:"description=请求服务的最大超时时间单位ms"`
}

// UntilEnd NewCtx方法的参数,用于设置ctx为不会超时
func WithConfig(config *FeastHttpClientConfig) optparams.Option[FeastHttpClientConfig] { //<- 2.定义可用的关键字参数项,一般命名上使用`with`开头
	return optparams.NewFuncOption(
		func(o *FeastHttpClientConfig) {
			o.Query_Address = config.Query_Address
			o.Query_Timeout = config.Query_Timeout
		})
}

// WithQueryAddress Init方法的参数,用于设置sdk请求的地址
// @params address ...string 连接服务的主机地址
func WithQueryAddress(address string) optparams.Option[FeastHttpClientConfig] {
	return optparams.NewFuncOption(
		func(o *FeastHttpClientConfig) {
			o.Query_Address = address
		})
}

// WithQueryTimeout sdk.Init方法的参数,用于设置sdk请求服务的最大超时时间
// @params wait int 请求服务的最大超时时间,单位ms
func WithQueryTimeout(timeout int) optparams.Option[FeastHttpClientConfig] {
	return optparams.NewFuncOption(
		func(o *FeastHttpClientConfig) {
			o.Query_Timeout = timeout
		})
}
