/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LagoonTaskSpec defines the desired state of LagoonTask
type LagoonTaskSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Key          string                  `json:"key,omitempty"`
	Task         LagoonTaskInfo          `json:"task,omitempty"`
	Project      LagoonTaskProject       `json:"project,omitempty"`
	Environment  LagoonTaskEnvironment   `json:"environment,omitempty"`
	Misc         *LagoonMiscInfo         `json:"misc,omitempty"`
	AdvancedTask *LagoonAdvancedTaskInfo `json:"advancedTask,omitempty"`
}

// LagoonTaskInfo defines what a task can use to communicate with Lagoon via SSH/API.
type LagoonTaskInfo struct {
	ID      string `json:"id"` // should be int, but the api sends it as a string :\
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
	Command string `json:"command,omitempty"`
	SSHHost string `json:"sshHost,omitempty"`
	SSHPort string `json:"sshPort,omitempty"`
	APIHost string `json:"apiHost,omitempty"`
}

// LagoonAdvancedTaskInfo defines what an advanced task can use for the creation of the pod.
type LagoonAdvancedTaskInfo struct {
	RunnerImage string `json:"runnerImage,omitempty"`
	JSONPayload string `json:"JSONPayload,omitempty"`
}

// LagoonMiscInfo defines the resource or backup information for a misc task.
type LagoonMiscInfo struct {
	ID           string                `json:"id"` // should be int, but the api sends it as a string :\
	Name         string                `json:"name,omitempty"`
	Backup       *LagoonMiscBackupInfo `json:"backup,omitempty"`
	MiscResource []byte                `json:"miscResource,omitempty"`
}

// LagoonMiscBackupInfo defines the information for a backup.
type LagoonMiscBackupInfo struct {
	ID       string `json:"id"` // should be int, but the api sends it as a string :\
	Source   string `json:"source"`
	BackupID string `json:"backupId"`
}

// LagoonTaskProject defines the lagoon project information.
type LagoonTaskProject struct {
	ID   string `json:"id"` // should be int, but the api sends it as a string :\
	Name string `json:"name"`
}

// LagoonTaskEnvironment defines the lagoon environment information.
type LagoonTaskEnvironment struct {
	ID                   string `json:"id"` // should be int, but the api sends it as a string :\
	Name                 string `json:"name"`
	Project              string `json:"project"` // should be int, but the api sends it as a string :\
	EnvironmentType      string `json:"environmentType"`
	OpenshiftProjectName string `json:"openshiftProjectName"`
}

// LagoonTaskStatus defines the observed state of LagoonTask
type LagoonTaskStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []LagoonConditions `json:"conditions,omitempty"`
	Log        []byte             `json:"log,omitempty"`
}

// +kubebuilder:object:root=true

// LagoonTask is the Schema for the lagoontasks API
type LagoonTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec           LagoonTaskSpec        `json:"spec,omitempty"`
	Status         LagoonTaskStatus      `json:"status,omitempty"`
	StatusMessages *LagoonStatusMessages `json:"statusMessages,omitempty"`
}

// +kubebuilder:object:root=true

// LagoonTaskList contains a list of LagoonTask
type LagoonTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LagoonTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LagoonTask{}, &LagoonTaskList{})
}
