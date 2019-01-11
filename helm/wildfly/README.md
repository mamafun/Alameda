# Wildfly Helm Chart

This Helm chart deploy wildfly with the image at https://hub.docker.com/r/elmehdi/wildfly

## Installing the Chart

To install the chart into `my-namespace` namespace:

```console
$ helm install --namespace my-namespace .
```

## Uninstalling the Chart

To uninstall/delete the my-release deployment:

```console
$ helm delete my-release --purge
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


