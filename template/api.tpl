package {{.PkgName}}

import (
	"{{.PackageSite}}"
)

{{range .Apis}}

{{$ApiGoName:=.GoName}}

/* {{.Desc}} */
type {{$ApiGoName}}Request struct {
	opentb.TaobaoMethodRequest
}

{{range .Request.Param}}
/* {{.Desc}} */
func (r *{{$ApiGoName}}Request) Set{{.GoName}}(value string) {
	r.SetValue("{{.Name}}", value)
}
{{end}}

func (r *{{$ApiGoName}}Request) GetResponse(accessToken string) (*{{$ApiGoName}}Response, []byte, error) {
	var resp {{$ApiGoName}}ResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "{{.Name}}", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type {{$ApiGoName}}Response struct {
{{range .Response.Param}}	{{.GoName}} {{.GoType}} `json:"{{.Name}}"`
{{end}}}

type {{$ApiGoName}}ResponseResult struct {
	Response *{{.GoName}}Response `json:"{{.JsonName}}_response"`
}

{{end}}

