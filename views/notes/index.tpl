{{ define "notes/index.tpl" }}
    {{ template "layouts/header.tpl" }}
        <h2>All Notes<h2>

        <div class="row row-cols-1 row-cols-md-6 g-4">
            {{ with .notes }}
                {{ range . }}
                    <div class="col">
                        <div class="card h-200">
                            <div class="card-body">
                                <h5 class="card-title">{{ .Name }}</h5>
                                <p class="card-text fs-6 fw-light">{{ .Content }}</p>
                                <a class="btn btn-outline-primary" href="/notes/{{.ID}}" role="button">View</a>
                            </div>
                        </div>
                    </div>
                {{ end }}
            {{ end }}
        </div>

        <p>
            <a class="btn btn-outline-primary" href="/notes/new" role="button">New Note</a>
        </p>

    {{ template "layouts/footer.tpl" }}
{{ end }}