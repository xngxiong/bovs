# Bovs

注意：如有疑问请参考[kratos官方文档](https://go-kratos.dev/docs/)
## 配置

1. 修改`internal/conf/config.proto`文件内容；
2. 修改`configs/config.yaml`文件。
3. 在根目录下执行`make config`,然后直接通过读取`internal/conf/config.pb.go`文件中的结构体即可获取配置;

## Orm框架Ent

1. 安装工具`go install entgo.io/ent/cmd/ent@latest`;
2. 创建实体Schema,例如User;
```
cd internal/data
ent new User
```
3. 为生成的User添加字段
```
    func (User) Fields() []ent.Field {
        return []ent.Field{
            field.Int("age").
                Positive(),
            field.String("name").
                Default("unknown"),
        }
    }
```
4. 自动生成ent文件;
```
    go generate ./ent
```
## 参数校验
1. 安装工具`go install github.com/envoyproxy/protoc-gen-validate@latest`;
2. 使用方法详见[proto-gen-validate](https://github.com/bufbuild/protoc-gen-validate)

## 错误处理
1. 安装工具`go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest`;
2. 在`api`目录下的pb文件中中定义错误；
3. 在根目录下执行`make errors`生成错误处理pb文件；

