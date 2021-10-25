package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {

	//设置为测试模式
	gin.SetMode(gin.TestMode)

	//由于直接调用待测试函数，会造成循环引用，所有使用以上代码自己造一个router路由
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	//传给路由的参数数据
	body := `{
		"community_id": 1,
		"title": "test",
		"content": "just a test"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body))) //模拟请求
	//[]byte用来转换为字节切片，bytes.NewReader用来转换为io.Reader类型
	w := httptest.NewRecorder()  //造一个可接收响应的对象
	r.ServeHTTP(w, req)          //启动http服务(把请求记录到响应对象中) （r为待测试的函数对象）
	assert.Equal(t, 200, w.Code) //判断响应状态码是否等于200

	// 判断响应的内容是不是按预期返回了需要登录的错误

	// 方法1：判断响应内容(w.Body.String()以json字符串形式返回)中是不是包含指定的字符串
	//assert.Contains(t, w.Body.String(), "需要登录")

	// 方法2：将响应的内容反序列化到ResponseData 然后判断字段与预期是否一致
	res := new(ResponseDate) //ResponseData是最终返回的类型，res是拿到的响应数据
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err) //解析没有通过(Fatalf会自动将测试用例报错并停下来)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
