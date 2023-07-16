## Please see the Dockerfile
❯ docker build -t my-golang-app .
```
[+] Building 3.3s (10/10) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                                                0.0s
 => => transferring dockerfile: 365B                                                                                                                                                                                0.0s
 => [internal] load .dockerignore                                                                                                                                                                                   0.0s
 => => transferring context: 2B                                                                                                                                                                                     0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                                                                                                                                    0.5s
 => [internal] load build context                                                                                                                                                                                   0.0s
 => => transferring context: 412B                                                                                                                                                                                   0.0s
 => [1/5] FROM docker.io/library/golang:latest@sha256:cfc9d1b07b1ef4f7a4571f0b60a99646a92ef76adb7d9943f4cb7b606c6554e2                                                                                              0.0s
 => CACHED [2/5] WORKDIR /app                                                                                                                                                                                       0.0s
 => [3/5] COPY . /app                                                                                                                                                                                               0.0s
 => [4/5] RUN go env -w GO111MODULE=off                                                                                                                                                                             0.2s
 => [5/5] RUN go build -o myapp                                                                                                                                                                                     2.3s
 => exporting to image                                                                                                                                                                                              0.1s
 => => exporting layers                                                                                                                                                                                             0.1s
 => => writing image sha256:abf49de5c7a8a4d276c74e09d4f17250a09607856498f9d696d2ecb29ba81adc                                                                                                                        0.0s
 => => naming to docker.io/library/my-golang-app                                                                                                                                                                    0.0s
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them

```

❯ docker run my-golang-app
```
Hello, world!
```
