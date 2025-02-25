package config

// Server configuration
type Server struct {
	Port int `mapstructure:"port"`
}

// VulnerabilityReports configuration
type VulnerabilityReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// ConfigAuditReports configuration
type ConfigAuditReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// CISKubeBenchReports configuration
type CISKubeBenchReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// ComplianceReports configuration
type ComplianceReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// RbacAssessmentReports configuration
type RbacAssessmentReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// ExposedSecretReports configuration
type ExposedSecretReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

type InfraAssessmentReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

type ClusterInfraAssessmentReports struct {
	Enabled     bool     `mapstructure:"enabled"`
	Timeout     int      `mapstructure:"timeout"`
	ApplyLabels []string `mapstructure:"applyLabels"`
}

// Config of the Tracee Adapter
type Config struct {
	Kubeconfig                    string                        `mapstructure:"kubeconfig"`
	Server                        Server                        `mapstructure:"server"`
	VulnerabilityReports          VulnerabilityReports          `mapstructure:"vulnerabilityReports"`
	ConfigAuditReports            ConfigAuditReports            `mapstructure:"configAuditReports"`
	CISKubeBenchReports           CISKubeBenchReports           `mapstructure:"cisKubeBenchReports"`
	ComplianceReports             ComplianceReports             `mapstructure:"complianceReports"`
	RbacAssessmentReports         RbacAssessmentReports         `mapstructure:"rbacAssessmentReports"`
	ExposedSecretReports          ExposedSecretReports          `mapstructure:"exposedSecretReports"`
	InfraAssessmentReports        InfraAssessmentReports        `mapstructure:"infraAssessmentReports"`
	ClusterInfraAssessmentReports ClusterInfraAssessmentReports `mapstructure:"clusterInfraAssessmentReports"`
}
