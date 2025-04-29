# Helm Values Diff

Utility cli tool to take a base helm chart values file, and another values file with changes based on the based,
and output just the changes in a diff that can be used with fluxcd, and other orchestration tools.

## Usage

```bash

# Install the tool

brew tap prom3theu5/tap
brew install helm-values-diff

# Or download the binary from the releases page
```

```bash
helm-values-diff base.yaml new.yaml > output.yaml
```

## Example

#### base.yaml
```yaml
# Default values for Traefik
# This is a YAML-formatted file.
# Declare variables to be passed into templates

image:  # @schema additionalProperties: false
  # -- Traefik image host registry
  registry: docker.io
  # -- Traefik image repository
  repository: traefik
  # -- defaults to appVersion
  tag:  # @schema type:[string, null]
  # -- Traefik image pull policy
  pullPolicy: IfNotPresent

# -- Add additional label to all resources
commonLabels: {}

deployment:
  # -- Enable deployment
  enabled: true
  # -- Deployment or DaemonSet
  kind: Deployment
  # -- Number of pods of the deployment (only applies when kind == Deployment)
  replicas: 1
  # -- Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
  revisionHistoryLimit:  # @schema type:[integer, null];minimum:0
  # -- Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
  terminationGracePeriodSeconds: 60
  # -- The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
  minReadySeconds: 0
  ## -- Override the liveness/readiness port. This is useful to integrate traefik
  ## with an external Load Balancer that performs healthchecks.
  ## Default: ports.traefik.port
  healthchecksPort:  # @schema type:[integer, null];minimum:0
  ## -- Override the liveness/readiness host. Useful for getting ping to respond on non-default entryPoint.
  ## Default: ports.traefik.hostIP if set, otherwise Pod IP
  healthchecksHost: ""
  ## -- Override the liveness/readiness scheme. Useful for getting ping to
  ## respond on websecure entryPoint.
  healthchecksScheme:   # @schema enum:[HTTP, HTTPS, null]; type:[string, null]; default: HTTP
  ## -- Override the readiness path.
  ## Default: /ping
  readinessPath: ""
  # -- Override the liveness path.
  # Default: /ping
  livenessPath: ""
  # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
  annotations: {}
  # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
  labels: {}
  # -- Additional pod annotations (e.g. for mesh injection or prometheus scraping)
  # It supports templating. One can set it with values like traefik/name: '{{ template "traefik.name" . }}'
  podAnnotations: {}
  # -- Additional Pod labels (e.g. for filtering Pod by custom labels)
  podLabels: {}
  # -- Additional containers (e.g. for metric offloading sidecars)
  additionalContainers: []
```

#### new.yaml
```yaml
# Default values for Traefik
# This is a YAML-formatted file.
# Declare variables to be passed into templates

image:  # @schema additionalProperties: false
  # -- Traefik image host registry
  registry: my-awesome-mirror.io
  # -- Traefik image repository
  repository: traefik
  # -- defaults to appVersion
  tag:  # @schema type:[string, null]
  # -- Traefik image pull policy
  pullPolicy: IfNotPresent

# -- Add additional label to all resources
commonLabels: {}

deployment:
  # -- Enable deployment
  enabled: true
  # -- Deployment or DaemonSet
  kind: DaemonSet
  # -- Number of pods of the deployment (only applies when kind == Deployment)
  replicas: 2
  # -- Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
  revisionHistoryLimit: 10
  # -- Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
  terminationGracePeriodSeconds: 60
  # -- The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
  minReadySeconds: 0
  ## -- Override the liveness/readiness port. This is useful to integrate traefik
  ## with an external Load Balancer that performs healthchecks.
  ## Default: ports.traefik.port
  healthchecksPort:  # @schema type:[integer, null];minimum:0
  ## -- Override the liveness/readiness host. Useful for getting ping to respond on non-default entryPoint.
  ## Default: ports.traefik.hostIP if set, otherwise Pod IP
  healthchecksHost: ""
  ## -- Override the liveness/readiness scheme. Useful for getting ping to
  ## respond on websecure entryPoint.
  healthchecksScheme:   # @schema enum:[HTTP, HTTPS, null]; type:[string, null]; default: HTTP
  ## -- Override the readiness path.
  ## Default: /ping
  readinessPath: ""
  # -- Override the liveness path.
  # Default: /ping
  livenessPath: ""
  # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
  annotations: {}
  # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
  labels:
    i_like: "chocolate"
  # -- Additional pod annotations (e.g. for mesh injection or prometheus scraping)
  # It supports templating. One can set it with values like traefik/name: '{{ template "traefik.name" . }}'
  podAnnotations: {}
  # -- Additional Pod labels (e.g. for filtering Pod by custom labels)
  podLabels: {}
  # -- Additional containers (e.g. for metric offloading sidecars)
  additionalContainers: []
```

#### output.yaml
```yaml
deployment:
  kind: DaemonSet
  labels:
    i_like: chocolate
  replicas: 2
  revisionHistoryLimit: 10
image:
  registry: my-awesome-mirror.io
```