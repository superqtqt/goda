package https

import (
	"github.com/go-resty/resty/v2"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		url     string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test_get", args{url: "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get(tt.args.url, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPost(t *testing.T) {
	type args struct {
		url     string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Post(tt.args.url, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithBody(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithBody(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHeader(t *testing.T) {
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHeader(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHeaders(t *testing.T) {
	type args struct {
		headers map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHttpCodeToError(t *testing.T) {
	type args struct {
		f func(response *resty.Response) error
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHttpCodeToError(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHttpCodeToError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPathParam(t *testing.T) {
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPathParam(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPathParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPathParams(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPathParams(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPathParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithQueryParam(t *testing.T) {
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithQueryParam(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithQueryParams(t *testing.T) {
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithQueryParams(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithResult(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithResult(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetryCondition(t *testing.T) {
	type args struct {
		f func(response *resty.Response, err error) bool
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRetryCondition(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRetryCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRequest(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name  string
		args  args
		want  *resty.Request
		want1 *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateRequest(tt.args.options...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("generateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
