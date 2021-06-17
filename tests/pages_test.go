package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPage(t *testing.T) {
	baseURL := "http://localhost:8000"

	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/4", 200},
		{"GET", "/articles/4/edit", 200},
		{"POST", "/articles/4", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/4/delete", 404},
	}

	for _, test := range tests {
		t.Logf("当前请求 URL: %v \n", test.url)
		var (
			resp *http.Response
			err  error
		)

		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}

		assert.NoError(t, err, "请求 "+test.url+" 报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" 应返回状态码"+strconv.Itoa(test.expected))
	}
}

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:8000"

	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/")

	assert.NoError(t, err, "有错误发生，err不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}

func TestAboutPage(t *testing.T) {
	baseURL := "http://localhost:8000"

	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/about")

	assert.NoError(t, err, "有错误发生，err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}
