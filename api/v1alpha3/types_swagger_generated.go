// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha3

func (VirtualMachineInstance) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineInstance is *the* VirtualMachineInstance Definition. It represents a virtual machine in the runtime environment of kubernetes.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"spec":   "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
		"status": "Status is the high level overview of how the VirtualMachineInstance is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineInstanceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceList is a list of VirtualMachines\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                              "VirtualMachineInstanceSpec is a description of a VirtualMachineInstance.\n\n+k8s:openapi-gen=true",
		"priorityClassName":             "If specified, indicates the pod's priority.\nIf not specified, the pod priority will be default or zero if there is no\ndefault.\n+optional",
		"domain":                        "Specification of the desired behavior of the VirtualMachineInstance on the host.",
		"nodeSelector":                  "NodeSelector is a selector which must be true for the vmi to fit on a node.\nSelector which must match a node's labels for the vmi to be scheduled on that node.\nMore info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/\n+optional",
		"affinity":                      "If affinity is specifies, obey all the affinity rules",
		"schedulerName":                 "If specified, the VMI will be dispatched by specified scheduler.\nIf not specified, the VMI will be dispatched by default scheduler.\n+optional",
		"tolerations":                   "If toleration is specified, obey all the toleration rules.",
		"evictionStrategy":              "EvictionStrategy can be set to \"LiveMigrate\" if the VirtualMachineInstance should be\nmigrated instead of shut-off in case of a node drain.\n\n+optional",
		"terminationGracePeriodSeconds": "Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.",
		"volumes":                       "List of volumes that can be mounted by disks belonging to the vmi.",
		"livenessProbe":                 "Periodic probe of VirtualMachineInstance liveness.\nVirtualmachineInstances will be stopped if the probe fails.\nCannot be updated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"readinessProbe":                "Periodic probe of VirtualMachineInstance service readiness.\nVirtualmachineInstances will be removed from service endpoints if the probe fails.\nCannot be updated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"hostname":                      "Specifies the hostname of the vmi\nIf not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly.\n+optional",
		"subdomain":                     "If specified, the fully qualified vmi hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\".\nIf not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi,\nno matter if the vmi itself can pick up a hostname.\n+optional",
		"networks":                      "List of networks that can be attached to a vm's virtual interface.",
		"dnsPolicy":                     "Set DNS policy for the pod.\nDefaults to \"ClusterFirst\".\nValid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.\nDNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.\nTo have DNS options set along with hostNetwork, you have to specify DNS policy\nexplicitly to 'ClusterFirstWithHostNet'.\n+optional",
		"dnsConfig":                     "Specifies the DNS parameters of a pod.\nParameters specified here will be merged to the generated DNS\nconfiguration based on DNSPolicy.\n+optional",
	}
}

func (VirtualMachineInstanceStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachineInstanceStatus represents information about the status of a VirtualMachineInstance. Status may trail the actual\nstate of a system.\n\n+k8s:openapi-gen=true",
		"nodeName":        "NodeName is the name where the VirtualMachineInstance is currently running.",
		"reason":          "A brief CamelCase message indicating details about why the VMI is in this state. e.g. 'NodeUnresponsive'\n+optional",
		"conditions":      "Conditions are specific points in VirtualMachineInstance's pod runtime.",
		"phase":           "Phase is the status of the VirtualMachineInstance in kubernetes world. It is not the VirtualMachineInstance status, but partially correlates to it.",
		"interfaces":      "Interfaces represent the details of available network interfaces.",
		"guestOSInfo":     "Guest OS Information",
		"migrationState":  "Represents the status of a live migration",
		"migrationMethod": "Represents the method using which the vmi can be migrated: live migration or block migration",
		"qosClass":        "The Quality of Service (QOS) classification assigned to the virtual machine instance based on resource requirements\nSee PodQOSClass type for available QOS classes\nMore info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md\n+optional",
		"activePods":      "ActivePods is a mapping of pod UID to node name.\nIt is possible for multiple pods to be running for a single VMI during migration.",
	}
}

