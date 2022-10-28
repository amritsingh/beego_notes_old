{{ define "layouts/navbar.tpl" }}
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">Gin Notes App</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        {{if .email}}
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <a class="nav-link active" aria-current="page" href="/notes">Notes</a>
            </li>
          </ul>
        {{end}}
      </div>
    </div>
    <div class="float-xl-right" id="navbarUserInfo">
      <ul class="navbar-nav">
        {{if .email}}
          <li class="nav-item disabled">
            <a class="nav-link">{{.email}}</a>
          </li>
          <li class="nav-item">
            <form action="/logout" method="POST">
              <button type="submit nav-item" class="btn btn-outline-danger">Logout</button>
            </form>

          </li>
        {{else}}
          <li class="nav-item">
            <a class="nav-link" href="/signup">SignUp</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/login">Login</a>
          </li>
        {{end}}
      </ul>
    </div>
  </nav>
{{ end }}