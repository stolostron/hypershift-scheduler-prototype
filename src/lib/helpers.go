package lib

import (
	"fmt"
	"os"

	"github.com/go-logr/logr"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

//Asserts an error. If error exists, log error message and quit
//Requires:
//	*error (type error)
//	*message to append to error (type string)
//	*logger 
func AssertErr(err error, message string, logger logr.Logger) {
	if err != nil {
		logger.Error(err, message)
		os.Exit(1)
	}
}


//Switches context to a context specified by context
//Requires:
//	*Client Config
//	*Name of string to switch to
func SwitchContext(k *clientcmd.ClientConfig, context string) error {
	conf, _ := (*k).RawConfig()
	if conf.Contexts[context] == nil {
		fmt.Println(context + " doesn't exist")
		return fmt.Errorf("context: " + context + " doesn't exist")
	}

	conf.CurrentContext = context
	return clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), conf, true)
}
