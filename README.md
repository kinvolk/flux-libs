# Flux CRD libraries

This repository has libraries to interact with Flux related CRDs like HelmRelease, GitRepository and HelmRepository.

## Example usage

```go
import (
	"fmt"

	helmrelease "github.com/kinvolk/flux-libs/lib/helm-release"
	gitrepocontroller "github.com/kinvolk/flux-libs/lib/source-controller/git-repo-controller"
	helmrepocontroller "github.com/kinvolk/flux-libs/lib/source-controller/helm-repo-controller"

	sourceapi "github.com/fluxcd/source-controller/api/v1beta1"
	helmreleaseapi "github.com/fluxcd/helm-controller/api/v2beta1"
)

func getHelmRelease() *helmreleaseapi.HelmRelease {}
func getGitRepository() *sourceapi.GitRepository {}
func getHelmRepository() *sourceapi.HelmRepository {}

// Note: The error handling is neglected to keep the code crisp.
func InstallComponent(kubeconfig []byte) error {
	gitRepoCfg, err = gitrepocontroller.NewGitRepoConfig(
		gitrepocontroller.WithKubeconfig(kubeconfig),
	)
	gitRepo := getGitRepository()
	err := gitRepoCfg.CreateOrUpdate(gitRepo)

	helmReleaseCfg, err = helmrelease.NewHelmReleaseConfig(
		helmrelease.WithKubeconfig(kubeconfig),
	)
	helmRelease := getHelmRelease()
	err := helmReleaseCfg.CreateOrUpdate(helmRelease)

	helmRepoCfg, err = helmrepocontroller.NewHelmRepoConfig(
		helmrepocontroller.WithKubeconfig(kubeconfig),
	)
	helmRepo := getHelmRepository()
	err := helmRepoCfg.CreateOrUpdate(helmRepo)

	return nil
}
```

## Contributing

Please check out the [contributing](./CONTRIBUTING.md) guide, and observe our [Code of Conduct](./CODE_OF_CONDUCT.md) when participating in this project.

## License

 This project is released under the [MIT license](./LICENSE.txt).
