// Package fleetapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package fleetapi

import (
	"time"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Defines values for AgentPolicyMonitoringEnabled.
const (
	AgentPolicyMonitoringEnabledLogs    AgentPolicyMonitoringEnabled = "logs"
	AgentPolicyMonitoringEnabledMetrics AgentPolicyMonitoringEnabled = "metrics"
)

// Defines values for AgentPolicyCreateRequestMonitoringEnabled.
const (
	AgentPolicyCreateRequestMonitoringEnabledLogs    AgentPolicyCreateRequestMonitoringEnabled = "logs"
	AgentPolicyCreateRequestMonitoringEnabledMetrics AgentPolicyCreateRequestMonitoringEnabled = "metrics"
)

// Defines values for AgentPolicyUpdateRequestMonitoringEnabled.
const (
	Logs    AgentPolicyUpdateRequestMonitoringEnabled = "logs"
	Metrics AgentPolicyUpdateRequestMonitoringEnabled = "metrics"
)

// Defines values for OutputType.
const (
	OutputTypeElasticsearch OutputType = "elasticsearch"
	OutputTypeLogstash      OutputType = "logstash"
)

// Defines values for PostOutputsJSONBodyType.
const (
	PostOutputsJSONBodyTypeElasticsearch PostOutputsJSONBodyType = "elasticsearch"
)

// Defines values for UpdateOutputJSONBodyType.
const (
	Elasticsearch UpdateOutputJSONBodyType = "elasticsearch"
)

// AgentPolicy defines model for agent_policy.
type AgentPolicy struct {
	AgentFeatures *[]struct {
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
	} `json:"agent_features,omitempty"`
	Agents             *float32                        `json:"agents,omitempty"`
	DataOutputId       *string                         `json:"data_output_id"`
	Description        *string                         `json:"description,omitempty"`
	DownloadSourceId   *string                         `json:"download_source_id"`
	FleetServerHostId  *string                         `json:"fleet_server_host_id"`
	Id                 string                          `json:"id"`
	InactivityTimeout  *float32                        `json:"inactivity_timeout,omitempty"`
	MonitoringEnabled  *[]AgentPolicyMonitoringEnabled `json:"monitoring_enabled,omitempty"`
	MonitoringOutputId *string                         `json:"monitoring_output_id"`
	Name               string                          `json:"name"`
	Namespace          string                          `json:"namespace"`

	// PackagePolicies This field is present only when retrieving a single agent policy, or when retrieving a list of agent policy with the ?full=true parameter
	PackagePolicies *[]PackagePolicy `json:"package_policies,omitempty"`
	Revision        *float32         `json:"revision,omitempty"`
	UnenrollTimeout *float32         `json:"unenroll_timeout,omitempty"`
	UpdatedBy       *string          `json:"updated_by,omitempty"`
	UpdatedOn       *time.Time       `json:"updated_on,omitempty"`
}

// AgentPolicyMonitoringEnabled defines model for AgentPolicy.MonitoringEnabled.
type AgentPolicyMonitoringEnabled string

// AgentPolicyCreateRequest defines model for agent_policy_create_request.
type AgentPolicyCreateRequest struct {
	AgentFeatures *[]struct {
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
	} `json:"agent_features,omitempty"`
	DataOutputId       *string                                      `json:"data_output_id"`
	Description        *string                                      `json:"description,omitempty"`
	DownloadSourceId   *string                                      `json:"download_source_id"`
	FleetServerHostId  *string                                      `json:"fleet_server_host_id"`
	Id                 *string                                      `json:"id,omitempty"`
	InactivityTimeout  *float32                                     `json:"inactivity_timeout,omitempty"`
	MonitoringEnabled  *[]AgentPolicyCreateRequestMonitoringEnabled `json:"monitoring_enabled,omitempty"`
	MonitoringOutputId *string                                      `json:"monitoring_output_id"`
	Name               string                                       `json:"name"`
	Namespace          string                                       `json:"namespace"`
	UnenrollTimeout    *float32                                     `json:"unenroll_timeout,omitempty"`
}

// AgentPolicyCreateRequestMonitoringEnabled defines model for AgentPolicyCreateRequest.MonitoringEnabled.
type AgentPolicyCreateRequestMonitoringEnabled string

// AgentPolicyUpdateRequest defines model for agent_policy_update_request.
type AgentPolicyUpdateRequest struct {
	AgentFeatures *[]struct {
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
	} `json:"agent_features,omitempty"`
	DataOutputId       *string                                      `json:"data_output_id"`
	Description        *string                                      `json:"description,omitempty"`
	DownloadSourceId   *string                                      `json:"download_source_id"`
	FleetServerHostId  *string                                      `json:"fleet_server_host_id"`
	InactivityTimeout  *float32                                     `json:"inactivity_timeout,omitempty"`
	MonitoringEnabled  *[]AgentPolicyUpdateRequestMonitoringEnabled `json:"monitoring_enabled,omitempty"`
	MonitoringOutputId *string                                      `json:"monitoring_output_id"`
	Name               string                                       `json:"name"`
	Namespace          string                                       `json:"namespace"`
	UnenrollTimeout    *float32                                     `json:"unenroll_timeout,omitempty"`
}

// AgentPolicyUpdateRequestMonitoringEnabled defines model for AgentPolicyUpdateRequest.MonitoringEnabled.
type AgentPolicyUpdateRequestMonitoringEnabled string

// EnrollmentApiKey defines model for enrollment_api_key.
type EnrollmentApiKey struct {
	Active    bool    `json:"active"`
	ApiKey    string  `json:"api_key"`
	ApiKeyId  string  `json:"api_key_id"`
	CreatedAt string  `json:"created_at"`
	Id        string  `json:"id"`
	Name      *string `json:"name,omitempty"`
	PolicyId  *string `json:"policy_id,omitempty"`
}

// FleetServerHost defines model for fleet_server_host.
type FleetServerHost struct {
	HostUrls        []string `json:"host_urls"`
	Id              string   `json:"id"`
	IsDefault       bool     `json:"is_default"`
	IsPreconfigured bool     `json:"is_preconfigured"`
	Name            *string  `json:"name,omitempty"`
}

// NewPackagePolicy defines model for new_package_policy.
type NewPackagePolicy struct {
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Inputs      []struct {
		Config     *map[string]interface{} `json:"config,omitempty"`
		Enabled    bool                    `json:"enabled"`
		Processors *[]string               `json:"processors,omitempty"`
		Streams    *[]interface{}          `json:"streams,omitempty"`
		Type       string                  `json:"type"`
		Vars       *map[string]interface{} `json:"vars,omitempty"`
	} `json:"inputs"`
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
	// Deprecated:
	OutputId *string `json:"output_id,omitempty"`
	Package  *struct {
		Name    string  `json:"name"`
		Title   *string `json:"title,omitempty"`
		Version string  `json:"version"`
	} `json:"package,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
}

// Output defines model for output.
type Output struct {
	CaSha256             *string                 `json:"ca_sha256,omitempty"`
	CaTrustedFingerprint *string                 `json:"ca_trusted_fingerprint,omitempty"`
	Config               *map[string]interface{} `json:"config,omitempty"`
	ConfigYaml           *string                 `json:"config_yaml,omitempty"`
	Hosts                *[]string               `json:"hosts,omitempty"`
	Id                   string                  `json:"id"`
	IsDefault            bool                    `json:"is_default"`
	IsDefaultMonitoring  *bool                   `json:"is_default_monitoring,omitempty"`
	Name                 string                  `json:"name"`
	ProxyId              *string                 `json:"proxy_id,omitempty"`
	Shipper              *struct {
		CompressionLevel            *float32 `json:"compression_level,omitempty"`
		DiskQueueCompressionEnabled *bool    `json:"disk_queue_compression_enabled,omitempty"`
		DiskQueueEnabled            *bool    `json:"disk_queue_enabled,omitempty"`
		DiskQueueEncryptionEnabled  *bool    `json:"disk_queue_encryption_enabled,omitempty"`
		DiskQueueMaxSize            *float32 `json:"disk_queue_max_size,omitempty"`
		DiskQueuePath               *string  `json:"disk_queue_path,omitempty"`
		Loadbalance                 *bool    `json:"loadbalance,omitempty"`
	} `json:"shipper,omitempty"`
	Ssl *struct {
		Certificate            *string   `json:"certificate,omitempty"`
		CertificateAuthorities *[]string `json:"certificate_authorities,omitempty"`
		Key                    *string   `json:"key,omitempty"`
	} `json:"ssl,omitempty"`
	Type OutputType `json:"type"`
}

// OutputType defines model for Output.Type.
type OutputType string

// PackagePolicy defines model for package_policy.
type PackagePolicy struct {
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Id          string  `json:"id"`
	Inputs      []struct {
		Config     *map[string]interface{} `json:"config,omitempty"`
		Enabled    bool                    `json:"enabled"`
		Processors *[]string               `json:"processors,omitempty"`
		Streams    *[]interface{}          `json:"streams,omitempty"`
		Type       string                  `json:"type"`
		Vars       *map[string]interface{} `json:"vars,omitempty"`
	} `json:"inputs"`
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
	// Deprecated:
	OutputId *string `json:"output_id,omitempty"`
	Package  *struct {
		Name    string  `json:"name"`
		Title   *string `json:"title,omitempty"`
		Version string  `json:"version"`
	} `json:"package,omitempty"`
	PolicyId *string `json:"policy_id,omitempty"`
	Revision float32 `json:"revision"`
}

// DeleteAgentPolicyJSONBody defines parameters for DeleteAgentPolicy.
type DeleteAgentPolicyJSONBody struct {
	AgentPolicyId string `json:"agentPolicyId"`
}

// PostFleetServerHostsJSONBody defines parameters for PostFleetServerHosts.
type PostFleetServerHostsJSONBody struct {
	HostUrls  []string `json:"host_urls"`
	Id        *string  `json:"id,omitempty"`
	IsDefault *bool    `json:"is_default,omitempty"`
	Name      string   `json:"name"`
}

// UpdateFleetServerHostsJSONBody defines parameters for UpdateFleetServerHosts.
type UpdateFleetServerHostsJSONBody struct {
	HostUrls  *[]string `json:"host_urls,omitempty"`
	IsDefault *bool     `json:"is_default,omitempty"`
	Name      *string   `json:"name,omitempty"`
}

// PostOutputsJSONBody defines parameters for PostOutputs.
type PostOutputsJSONBody struct {
	CaSha256            *string                 `json:"ca_sha256,omitempty"`
	ConfigYaml          *string                 `json:"config_yaml,omitempty"`
	Hosts               *[]string               `json:"hosts,omitempty"`
	Id                  *string                 `json:"id,omitempty"`
	IsDefault           *bool                   `json:"is_default,omitempty"`
	IsDefaultMonitoring *bool                   `json:"is_default_monitoring,omitempty"`
	Name                string                  `json:"name"`
	Type                PostOutputsJSONBodyType `json:"type"`
}

// PostOutputsJSONBodyType defines parameters for PostOutputs.
type PostOutputsJSONBodyType string

// UpdateOutputJSONBody defines parameters for UpdateOutput.
type UpdateOutputJSONBody struct {
	CaSha256             *string                  `json:"ca_sha256,omitempty"`
	CaTrustedFingerprint *string                  `json:"ca_trusted_fingerprint,omitempty"`
	ConfigYaml           *string                  `json:"config_yaml,omitempty"`
	Hosts                *[]string                `json:"hosts,omitempty"`
	IsDefault            *bool                    `json:"is_default,omitempty"`
	IsDefaultMonitoring  *bool                    `json:"is_default_monitoring,omitempty"`
	Name                 string                   `json:"name"`
	Type                 UpdateOutputJSONBodyType `json:"type"`
}

// UpdateOutputJSONBodyType defines parameters for UpdateOutput.
type UpdateOutputJSONBodyType string

// CreateAgentPolicyJSONRequestBody defines body for CreateAgentPolicy for application/json ContentType.
type CreateAgentPolicyJSONRequestBody = AgentPolicyCreateRequest

// DeleteAgentPolicyJSONRequestBody defines body for DeleteAgentPolicy for application/json ContentType.
type DeleteAgentPolicyJSONRequestBody DeleteAgentPolicyJSONBody

// UpdateAgentPolicyJSONRequestBody defines body for UpdateAgentPolicy for application/json ContentType.
type UpdateAgentPolicyJSONRequestBody = AgentPolicyUpdateRequest

// PostFleetServerHostsJSONRequestBody defines body for PostFleetServerHosts for application/json ContentType.
type PostFleetServerHostsJSONRequestBody PostFleetServerHostsJSONBody

// UpdateFleetServerHostsJSONRequestBody defines body for UpdateFleetServerHosts for application/json ContentType.
type UpdateFleetServerHostsJSONRequestBody UpdateFleetServerHostsJSONBody

// PostOutputsJSONRequestBody defines body for PostOutputs for application/json ContentType.
type PostOutputsJSONRequestBody PostOutputsJSONBody

// UpdateOutputJSONRequestBody defines body for UpdateOutput for application/json ContentType.
type UpdateOutputJSONRequestBody UpdateOutputJSONBody
