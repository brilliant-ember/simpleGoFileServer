A simple file server written in Go

- Runs on port 5001 by default
- it has 3 routes:
   1. `ping` so calling GET 127.0.0.1/ping will return the respone "pong". This route is for checking that the server is up and running
   2. `upload` POST request allows us to upload a file to the server, you must use "myFile" key for example `curl -F myFile=@file_name 127.0.0.1:5001/upload`
   3. `download` will allow the clinet to download the file previously uploaded using the `upload` route. 
   
Note that this server can upload and receive one single file. It serves the same file it received. If you upload multiple files the last file will overwrite any previous uploads.


### Steps for running
1. `go run ./cmd/.`
2. in another terminal do `curl 127.0.0.1:5001/ping` to see it working