# 本文将用于快速编译并生成webook镜像，使用方法 make docker
# 声明了一个名为 "docker" 的伪目标，
# 意味着当你在命令行中执行 `make docker` 时，
# 不会检查是否存在名为 "docker" 的文件或目录。而是执行 Makefile 中与 "docker" 目标相关联的命令。
.PHONY: docker
docker:
	# 删除上次编译
	@rm webook || true
	# 运行依赖
	@go mod tidy
	# 打包webook
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook .
	# 删除镜像
	@docker rmi -f zhao/webook:v0.0.1
	# 生成镜像
	@docker build -t zhao/webook:v0.0.1 .