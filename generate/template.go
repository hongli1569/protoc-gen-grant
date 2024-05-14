package generate

import (
	"bytes"
	"strings"
	"text/template"
)

func execTpl(s *ServiceGrant) string {
	tpl := `

func {{ .ServiceName }}Grant() []*annotations.Grant {
	return []*annotations.Grant{
		{{- range .Grant }}
		{
			Action: "{{ .Action }}",
			FullName: {{ .FullName }},
			Desc:"{{ .Desc }}",
		},
		{{ end }}
	}
}

`

	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(tpl)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
