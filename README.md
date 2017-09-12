
### 对 alertover 的api 封装, 主要为了方便服务器端各种脚本调用, 其实就是个POST请求 XD

## 使用:
### 修改为自己的信息
修改 `main.go` 中 `sourceID`, `receiverID` 为自己的信息, 需要按照 [alerover 配置指南](https://www.alertover.com/pages/api) 来获取
### 编译 
`go build -o sendmsg main.go`
### 用法
- ./sendmsg CONTENT [TITLE]
- echo hello |./sendmsg
- ./sendmsg < file.txt