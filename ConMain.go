package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

var RedirectUriRegexp = regexp.MustCompile("^http://[\\w.:]+/auth/callback$")

type ConfMainT struct {
	AppKey      string
	AppSecret   string
	PackageSite string
	ConfFile    string `json:"-"` // 不需要输出

}

func NewConfMain(confFile string) (*ConfMainT, error) {
	fmt.Println("read", confFile)
	var confMain ConfMainT
	confMain.ConfFile = confFile

	body, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &confMain)
	if err != nil {
		return nil, err
	}

	fmt.Printf("conf :%+v", confMain)

	return &confMain, nil
}

func (c *ConfMainT) CheckPara() error {
	if c.AppKey == "" {
		return errors.New("App Key 不能为空!")
	}
	if c.AppSecret == "" {
		return errors.New("App Secret 不能为空!")
	}

	if c.PackageSite == "" {
		return errors.New("package site 不能为空!")
	}

	return nil
}

func (c *ConfMainT) SavePara(appKey, appSecret, packageSite string) error {
	c.AppKey = appKey
	c.AppSecret = appSecret
	c.PackageSite = packageSite

	body, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.ConfFile, body, 0600)
}
