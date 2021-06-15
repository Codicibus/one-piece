# OPiece

# 编译后端
```bash
git clone https://github.com/Codicibus/one-piece.git
cd one-piece && git switch server
GOPROXY=https://goproxy.cn,direct go build .
```

# 配置文件
```yaml
jwt:
  secret_key: abc
  expire_time: 604800
  buffer_time: 86400

gin:
  mode: debug
  port: 8081

db:
  type: postgres

postgres:
  db_host: localhost-1
  db_port: 54321
  db_name: opiecedb
  db_username: opiece
  db_password: postgres
  ssl_mode: disable
  timezone: Asia/Shanghai
```