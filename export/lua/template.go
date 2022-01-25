package lua

const templateLua string = `local {{.Name}} = {
{{ template "lua" .Content -}}
}

return {{.Name}}

{{- define "lua" -}}
{{- if IsString $ -}}
{{- ToString $ -}}
{{- else -}}
{{- range $k, $v := $ -}}[{{$k}}] = { {{- template "lua" $v -}} },
{{ end -}}
{{- end -}}
{{- end -}}

`
