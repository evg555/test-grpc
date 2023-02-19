#  Тестовое приложение

Пример сервера и клиента, взаимодействующих по gRPC

###  Генерация файлов и сборка приложения:

`make build`

### Запуск сервера:

`make run server`

### Запуск клиента:

`./bin/client [Name] [Age]`

Например,

`./bin/client Иван 33`

### Замечаниe:

Для генерации файлов необходимо установить библиотеки:

>go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
> 
>go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest