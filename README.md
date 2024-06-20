### Docker 

Structure of project :

Quit docker terminal 🔍

```
Ctrl + D
```

```
ascii-art-web/
├── main.go
├── status.go
├── template/
│   ├── index.html
│   ├── other.html
├── Dockerfile
├── go.mod
└── go.sum
```

Create image

```
docker build -t myproject .
```

Launch server

```
docker run -p 8080:8080 myproject
```

rename image

```
docker tag oldimage:latest newimage:v1
```

show containers

```
docker -p
```

show images

```
docker images
```

rename container

```
docker rename <container_name_or_id> <new_container_name>

```

add metada 

```
docker build -t ascii-art-web/server:latest .
```

Expected output : 

```
MacBook-Pro:ascii-art-web pierrecaboor$ docker build -t ascii-art-web/server:latest .
[+] Building 19.8s (12/12) FINISHED                                                                                                          docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                         0.0s
 => => transferring dockerfile: 780B                                                                                                                         0.0s
 => [internal] load metadata for docker.io/library/golang:1.22-alpine                                                                                        2.0s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                                                0.0s
 => [internal] load .dockerignore                                                                                                                            0.0s
 => => transferring context: 2B                                                                                                                              0.0s
 => [1/6] FROM docker.io/library/golang:1.22-alpine@sha256:6522f0ca555a7b14c46a2c9f50b86604a234cdc72452bf6a268cae6461d9000b                                  0.0s
 => [internal] load build context                                                                                                                            0.0s
 => => transferring context: 16.54kB                                                                                                                         0.0s
 => CACHED [2/6] WORKDIR /app                                                                                                                                0.0s
 => CACHED [3/6] COPY go.mod ./                                                                                                                              0.0s
 => CACHED [4/6] RUN go mod download                                                                                                                         0.0s
 => [5/6] COPY . .                                                                                                                                           0.1s
 => [6/6] RUN go build -o main .                                                                                                                            16.7s
 => exporting to image                                                                                                                                       0.8s
 => => exporting layers                                                                                                                                      0.7s
 => => writing image sha256:45582c64deaac18986b67171caead70469d4cee0c53213f4de6508ae495cb7f9                                                                 0.0s
 => => naming to docker.io/ascii-art-web/server:latest                                                                                                       0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/xshf8u6dozi4j626xx3wyma6j

What's Next?
  View a summary of image vulnerabilities and recommendations → docker scout quickview
  ```
  
show metada 

```
docker inspect --format='{{json .Config.Labels}}' ascii-art-web/server:latest
```

>tips install jq to show clean output ✨

use jq 

```
brew install jq
```

after jq 

```sh
MacBook-Pro:ascii-art-web pierrecaboor$ docker inspect --format='{{json .Config.Labels}}' ascii-art-web/server:latest | jq .
{
  "com.example.department": "IT",
  "description": "Image Docker in Go 🐳",
  "maintainer": "Pierre Caboor pierre.caboor59@gmail.com",
  "version": "1.0"
}
```
