# how to use

## prereqs

- skaffold, docker, kind

## setup
Create a KIND cluster:
```
$ kind cluster create
```

Setup skaffold to work with your kind cluster context:
```
$ skaffold config set --kube-context kubernetes-admin@kind local-cluster true
$ export KUBECONFIG="$(kind get kubeconfig-path)"
```

> Skaffold is smart enough to figure out your local cluster is using KIND, and
> will use `kind load` automatically.

If you aren't using the default kube context for kind, you can do something like:
```
$ skaffold config set --kube-context "$(kind get kubeconfig | yq '."current-context"' -r)" local-cluster true
```

## run
From here, just use normal skaffold commands:

```
$ skaffold run --tail
```
