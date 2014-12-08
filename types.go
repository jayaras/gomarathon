package gomarathon

import "encoding/json"

// RequestOptions passed for query api
type RequestOptions struct {
	Method string
	Path   string
	Datas  interface{}
	Params *Parameters
}

// Parameters to build url query
type Parameters struct {
	Cmd         string
	Host        string
	Scale       bool
	CallbackURL string
	Force       bool `json:"force"`
}

// is there a better way to handle the return code storage?
// I don't think an interface can specify a field
type Response interface {
	SetCode(int)
	GetCode() int
	UnmarshalJSON([]byte) error
}

// Response representation of a full marathon response
type DefaultResponse struct {
	code     int
	Apps     []*Application `json:"apps,omitempty"`
	App      *Application   `json:"app,omitempty"`
	Versions []string       `json:",omitempty"`
	Tasks    []*Task        `json:"tasks,omitempty"`
}

func (r *DefaultResponse) SetCode(code int) {
	r.code = code
}

func (r *DefaultResponse) GetCode() int {
	return r.code
}

func (r *DefaultResponse) UnmarshalJSON(data []byte) error {

	return json.Unmarshal(data, r)

}

type DeploymentResponse struct {
	Deployments []Deployment
	code        int
}

func (r *DeploymentResponse) SetCode(code int) {
	r.code = code
}

func (r *DeploymentResponse) GetCode() int {
	return r.code
}

func (r *DeploymentResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Deployments)
}

type Deployment struct {
	AffectedApps   []string           `json:"affectedApps,omitempty"`
	ID             string             `json:"id,omitempty"`
	Steps          []DeploymentAction `json:"steps,omitempty"`
	CurrentActions []DeploymentAction `json:"currentActions,omitempty"`
	Version        string             `json:"version,omitempty"`
	CurrentStep    int                `json:"currentStep,omitempty"`
	TotalSteps     int                `json:"totalSteps,omitempty"`
}

type DeploymentAction struct {
	Action string `json:"action,omitempty"`
	App    string `json:"app,omitempty"`
}

//
type Group struct {
	ID           string         `json:"id"`
	Version      string         `json:"version,omitempty"`
	Apps         []*Application `json:"apps,omitempty"`
	Dependencies []string       `json:"dependencies,omitempty"`
	Groups       []*Group       `json:"groups,omitempty"`
}

// Application marathon application see :
// https://mesosphere.github.io/marathon/docs/rest-api.html#apps
type Application struct {
	ID              string            `json:"id"`
	Cmd             string            `json:"cmd,omitempty"`
	Constraints     [][]string        `json:"constraints,omitempty"`
	Container       *Container        `json:"container,omitempty"`
	CPUs            float32           `json:"cpus,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
	Executor        string            `json:"executor,omitempty"`
	HealthChecks    []*HealthCheck    `json:"healthChecks,omitempty"`
	Instances       int               `json:"instances,omitemptys"`
	Mem             float32           `json:"mem,omitempty"`
	Tasks           []*Task           `json:"tasks,omitempty"`
	Ports           []int             `json:"ports,omitempty"`
	RequirePorts    bool              `json:"requirePorts,omitempty"`
	BackoffFactor   float32           `json:"backoffFactor,omitempty"`
	TasksRunning    int               `json:"tasksRunning,omitempty"`
	TasksStaged     int               `json:"tasksStaged,omitempty"`
	UpgradeStrategy *UpgradeStrategy  `json:"upgradeStrategy,omitempty"`
	Uris            []string          `json:"uris,omitempty"`
	Version         string            `json:"version,omitempty"`
	Dependencies    []string          `json:"dependencies,omitempty"`
}

// Container is docker parameters
type Container struct {
	Type    string    `json:"type,omitempty"`
	Docker  *Docker   `json:"docker,omitempty"`
	Volumes []*Volume `json:"volumes,omitempty"`
}

// Docker options
type Docker struct {
	Image        string         `json:"image,omitempty"`
	Network      string         `json:"network,omitempty"`
	PortMappings []*PortMapping `json:"portMappings,omitempty"`
}

//Docker portmapping
type PortMapping struct {
	ContainerPort int    `json:"containerPort"`
	HostPort      int    `json:"hostPort"`
	Protocol      string `json:"protocol,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
}

// Container volumes
type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

// Upgrade strategy
type UpgradeStrategy struct {
	MinimumHealthCapacity float32 `json:"minimumHealthCapacity,omitempty"`
}

// HealthChecks are described here:
// https://mesosphere.github.io/marathon/docs/health-checks.html
type HealthCheck struct {
	Protocol               string `json:"protocol,omitempty"`
	Path                   string `json:"path,omitempty"`
	GracePeriodSeconds     int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        int    `json:"intervalSeconds,omitempty"`
	PortIndex              int    `json:"portIndex,omitempty"`
	TimeoutSeconds         int    `json:"timeoutSeconds,omitempty"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures,omitempty"`
}

// Task is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#tasks
type Task struct {
	AppID     string `json:"appId"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Ports     []int  `json:"ports"`
	StagedAt  string `json:"stagedAt"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

// EventSubscription are described here :
// https://mesosphere.github.io/marathon/docs/rest-api.html#event-subscriptions
type EventSubscription struct {
	CallbackURL  string   `json:"CallbackUrl"`
	ClientIP     string   `json:"ClientIp"`
	EventType    string   `json:"eventType"`
	CallbackURLs []string `json:"CallbackUrls"`
}
