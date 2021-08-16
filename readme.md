# clusterfile

describe contents of kubernetes clusters. and sync them afterwards with helmfile.

## why? 

helmfile offers great possibilities to describe environments. The possibilities to describe multple clusters are somehow mingled via a `kubeContext:` attribute. 

Clusterfile builds a layer on top that makes it easy to describe kubernetes cluster contents top-down.

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

