package helmrepocontroller

import (
	"context"
	"fmt"

	api "github.com/fluxcd/source-controller/api/v1beta1"
	"github.com/kinvolk/fluxlib/lib"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type HelmRepoConfig struct {
	c          client.Client
	kubeconfig []byte
}

type helmRepoConfigOpt func(*HelmRepoConfig)

var scheme *runtime.Scheme

func init() {
	scheme = runtime.NewScheme()
	_ = api.AddToScheme(scheme)
}

func NewHelmRepoConfig(fns ...helmRepoConfigOpt) (*HelmRepoConfig, error) {
	var ret HelmRepoConfig

	for _, fn := range fns {
		fn(&ret)
	}

	var err error

	ret.c, err = lib.GetKubernetesClient(ret.kubeconfig, scheme)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func WithKubeconfig(kubeconfig []byte) helmRepoConfigOpt {
	return func(hr *HelmRepoConfig) {
		hr.kubeconfig = kubeconfig
	}
}

func (h *HelmRepoConfig) Get(name, ns string) (*api.HelmRepository, error) {
	var got api.HelmRepository

	if err := h.c.Get(context.Background(), types.NamespacedName{
		Namespace: ns,
		Name:      name,
	}, &got); err != nil {
		return nil, fmt.Errorf("getting HelmRepository: %w", err)
	}

	return &got, nil
}

func (h *HelmRepoConfig) List(listOpts *client.ListOptions) (*api.HelmRepositoryList, error) {
	var got api.HelmRepositoryList

	if err := h.c.List(context.TODO(), &got, listOpts); err != nil {
		return nil, fmt.Errorf("listing HelmRepositories: %w", err)
	}

	return &got, nil
}

func (h *HelmRepoConfig) CreateOrUpdate(hr *api.HelmRepository) error {
	var got api.HelmRepository

	if err := h.c.Get(context.Background(), types.NamespacedName{
		Namespace: hr.GetNamespace(),
		Name:      hr.GetName(),
	}, &got); err != nil {
		if errors.IsNotFound(err) {
			// Create the object since it does not exists.
			if err := h.c.Create(context.Background(), hr); err != nil {
				return fmt.Errorf("creating HelmRepository: %w", err)
			}

			return nil
		}

		return fmt.Errorf("looking up HelmRepository: %w", err)
	}

	hr.ResourceVersion = got.ResourceVersion

	if err := h.c.Update(context.Background(), hr); err != nil {
		return fmt.Errorf("updating HelmRepository: %w", err)
	}

	return nil
}

func (h *HelmRepoConfig) Delete(hr *api.HelmRepository) error {
	if err := h.c.Delete(context.Background(), hr); err != nil {
		return fmt.Errorf("deleting HelmRepository: %w", err)
	}

	return nil
}