func (VirtualMachineInstanceCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "+k8s:openapi-gen=true",
		"lastProbeTime":      "+nullable",
		"lastTransitionTime": "+nullable",
	}
}

func (VirtualMachineInstanceMigrationCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "+k8s:openapi-gen=true",
		"lastProbeTime":      "+nullable",
		"lastTransitionTime": "+nullable",
	}
}

func (VirtualMachineInstanceNetworkInterface) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "+k8s:openapi-gen=true",
		"ipAddress":     "IP address of a Virtual Machine interface. It is always the first item of\nIPs",
		"mac":           "Hardware address of a Virtual Machine interface",
		"name":          "Name of the interface, corresponds to name of the network assigned to the interface",
		"ipAddresses":   "List of all IP addresses of a Virtual Machine interface",
		"interfaceName": "The interface name inside the Virtual Machine",
	}
}

func (VirtualMachineInstanceGuestOSInfo) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "+k8s:openapi-gen=true",
		"name":          "Name of the Guest OS",
		"kernelRelease": "Guest OS Kernel Release",
		"version":       "Guest OS Version",
		"prettyName":    "Guest OS Pretty Name",
		"versionId":     "Version ID of the Guest OS",
		"kernelVersion": "Kernel version of the Guest OS",
		"machine":       "Machine type of the Guest OS",
		"id":            "Guest OS Id",
	}
}

func (VirtualMachineInstanceMigrationState) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                               "+k8s:openapi-gen=true",
		"startTimestamp":                 "The time the migration action began",
		"endTimestamp":                   "The time the migration action ended",
		"targetNodeDomainDetected":       "The Target Node has seen the Domain Start Event",
		"targetNodeAddress":              "The address of the target node to use for the migration",
		"targetDirectMigrationNodePorts": "The list of ports opened for live migration on the destination node",
		"targetNode":                     "The target node that the VMI is moving to",
		"targetPod":                      "The target pod that the VMI is moving to",
		"sourceNode":                     "The source node that the VMI originated on",
		"completed":                      "Indicates the migration completed",
		"failed":                         "Indicates that the migration failed",
		"abortRequested":                 "Indicates that the migration has been requested to abort",
		"abortStatus":                    "Indicates the final status of the live migration abortion",
		"migrationUid":                   "The VirtualMachineInstanceMigration object associated with this migration",
	}
}

func (VMISelector) SwaggerDoc() map[string]string {
	return map[string]string{
		"name": "Name of the VirtualMachineInstance to migrate",
	}
}

func (VirtualMachineInstanceReplicaSet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineInstance is *the* VirtualMachineInstance Definition. It represents a virtual machine in the runtime environment of kubernetes.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"spec":   "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
		"status": "Status is the high level overview of how the VirtualMachineInstance is doing. It contains information available to controllers and users.\n+nullable",
	}
}

func (VirtualMachineInstanceReplicaSetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VMIList is a list of VMIs\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceReplicaSetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"replicas": "Number of desired pods. This is a pointer to distinguish between explicit\nzero and not specified. Defaults to 1.\n+optional",
		"selector": "Label selector for pods. Existing ReplicaSets whose pods are\nselected by this will be the ones affected by this deployment.",
		"template": "Template describes the pods that will be created.",
		"paused":   "Indicates that the replica set is paused.\n+optional",
	}
}

func (VirtualMachineInstanceReplicaSetStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "+k8s:openapi-gen=true",
		"replicas":      "Total number of non-terminated pods targeted by this deployment (their labels match the selector).\n+optional",
		"readyReplicas": "The number of ready replicas for this replica set.\n+optional",
		"labelSelector": "Canonical form of the label selector for HPA which consumes it through the scale subresource.",
	}
}

func (VirtualMachineInstanceReplicaSetCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "+k8s:openapi-gen=true",
		"lastProbeTime":      "+nullable",
		"lastTransitionTime": "+nullable",
	}
}

func (VirtualMachineInstanceTemplateSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"metadata": "+nullable",
		"spec":     "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
	}
}

