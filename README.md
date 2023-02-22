# bad-swaggo-example

I don't know why building swagger documentation causes the following error:
```shell
$ swag init -g cmd/api/main.go
2023/02/22 15:58:41 Generate swagger docs....
2023/02/22 15:58:41 Generate general API Info, search dir:./
2023/02/22 15:58:41 warning: failed to get package name in dir: ./, error: execute go list command, exit status 1, stdout:, stderr:no Go files in <<truncated>>/bad-swaggo-example
2023/02/22 15:58:41 ParseComment error in file <<truncated>>/bad-swaggo-example/internal/users/routes.go :cannot find type definition: api.PaginatedResponse[customerResponse]
```