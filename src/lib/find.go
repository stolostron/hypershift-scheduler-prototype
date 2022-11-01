package lib

import (
	"context"
	"sort"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/cluster-api/api/v1alpha4"

	hyper "github.com/openshift/hypershift/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterLoad struct {
	Load    int
	Cluster v1alpha4.Cluster
}

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
		//fmt.Println("-----" + unfiltered[cluster].Name + "-----")
		found = false
		for key, val := range unfiltered[cluster].Labels {

			//fmt.Println(key + ":" + val)

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

// Enumerate the number of hosted clusters on a hosting cluster
// Requires:
//
//	*A hosting cluster
//	*A dynamic client
//
// Note: This assumes the passed in cluster is a hosting cluster, does no checking at the moment
func EnumHostedClusters(hostingCluster v1alpha4.Cluster, client dynamic.Interface) (int, error) {
	hostedclusters := schema.GroupVersionResource{Group: "hypershift.openshift.io", Version: "v1alpha1", Resource: "hostedclusters"}

	res, err := client.Resource(hostedclusters).Namespace("clusters").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		return -1, err
	}

	unstructured := res.UnstructuredContent()
	var clusterlist hyper.HostedClusterList

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &clusterlist)

	if err != nil {
		return -1, err
	}

	return len(clusterlist.Items), nil
}

// Given a set of hosting clusters, return a sorted list of hosting clusters based on current hosted cluster load
// Requires:
//
//	*List of hosting clusters
//	*A boolean to list in ascending or descending order
//	true = asc
//	false = dsc
func SortHostingCluster(hostingClusterList []v1alpha4.Cluster, dynClient dynamic.Interface, k *clientcmd.ClientConfig, asc bool) ([]ClusterLoad, error) {
	var sortedList []ClusterLoad

	//Populate list

	for cluster := range hostingClusterList {
		client := dynClient
		var err error
		if hostingClusterList[cluster].Name != "local-cluster" {
			client, err = SwitchContext(k, hostingClusterList[cluster].Name)
			if err != nil {
				return nil, err
			}
		}

		l, err := EnumHostedClusters(hostingClusterList[cluster], client)
		if err != nil {
			return nil, err
		}
		toAppend := ClusterLoad{
			Load:    l,
			Cluster: hostingClusterList[cluster],
		}
		sortedList = append(sortedList, toAppend)
	}

	//Sort list
	if asc {
		sort.Slice(sortedList[:], func(i, j int) bool { return sortedList[i].Load < sortedList[j].Load})
	} else {
		sort.Slice(sortedList[:], func(i, j int) bool { return sortedList[i].Load > sortedList[j].Load })
	}

	return sortedList, nil
}
