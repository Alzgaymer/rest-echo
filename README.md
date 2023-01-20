# rest-echo

## Docker 

U can run `Dockerfile`, but u need before it run mongo database image and set port mapping

### docker-compose

Using the `docker-compose.yml`, u can run 2 containers as mongo database and **rest-echo** in one time

## Debug

### Start 

gin uses as perfect backend debugger that reloads server

```console
gin --all -i server.go
```


## API

### Adding item

Using the Postman program set url as `http://localhost:8080/products`, then select `body` and type in `key` - *product_name*, in value - whenever u want and click send
