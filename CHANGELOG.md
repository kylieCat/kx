# kx Changelog

### 0.4.1 - 2019-06-26

- Fix bug in rename command

### 0.4.0 - 2019-03-31

- New feature: ability to rename a context using `-r`. See README for example
- New feature: ability to swap back to a previous context/namespace. See README for example
- New feature: ability to list available color options and set them via a command
- General code cleanup

### 0.3.1 - 2019-01-05T02:37:13Z

- Bugfix: minikube config struct was wrong, caused broken files
- Stopped unnecessarily writing file after read-only ops.
- omit empty fields in kube conf

### 0.3.0- 2018-12-21T21:00:00Z

- Move some utility functions from the `cmd` package to `pkg`
- Expose their of their functionality via:
  
  - `GetKubeConfigPath`: Checks `$KUBECONFIG` for a kubeconfig path, if it's not found the default is used (`~/.kube/config`).
  - `GetKxConfigPath`: Checks `$KXONFIG` for a kxconfig path, if it's not found the default is used (`~/.kx.yaml`).
  - `GetDefaultKubeConfig`: Gets a `KubeConfig` using the kubeconfig in the default location.

- Moved constants to `pkg` and made them public:

  - `KubeConfigEnvVar`      = "KUBECONFIG"
  - `DefaultKubeConfigPath` = "~/.kube/config"
  - `KxConfigEnvVar`        = "KXCONFIG"
  - `DefaultKxConfigPath`   = "~/.kx.yaml"


### 0.2.1 - 2018-12-21T06:53:10Z

- Fix bug when users don't have a `.kx.yaml` file.
- Minor refactoring to make code a bit neater and more reusable.

### 0.2.0 - 2018-12-20T17:0:00Z 

- Add bash completion for contexts found in kubeconfig.

