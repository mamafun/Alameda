package k8swhsrv

import (
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	operatorwebhook "github.com/containers-ai/alameda/operator/pkg/webhook"
	"github.com/containers-ai/alameda/pkg/utils/kubernetes"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	osappsapi "github.com/openshift/api/apps/v1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	extensionsv1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/builder"
)

var scope = logUtil.RegisterScope("k8s_webhook_server", "K8S Webhook Server.", 0)

type K8SWebhookServer struct {
	manager *manager.Manager
	config  *Config
}

func NewK8SWebhookServer(mgr *manager.Manager, config *Config) *K8SWebhookServer {
	return &K8SWebhookServer{
		manager: mgr,
		config:  config,
	}
}

func (srv *K8SWebhookServer) Launch() {
	srv.registerWebhooks()
}

func (srv *K8SWebhookServer) registerWebhooks() {
	webhooks := []webhook.Webhook{}
	svr, err := webhook.NewServer("operator-k8s-admission-server", *srv.manager, webhook.ServerOptions{
		Port:    srv.config.Port,
		CertDir: srv.config.CertDir,
		BootstrapOptions: &webhook.BootstrapOptions{
			ValidatingWebhookConfigName: srv.config.ValidatingWebhookConfigName,
			MutatingWebhookConfigName:   srv.config.MutatingWebhookConfigName,
			/*
				Secret: &types.NamespacedName{
					Namespace: srv.config.Secret.Namespace,
					Name:      srv.config.Secret.Name,
				},
			*/
			Service: &webhook.Service{
				Namespace: srv.config.Service.Namespace,
				Name:      srv.config.Service.Name,
				// Selectors should select the pods that runs this webhook server.
				Selectors: map[string]string{
					"app": "operator",
				},
			},
		},
	})
	if err != nil {
		scope.Errorf(err.Error())
		return
	}
	wh, err := builder.NewWebhookBuilder().Name("deployment.validating.containers.ai").
		NamespaceSelector(&metav1.LabelSelector{}).Validating().
		Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update).
		ForType(&extensionsv1.Deployment{}).
		Handlers(operatorwebhook.GetDeploymentHandler()).
		WithManager(*srv.manager).
		Build()
	if err != nil {
		scope.Errorf(err.Error())
	} else {
		webhooks = append(webhooks, wh)
	}

	okdCluster := false
	if okdCluster, err = kubernetes.IsOKDCluster(); err != nil {
		scope.Error(err.Error())
	}
	if okdCluster {
		wh, err = builder.NewWebhookBuilder().Name("deploymentconfig.validating.containers.ai").
			NamespaceSelector(&metav1.LabelSelector{}).Validating().
			Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update).
			ForType(&osappsapi.DeploymentConfig{}).
			Handlers(operatorwebhook.GetDeploymentConfigHandler()).
			WithManager(*srv.manager).
			Build()
		if err != nil {
			scope.Errorf(err.Error())
		} else {
			webhooks = append(webhooks, wh)
		}
	}

	wh, err = builder.NewWebhookBuilder().Name("alamedascaler.validating.containers.ai").
		NamespaceSelector(&metav1.LabelSelector{}).Validating().
		Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update).
		ForType(&autoscalingv1alpha1.AlamedaScaler{}).
		Handlers(operatorwebhook.GetAlamedaScalerHandler()).
		WithManager(*srv.manager).
		Build()
	if err != nil {
		scope.Errorf(err.Error())
	} else {
		webhooks = append(webhooks, wh)
	}
	svr.Register(webhooks...)
}