# GoExpert - Clean Arch

```shell
docker-compose up  
```

```shell
go run main.go wire_gen.go     
```

## gRPC

Gerar os arquivos do protobuf
```shell
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

### Evans
```shell
evans -r repl
```
```
service order_service
call ListOrders
```

## GraphQL

http://localhost:8080/#

```graphql
mutation createOrder {
  createOrder(input: {id:"ccc", Price: 12.2,Tax:2.0}){
    id
    Price
    Tax
    FinalPrice
  }
}

query Orders {
        orders {
            id
            Price
            Tax
            FinalPrice
        }
    }
```
