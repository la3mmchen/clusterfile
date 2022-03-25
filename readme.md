# clusterfile

describe contents of kubernetes clusters. and sync them afterwards with helm and helmfile.

```
version: 1
clusters:
  - name: dev
    context: kubernetes-dev
    releases:
      - name: nginx
        version: 8.9.1
    envs:
      - name: web-apps
        location: helmfile/web-apps.yaml
      - name: backend
        location: helmfile/backend.yaml
  - name: test
    context: kubernetes-test
    envs:
      - name: backend
        location: helmfile/backend.yaml
```

/!\ work in progress /!\

## why? 

helmfile offers great possibilities to describe environments. The possibilities to describe multple clusters are somehow mingled via a `kubeContext:` attribute. 

Clusterfile builds a layer on top that makes it easy to describe kubernetes cluster contents top-down.

Helmfile builds environments that deploys versioned helm charts together. Sometimes there is more than one environment in a kubernetes cluster, e.g. multiple web-tiers, middle-warez and backends. Thats the point where bash scripts or ci/cd takes over and syncs all the environments. 

Clusterfile let you describe the environments per kubernetes cluster declarative and provides a binary to handle the helmfile subapp.

_to be continued_

## current state

- load a local clusterfile and its dependencies
- binary supports these helmfile subcommand:
  - lint
  - build
  - template
  - sync
  - one example clusterfile

lot more to do.

