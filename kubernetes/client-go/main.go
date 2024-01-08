package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		// handle error
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
	}

	clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions) // lists all pods from a specific namespace : the pods can be accessed using CoreV1

	fmt.Println(config)
	fmt.Println(clientset)
}
