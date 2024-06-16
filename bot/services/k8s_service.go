package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/h3mmy/bloopyboi/bot/k8s"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type BotK8sMeta struct {
}

type K8sService struct {
	botMeta       *BotK8sMeta
	kubeClient    kubernetes.Interface
	metricsClient versioned.Clientset
	logger        *zap.Logger
}

func NewK8sService() *K8sService {
	return &K8sService{
		botMeta:    &BotK8sMeta{},
		kubeClient: k8s.GetClient(),
		logger: log.NewZapLogger().With(zapcore.Field{
			Key:    ServiceLoggerFieldKey,
			Type:   zapcore.StringType,
			String: "k8s",
		}),
	}
}

func (ks *K8sService) ListNamespaces(ctx context.Context) []coreV1.Namespace {
	nsList, err := ks.kubeClient.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		ks.logger.Sugar().Error("Error getting Namespaces: ", err)
		return nil
	}
	ks.logger.Sugar().Debug(fmt.Sprintf("List: %v", nsList.Items))
	return nsList.Items
}

func (ks *K8sService) GetPodMetrics(ctx context.Context) []v1beta1.PodMetrics {
	// ks.kubeClient.CoreV1().ServiceAccounts().
	namespace := "default"
	metricsList, err := ks.metricsClient.MetricsV1beta1().PodMetricses(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		ks.logger.Sugar().Error("Error getting podMetrics: ", err)
		return nil
	}
	ks.logger.Sugar().Debug(fmt.Sprintf("List: %v", metricsList.Items))
	return metricsList.Items
}

func (ks *K8sService) GetPodNamespace(ctx context.Context) (string, error) {
	ks.logger.Debug("Getting pod namespace")
	if ns := os.Getenv("POD_NAMESPACE"); ns != "" {
		ks.logger.Sugar().Debug("Found in os.env: ", ns)
		return ns, nil
	}
	if data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			ks.logger.Sugar().Debug("Found in serviceaccount: ", ns)
			return ns, nil
		}
		return "", err
	}
	return "", nil
}
