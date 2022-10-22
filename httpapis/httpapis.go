package httpapis

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Golang-Tools/optparams"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type FeastHttpClient struct {
	SDKConfig *FeastHttpClientConfig

	inited bool
	client *http.Client
}

func New(opts ...optparams.Option[FeastHttpClientConfig]) *FeastHttpClient {
	c := new(FeastHttpClient)
	c.client = &http.Client{}
	if len(opts) > 0 {
		c.Init(opts...)
	}
	return c
}

// Init 初始化
func (c *FeastHttpClient) Init(opts ...optparams.Option[FeastHttpClientConfig]) error {
	if c.inited {
		return ErrClientInited
	}
	optparams.GetOption(c.SDKConfig, opts...)
	c.inited = true
	return nil
}

func (c *FeastHttpClient) query(method, path string, body io.Reader, jsonbody bool) (*http.Response, error) {
	if !c.inited {
		return nil, ErrClientNotInit
	}
	requrl, err := url.JoinPath(c.SDKConfig.Query_Address, path)
	if err != nil {
		return nil, err
	}
	var ctx context.Context
	var cancel context.CancelFunc
	if c.SDKConfig.Query_Timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(c.SDKConfig.Query_Timeout)*time.Millisecond)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, method, requrl, body)
	if err != nil {
		return nil, err
	}
	if jsonbody {
		req.Header.Add("Content-Type", "application/json")
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, ErrHttpQueryNotOK
	}
	return res, nil
}

// Health 健康检查接口
func (c *FeastHttpClient) Health() error {
	res, err := c.query("GET", "get-online-features", nil, false)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return ErrHttpQueryNotOK
	}
	return nil
}

// Push 推送数据接口
// @params q *PushReqest push请求
func (c *FeastHttpClient) Push(q *PushReqest) error {
	pushrequestbytes, err := json.Marshal(q)
	if err != nil {
		return err
	}
	res, err := c.query("POST", "push", bytes.NewBuffer(pushrequestbytes), true)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return ErrHttpQueryNotOK
	}
	return nil
}

// GetOnlineFeatures 获取数据接口
// @params q *GetOnlineFeaturesHttpReq 请求数据
// @returns *GetOnlineFeaturesHttpResp 返回数据
func (c *FeastHttpClient) GetOnlineFeatures(q *GetOnlineFeaturesHttpReq) (*GetOnlineFeaturesHttpResp, error) {
	getrequeststr, err := json.MarshalToString(q)
	if err != nil {
		return nil, err
	}
	res, err := c.query("POST", "get-online-features", bytes.NewBufferString(getrequeststr), true)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, ErrHttpQueryNotOK
	}
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resjson := new(GetOnlineFeaturesHttpResp)
	err = json.Unmarshal(resbody, resjson)
	if err != nil {
		return nil, err
	}
	return resjson, nil
}

// DefaultSDK 默认的sdk,需要使用Init方法初始化
var DefaultSDK = New()
