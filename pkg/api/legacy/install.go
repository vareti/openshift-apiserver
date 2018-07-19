package legacy

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupName            = ""
	GroupVersion         = schema.GroupVersion{Group: GroupName, Version: "v1"}
	internalGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
)

// DEPRECATED
func Kind(kind string) schema.GroupKind {
	return schema.GroupKind{Group: GroupName, Kind: kind}
}

// DEPRECATED
func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

func InstallLegacyInternalAll(scheme *runtime.Scheme) {
	InstallInternalLegacyApps(scheme)
	InstallInternalLegacyAuthorization(scheme)
	InstallInternalLegacyBuild(scheme)
	InstallInternalLegacyImage(scheme)
	InstallInternalLegacyNetwork(scheme)
	InstallInternalLegacyOAuth(scheme)
	InstallInternalLegacyProject(scheme)
	InstallInternalLegacyQuota(scheme)
	InstallInternalLegacyRoute(scheme)
	InstallInternalLegacySecurity(scheme)
	InstallInternalLegacyTemplate(scheme)
	InstallInternalLegacyUser(scheme)
}

func InstallLegacyExternalAll(scheme *runtime.Scheme) {
	InstallExternalLegacyApps(scheme)
	InstallExternalLegacyAuthorization(scheme)
	InstallExternalLegacyBuild(scheme)
	InstallExternalLegacyImage(scheme)
	InstallExternalLegacyNetwork(scheme)
	InstallExternalLegacyOAuth(scheme)
	InstallExternalLegacyProject(scheme)
	InstallExternalLegacyQuota(scheme)
	InstallExternalLegacyRoute(scheme)
	InstallExternalLegacySecurity(scheme)
	InstallExternalLegacyTemplate(scheme)
	InstallExternalLegacyUser(scheme)
}
