package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var responseMessage = map[int]string{
	0:     "成功",
	10000: "内部错误",

	// 输入错误
	10001: "参数不正确",
	10002: "认证不通过",

	// 处理错误
	11001: "数据不存在",
	11002: "数据有冲突",
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func resp(w http.ResponseWriter, status int, data interface{}) {
	msg, exist := responseMessage[status]
	if !exist {
		msg = "未知错误类型"
	}
	if err, ok := data.(error); ok {
		data = err.Error()
	}

	ret := response{
		status,
		msg,
		data,
	}
	rt, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rt)
}

func respRaw(w http.ResponseWriter, status int, data interface{}) {
	var rt []byte
	if err, ok := data.(error); ok {
		data = err.Error()
	}

	rt, _ = json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rt)
}

func bind(r *http.Request, tar interface{}) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "读取请体求错误"
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, tar); err != nil {
		return "请求体为非法JSON"
	}
	return nil
}

func (s *Server) handleAPI(path string, handler http.HandlerFunc) {
	s.handler.HandleFunc("/api"+path, func(w http.ResponseWriter, r *http.Request) {
		status := "SUCC"
		// 请求结束，打印日志
		defer func(start time.Time) {
			// 打印请求日志
			// 格式样例：
			// 2019/09/16 15:53:04 [Access] success 127.0.0.1 /todo/add 0.233
			s.log.Printf(
				"[Access] %s %s %s %.3f\n",
				status, // 接口响应状态 success failed
				strings.Split(r.RemoteAddr, ":")[0],
				path, // 接口路径
				float64(time.Now().Sub(start).Nanoseconds())/1e6, // 接口响应时间，毫秒
			)
		}(time.Now())

		// 拦截接口处理逻辑中的panic，并标记请求失败
		defer func() {
			if err := recover(); err != nil {
				s.log.Printf("[Error] panic: %v\n%v\n", err, callstack())
				status = "FAIL"
				resp(w, 10000, nil)
			}
		}()

		handler(w, r)
	})
}

func callstack() []interface{} {
	var cs []interface{}
	cs = append(cs, "Callstack:\n") // start with a new line
	for skip := 0; ; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		// remove golang std package callstack
		if !strings.Contains(file, "/golang/src/") {
			cs = append(cs, fmt.Sprintf("%s:%d\n", file, line))
		}
	}
	return cs
}
