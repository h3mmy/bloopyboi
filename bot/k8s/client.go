package k8s

import (
	"os"

	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

var (
	logger = log.NewZapLogger().Named("k8s")
)

// GetClient returns a k8s clientset
func GetClient() kubernetes.Interface {
	var kubeClient kubernetes.Interface
	_, err := rest.InClusterConfig()
	if err != nil {
		logger.Warn("error building in-cluster client", zap.Error(err))
		kubeClient = getClientOutOfCluster()
	} else {
		kubeClient = getClientInCluster()
	}

	return kubeClient
}

func GetMetricsClient() versioned.Clientset {
	var kubeClient versioned.Clientset
	_, err := rest.InClusterConfig()
	if err != nil {
		logger.Warn("error building in-cluster metrics-client", zap.Error(err))
		kubeClient = *getMetricsClientOutOfCluster()
	} else {
		kubeClient = *getMetricsClientInCluster()
	}

	return kubeClient
}

// GetClientInCluster returns a k8s clientset to the request from inside of cluster
func getClientInCluster() kubernetes.Interface {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Sugar().Errorf("Can not get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Sugar().Errorf("Can not create kubernetes client: %v", err)
	}

	return clientset
}

// GetClientInCluster returns a k8s clientset to the request from inside of cluster
func getMetricsClientInCluster() *versioned.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Sugar().Errorf("Can not get kubernetes config: %v", err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		logger.Sugar().Errorf("Can not create metrics client: %v", err)
	}

	return clientset
}

func buildOutOfClusterConfig() (*rest.Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
	}
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

// GetClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func getClientOutOfCluster() kubernetes.Interface {
	config, err := buildOutOfClusterConfig()
	if err != nil {
		logger.Sugar().Errorf("Cannot get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		logger.Sugar().Errorf("Cannot create new kubernetes client from config: %v", err)
	}

	return clientset
}

// GetClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func getMetricsClientOutOfCluster() *versioned.Clientset {
	config, err := buildOutOfClusterConfig()
	if err != nil {
		logger.Sugar().Errorf("Cannot get kubernetes config: %v", err)
	}

	clientset, err := versioned.NewForConfig(config)

	if err != nil {
		logger.Sugar().Errorf("Cannot create new metrics client from config: %v", err)
	}

	return clientset
}
