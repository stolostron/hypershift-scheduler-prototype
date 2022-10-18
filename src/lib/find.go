package lib

import (
	"context"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/cluster-api/api/v1alpha4"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Returns a list of managed clusters from current kubeconfig.
//
// Requires:
//
//	*Namespace (set to "" for all namespaces)
//	*Dynamic Client
func GetManagedClusters(namespace string, client dynamic.Interface) []v1alpha4.Cluster {
	managedClusters := schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1", Resource: "managedclusters"}
	res, err := client.Resource(managedClusters).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	//res, err := client.Resource(managedClusters).Namespace(namespace).Get(context.Background(), "local-cluster", metav1.GetOptions{})

	if err != nil {
		//log.Error(err, "unable to get resources")
		os.Exit(1)
	}

	unstructured := res.UnstructuredContent()
	var clusterlist v1alpha4.ClusterList

	var clusters []v1alpha4.Cluster

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &clusterlist)

	for c := range clusterlist.Items {
		clusters = append(clusters, clusterlist.Items[c])
	}

	if err != nil {
		//log.Error(err, "unable to get resources")
		os.Exit(1)
	}

	return clusters

}
