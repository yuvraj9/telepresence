package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/datawire/k8sapi/pkg/k8sapi"
)

func RecordWorkloadInfo(span trace.Span, wl k8sapi.Workload) {
	if wl == nil {
		return
	}
	span.SetAttributes(
		attribute.String("tel2.workload-name", wl.GetName()),
		attribute.String("tel2.workload-namespace", wl.GetNamespace()),
		attribute.String("tel2.workload-kind", wl.GetKind()),
	)
}

// GetWorkload returns a workload for the given name, namespace, and workloadKind. The workloadKind
// is optional. A search is performed in the following order if it is empty:
//
//  1. Deployments
//  2. ReplicaSets
//  3. StatefulSets
//
// The first match is returned.
func GetWorkload(c context.Context, name, namespace, workloadKind string) (obj k8sapi.Workload, err error) {
	c, span := otel.GetTracerProvider().Tracer("").Start(c, "k8sapi.GetWorkload",
		trace.WithAttributes(
			attribute.String("tel2.workload-name", name),
			attribute.String("tel2.workload-namespace", namespace),
			attribute.String("tel2.workload-kind", workloadKind),
		),
	)
	defer EndAndRecord(span, err)

	return k8sapi.GetWorkload(c, name, namespace, workloadKind)
}
