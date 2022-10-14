{{ define "notes/show.tpl" }}
    {{ template "layouts/header.tpl" }}
        <div class="card h-200">
            <div class="card-body">
                <h5 class="card-title">{{ .note.Name }}</h5>
                <p class="card-text fs-6 fw-light">{{ .note.Content }}</p>
            </div>
        </div>
    {{ template "layouts/footer.tpl" }}
{{ end }}