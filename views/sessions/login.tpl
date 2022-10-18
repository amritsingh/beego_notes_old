{{ define "sessions/login.tpl"}}                                                                                                                                                                                      
 
{{ template "layouts/header.tpl" .}} 
 
<div class="container-xl">
<h1>Login</h1>
 
<form  action="/login" method="POST">
    <div class="mb-3">
        <label for="email" class="form-label">Email address</label>
        <input type="email" class="form-control" id="email" aria-describedby="emailHelp" name="email">
        <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
    </div>
    <div class="mb-3">
        <label for="password" class="form-label" name="password">Password</label>
        <input type="password" class="form-control" id="password" name="password">
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>
</div>
{{ template "layouts/footer.tpl" }} 
 
{{ end }}