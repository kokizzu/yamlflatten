
# YAML Flatten

converts this kind of structure:

```yaml
foo:
  bar:
    baz: true
```

to this kind of structure:

```
foo.bar.baz = true
```

## Usage
```
go mod tidy
go run yamlflatten.go example.yaml

# example:
go run yamlflatten.go example.yaml | grep 'enabled = true'
kubelet.enabled = true
coreDns.enabled = true
kubeStateMetrics.enabled = true
prometheus.enabled = true
kubeScheduler.service.enabled = true
kubeScheduler.serviceMonitor.enabled = true
kubeScheduler.enabled = true
kube-state-metrics.prometheus.monitor.enabled = true
kubeProxy.serviceMonitor.enabled = true
kubeProxy.enabled = true
kubeProxy.service.enabled = true
kubeEtcd.enabled = true
kubeEtcd.service.enabled = true
kubeEtcd.serviceMonitor.enabled = true
grafana.enabled = true
grafana.sidecar.dashboards.enabled = true
grafana.sidecar.datasources.alertmanager.enabled = true
grafana.sidecar.datasources.enabled = true
grafana.serviceMonitor.enabled = true
kubeApiServer.enabled = true
kubeControllerManager.enabled = true
kubeControllerManager.service.enabled = true
kubeControllerManager.serviceMonitor.enabled = true
nodeExporter.enabled = true
prometheus-node-exporter.prometheus.monitor.enabled = true
prometheusOperator.tls.enabled = true
prometheusOperator.kubeletService.enabled = true
prometheusOperator.admissionWebhooks.patch.enabled = true
prometheusOperator.admissionWebhooks.enabled = true
prometheusOperator.enabled = true
alertmanager.enabled = true
kubernetesServiceMonitors.enabled = true
```
