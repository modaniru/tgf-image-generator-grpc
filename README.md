
# Twitch general follows
A project that drawing an image in runtime and then returns as []byte. Needs [twitch-general-follows-gRPC](https://github.com/modaniru/twitch-general-follows-gRPC) server

## Content

1. [Run Locally](https://github.com/modaniru/tgf-image-generator-grpc#run-locally)
2. [Docker](https://github.com/modaniru/tgf-image-generator-grpc#docker)
3. [Environment variables](https://github.com/modaniru/tgf-image-generator-grpc#environment-variables)
4. [Tasks](https://github.com/modaniru/tgf-image-generator-grpc#tasks)

## Run Locally

Clone the project

~~~bash
  git clone https://github.com/modaniru/tgf-image-generator-grpc
~~~

Go to the project directory

~~~bash
  cd tgf-image-generator-grpc
~~~

Create .env file

~~~bash
  touch .env
~~~

Write secrets in .env ([more](https://github.com/modaniru/tgf-image-generator-grpc#environment-variables))

~~~bash
  TGF_SERVICE_HOST=your.tgf.server
~~~

If you can run "make" commands

~~~bash
  make
~~~

Else: \
Install dependencies

~~~bash
go get ./...
~~~

Start the server

~~~bash
go run src/main.go
~~~

the server will run on **8080** port\
You can change port *check* [about .env](https://github.com/modaniru/tgf-image-generator-grpc#environment-variables)

## Docker
soon
<!-- run from **Docker Hub**
~~~bash
docker run -p 8080:8080 -e TWITCH_CLIENT_ID=clientId -e TWITCH_CLIENT_SECRET=clientSecert modaniru/tgf
~~~
or
~~~bash
docker run -p 8080:8080 --env-file path modaniru/tgf
~~~
**build** and run docker container
~~~bash
docker build -t imageName .
docker run --env-file path -p 8080:8080 imageName
~~~ -->

## Environment variables

~~~bash
  TGF_SERVICE_HOST=localhost:8080 // your twitch-general-follows-grpc server url
  PORT=80 // application running port (optional, default: 8080)
~~~

## Tasks
- [ ] Docker
- [x] CI/CD
- [ ] Tests
- [ ] Crossplatform working with files (now: Linux)