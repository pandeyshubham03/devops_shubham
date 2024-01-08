package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		// handle error
		fmt.Printf("error %s building config from flags\n", err.Error())
		_, err := rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}

	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}

	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	podinformer := informerfactory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			fmt.Println("add was called")
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("updated was called")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete was called")
		},
	})

	informerfactory.Start(wait.NeverStop)
	informerfactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podinformer.Lister().Pods("default").Get("default")
	fmt.Println(pod)

	// ctx := context.Background()
	// pod, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions)
	// // lists all pods from a specific namespace : the pods can be accessed using CoreV1

	// // fmt.Println(config)
	// // fmt.Println(clientset)

	// fmt.Println("Pods from default namespace: ")
	// for _, pod := range pod.Items {
	// 	fmt.Printf("%s", pod.Name)
	// }

	// fmt.Println("Deployments from default namespace")
	// deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	// // lists all deployments from a the default namespace
	// for _, d := range deployments.Items {
	// 	fmt.Printf("%s", d.Name)
	// }

}
