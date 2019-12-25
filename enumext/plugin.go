package enumext

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func EnumExt() *module { return &module{ModuleBase: &pgs.ModuleBase{}} }

func (p *module) Name() string {
	return "enumext"
}

func (p *module) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("enumext").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
		"extVal":  p.extVal,
	})

	p.tpl = template.Must(tpl.Parse(EnumExtTpl))
}

func (p *module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}
	return p.Artifacts()
}

func (p *module) generate(f pgs.File) {
	if len(f.Enums()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".enumext.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

func (p *module) extVal(ev pgs.EnumValue) string {
	var val string
	ok, err := ev.Extension(E_ExtVal, &val)
	if !ok || err != nil {
		return ev.Name().String()
	}
	return val
}

const EnumExtTpl = `package {{ package . }}

{{ range .AllEnums }}
// -------------------------------------------------------------------
// {{ name . }}
// -------------------------------------------------------------------
func (e {{ name . }}) IsValid() bool {
	_, ok := {{ name . }}_name[int32(e)]
	return ok
}

var gql_{{ name . }}_name = map[int32]string{
{{ range .Values -}}
	{{ .Value }}: "{{ extVal . }}",
{{ end }}
}
var gql_{{ name . }}_value = map[string]int32{
{{ range .Values -}}
	"{{ extVal . }}": {{ .Value }},
{{ end }}
}

func (e {{ name . }}) ExtValue() (string, error) {
	if !e.IsValid() {
		return "", fmt.Errorf("invalid {{ name . }} '%s'", e)
	}
	return gql_{{ name . }}_name[int32(e)], nil
}

{{ end }}
`
