package lib_test

import (
	"testing"

	"hypershift-scheduler-prototype/src/lib"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)


func Test_SwitchContext(t *testing.T) {
	
}

func Test_GetManagedClusters(t *testing.T) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, _ := kubeconfig.ClientConfig()
	dynClient, _ := dynamic.NewForConfig(config)
	
	t.Log(lib.GetManagedClusters("", dynClient))

}

func Test_Filter(t *testing.T) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, _ := kubeconfig.ClientConfig()
	dynClient, _ := dynamic.NewForConfig(config)
	
	t.Log(lib.GetManagedClusters("", dynClient))

}

func Test_EnumHostedClusters(t *testing.T) {

}

func Test_SortHostingCluster(t *testing.T) {

}

