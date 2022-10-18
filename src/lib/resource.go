package lib

import (
	clientcmd "k8s.io/client-go/tools/clientcmd"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd/api"
)


//Returns the current kube context located at pathToKubeconfig
func GetContext(pathToKubeConfig string) (api.Config, error) {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
        &clientcmd.ClientConfigLoadingRules{ExplicitPath: pathToKubeConfig},
        &clientcmd.ConfigOverrides{
            CurrentContext: "",
        }).RawConfig();
	
    
	return config, err;
}

//Returns the following information on CPUs in the specified node
//1. Number of CPUs
//2. Utilization of each CPU (as a percentage)
//3. Speed of each CPU
// func GetCPU(*corev1.Node) {
// 	config, err := GetContext("/home/ofarag/.kube/config")
// 	clientcmd.BuildConfigFromFlags()
// }

func GetRAM(*corev1.Node) {
	
}

func GetGPU(*corev1.Node) {

}