package {{.PkgName}}

const VersionNo = "{{.MetaVersionNo}}"

{{range .Structs }}
/* {{.Desc}} */
type {{.Name}} struct {
{{range .Props.Prop }}	{{.GoName}} {{.GoType}} `json:"{{.Name}}"` // {{.Desc}}
{{end}}
}
{{end}}
