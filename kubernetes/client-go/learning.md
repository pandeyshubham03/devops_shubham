# Learning Kubernetes client-go

To lists all pods in "default" namespace: 

``` sh
pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions)
```

To list all deployments in "default" namespace:
``` sh
ctx := context.Background()
deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
```