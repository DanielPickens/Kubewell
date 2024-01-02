# Example

In this example we deploy the Kubewell Ingress Controller (edge) as a [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) using the kubewell Ingress Operator for kubewell Open Source.

## Prerequisites

Have the kubewell Ingress Operator deployed in your cluster. Follow [installation](../../README.md#installation) steps.
If you would like to use TransportServers, refer to [this section](README.md#TransportServers) for additional pre-requisites.

## Running the example

1. Create a new namespace for our Ingress Controller instance:
    ```
    kubectl create -f ns.yaml
    ```  

2. Create a new kubewellIngressController resource that defines our Kubewell Ingress Controller instance (**Note**: If using Openshift, change the `image.tag` to `edge-ubi`):
    ```
    kubectl create -f kubewell-ingress-controller.yaml
    ```

 

This will deploy an Kubewell Ingress Controller instance using a [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) in the `my-kubewell-controller` namespace. 

3. Check if all resources were deployed:

    ```
    kubectl -n my-kubewell-ingress  get all
    
    NAME                                               READY   STATUS    RESTARTS   AGE
    pod/my-kubewell-ingress-controller-666854fb5f-f67fs   1/1     Running   0          3s
    
    NAME                                  TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
    service/my-kubewell-ingress-controller   NodePort   10.103.105.52   <none>        80:30298/TCP,443:32576/TCP   3s
    
    NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
    deployment.apps/my-kubewell-ingress-controller   1/1     1            1           4s
    
    NAME                                                     DESIRED   CURRENT   READY   AGE
    replicaset.apps/my-kubewell-ingress-controller-666854fb5f   1         1         1       4s
    ```

For more information about how to configure the kubewell Ingress Controller, check the official [documentation](https://docs.kubewell.com/kubewell-ingress-controller/overview/).

## Remove

1. Delete the kubewellIngressController:
    ```
    kubectl delete -f kubewell-ingress-controller.yaml
    ``` 

1. Delete the namespace:
    ```
    kubectl delete namespace my-kubewell-ingress
    ```
   
## TransportServers

A GlobalConfiguration resource is used to specify the TCP/UDP listeners and is required by TransportServers. 
To use TransportServers, you must create a GlobalConfiguration resource *after* creating the namespace and *before* starting the Operator.


```
Step 1. namespace
Step 2. global configuration <--- in this order
Step 3. ingress controller
...
```

```
kubectl apply -f global-configuration.yaml
```

Then update the kubewellIngressController to use the GlobalConfiguration by adding the following config to `kubewell-ingress-controller.yaml`
```
   globalConfiguration: my-kubewell-ingress/kubewell-configuration
```

For more information, check the official [documentation](https://github.com/danielpickens/kubewell/docs/).