func (VirtualMachineInstanceMigration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigration represents the object tracking a VMI's migration\nto another host in the cluster\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceMigrationList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigrationList is a list of VirtualMachineMigrations\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceMigrationSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "+k8s:openapi-gen=true",
		"vmiName": "The name of the VMI to perform the migration on. VMI must exist in the migration objects namespace",
	}
}

func (VirtualMachineInstanceMigrationStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigration reprents information pertaining to a VMI's migration.\n\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstancePreset) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"spec": "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
	}
}

func (VirtualMachineInstancePresetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstancePresetList is a list of VirtualMachinePresets\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstancePresetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"selector": "Selector is a label query over a set of VMIs.\nRequired.",
		"domain":   "Domain is the same object type as contained in VirtualMachineInstanceSpec",
	}
}

func (VirtualMachine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachine handles the VirtualMachines that are not running\nor are in a stopped state\nThe VirtualMachine contains the template to create the\nVirtualMachineInstance. It also mirrors the running state of the created\nVirtualMachineInstance in its status.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"spec":   "Spec contains the specification of VirtualMachineInstance created",
		"status": "Status holds the current state of the controller and brief information\nabout its associated VirtualMachineInstance",
	}
}

func (VirtualMachineList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineList is a list of virtualmachines\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "VirtualMachineSpec describes how the proper VirtualMachine\nshould look like\n\n+k8s:openapi-gen=true",
		"running":             "Running controls whether the associatied VirtualMachineInstance is created or not\nMutually exclusive with RunStrategy",
		"runStrategy":         "Running state indicates the requested running state of the VirtualMachineInstance\nmutually exclusive with Running",
		"template":            "Template is the direct specification of VirtualMachineInstance",
		"dataVolumeTemplates": "dataVolumeTemplates is a list of dataVolumes that the VirtualMachineInstance template can reference.\nDataVolumes in this list are dynamically created for the VirtualMachine and are tied to the VirtualMachine's life-cycle.",
	}
}

func (VirtualMachineStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "VirtualMachineStatus represents the status returned by the\ncontroller to describe how the VirtualMachine is doing\n\n+k8s:openapi-gen=true",
		"snapshotInProgress":  "SnapshotInProgress is the name of the VirtualMachineSnapshot currently executing",
		"created":             "Created indicates if the virtual machine is created in the cluster",
		"ready":               "Ready indicates if the virtual machine is running and ready",
		"conditions":          "Hold the state information of the VirtualMachine and its VirtualMachineInstance",
		"stateChangeRequests": "StateChangeRequests indicates a list of actions that should be taken on a VMI\ne.g. stop a specific VMI then start a new one.",
	}
}

func (VirtualMachineStateChangeRequest) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "+k8s:openapi-gen=true",
		"action": "Indicates the type of action that is requested. e.g. Start or Stop",
		"data":   "Provides additional data in order to perform the Action",
		"uid":    "Indicates the UUID of an existing Virtual Machine Instance that this change request applies to -- if applicable",
	}
}

func (VirtualMachineCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "VirtualMachineCondition represents the state of VirtualMachine\n\n+k8s:openapi-gen=true",
		"lastProbeTime":      "+nullable",
		"lastTransitionTime": "+nullable",
	}
}

func (Handler) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "Handler defines a specific action that should be taken",
		"httpGet":   "HTTPGet specifies the http request to perform.\n+optional",
		"tcpSocket": "TCPSocket specifies an action involving a TCP port.\nTCP hooks not yet supported\n+optional",
	}
}

func (Probe) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "Probe describes a health check to be performed against a VirtualMachineInstance to determine whether it is\nalive or ready to receive traffic.\n+k8s:openapi-gen=true",
		"initialDelaySeconds": "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"timeoutSeconds":      "Number of seconds after which the probe times out.\nDefaults to 1 second. Minimum value is 1.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"periodSeconds":       "How often (in seconds) to perform the probe.\nDefault to 10 seconds. Minimum value is 1.\n+optional",
		"successThreshold":    "Minimum consecutive successes for the probe to be considered successful after having failed.\nDefaults to 1. Must be 1 for liveness. Minimum value is 1.\n+optional",
		"failureThreshold":    "Minimum consecutive failures for the probe to be considered failed after having succeeded.\nDefaults to 3. Minimum value is 1.\n+optional",
	}
}

