{{ template "../global/main.tpl" . }}
	{{ define "css" }}
    		<link rel="stylesheet" href="/static/css/current.css">
	{{ end}}


	{{ define "content" }}
    		<h2>{{ .Title }}</h2>
    		<p> This is SomeVar: {{ .SomeVar }}</p>
	{{ end }}

	{{ define "js" }}
		<script src="/static/js/current.js"></script>
	{{ end}}