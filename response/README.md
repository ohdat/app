### response 
修改 errors.go 文件，增加一个错误类型 提交代码 ci执行
```shell
# go get golang.org/x/tools/cmd/stringer
go generate ./...
```

**如何使用**
```
#git config --global url."https://yanghao:test_.123@codeup.aliyun.com".insteadOf "https://codeup.aliyun.com"
git config --global url."https://username:password@codeup.aliyun.com".insteadOf "https://codeup.aliyun.com"
#执行
go env -w GOPRIVATE=github.com/ohdat 
go get -u github.com/ohdat/response
```
