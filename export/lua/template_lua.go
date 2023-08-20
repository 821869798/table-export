package lua

const templateLua string = `local {{.Name}} = {
{{ template "lua" .Content -}}
}

return {{.Name}}

{{- consts "lua" -}}
{{- if IsString . -}}
{{- ToString . -}}
{{- else -}}
{{- range $k, $v := . -}}[{{$k}}] = { {{- template "lua" $v -}} },
{{ end -}}
{{- end -}}
{{- end -}}

`
