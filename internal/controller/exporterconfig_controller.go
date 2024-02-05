/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	networkingv1alpha1 "git.stiil.dk/traefik-external-config/api/v1alpha1"
	"k8s.io/apimachinery/pkg/fields" // Required for Watching
	"k8s.io/apimachinery/pkg/types"  // Required for Watching

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"   // Required for Watching
	"sigs.k8s.io/controller-runtime/pkg/handler"   // Required for Watching
	"sigs.k8s.io/controller-runtime/pkg/predicate" // Required for Watching
	"sigs.k8s.io/controller-runtime/pkg/reconcile" // Required for Watching
)

// ExporterConfigReconciler reconciles a ExporterConfig object
type ExporterConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	//Recorder is importaint for events.
	Recorder record.EventRecorder
	Config   *LockableConfig
}

const (
	//Where to find secret name in spec (for watcher)
	serviceNameField = ".spec.cluster.serviceName"
)

//+kubebuilder:rbac:groups=networking.stiil.dk,resources=exporterconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.stiil.dk,resources=exporterconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.stiil.dk,resources=exporterconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ExporterConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *ExporterConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// Watcher handler for Secrets
func (r *ExporterConfigReconciler) findObjectForService(ctx context.Context, service client.Object) []reconcile.Request {
	attachedExporterConfigs := &networkingv1alpha1.ExporterConfigList{}
	listOps := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(serviceNameField, service.GetName()),
		Namespace:     "",
	}
	err := r.List(ctx, attachedExporterConfigs, listOps)
	if err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(attachedExporterConfigs.Items))
	for i, item := range attachedExporterConfigs.Items {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: "",
			},
		}
	}
	return requests
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExporterConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1alpha1.ExporterConfig{}).
		Watches(
			&corev1.Service{},
			handler.EnqueueRequestsFromMapFunc(r.findObjectForService),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}
