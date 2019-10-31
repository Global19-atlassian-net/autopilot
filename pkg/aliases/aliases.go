package aliases

import appsv1 "k8s.io/api/apps/v1"
import corev1 "k8s.io/api/core/v1"
import trafficsplitv1alpha2 "github.com/deislabs/smi-sdk-go/pkg/apis/split/v1alpha2"

type (
	Deployment = appsv1.Deployment
	Service = corev1.Service
	TrafficSplit = trafficsplitv1alpha2.TrafficSplit

	Deployments = []*Deployment
	Services = []*Service
	TrafficSplits = []*TrafficSplit
)