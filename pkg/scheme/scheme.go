package scheme

import (
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	scyllav1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	monitoringv1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	cqlclientv1alpha1 "github.com/scylladb/scylla-operator/pkg/scylla/api/cqlclient/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	Scheme                = runtime.NewScheme()
	Codecs                = serializer.NewCodecFactory(Scheme, serializer.EnableStrict)
	DefaultYamlSerializer = json.NewSerializerWithOptions(
		json.DefaultMetaFactory,
		Scheme,
		Scheme,
		json.SerializerOptions{
			Yaml:   true,
			Pretty: false,
			Strict: true,
		},
	)
)

func init() {
	utilruntime.Must(kscheme.AddToScheme(Scheme))

	utilruntime.Must(scyllav1.Install(Scheme))
	utilruntime.Must(scyllav1alpha1.Install(Scheme))

	utilruntime.Must(cqlclientv1alpha1.Install(Scheme))

	utilruntime.Must(monitoringv1.Install(Scheme))
}
