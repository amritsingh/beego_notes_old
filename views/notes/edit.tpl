{{ define "notes/edit.tpl" }}
    {{ template "layouts/header.tpl" }}

        <form class="row g-3" action="/notes/{{.note.ID}}" method="POST">
            <div class="mb-3">
                <label for="name" class="form-label">Title</label>
                <input type="text" class="form-control" id="name" aria-describedby="nameHelp" name="name" value="{{.note.Name}}">
                <div id="nameHelp" class="form-text">Title of the note</div>
            </div>
            <div class="mb-3">
                <label for="content" class="form-label">Content</label>
                <textarea class="form-control" id="content" rows="5" name="content">{{.note.Content}}</textarea>
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
        </form>
    {{ template "layouts/footer.tpl" }}
{{ end }}
