package main

// import (
// 	"hypershift-scheduler-prototype/src/lib"
// 	"log"

// 	clusterv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
// 	appsv1 "k8s.io/api/apps/v1"
// 	corev1 "k8s.io/api/core/v1"
// 	"open-cluster-management.io/addon-framework/pkg/addonfactory"
// 	"open-cluster-management.io/addon-framework/pkg/addonmanager"
// 	"open-cluster-management.io/addon-framework/pkg/agent"
// 	"open-cluster-management.io/addon-framework/pkg/utils"
// 	kubecr "sigs.k8s.io/controller-runtime"
// )

import (
	//ctrl "sigs.k8s.io/controller-runtime"

	"os"
	"context"
	
	corev1 "k8s.io/api/core/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	clusterv1beta1 "open-cluster-management.io/api/cluster/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	
)

var log = logf.Log.WithName("main")

func main() {

	logf.SetLogger(zap.New())

	//Creating a new manager
	
	//Gets current kubeconfig
	cfg, err := config.GetConfig();
	if err != nil {
		//Usually occurs when container is not running or paused
		log.Error(err, "unable to get kubeconfig");
		os.Exit(1);
	}

	//Creates a new manager. Can create controllers with this manager
	mgr, err := manager.New(cfg, manager.Options{});
	if err != nil {
		log.Error(err, "unable to set up manager");
		os.Exit(1);
	}
	log.Info("created manager", "manager", mgr);


	//Create a new controller
	c, err := controller.New("pod-controller", mgr, controller.Options{
		Reconciler: reconcile.Func(func(context.Context, reconcile.Request) (reconcile.Result, error) {
			// Your business logic to implement the API by creating, updating, deleting objects goes here.
			return reconcile.Result{}, nil
		}),
	})

	if err != nil {
		log.Error(err, "unable to set up pod-controller");
		os.Exit(1);
	}

	// Watch for Pod create / update / delete events and call Reconcile
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		log.Error(err, "unable to watch pods")
		os.Exit(1)
	}

	// Start the Controller through the manager.
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "unable to continue running manager")
		os.Exit(1)
	}

	
	
	
}