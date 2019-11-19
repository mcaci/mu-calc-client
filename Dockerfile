FROM mcaci/docker-dojo-golang

RUN ["go", "get", "-u", "github.com/mcaci/mu-calc-client"]
ENTRYPOINT ["go", "run", "github.com/mcaci/mu-calc-client"]
EXPOSE 8080
