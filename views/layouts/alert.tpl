{{ define "layouts/alert.tpl"}}
  {{if .flash}}
    {{if .flash.error}}
      <div class="alert alert-danger" role="alert">{{.flash.error}}</div>
    {{end}}
    {{if .flash.warning}}
      <div class="alert alert-warning" role="alert">{{.flash.warning}}</div>
    {{end}}
    {{if .flash.success}}
      <div class="alert alert-primary" role="alert">{{.flash.success}}</div>
    {{end}}
    {{if .flash.notice}}
      <div class="alert alert-secondary" role="alert">{{.flash.notice}}</div>
    {{end}}
  {{end}}
{{ end }}