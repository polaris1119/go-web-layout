package controller

import (
	"context"
	"fmt"
	"github.com/polaris1119/go-web-layout/demo"
	"html/template"
	"net/http"
	"strconv"
)

type DemoHttpHandler struct {
	logic demo.Logic
}

func NewDemoHttpHandler(demoLogic demo.Logic) *DemoHttpHandler {
	handler := &DemoHttpHandler{
		logic: demoLogic,
	}

	http.HandleFunc("/demo/detail", handler.Fetch)
	http.HandleFunc("/demo/create", handler.Create)
	http.HandleFunc("/", handler.Index)

	return handler
}

func (dh *DemoHttpHandler) Index(writer http.ResponseWriter, req *http.Request) {
	htmlTxt := `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Demo</title>
	</head>
	<body>
		<form action="/demo/create" method="post">
			JSON 数据：<textarea rows="4" name="data"></textarea>
			<input type="submit" value="提交">
		</form>
		<a href="/demo/detail?id=1">获取数据</a>
	</body>
</html>`

	tpl, err := template.New("demo").Parse(htmlTxt)
	if err != nil {
		fmt.Fprintln(writer, "parse error:", err)
		return
	}

	err = tpl.Execute(writer, nil)
	if err != nil {
		fmt.Fprintln(writer, "execute error:", err)
		return
	}
}

func (dh *DemoHttpHandler) Fetch(writer http.ResponseWriter, req *http.Request) {
	id := mustInt(req.FormValue("id"))

	ctx := context.Background()

	demo, err := dh.logic.Fetch(ctx, id)
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	fmt.Fprintln(writer, demo)
}

func (dh *DemoHttpHandler) Create(writer http.ResponseWriter, req *http.Request) {
	data := req.FormValue("data")

	ctx := context.Background()

	err := dh.logic.Create(ctx, data)
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	fmt.Fprintln(writer, "创建成功")
}

func mustInt(s string, defVal ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
	}

	return i
}
