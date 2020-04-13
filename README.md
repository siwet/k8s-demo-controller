# k8s-demo-controller

## STEP1. module initialize

Run the following command to initialize go module
```bash
go mod init 
```
Since this is a git repository, module path `github.com/cntsw/k8s-demo-controller` while be auto-detected.

Add dependencies: 
```
require (
	k8s.io/api kubernetes-1.13.12
	k8s.io/apimachinery kubernetes-1.13.12
	k8s.io/client-go v10.0.0
)
```

## STEP2. generate code

Generate code on the remote server because code-generator does not support go mod

```bash
# install go 1.12
sudo snap install --classic --channel=1.12/stable go

# config env
export GOPATH=/root/.gopath
export GO111MODULE=off

# clone code
mkdir -p "$HOME/.gopath/src/github.com/cntsw"
cd "$HOME/.gopath/src/github.com/cntsw"
git clone https://github.com/cntsw/k8s-demo-controller.git

# get code-generator
go get -u k8s.io/code-generator/...
go get -u k8s.io/apimachinery
cd $GOPATH/src/k8s.io/code-generator

# generate code
MODULE_PATH="github.com/cntsw/k8s-demo-controller"
CLIENT_PATH="$MODULE_PATH/pkg/client"
APIS_PATH="$MODULE_PATH/pkg/apis"
./generate-groups.sh all "$CLIENT_PATH" "$APIS_PATH" "example:v1alpha1" -v 5

```

apk add git bash


export GO111MODULE=on

go get k8s.io/apimachinery
cd /go/src/k8s.io/apimachinery && git checkout kubernetes-1.16.6

go get k8s.io/code-generator
cd /go/src/k8s.io/code-generator && git checkout kubernetes-1.16.6

go get -u ./...

go install ./cmd/defaulter-gen
go install ./cmd/client-gen
go install ./cmd/lister-gen
go install ./cmd/informer-gen
go install ./cmd/deepcopy-gen

,,,,

https://github.com/trstringer/k8s-controller-custom-resource

apk add git
cd
git clone https://github.com/kubernetes/code-generator.git -b kubernetes-1.16.6 --depth 1
cd code-generator


git clone https://github.com/kubernetes/apimachinery.git -b kubernetes-1.16.6 --depth 1
s

    go install ./cmd/defaulter-gen && \
    go install ./cmd/client-gen && \
    go install ./cmd/lister-gen && \
    go install ./cmd/informer-gen && \
    go install ./cmd/deepcopy-gen && \

