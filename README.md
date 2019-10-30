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




