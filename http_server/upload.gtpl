<html>
    <head>
        <title>Login</title>
    </head>
    <body>
        <form action="/upload" method="POST" enctype="multipart/form-data">
         <input type="file" name="uploadfile">
         <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="上传"/>
        </form>
    </body>
</html>