package apis

import (
	"github.com/che-incubator/che-workspace-operator/internal/cluster"
	"github.com/che-incubator/che-workspace-operator/pkg/apis/controller/v1alpha1"
	devworkspace "github.com/devfile/kubernetes-api/pkg/apis/workspaces/v1alpha1"
	oauthv1 "github.com/openshift/api/oauth/v1"
	routeV1 "github.com/openshift/api/route/v1"
	templateV1 "github.com/openshift/api/template/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
	AddToSchemes = append(AddToSchemes, devworkspace.SchemeBuilder.AddToScheme)
	if isOS, err := cluster.IsOpenShift(); isOS && err == nil {
		AddToSchemes = append(AddToSchemes,
			routeV1.Install,
			templateV1.Install,
			oauthv1.Install,
		)
	}
}