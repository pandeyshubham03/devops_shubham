# Learing Kubernetes client-go

To lists all pods in "default" namespace: \n
```clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions)```
