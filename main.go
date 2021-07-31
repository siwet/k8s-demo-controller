package main

import (
	"flag"
	"fmt"
	"github.com/cntsw/k8s-demo-controller/pkg/controller"
	clientset "github.com/cntsw/k8s-demo-controller/pkg/generated/client/clientset/versioned"
	informers "github.com/cntsw/k8s-demo-controller/pkg/generated/client/informers/externalversions"
	signals "github.com/cntsw/k8s-demo-controller/pkg/singles"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	homeKubeconfigFile := filepath.Join(homeDir(), ".kube", "config")
	if !fileExists(homeKubeconfigFile) {
		homeKubeconfigFile = ""
	}

	flag.StringVar(&kubeconfig, "kubeconfig", homeKubeconfigFile, "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Set("v", "2")

	flag.Parse()
}

func main() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	test(kubeClient) // test k8s client

	userClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	userInformerFactory := informers.NewSharedInformerFactory(userClient, time.Second*30)

	controller := controller.NewController(kubeClient, userClient,
		userInformerFactory.Example().V1alpha1().Users())

	go userInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func test(client *kubernetes.Clientset) {
	podList, err := client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic("xxx")
	}

	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
	}
}
