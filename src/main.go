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

	"hypershift-scheduler-prototype/src/lib"

	//"hypershift-scheduler-prototype/src/lib"

	"k8s.io/client-go/tools/clientcmd"

	//"k8s.io/client-go/scale/scheme/appsv1beta1"

	logf "sigs.k8s.io/controller-runtime/pkg/log"

	//"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var logger = logf.Log.WithName("main")

func main() {

	logf.SetLogger(zap.New())

	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()

	lib.AssertErr(err, "unable to get kubeconfig", logger)

	dynClient, err := dynamic.NewForConfig(config)

	lib.AssertErr(err, "unable to create client", logger)

	managedclusters, err := lib.GetManagedClusters("", dynClient)

	lib.AssertErr(err, "unable to get managed clusters", logger)

	lib.Filter("feature.open-cluster-management.io/addon-hypershift-addon", "available", managedclusters, true)

	lib.AssertErr(err, "unable to filter clusters", logger)

	//mcl := ocm.NewManagedClusterLister(cache.NewIndexer())

	// //Creating a new manager

	// //Creates a new manager. Can create controllers with this manager
	// mgr, err := manager.New(config, manager.Options{})
	// if err != nil {
	// 	log.Error(err, "unable to set up manager")
	// 	os.Exit(1)
	// }
	// log.Info("created manager", "manager", mgr)

	// //Create a new controller
	// controller, err := controller.New("pod-controller", mgr, controller.Options{
	// 	Reconciler: reconcile.Func(func(context.Context, reconcile.Request) (reconcile.Result, error) {
	// 		// Your business logic to implement the API by creating, updating, deleting objects goes here.
	// 		return reconcile.Result{}, nil
	// 	}),
	// })

	// if err != nil {
	// 	log.Error(err, "unable to set up pod-controller")
	// 	os.Exit(1)
	// }

	// // Watch for Pod create / update / delete events and call Reconcile
	// err = controller.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{})
	// if err != nil {
	// 	log.Error(err, "unable to watch pods")
	// 	os.Exit(1)
	// }

	// // Start the Controller through the manager.
	// if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
	// 	log.Error(err, "unable to continue running manager")
	// 	os.Exit(1)
	// }

	//Get hosting clusters (from where????)
	//fmt.Println("Poggers we got here")

}