func (KubeVirt) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirt represents the object deploying all KubeVirt resources\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (KubeVirtList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtList is a list of KubeVirts\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (KubeVirtSelfSignConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (KubeVirtCertificateRotateStrategy) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (KubeVirtSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"imageTag":          "The image tag to use for the continer images installed.\nDefaults to the same tag as the operator's container image.",
		"imageRegistry":     "The image registry to pull the container images from\nDefaults to the same registry the operator's container image is pulled from.",
		"imagePullPolicy":   "The ImagePullPolicy to use.",
		"monitorNamespace":  "The namespace Prometheus is deployed in\nDefaults to openshift-monitor",
		"monitorAccount":    "The name of the Prometheus service account that needs read-access to KubeVirt endpoints\nDefaults to prometheus-k8s",
		"uninstallStrategy": "Specifies if kubevirt can be deleted if workloads are still present.\nThis is mainly a precaution to avoid accidental data loss",
		"configuration":     "holds kubevirt configurations.\nsame as the virt-configMap",
	}
}

func (KubeVirtStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtStatus represents information pertaining to a KubeVirt deployment.\n\n+k8s:openapi-gen=true",
	}
}

func (KubeVirtCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "KubeVirtCondition represents a condition of a KubeVirt deployment\n\n+k8s:openapi-gen=true",
		"lastProbeTime":      "+optional\n+nullable",
		"lastTransitionTime": "+optional\n+nullable",
	}
}

func (RestartOptions) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "RestartOptions may be provided when deleting an API object.\n\n+k8s:openapi-gen=true",
		"gracePeriodSeconds": "The duration in seconds before the object should be force-restared. Value must be non-negative integer.\nThe value zero indicates, restart immediately. If this value is nil, the default grace period for deletion of the corresponding VMI for the\nspecified type will be used to determine on how much time to give the VMI to restart.\nDefaults to a per object value if not specified. zero means restart immediately.\nAllowed Values: nil and 0\n+optional",
	}
}

func (VirtualMachineInstanceGuestAgentInfo) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "VirtualMachineInstanceGuestAgentInfo represents information from the installed guest agent\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"guestAgentVersion": "GAVersion is a version of currently installed guest agent",
		"hostname":          "Hostname represents FQDN of a guest",
		"os":                "OS contains the guest operating system information",
		"timezone":          "Timezone is guest os current timezone",
		"userList":          "UserList is a list of active guest OS users",
		"fsInfo":            "FSInfo is a guest os filesystem information containing the disk mapping and disk mounts with usage",
	}
}

func (VirtualMachineInstanceGuestOSUserList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceGuestOSUserList comprises the list of all active users on guest machine\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceGuestOSUser) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineGuestOSUser is the single user of the guest os\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceFileSystemInfo) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceFileSystemInfo represents information regarding single guest os filesystem\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceFileSystemList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceFileSystemList comprises the list of all filesystems on guest machine\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineInstanceFileSystem) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceFileSystem represents guest os disk\n+k8s:openapi-gen=true",
	}
}

func (RenameOptions) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Options for a rename operation",
	}
}

func (KubeVirtConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtConfiguration holds all kubevirt configurations\n+k8s:openapi-gen=true",
	}
}

func (SMBiosConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (MigrationConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "MigrationConfiguration holds migration options\n+k8s:openapi-gen=true",
	}
}

func (DeveloperConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DeveloperConfiguration holds developer options\n+k8s:openapi-gen=true",
	}
}

func (NetworkConfiguration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "NetworkConfiguration holds network options\n+k8s:openapi-gen=true",
	}
}
