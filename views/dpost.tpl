{{template "header" .}}



    <div style="color:white;margin-top:5%;margin-left:5%" >

            <p>
            <img src="/static/img/user.png" class="w3-round" alt="img" style="width:2%;height:26px">
               
                <a style="font-size:11px; font-style:italic;color:gray" >posted by</a> <a style="font-weight:700;font-size:20px"> {{.post.Author}}</a>
                <br>
				<a style="font-size:11px; font-style:italic;color:gray" >{{.post.Date}}</a>
				
            </p>

        <div class="container w3-border w3-round-large w3-bar-block" style="color:white;
         margin-top:1%;margin-left:0%;width:50%;height:250px" >

        
            <h2 style="text-align:center">{{.post.Title}}</h2>

            <hr>
           
            
               <a style="font-family: 'Maven Pro', sans-serif;"> {{.post.Description}}</a>
            

            
        </div>

        <div>
        <img src="/static/img/rocket.jpg" class="w3-circle" alt="img" style="width:40%;height:500px;margin-left:55%;margin-top:-20%">
        </div>
        
    </div>

{{template "footer" .}}