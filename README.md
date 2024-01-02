
# Kubewell Ingress Operator


The Kubewell Ingress Operator is a custom Kubernetes component which deploys and manages one or more operations, which in turn handles Ingress traffic for applications running in a cluster.


To install a specific version of the Kubewell Ingress Controller with the operator, a specific version of the Kubewell Ingress Operator is required.

The following table shows the relation between the versions of the two projects:

| Kubewell Ingress Controller | Kubewell Ingress Operator |
| --- | --- |
| 2.1.x | 0.5.1 |
| 2.0.x | 0.4.0 |
| 1.12.x | 0.3.0 |
| 1.11.x | 0.2.0 |
| 1.10.x | 0.1.0 |
| 1.9.x | 0.0.7 |
| 1.8.x | 0.0.6 |
| 1.7.x | 0.0.4 |
| < 1.7.0 | N/A |

Note: The Kubewell Ingress Operator works only for Kubewell Ingress Controller versions after `1.0.0`.

## Getting Started

1. Install the Kubewell Ingress Operator. See [docs](./docs/installation.md).
    NOTE: To use TransportServers as part of your Kubewell Ingress Controller configuration, a GlobalConfiguration resource must be created *before* starting the Operator - [see the notes](./examples/deployment-oss-min/README.md#TransportServers)
1. Deploy a new Kubewell Ingress Controller using the [KubewellIngressController](docs/Kubewell-ingress-controller.md) Custom Resource:
    * For an Kubewell installation see the [Kubewell example](./examples/deployment-oss-min).
    * For an Kubewell Plus installation see the [Kubewell Plus example](./examples/deployment-plus-min).

## Upgrades

See [upgrade docs](./docs/upgrades)

## Kubewell Ingress Operator Releases
There are newly published Kubewell Ingress Operator releases on GitHub. See the [releases page](https://github.com/danielpickens/Kubewell-ingress-operator/releases).

The latest stable release is [1.0.0](https://github.com/danielpickens/Kubewell-ingress-operator/). For production use, it's recommend that you choose the latest stable release.

## Development

It is possible to run the operator in your local machine. This is useful for testing or during development.

### Run Operator locally

1. Have access to a Kubernetes/Openshift cluster.
1. Apply the latest CRDs:
   ```
    make install
    kubectl apply -f config/crd/
    ```
2. Run `make run`.

The operator will run in your local machine but will be communicating with the cluster.

### Update CRD

If any change is made in the CRD in the go code, run the following commands to update the changes in the CRD yaml:

1. `make manifests`
2. Apply the new CRD definition again in your cluster `make install`.

### Run tests

Run `make test` to run the full test suite including envtest, or `make unit-test` to run just the unit tests locally.

## Contributing

If you'd like to contribute to the project, please read the [Contributing](./CONTRIBUTING.md) guide.
