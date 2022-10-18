{{ define "index.tpl"}}                                                                                                                                                                                      
 
{{ template "layouts/header.tpl" .}} 
 
<div class="container-xl">
 
<h1>Notes App</h1>
 
<p>
    <div class="btn-group" role="group">
        {{if .email}}
            <form action="/logout" method="POST">
                <button type="submit" class="btn btn-outline-danger">Logout</button>
            </form>
        {{else}}
            <a class="btn btn-outline-primary" href="/signup" role="button">Signup</a>
            <a class="btn btn-outline-primary" href="/login" role="button">Login</a>
        {{end}}
    </div>
</p>
 
</div>
 
{{ template "layouts/footer.tpl" .}} 
 
{{ end }}
