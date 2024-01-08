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
	ctx := context.Background()
	pod, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions)
	// lists all pods from a specific namespace : the pods can be accessed using CoreV1

	// fmt.Println(config)
	// fmt.Println(clientset)

	fmt.Println("Pods from default namespace: ")
	for _, pod := range pod.Items {
		fmt.Printf("%s", pod.Name)
	}

	fmt.Println("Deployments from default namespace")
	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	// lists all deployments from a the default namespace
	for _, d := range deployments.Items {
		fmt.Printf("%s", d.Name)
	}

}
