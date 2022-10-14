{{ define "notes/show.tpl" }}
    {{ template "layouts/header.tpl" }}
        <div class="card h-200">
            <div class="card-body">
                <h5 class="card-title">{{ .note.Name }}</h5>
                <p class="card-text fs-6 fw-light">{{ .note.Content }}</p>
            </div>
        </div>

        <p>
            <a class="btn btn-outline-primary" href="/notes/edit/{{ .note.ID }}" role="button">Edit</a>
        </p>
    {{ template "layouts/footer.tpl" }}
{{ end }}