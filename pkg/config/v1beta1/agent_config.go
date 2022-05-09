package v1beta1

import (
	"github.com/rancher/opni/pkg/config/meta"
)

type AgentConfig struct {
	meta.TypeMeta `json:",inline"`

	Spec AgentConfigSpec `json:"spec,omitempty"`
}

type TrustStrategyKind string

const (
	TrustStrategyPKP      TrustStrategyKind = "pkp"
	TrustStrategyCACerts  TrustStrategyKind = "cacerts"
	TrustStrategyInsecure TrustStrategyKind = "insecure"
)

type AgentConfigSpec struct {
	// The address which the agent will listen on for incoming connections.
	// This should be in the format "host:port" or ":port", and must not
	// include a scheme.
	ListenAddress string `json:"listenAddress,omitempty"`
	// The address of the gateway's public HTTP API. This should be of the format
	// "https://host:port". The scheme must be "https".
	GatewayAddress string `json:"gatewayAddress,omitempty"`
	// The name of the identity provider to use. Defaults to "kubernetes".
	IdentityProvider string `json:"identityProvider,omitempty"`
	// The type of trust strategy to use for verifying the authenticity of the
	// gateway server. Defaults to "pkp".
	TrustStrategy TrustStrategyKind `json:"trustStrategy,omitempty"`
	// Configuration for agent keyring storage.
	Storage   StorageSpec    `json:"storage,omitempty"`
	Rules     *RulesSpec     `json:"rules,omitempty"`
	Bootstrap *BootstrapSpec `json:"bootstrap,omitempty"`
}

type BootstrapSpec struct {
	// Bootstrap token
	Token string `json:"token,omitempty"`
	// List of public key pins. Used when the trust strategy is "pkp".
	Pins []string `json:"pins,omitempty"`
	// List of paths to CA Certs. Used when the trust strategy is "pkp".
	// If empty, the system certs will be used.
	CACerts []string `json:"caCerts,omitempty"`
}

func (s *AgentConfigSpec) SetDefaults() {
	if s == nil {
		return
	}
	if s.IdentityProvider == "" {
		s.IdentityProvider = "kubernetes"
	}
	if s.ListenAddress == "" {
		s.ListenAddress = ":8080"
	}
	if s.TrustStrategy == "" {
		s.TrustStrategy = "pkp"
	}
}

type RulesSpec struct {
	Discovery DiscoverySpec `json:"discovery,omitempty"`
}

type DiscoverySpec struct {
	PrometheusRules *PrometheusRulesSpec `json:"prometheusRules,omitempty"`
	// Search interval. Defaults to "15m"
	Interval string `json:"interval,omitempty"`
}

type PrometheusRulesSpec struct {
	// Namespaces to search for rules in. If empty, will search all accessible
	// namespaces.
	SearchNamespaces []string `json:"searchNamespaces,omitempty"`
	// Kubeconfig to use for rule discovery. If nil, will use the in-cluster
	// kubeconfig.
	Kubeconfig *string `json:"kubeconfig,omitempty"`
}
