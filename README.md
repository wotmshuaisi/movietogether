# Movie Together

based on golang, rtmp websocket project for chating and watch movie together

## preview

[![youtube](http://img.youtube.com/vi/zyHoc-2rhN4/0.jpg)](http://www.youtube.com/watch?v=zyHoc-2rhN4)

## how to start

1. `dep ensure -v`  download related depend
2. edit `config/config.go` to modify publish secret if you want
3. run `cd cmd && go build -o movietogether` 
4. `chmod +x movietogether && ./movietogether`
5. open http://localhost:port/mt/

## todo

1. add chat history
2. write dockerfile
3. online list