simple-api-go
=============

This is a simple REST API for exploring Go. The library is currently using gorilla/mux, gorp, goose, and godep.

To install

1. Copy db/dbconf.yml.sample to db/dbconf.yml with the updated information
2. Run ```goose up```
3. Run ```go build```
4. Run ```./simple_api_go.```
5. Go to __localhost:3000__

Hope this helps!