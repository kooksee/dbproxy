.PHONY: version build build_linux docker_login docker_build docker_push_dev docker_push_pro
.PHONY: rm_stop

_Version = v0.0.1
_VersionFile = version/version.go
_VersionCheckFile = version/version.md
_CommitVersion = $(shell git rev-parse --short=8 HEAD)
_BuildVersion = $(shell date "+%F %T")
_GOBIN = $(shell pwd)

_ImageName = dbproxy
_ProjectPath = github.com/kooksee/$(_ImageName)
_ImagesPrefix = registry.cn-hangzhou.aliyuncs.com/ybase/
_ImageVersionName = $(_ImagesPrefix)$(_ImageName):$(_Version)

_version:
	@echo "package version" > $(_VersionFile)
	@echo "const Version = "\"$(_Version)\" >> $(_VersionFile)
	@echo "const BuildVersion = "\"$(_BuildVersion)\" >> $(_VersionFile)
	@echo "const CommitVersion = "\"$(_CommitVersion)\" >> $(_VersionFile)

	@echo	"## Version: $(_Version)" >> $(_VersionCheckFile)
	@echo	"	1. BuildVersion: $(_BuildVersion)" >> $(_VersionCheckFile)
	@echo	"	2. CommitVersion: $(_CommitVersion)" >> $(_VersionCheckFile)

_build_linux: _version
	@echo "交叉编译linux"
	docker run --rm -v $(GOPATH):/go golang:latest go build -o /go/src/$(_ProjectPath)/main /go/src/$(_ProjectPath)/main.go

_docker_build:_build_linux
	@echo "docker build"
	sudo docker build -t $(_ImageVersionName) .

_docker_push:
	@echo "docker push"
	sudo docker push $(_ImageVersionName)

version:_version
	@echo $(_Version)

b:
	@echo "开始编译"
	GOBIN=$(_GOBIN) go install main.go

r:
	@echo "运行"
	go run main.go

docker:_docker_build _build_linux