FROM mcaci/docker-dojo-golang

RUN ["go", "get", "-u", "github.com/mcaci/mu-calc-service"]
ENTRYPOINT ["go", "run", "github.com/mcaci/mu-calc-service"]
EXPOSE 8080
