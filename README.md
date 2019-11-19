# Calc microservice

## Direct usage

* On one terminal page
```shell
go run main.go
```
 Press `Ctrl + C` to exit

 * On a separate terminal page
```shell
curl -XPOST -d '{"expr": "4+5"}' http://localhost:8080/mu-calc-client
```
Where `expr` is any arithmetical expression with integers and the operators +,-,*,/

## Usage via docker

* On one terminal page
```shell
docker build -t mu-calc-client . 
# Any other tag will do inside docker build but must be the same used in docker run
docker run --rm -p 4000:8080 mu-calc-client
```
* On a separate terminal page
```shell
curl -XPOST -d '{"expr": "4+5"}' http://localhost:4000/mu-calc-client
```
Where `expr` is any arithmetical expression with integers and the operators +,-,*,/

For now `Ctrl + C` doesn't work so the `docker run` command must be killed manually until I find a better solution
