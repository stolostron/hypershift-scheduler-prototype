package lib

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"k8s.io/client-go/dynamic"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

// Asserts an error. If error exists, log error message and quit
// Requires:
//
//	*error (type error)
//	*message to append to error (type string)
//	*logger
func AssertErr(err error, message string, logger logr.Logger) {
	if err != nil {
		logger.Error(err, message)
		os.Exit(1)
	}
}

// Switches context to a context specified by context
// Requires:
//
//	*Client Config
//	*Name of string to switch to
func SwitchContext(k *clientcmd.ClientConfig, context string) (dynamic.Interface, error) {
	conf, _ := (*k).RawConfig()
	if conf.Contexts[context] == nil {
		fmt.Println(context + " doesn't exist")
		return nil, fmt.Errorf("context: " + context + " doesn't exist")
	}

	conf.CurrentContext = context

	kubeConf, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}).ClientConfig()

	if err != nil {
		return nil, err
	}
	client, err := dynamic.NewForConfig(kubeConf)
	if err != nil {
		return nil, err
	}
	return client, nil

}
