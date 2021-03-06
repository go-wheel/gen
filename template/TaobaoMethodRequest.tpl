package opentb

import (
	"errors"
)

type TaobaoMethodRequest struct {
	TaobaoRequest
}

func (t *TaobaoMethodRequest) GetResponse(accessToken, apiMethodName string, resp interface{}) ([]byte, error) {
	if accessToken == "" {
		return nil, errors.New("[" + apiMethodName + "] AccessToken is null")
	}
	t.SetReqUrl("https://eco.taobao.com/router/rest")

	t.SetValue("method", apiMethodName)
	t.SetValue("format", "json")
	t.SetValue("access_token", accessToken)
	t.SetValue("v", "2.0")

	return executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
}
