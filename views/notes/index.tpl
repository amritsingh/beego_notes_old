{{ define "notes/index.tpl" }}
    {{ template "layouts/header.tpl" }}
        <h1>All Notes<h1>

        <div class="row row-cols-1 row-cols-md-6 g-4">
            {{ with .notes }}
                {{ range . }}
                    <div class="col">
                        <div class="card h-200">
                            <div class="card-body">
                                <h5 class="card-title">{{ .Name }}</h5>
                                <p class="card-text">{{ .Content }}</p>
                                <a class="btn btn-outline-primary" href="/notes/{{.ID}}" role="button">View</a>
                            </div>
                        </div>
                    </div>
                {{ end }}
            {{ end }}
        </div>

    {{ template "layouts/footer.tpl" }}
{{ end }}