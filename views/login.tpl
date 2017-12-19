{{ template "header" .}}


<div class="login">
	<h1 style="font-family:'Pacifico', cursive; color:#42e2f4">Login</h1>
    &nbsp;
{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}
{{if .Errors}}
{{range $rec := .Errors}}
<h3>{{$rec}}</h3>
{{end}}
&nbsp;
{{end}}
    <form action="/login" method="post">
    	<input type="text" name="username" placeholder="Username" required="required" />
        <br>
        <br>
        
        <input type="password" name="password" placeholder="Password" required="required" />
        <br><br>
        <button type="submit" class="btn btn-primary btn-block btn-large">Let me in.</button>
    </form>
</div>


{{ template "footer" }}