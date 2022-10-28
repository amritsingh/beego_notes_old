{{ define "notes/show.tpl" }}
    {{ template "layouts/header.tpl" . }}
        <div class="card h-200">
            <div class="card-body">
                <h5 class="card-title">{{ .note.Name }}</h5>
                <p class="card-text fs-6 fw-light">{{ .note.Content }}</p>
            </div>
        </div>

        <script>                                                                                                                                                                                                           
            function sendDelete(event, href){
                var xhttp = new XMLHttpRequest();
                event.preventDefault();
                xhttp.onreadystatechange = function() {
                    // return if not ready state 4
                    if (this.readyState !== 4) {
                        return;
                    }
        
                    if (this.readyState === 4) {
                        // Redirect the page
                        window.location.replace(this.responseURL);
                    }
                };
                xhttp.open("DELETE", href, true);
                xhttp.send();
            }
        </script>
        
        <p>
            <div class="btn-group" role="group">
                <a class="btn btn-outline-primary" href="/notes/edit/{{ .note.ID }}" role="button">Edit</a>
                <a class="btn btn-outline-danger" href="/notes/{{ .note.ID }}" onclick="sendDelete(event, this.href)">Delete</a>
            </div>
        </p>

    {{ template "layouts/footer.tpl" . }}
{{ end }}