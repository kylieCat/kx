# kx

A more convenient method for changing contexts and name spaces for kubectl

## Usage:

    kx [context [namespace]] [-a|--all] [-s|--set] [--no-color]

To add to prompt:

    $ export PROMPT_COMMAND=$(kx prompt)

## Examples:

Change context and namespace

    $ kx new-context new-namespace

Use a dash as the context name to stay in the same context but change namespace

    $ kx - new-namespace

Change context, keep the namespace set in .kube/config

    $ kctx new-context

Show current context and namespace

    $ kx
    current-context current-namespace

Show all contexts and namespaces

```
$ kx -a
Current	Context          	Namespace
-->    	foo-eks-prd  	        foo
        foo-gke-prd      	foo
        foo-gke-stg      	foo
```

## Todo

- Work with custom locations of kubeconfig
- Code clean up
- Better error messaging
- Add context names to bash completion