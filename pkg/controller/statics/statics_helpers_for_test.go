package statics

import (
	"testing"

	util "2uasimojo/efs-csi-operator/pkg/util"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"golang.org/x/net/context"

	securityv1 "github.com/openshift/api/security/v1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	expectedNumStatics = 5
)

// Common options we're passing into cmp.Diff.
var diffOpts cmp.Options

func init() {
	diffOpts = cmp.Options{
		// We want to ignore TypeMeta in all cases, because it's a given of the type itself.
		cmpopts.IgnoreTypes(metav1.TypeMeta{}),
		// We ignore the ResourceVersion because it gets set by the server and is unpredictable/opaque.
		// We ignore labels *in cmp.Diff* because sometimes we're checking a virgin resource definition
		// from a getter (label validation is done separately).
		cmpopts.IgnoreFields(metav1.ObjectMeta{}, "ResourceVersion", "Labels"),
	}
}

// checkNumStatics is a helper to guard against static resources being added in the future without tests
// being updated. Use it from any test that would need to be fixed if new statics are added.
func checkNumStatics(t *testing.T) {
	if numStatics := len(staticResources); numStatics != expectedNumStatics {
		t.Fatalf("Test update needed! Expected %d static resources but found %d.",
			expectedNumStatics, numStatics)
	}
}

// checkStatics queries the client for all the known static resources, verifying that they exist
// and have the expected content. It returns a map, keyed by the short name of the resource type
// (e.g. "SecurityContextConstraints") of the runtime.Object returned by the client for each resource.
func checkStatics(t *testing.T, client crclient.Client) map[string]runtime.Object {
	ret := make(map[string]runtime.Object)
	ctx := context.TODO()

	for _, i := range []struct {
		name   string
		obj    runtime.Object
		nsname types.NamespacedName
	}{
		{
			"ServiceAccount",
			&corev1.ServiceAccount{},
			types.NamespacedName{Name: serviceAccountName, Namespace: namespaceName},
		},
		{
			"SecurityContextConstraints",
			&securityv1.SecurityContextConstraints{},
			types.NamespacedName{Name: sccName},
		},
		{
			"DaemonSet",
			&appsv1.DaemonSet{},
			types.NamespacedName{Name: daemonSetName, Namespace: namespaceName},
		},
		{
			"CSIDriver",
			&storagev1beta1.CSIDriver{},
			types.NamespacedName{Name: CSIDriverName},
		},
		{
			"StorageClass",
			&storagev1.StorageClass{},
			types.NamespacedName{Name: StorageClassName},
		},
	} {
		if err := client.Get(ctx, i.nsname, i.obj); err != nil {
			t.Fatalf("Couldn't get %s: %v", i.name, err)
		}
		diff := cmp.Diff(findStatic(i.nsname).(*util.EnsurableImpl).Definition, i.obj, diffOpts...)
		if diff != "" {
			t.Fatal("Objects differ: -expected, +actual\n", diff)
		}
		if !util.DoICare(i.obj) {
			t.Fatalf("Missing label for %s", i.name)
		}
		ret[i.name] = i.obj
	}

	return ret
}