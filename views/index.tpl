{{template "header" .}}


 {{range $index,$post:=.post_list}}

<div class="container w3-light-grey w3-round-large w3-bar-block" style="width:40% ;margin-left:2%;margin-top:5%">
 
  
         <a href="/post/{{$post.Id}}"><h1 style="color:black;font-weight:700;"> {{$post.Title}}</h1></a>

<img src="static/img/ms.png" class="w3-round-xxlarge" alt="Norway" style="width:10%; height:50px; margin-top:-20%; margin-left:70%">

           <div class="data" style="color:black; margin-top:-6%">
       <p style="font-size:11px; font-style:italic;color:grey">Posted by {{$post.Author}}</p>
          <br>
         <p style="font-size:11px; font-style:italic;color:grey"> {{$post.Date}}</p>
<br>
                {{ if $.user}}
                    {{if eq $post.Users.Id $.user.Id}}
							
    &nbsp<a href="/post/edit/{{$post.Id}}"><button class="w3-button w3-round w3-green">edit</button></a>
	 &nbsp<a href="/post/delete/{{$post.Id}}?from=index"> <button class="w3-button w3-round w3-red">delete</button></a>
							
                    {{end}}
                    {{end}}



        </div>
<br>
</div>

{{end}}



{{template "footer" .}}