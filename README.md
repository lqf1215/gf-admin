# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.


这里新项目是通过
- `https://github.com/SupenBysz/gf-admin-community`
  `https://github.com/SupenBysz/gf-admin-company-modules`
- 引入这两个git库的骨架


原数据库文件位置：
- `db/latest.sql`

新增数据库文件位置：
 dept: 部门、post: 岗位
- `db/co_company_dept_post.sql`

进行vendor的命令：
- `go mod vendor`

运行位置：
- `internal/cmd/cmd.go`