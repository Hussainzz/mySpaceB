{{define "header"}}
{{ template "topheader" . }}

<div class="header_loggedin">

      {{if .InSession}}

          
              <div class="left">
                <a href="/">MySpaceBlog</a>
                <a href="/newpost">NewPost</a>
              </div>
              <div class="right">
              <a href="#">Hii {{.Username}}</a>
              <a href="/logout">Logout</a>
              </div>

      {{else}}

              <div class="left">
                <a href="/">MySpaceBlog</a>
              
              </div>
              <div class="right">
              <a href="/login">Login</a>
                <a href="/register">Register</a>
              </div>
            </div>
      {{end}}


  </div>


{{end}}