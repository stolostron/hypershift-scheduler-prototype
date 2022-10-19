package lib

import (
	"context"
	"fmt"

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
//	*A dynamic Client
func GetManagedClusters(namespace string, client dynamic.Interface) ([]v1alpha4.Cluster, error) {
	managedClusters := schema.GroupVersionResource{Group: "cluster.open-cluster-management.io", Version: "v1", Resource: "managedclusters"}
	res, err := client.Resource(managedClusters).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	//res, err := client.Resource(managedClusters).Namespace(namespace).Get(context.Background(), "local-cluster", metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	unstructured := res.UnstructuredContent()
	var clusterlist v1alpha4.ClusterList

	var clusters []v1alpha4.Cluster

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &clusterlist)

	if err != nil {
		return nil, err
	}

	for c := range clusterlist.Items {
		clusters = append(clusters, clusterlist.Items[c])
	}

	return clusters, nil

}

// Given a set, return a subset filtered by a given label
//
// Requires:
//
//	*A label used for filtering. Need both key and value
//	*A set to filter from
//	*intersect (a boolean). If interset is true, only objects INCLUDING the label will be returned. Otherwise, only objects EXCLUDING the value will be returned
func Filter(labelKey string, labelVal string, unfiltered []v1alpha4.Cluster, intersect bool) ([]v1alpha4.Cluster, error) {
	var filteredClusters []v1alpha4.Cluster
	var found bool
	for cluster := range unfiltered {
		fmt.Println("-----" + unfiltered[cluster].Name + "-----")
		found = false
		for key, val := range unfiltered[cluster].Labels {

			fmt.Println(key + ":" + val)

			if key == labelKey && val == labelVal {
				found = true
				if intersect {
					filteredClusters = append(filteredClusters, unfiltered[cluster])

				}

				break
			}

		}

		if !intersect && !found {
			filteredClusters = append(filteredClusters, unfiltered[cluster])
		}
	}
	return filteredClusters, nil
}
