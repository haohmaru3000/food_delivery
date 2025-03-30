<h1 align="center">Food Delivery</h1>

## Prerequisites
- Go >= 1.19.2
- Docker

## Commands
### `Build`
```bash
# build cross platform
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
```

### `Docker`

```bash
# Setup Database
$ docker run -d --name mysql --privileged=true \
    -e MYSQL_ROOT_PASSWORD="admin" \
    -e MYSQL_USER="food_delivery" \
    -e MYSQL_PASSWORD="12345678" \
    -e MYSQL_DATABASE="food_delivery" \
    -p 3307:3306 mysql:latest

# docker build
$ docker build -t food_delivery

# create a network
$ docker network create food_delivery_network

# docker network connect
$ docker network connect food_delivery_network mysql

# run 
$ docker run -d --name food_delivery \ 
    -e DBConnectionStr="food_delivery:12345678@tcp(mysql:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" \ 
    --network=food_delivery_network \ 
    -p 3001:8080 \
    food_delivery_app
```

### `Reverse Proxy with Nginx`

```bash
$ docker run -d -p 80:80 -p 443:443 \
    --network=food_delivery_network --name nginx-proxy \
    -e ENABLE_IPV6=true \
    --privileged=true \
    -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
    -v ~/nginx-certs:/etc/nginx/certs:ro \
    -v ~/nginx-conf:/etc/nginx/conf.d \
    -v ~/nginx-logs:/var/log/nginx \
    -v /usr/share/nginx/html \
    -v /var/run/docker.sock:/tmp/docker.socker:ro \
    --label nginx_proxy jwilder/nginx-proxy
    
$ docker run -d --network=food_delivery_network \
    -v ~/nginx/vhost.d:/etc/nginx/vhost.d \
    -v ~/nginx-certs:/etc/nginx/certs:rw \
    -v /var/run/docker.sock:/tmp/docker.socker:ro \
    --volumes-from nginx-proxy \
    --privileged=true \
    jrcs/letsencrypt-nginx-proxy-companion
```
