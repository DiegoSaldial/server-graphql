{{ template "../global/main.tpl" . }}
{{ define "css" }}
    <link rel="stylesheet" href="/static/auth.css">
{{ end}}


{{ define "content" }}
    <h2>{{ .Title }}</h2>
    <p> This is SomeVar: {{ .SomeVar }}</p>
{{ end }}

{{ define "js" }}
    <script src="/static/auth.js"></script>
{{ end}}