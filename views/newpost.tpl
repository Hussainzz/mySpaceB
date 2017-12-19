{{ template "header" .}}



<div class="container">
  <div class="row">
   
 </div>
</div>

<div class="container" style="margin-top:5%">
  <div class="row">
    
    

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}

    <div style="margin-left:45%">
    <h1 style="font-family:'Pacifico', cursive; color:#42e2f4">New Post</h1>
    </div>
             
      <form  class="w3-border w3-round-xxlarge" role="form" id="postfrm" method="POST" style="color:white;width:50%;margin-left:25%">
       <br>
      <div style="margin-left:15%"> 
       
        Title :<input type="text" name="title"  style="width:30%" required /><br><br>

 <div class="message">
  <p>	
    <label>Description</label><br>
    <textarea class="w3-border w3-round-large" name="description" placeholder="Write something.." style="width:83%; height:200px" required></textarea>
	</p>
	</div>
        <br>
        <p><button class="w3-btn w3-blue w3-border w3-round" style="margin-left:30%; margin-top:-1%; width:20%">Post</button></p>
        
<br>
    </div>    
      </form>
     
    
  </div>
</div>

{{ template "footer" }}