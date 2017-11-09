package main

import (
	"fmt"
	"net/http"
	"strings"
)

type CtrlAuth struct {
	UtilPage
}

func (c *CtrlAuth) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	if r.RequestURI == "/auth/" {
		c.ShowPage(w, c, "auth.tpl")
		return
	}
	needRedirect := true

	if strings.HasPrefix(r.RequestURI, "/auth/save") {
		c.save(w, r)
	} else {
		c.SetErr("不能识别的请求:" + r.RequestURI)
	}

	if needRedirect {
		c.Redirect(w, "/auth/")
	}
}

func (c *CtrlAuth) save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	err = c.GetConfMain().SavePara(
		r.Form.Get("AppKey"),
		r.Form.Get("AppSecret"),
		r.Form.Get("PackgeSite"))
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	c.SetInfo("save ok!")
}
