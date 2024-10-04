// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/user/v1"
	userv1 "github.com/openshift/client-go/user/applyconfigurations/user/v1"
	scheme "github.com/openshift/client-go/user/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// IdentitiesGetter has a method to return a IdentityInterface.
// A group's client should implement this interface.
type IdentitiesGetter interface {
	Identities() IdentityInterface
}

// IdentityInterface has methods to work with Identity resources.
type IdentityInterface interface {
	Create(ctx context.Context, identity *v1.Identity, opts metav1.CreateOptions) (*v1.Identity, error)
	Update(ctx context.Context, identity *v1.Identity, opts metav1.UpdateOptions) (*v1.Identity, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Identity, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IdentityList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Identity, err error)
	Apply(ctx context.Context, identity *userv1.IdentityApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Identity, err error)
	IdentityExpansion
}

// identities implements IdentityInterface
type identities struct {
	*gentype.ClientWithListAndApply[*v1.Identity, *v1.IdentityList, *userv1.IdentityApplyConfiguration]
}

// newIdentities returns a Identities
func newIdentities(c *UserV1Client) *identities {
	return &identities{
		gentype.NewClientWithListAndApply[*v1.Identity, *v1.IdentityList, *userv1.IdentityApplyConfiguration](
			"identities",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.Identity { return &v1.Identity{} },
			func() *v1.IdentityList { return &v1.IdentityList{} }),
	}
}
