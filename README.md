# how to use

## prereqs

- skaffold, docker, kind

## setup
```
$ kind cluster create
$ skaffold config set --kube-context kubernetes-admin@kind local-cluster true
$ # skaffold config set --kube-context "$(kind get kubeconfig | yq '."current-context"' -r)" local-cluster true
$ export KUBECONFIG="$(kind get kubeconfig-path)"
```

> Skaffold is smart enough to figure out your local cluster is using KIND, and
> will use `kind load` automatically.

## run
```
$ skaffold run
```
