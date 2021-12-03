// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeKubermaticV1
}

var usersResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "users"}

var usersKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "User"}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *FakeUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(usersResource, name), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(usersResource, usersKind, opts), &kubermaticv1.UserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.UserList{ListMeta: obj.(*kubermaticv1.UserList).ListMeta}
	for _, item := range obj.(*kubermaticv1.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(usersResource, opts))
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Create(ctx context.Context, user *kubermaticv1.User, opts v1.CreateOptions) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(usersResource, user), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Update(ctx context.Context, user *kubermaticv1.User, opts v1.UpdateOptions) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(usersResource, user), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUsers) UpdateStatus(ctx context.Context, user *kubermaticv1.User, opts v1.UpdateOptions) (*kubermaticv1.User, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(usersResource, "status", user), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(usersResource, name), &kubermaticv1.User{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(usersResource, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.UserList{})
	return err
}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(usersResource, name, pt, data, subresources...), &kubermaticv1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.User), err
}
