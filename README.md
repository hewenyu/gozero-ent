# gozero-ent
go zero 使用ent组件作为数据查询



## proto 文件生成

```bash
# 生成文件
goctl rpc protoc rpc/pb/*.proto --go_out=rpc/ --go-grpc_out=rpc/  --zrpc_out=rpc/  --style go_zero
# 去除不必要的文件
sed -i  's/,omitempty//g' rpc/pb/*.pb.go
```


## 使用生成好的ent 库

```bash
go get -u github.com/hewenyu/ent
```

## 修改相关配置简单操作