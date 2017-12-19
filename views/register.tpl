<!DOCTYPE html>
<html>
<title>Register</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<link href="https://fonts.googleapis.com/css?family=Pacifico" rel="stylesheet">

<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>



form {width:500px; margin-left:500px; margin-top:5%}

html,h3,p{
    font-family: "Century Gothic"; font-style: normal; font-weight: 700;
    height:100%
	}
	
body{
	

	background:  #374954;
	
	 
}
label{
  color:white;
}

h2{ margin-top:80px;
font-family:'Pacifico', cursive; color:#42e2f4;
}

.data{
	color:grey;
}




</style>

<body>



<div class="w3-container" style="margin-left:43%;">
  <h2><b>Lets get Started..</b></h2>
</div>



 

 
 
 <div style="width:100%;height:70px;bottom:0%; margin-top:-9.5%;margin-left:0%">
 <br>
 <br>
<a href="/" style="margin-left:2%;"><span class="glyphicon glyphicon-home"></a>
 <hr>
</div> 
 
 &nbsp;
{{if .flash.error}}
<h3>{{.flash.error}}</h3>
&nbsp;
{{end}}
{{if .Errors}}
{{range $rec := .errors}}
<h3>{{$rec}}</h3>
{{end}}
&nbsp;
{{end}}

<form action="/register" class="w3-container w3-border w3-round-xxlarge" style="margin-left:35%" method="POST">
<br>
  <p class="data">Tell us about yourself so we can create your account ...</p>
   
  <p>
  <label>UserName</label>
  <input class="w3-input w3-border w3-round-large" name="username" type="text" required></p>
  
   <p>
  <label>Password</label>
  <input class="w3-input w3-border w3-round-large" name="pwd" type="password" required></p>
  
  <p>
  <label>Email Address</label>
  <input class="w3-input w3-border w3-round-large" name="email" type="text" required></p>
 
  <p><button class="w3-btn w3-blue w3-border w3-round">Register</button></p>
  <br>
</form>

</body>
</html>