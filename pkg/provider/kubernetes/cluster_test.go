/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes_test

import (
	"reflect"
	"testing"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"
	"k8c.io/kubermatic/v2/pkg/provider/kubernetes"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"

	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCreateCluster(t *testing.T) {
	// test data
	testcases := []struct {
		name                      string
		workerName                string
		existingKubermaticObjects []ctrlruntimeclient.Object
		project                   *kubermaticv1.Project
		userInfo                  *provider.UserInfo
		spec                      *kubermaticv1.ClusterSpec
		clusterType               string
		expectedCluster           *kubermaticv1.Cluster
		expectedError             string
		shareKubeconfig           bool
	}{
		{
			name:            "scenario 1, create kubernetes cluster",
			shareKubeconfig: false,
			workerName:      "test-kubernetes",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec:            genClusterSpec("test-k8s"),
			clusterType:     "kubernetes",
			existingKubermaticObjects: []ctrlruntimeclient.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedCluster: func() *kubermaticv1.Cluster {
				cluster := genCluster("test-k8s", "kubernetes", "my-first-project-ID", "test-kubernetes", "john@acme.com")
				cluster.ResourceVersion = "1"
				return cluster
			}(),
		},
		{
			name:            "scenario 2, create OpenShift cluster",
			shareKubeconfig: false,
			workerName:      "test-openshift",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec:            genClusterSpec("test-openshift"),
			clusterType:     "openshift",
			existingKubermaticObjects: []ctrlruntimeclient.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedCluster: func() *kubermaticv1.Cluster {
				cluster := genCluster("test-openshift", "openshift", "my-first-project-ID", "test-openshift", "john@acme.com")
				cluster.ResourceVersion = "1"
				return cluster
			}(),
		},
		{
			name:            "scenario 3, create kubernetes cluster when share kubeconfig is enabled and OIDC is set",
			shareKubeconfig: true,
			workerName:      "test-kubernetes",
			userInfo:        &provider.UserInfo{Email: "john@acme.com", Group: "owners-abcd"},
			project:         genDefaultProject(),
			spec: func() *kubermaticv1.ClusterSpec {
				spec := genClusterSpec("test-k8s")
				spec.OIDC = kubermaticv1.OIDCSettings{
					IssuerURL: "http://test",
					ClientID:  "test",
				}
				return spec
			}(),
			clusterType: "kubernetes",
			existingKubermaticObjects: []ctrlruntimeclient.Object{
				createAuthenitactedUser(),
				genDefaultProject(),
			},
			expectedError: "can not set OIDC for the cluster when share config feature is enabled",
		},
	}

	versions := kubermatic.NewDefaultVersions()

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			fakeClient := fakectrlruntimeclient.
				NewClientBuilder().
				WithScheme(scheme.Scheme).
				WithObjects(tc.existingKubermaticObjects...).
				Build()

			fakeImpersonationClient := func(impCfg restclient.ImpersonationConfig) (ctrlruntimeclient.Client, error) {
				return fakeClient, nil
			}

			// act
			target := kubernetes.NewClusterProvider(&restclient.Config{}, fakeImpersonationClient, nil, tc.workerName, nil, nil, nil, tc.shareKubeconfig, versions)
			partialCluster := &kubermaticv1.Cluster{}
			partialCluster.Spec = *tc.spec
			if tc.clusterType == "openshift" {
				partialCluster.Annotations = map[string]string{
					"kubermatic.io/openshift": "true",
				}
			}
			if tc.expectedCluster != nil {
				partialCluster.Finalizers = tc.expectedCluster.Finalizers
			}

			cluster, err := target.New(tc.project, tc.userInfo, partialCluster)
			if len(tc.expectedError) > 0 {
				if err == nil {
					t.Fatalf("expected error: %s", tc.expectedError)
				}
				if tc.expectedError != err.Error() {
					t.Fatalf("expected error: %s got %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}

				// override autogenerated field
				cluster.Name = tc.expectedCluster.Name
				cluster.Status.NamespaceName = tc.expectedCluster.Status.NamespaceName

				if !reflect.DeepEqual(cluster, tc.expectedCluster) {
					t.Fatalf("%v", diff.ObjectGoPrintSideBySide(tc.expectedCluster, cluster))
				}
			}

		})
	}
}
