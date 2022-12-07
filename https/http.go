package https

import (
	"github.com/go-resty/resty/v2"
	"github.com/superqtqt/goda/cast"
)

var GlobalDefaultConfig *Config

type Config struct {
	headers         map[string]string
	body            interface{}
	result          interface{}
	httpCodeToError func(response *resty.Response) error
	queryParams     map[string]string
	pathParams      map[string]string
	retryConditions []func(response *resty.Response, err error) bool
}

type Option func(c *Config)

func generateRequest(options ...Option) (*resty.Request, *Config) {
	c := &Config{}
	for _, option := range options {
		option(c)
	}
	request := resty.New().R()
	if c.headers != nil {
		request = request.SetHeaders(c.headers)
	}
	if c.body != nil {
		request = request.SetBody(c.body)
	}
	if c.result != nil {
		request = request.SetResult(c.result)
	}
	if c.queryParams != nil {
		request = request.SetQueryParams(c.queryParams)
	}
	if c.pathParams != nil {
		request = request.SetPathParams(c.pathParams)
	}
	if c.retryConditions != nil {
		for _, f := range c.retryConditions {
			request = request.AddRetryCondition(f)
		}
	}
	return request, c
}

func Get(url string, options ...Option) error {
	request, c := generateRequest(options...)

	response, err := request.Get(url)
	if err != nil {
		return err
	}
	if c.httpCodeToError != nil {
		err = c.httpCodeToError(response)
		if err != nil {
			return err
		}
	}
	return nil
}

func Post(url string, options ...Option) error {
	request, c := generateRequest(options...)

	response, err := request.Post(url)
	if err != nil {
		return err
	}
	if c.httpCodeToError != nil {
		err = c.httpCodeToError(response)
		if err != nil {
			return err
		}
	}
	return nil
}

func WithHeader(k string, v interface{}) Option {
	return func(c *Config) {
		if c.headers == nil {
			c.headers = make(map[string]string)
		}
		c.headers[k] = cast.ToString(v)
	}
}

func WithHeaders(headers map[string]interface{}) Option {
	return func(c *Config) {
		if c.headers == nil {
			c.headers = make(map[string]string)
		}
		for k, v := range headers {
			c.headers[k] = cast.ToString(v)
		}
	}
}
func WithBody(v interface{}) Option {
	return func(c *Config) {
		c.body = v
	}
}

func WithResult(v interface{}) Option {
	return func(c *Config) {
		c.result = v
	}
}

func WithHttpCodeToError(f func(response *resty.Response) error) Option {
	return func(c *Config) {
		c.httpCodeToError = f
	}
}

func WithQueryParam(k string, v interface{}) Option {
	return func(c *Config) {
		if c.queryParams == nil {
			c.queryParams = make(map[string]string)
		}
		c.queryParams[k] = cast.ToString(v)
	}
}
func WithQueryParams(params map[string]interface{}) Option {
	return func(c *Config) {
		if c.queryParams == nil {
			c.queryParams = make(map[string]string)
		}
		for k, v := range params {
			c.queryParams[k] = cast.ToString(v)
		}
	}
}

func WithPathParam(k string, v interface{}) Option {
	return func(c *Config) {
		if c.pathParams == nil {
			c.pathParams = make(map[string]string)
		}
		c.pathParams[k] = cast.ToString(v)
	}
}

func WithPathParams(m map[string]interface{}) Option {
	return func(c *Config) {
		if c.pathParams == nil {
			c.pathParams = make(map[string]string)
		}
		for k, v := range m {
			c.pathParams[k] = cast.ToString(v)
		}
	}
}

func WithRetryCondition(f func(response *resty.Response, err error) bool) Option {
	return func(c *Config) {
		if c.retryConditions == nil {
			c.retryConditions = make([]func(response *resty.Response, err error) bool, 0)
		}
		c.retryConditions = append(c.retryConditions, f)
	}
}
