# API Reference

## Packages
- [weka.weka.io/v1alpha1](#wekawekaiov1alpha1)


## weka.weka.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the weka v1alpha1 API group

### Resource Types
- [DriveClaim](#driveclaim)
- [WekaClient](#wekaclient)
- [WekaCluster](#wekacluster)
- [WekaContainer](#wekacontainer)
- [WekaManualOperation](#wekamanualoperation)
- [WekaPolicy](#wekapolicy)





#### AdditionalMemory







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _integer_ |  |  |  |
| `drive` _integer_ |  |  |  |
| `s3` _integer_ |  |  |  |
| `nfs` _integer_ |  |  |  |
| `envoy` _integer_ |  |  |  |
| `smbw` _integer_ |  |  |  |
| `dataServices` _integer_ |  |  |  |


#### AdvancedCsiConfig







_Appears in:_
- [ClientCsiConfig](#clientcsiconfig)
- [CsiConfig](#csiconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enforceTrustedHttps` _boolean_ |  |  |  |
| `nodeLabels` _object (keys:string, values:string)_ |  |  |  |
| `nodeTolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `controllerLabels` _object (keys:string, values:string)_ |  |  |  |
| `controllerTolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `skipGarbageCollection` _boolean_ |  |  |  |


#### BlockDrivesPayload







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `serialIDs` _string array_ |  |  |  |
| `physicalUUIDs` _string array_ |  |  |  |
| `node` _string_ |  |  |  |


#### CapacityMetrics







_Appears in:_
- [ClusterMetrics](#clustermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `totalBytes` _[IntMetric](#intmetric)_ |  |  |  |
| `unprovisionedBytes` _[IntMetric](#intmetric)_ |  |  |  |
| `unavailableBytes` _[IntMetric](#intmetric)_ |  |  |  |
| `hotSpareBytes` _[IntMetric](#intmetric)_ |  |  |  |


#### CatalogConfig



CatalogConfig defines configuration for the data catalog service



_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `indexInterval` _string_ | IndexInterval specifies how often the catalog index is updated (e.g., "1d", "1m") | 1d |  |
| `retentionPeriod` _string_ | RetentionPeriod specifies how long catalog data is retained (e.g., "30d", "10m") | 30d |  |




#### ClientCsiConfig







_Appears in:_
- [WekaClientSpec](#wekaclientspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `csiGroup` _string_ |  |  |  |
| `disableControllerCreation` _boolean_ |  |  |  |
| `advanced` _[AdvancedCsiConfig](#advancedcsiconfig)_ |  |  |  |


#### ClientMetrics







_Appears in:_
- [WekaClientStatus](#wekaclientstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `containers` _[EntityStatefulNum](#entitystatefulnum)_ |  |  |  |


#### ClientPrinterColumns







_Appears in:_
- [WekaClientStatus](#wekaclientstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `containers` _[StringMetric](#stringmetric)_ |  |  |  |


#### ClusterMetrics







_Appears in:_
- [WekaClusterStatus](#wekaclusterstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `containers` _[ContainersMetrics](#containersmetrics)_ |  |  |  |
| `ioStats` _[IoStats](#iostats)_ |  |  |  |
| `drives` _[DriveMetrics](#drivemetrics)_ |  |  |  |
| `alertsCount` _[IntMetric](#intmetric)_ |  |  |  |
| `clusterStatus` _[StringMetric](#stringmetric)_ |  |  |  |
| `capacity` _[CapacityMetrics](#capacitymetrics)_ |  |  |  |
| `numFailures` _object (keys:string, values:[FloatMetric](#floatmetric))_ |  |  |  |
| `lastUpdate` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ |  |  |  |
| `filesystem` _[FilesystemMetrics](#filesystemmetrics)_ |  |  |  |


#### ClusterPorts







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)
- [WekaClusterStatus](#wekaclusterstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `basePort` _integer_ | We should not be updating Spec, as it's a user interface and we should not break ability to update spec file<br />Therefore, when BasePort is 0, and Range as 0, we have application level defaults that will be written in here |  |  |
| `portRange` _integer_ |  |  |  |
| `lbPort` _integer_ |  |  |  |
| `lbAdminPort` _integer_ |  |  |  |
| `s3Port` _integer_ |  |  |  |
| `managementProxyPort` _integer_ |  |  |  |
| `dataServicesPort` _integer_ |  |  |  |


#### ClusterPrinterColumns







_Appears in:_
- [WekaClusterStatus](#wekaclusterstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `computeContainers` _[StringMetric](#stringmetric)_ |  |  |  |
| `driveContainers` _[StringMetric](#stringmetric)_ |  |  |  |
| `drives` _[StringMetric](#stringmetric)_ |  |  |  |
| `throughput` _[StringMetric](#stringmetric)_ |  |  |  |
| `iops` _[StringMetric](#stringmetric)_ |  |  |  |
| `filesystemCapacity` _[StringMetric](#stringmetric)_ | Information about filesystem capacity: Available/Used |  |  |


#### ContainerAllocations







_Appears in:_
- [WekaContainerStatus](#wekacontainerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `drives` _string array_ |  |  |  |
| `wekaPort` _integer_ |  |  |  |
| `agentPort` _integer_ |  |  |  |
| `failureDomain` _string_ | value of the failure domain label of the node where the container is running |  |  |
| `machineIdentifier` _string_ |  |  |  |
| `netDevices` _string array_ |  |  |  |
| `virtualDrives` _[VirtualDrive](#virtualdrive) array_ | VirtualDrives contains virtual drive allocations for drive sharing mode.<br />Each VirtualDrive maps a virtual UUID to a physical drive UUID with allocated capacity. |  |  |


#### ContainerMetrics







_Appears in:_
- [ContainersMetrics](#containersmetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `numContainers` _[EntityStatefulNum](#entitystatefulnum)_ |  |  |  |
| `processes` _[EntityStatefulNum](#entitystatefulnum)_ |  |  |  |
| `cpuUtilization` _[FloatMetric](#floatmetric)_ |  |  |  |


#### ContainerPrinterColumns







_Appears in:_
- [WekaContainerStatus](#wekacontainerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `processes` _[StringMetric](#stringmetric)_ |  |  |  |
| `drives` _[StringMetric](#stringmetric)_ |  |  |  |
| `activeMounts` _[StringMetric](#stringmetric)_ |  |  |  |
| `managementIPs` _string_ | pretty-printed management IPs |  |  |
| `nodeAffinity` _string_ | node name where the container is running |  |  |


#### ContainerState

_Underlying type:_ _string_





_Appears in:_
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description |
| --- | --- |
| `active` |  |
| `paused` | Paused state indicates that the cluster which owns the container is in grace period before being destroyed<br /> |
| `destroying` | Destroying state indicates that the container can be destroyed without any additional steps (such as deactivation)<br /> |
| `deleting` | Deleting state indicates that the container is being deleted, normally such container should go through deactivation flow<br /> |


#### ContainerStatus

_Underlying type:_ _string_





_Appears in:_
- [WekaContainerStatus](#wekacontainerstatus)

| Field | Description |
| --- | --- |
| `Init` |  |
| `PodNotRunning` |  |
| `PodRunning` |  |
| `PodTerminating` |  |
| `WaitForDrivers` |  |
| `Running` |  |
| `Stopped` |  |
| `Starting` |  |
| `Deleting` |  |
| `Destroying` |  |
| `Paused` |  |
| `Degraded` |  |
| `Unhealthy` |  |
| `Error` |  |
| `DrivesAdding` |  |
| `Draining` |  |
| `Completed` | for drivers-build and adhoc-op-with-container (sign-dives) container<br /> |
| `Building` |  |
| `StoppingAttempt` |  |


#### ContainersMetrics







_Appears in:_
- [ClusterMetrics](#clustermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `drive` _[ContainerMetrics](#containermetrics)_ |  |  |  |
| `compute` _[ContainerMetrics](#containermetrics)_ |  |  |  |
| `s3` _[ContainerMetrics](#containermetrics)_ |  |  |  |
| `nfs` _[ContainerMetrics](#containermetrics)_ |  |  |  |


#### CpuPolicy

_Underlying type:_ _string_





_Appears in:_
- [WekaClientSpec](#wekaclientspec)
- [WekaClusterSpec](#wekaclusterspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description |
| --- | --- |
| `auto` |  |
| `shared` |  |
| `dedicated` |  |
| `dedicated_ht` |  |
| `manual` |  |


#### CsiConfig







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `endpointsSubnets` _string array_ |  |  |  |
| `csiGroup` _string_ |  |  |  |
| `advanced` _[AdvancedCsiConfig](#advancedcsiconfig)_ |  |  |  |


#### DataServicesConfig







_Appears in:_
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `dataServicesFeCores` _integer_ |  |  |  |


#### DiscoverDrivesPayload







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)
- [PolicyPayload](#policypayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |


#### DistServiceStatus



DistServiceStatus holds the status for the enable-local-drivers-distribution policy



_Appears in:_
- [TypedPolicyStatus](#typedpolicystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `serviceUrl` _string_ |  |  |  |


#### DpdkBaseMemoryMbOverride



DpdkBaseMemoryMbOverride specifies DPDK base memory overrides (in MiB) per container mode.
Used for hugepages calculation and resources.json configuration. Default value is 64 MiB per core.
Only positive values are applied; zero or unset means use default.



_Appears in:_
- [WekaClusterSpecOverrides](#wekaclusterspecoverrides)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `drive` _integer_ |  |  |  |
| `compute` _integer_ |  |  |  |
| `s3` _integer_ |  |  |  |
| `nfs` _integer_ |  |  |  |
| `smbw` _integer_ |  |  |  |
| `dataServices` _integer_ |  |  |  |


#### Drive







_Appears in:_
- [WekaContainerStatus](#wekacontainerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `uuid` _string_ |  |  |  |
| `added_time` _string_ |  |  |  |
| `device_path` _string_ |  |  |  |
| `serial_number` _string_ |  |  |  |
| `size_bytes` _integer_ |  |  |  |
| `status` _string_ |  |  |  |


#### DriveClaim









| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `DriveClaim` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[DriveClaimSpec](#driveclaimspec)_ |  |  |  |
| `status` _[DriveClaimStatus](#driveclaimstatus)_ |  |  |  |


#### DriveClaimSpec







_Appears in:_
- [DriveClaim](#driveclaim)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `owner` _string_ |  |  |  |
| `device` _string_ |  |  |  |
| `driveUuid` _string_ |  |  |  |


#### DriveClaimStatus







_Appears in:_
- [DriveClaim](#driveclaim)



#### DriveFailures







_Appears in:_
- [DriveMetrics](#drivemetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `serialId` _string_ |  |  |  |
| `wekaDriveId` _string_ |  |  |  |


#### DriveMetrics







_Appears in:_
- [ClusterMetrics](#clustermetrics)
- [WekaContainerMetrics](#wekacontainermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `counters` _[EntityStatefulNum](#entitystatefulnum)_ |  |  |  |
| `failures` _[DriveFailures](#drivefailures) array_ |  |  |  |


#### DriveTypesRatio







_Appears in:_
- [WekaClusterTemplate](#wekaclustertemplate)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `tlc` _integer_ |  | 0 | Minimum: 0 <br /> |
| `qlc` _integer_ |  | 0 | Minimum: 0 <br /> |


#### DriverDistPayload



DriverDistPayload defines the parameters for the enable-local-drivers-distribution policy



_Appears in:_
- [PolicyPayload](#policypayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ensureImages` _string array_ | EnsureImages is a list of Weka images for which drivers should be proactively built. |  |  |
| `nodeSelectors` _object array_ | NodeSelectors is a list of node selectors. Nodes matching any of these selectors will be considered for driver building.<br />If empty, all nodes in the cluster are considered. |  |  |
| `kernelLabelKey` _string_ | KernelLabelKey is the custom label key to use for storing the node's kernel version.<br />If not specified, "weka.io/kernel" will be used. |  |  |
| `architectureLabelKey` _string_ | ArchitectureLabelKey is the custom label key to use for storing the node's architecture.<br />If not specified, "weka.io/architecture" will be used. |  |  |
| `osLabelKey` _string_ | OsLabelKey is the custom label key to use for storing the node's os.<br />If not specified, "weka.io/os" will be used. |  |  |
| `builderImageOverride` _string_ | BuilderImageOverride is an optional image that you can specify for the builder |  |  |
| `builderPreRunScript` _string_ | BuilderPreRunScript is an optional script to run on builder containers after kernel validation. |  |  |
| `distNodeSelector` _object (keys:string, values:string)_ | DistNodeSelector is the node selector for the drivers distribution (dist) container.<br />If not specified, the dist container will be scheduled on any available node. |  |  |




#### DumperConfigMode

_Underlying type:_ _string_





_Appears in:_
- [TracesConfiguration](#tracesconfiguration)

| Field | Description |
| --- | --- |
| `auto` |  |
| `override` |  |
| `partial-override` |  |
| `cluster` |  |


#### EncryptionConfig







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `vault` _[VaultConfig](#vaultconfig)_ |  |  |  |
| `internal` _[InternalEncryptionConfig](#internalencryptionconfig)_ | InternalConfig defines internal encryption settings, encryption key stored in weka configuration, for production systems use real KMS, however this mode can be useful to evaluate performance of encrypted filesystems |  |  |


#### EnsureNICsPayload







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)
- [PolicyPayload](#policypayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `type` _[NICType](#nictype)_ |  |  |  |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |
| `dataNICsNumber` _integer_ |  |  |  |


#### EntityStatefulNum







_Appears in:_
- [ClientMetrics](#clientmetrics)
- [ContainerMetrics](#containermetrics)
- [DriveMetrics](#drivemetrics)
- [WekaContainerMetrics](#wekacontainermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _[IntMetric](#intmetric)_ |  |  |  |
| `created` _[IntMetric](#intmetric)_ |  |  |  |
| `desired` _[IntMetric](#intmetric)_ |  |  |  |


#### FailureDomain







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `label` _string_ | label used for spreading the weka containers across different failure domains (if set)<br />nodes that have the same value for the label will be considered as a single failure domain |  |  |
| `skew` _integer_ | skew for the failure domain, if set, the weka containers will be spread with the skew in mind<br />(only applicable if `label` is set) |  |  |
| `compositeLabels` _string array_ | If multiple labels are specified, the failure domain will be the combination of the labels.<br />If `compositeLabels` is set, `label` and `skew` will be ignored.<br />When using compositeLabels, weka containers will be spread considering all labels<br />with best effort, but even distribution is not guaranteed |  |  |


#### FilesystemMetrics



FilesystemMetrics contains metrics about filesystem usage



_Appears in:_
- [ClusterMetrics](#clustermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `totalProvisionedCapacity` _[IntMetric](#intmetric)_ | TotalProvisionedCapacity is the sum of total_budget for all filesystems |  |  |
| `totalUsedCapacity` _[IntMetric](#intmetric)_ | TotalUsedCapacity is the sum of used_total for all filesystems |  |  |
| `totalAvailableCapacity` _[IntMetric](#intmetric)_ | TotalAvailableCapacity is the difference between TotalProvisionedCapacity and TotalUsedCapacity |  |  |
| `totalProvisionedSSDCapacity` _[IntMetric](#intmetric)_ | SSD-specific metrics |  |  |
| `totalUsedSSDCapacity` _[IntMetric](#intmetric)_ |  |  |  |
| `totalAvailableSSDCapacity` _[IntMetric](#intmetric)_ |  |  |  |
| `hasTieredFilesystems` _boolean_ | Object Store metrics |  |  |
| `totalObsCapacity` _[IntMetric](#intmetric)_ |  |  |  |
| `obsBucketCount` _[IntMetric](#intmetric)_ |  |  |  |
| `activeObsBucketCount` _[IntMetric](#intmetric)_ |  |  |  |


#### FloatMetric

_Underlying type:_ _string_





_Appears in:_
- [ClusterMetrics](#clustermetrics)
- [ContainerMetrics](#containermetrics)
- [MinMaxAvgPercent](#minmaxavgpercent)
- [WekaContainerMetrics](#wekacontainermetrics)



#### ForceResignDrivesPayload







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nodeName` _[NodeName](#nodename)_ |  |  |  |
| `deviceSerials` _string array_ |  |  |  |
| `devicePaths` _string array_ |  |  |  |




#### Instructions







_Appears in:_
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `type` _string_ |  |  |  |
| `payload` _string_ |  |  |  |


#### IntMetric

_Underlying type:_ _integer_





_Appears in:_
- [CapacityMetrics](#capacitymetrics)
- [ClusterMetrics](#clustermetrics)
- [EntityStatefulNum](#entitystatefulnum)
- [FilesystemMetrics](#filesystemmetrics)
- [StatusIops](#statusiops)
- [StatusThroughput](#statusthroughput)
- [WekaContainerMetrics](#wekacontainermetrics)



#### InternalEncryptionConfig







_Appears in:_
- [EncryptionConfig](#encryptionconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |


#### IoStats







_Appears in:_
- [ClusterMetrics](#clustermetrics)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `throughput` _[StatusThroughput](#statusthroughput)_ |  |  |  |
| `iops` _[StatusIops](#statusiops)_ |  |  |  |


#### ManualOperatorPayload







_Appears in:_
- [WekaManualOperationSpec](#wekamanualoperationspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `signDrivesPayload` _[SignDrivesPayload](#signdrivespayload)_ |  |  |  |
| `blockDrivesPayload` _[BlockDrivesPayload](#blockdrivespayload)_ |  |  |  |
| `discoverDrivesPayload` _[DiscoverDrivesPayload](#discoverdrivespayload)_ |  |  |  |
| `ensureNICsPayload` _[EnsureNICsPayload](#ensurenicspayload)_ |  |  |  |
| `forceResignDrivesPayload` _[ForceResignDrivesPayload](#forceresigndrivespayload)_ |  |  |  |
| `remoteTracesSessionPayload` _[RemoteTracesSessionConfig](#remotetracessessionconfig)_ |  |  |  |




#### NICType

_Underlying type:_ _string_





_Appears in:_
- [EnsureNICsPayload](#ensurenicspayload)

| Field | Description |
| --- | --- |
| `aws` |  |


#### Network







_Appears in:_
- [RoleNetworkSelector](#rolenetworkselector)
- [WekaClientSpec](#wekaclientspec)
- [WekaClusterSpec](#wekaclusterspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ethDevice` _string_ | The name of a single network interface (for example, eth1) to be used by every backend container.<br />This is for clusters that use only one dedicated NIC for the data path.<br />You cannot use this field with ethDevices.<br />If you leave this empty, the system automatically uses the node’s interface associated with the first subnet defined in deviceSubnets. |  |  |
| `ethDevices` _string array_ | A list of network interface names to be used by backend containers when you have multiple dedicated NICs.<br />The order of interfaces in this list is important, as it maps directly to the ethSlots index (the first interface maps to slot-0, the second to slot-1, and so on).<br />You cannot use this field with ethDevice. Ensure that every interface listed here exists on all nodes that are part of the cluster. |  |  |
| `gateway` _string_ | The default gateway IPv4 address for the backend containers’ data-path network.<br />This is only necessary if backend subnets need to communicate with destinations outside of their local network (L2 segment).<br />If you have a flat, non-routed backend network, you can leave this field empty. |  |  |
| `udpMode` _boolean_ | A setting that enables or disables UDP encapsulation for backend traffic.<br />- false (default): Uses standard raw Ethernet frames. true: Wraps data-path traffic in UDP packets.<br />This is required if your network infrastructure or CNI (Container Network Interface) blocks traffic that isn’t IP-based. |  |  |
| `deviceSubnets` _string array_ | A list of backend subnets in CIDR notation (for example, 192.168.10.0/24).<br />The operator assigns IP addresses from these subnets to the backend containers for their data path network |  | items:Pattern: `^([0-9]\{1,3\}\.)\{3\}[0-9]\{1,3\}\/[0-9]\{1,2\}$` <br /> |
| `selectors` _[NetworkSelector](#networkselector) array_ |  |  |  |
| `managementIpsSelectors` _[NetworkSelector](#networkselector) array_ |  |  |  |
| `bindManagementAll` _boolean_ | BindManagementAll controls whether Weka containers bind to all network interfaces or only to specific management interfaces.<br />When set to false (default), containers will only listen on the management ips interfaces (restrict_listen mode).<br />When set to true, containers will listen on all ips (0.0.0.0) instead of specific IP addresses. |  |  |
| `nvidiaVfSingleIp` _boolean_ | NvidiaVfSingleIp indicates whether NVIDIA virtual functions (VFs) should be configured to use a single-ip weka mode, where multiple weka processes can share same VF<br />When not set defaults to false, in future releases, when auto-discovery of capabilities will be implemented not set might translate to true on supported setups |  |  |
| `allocateVfPerIoNode` _boolean_ |  |  |  |


#### NetworkSelector







_Appears in:_
- [Network](#network)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `subnet` _string_ |  |  |  |
| `min` _integer_ |  |  |  |
| `max` _integer_ |  |  |  |
| `deviceNames` _string array_ |  |  |  |
| `rdmaOnly` _boolean_ |  |  |  |
| `disableRdma` _boolean_ |  |  |  |


#### NfsConfig







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `interfaces` _string array_ |  |  |  |
| `ipRanges` _string array_ |  |  |  |


#### NodeName

_Underlying type:_ _[NodeName](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#nodename-types-pkg)_





_Appears in:_
- [ForceResignDrivesPayload](#forceresigndrivespayload)
- [WekaContainerSpec](#wekacontainerspec)
- [WekaContainerStatus](#wekacontainerstatus)



#### ObjectReference







_Appears in:_
- [RemoteTracesSessionConfig](#remotetracessessionconfig)
- [WekaClientSpec](#wekaclientspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ |  |  |  |
| `namespace` _string_ |  |  |  |


#### PCIDevices







_Appears in:_
- [SignDrivesPayload](#signdrivespayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `vendorId` _string_ | VendorId is the 4-digit hexadecimal vendor ID<br />default for AWS: `1d0f` (Amazon.com, Inc.) |  | Pattern: `^[0-9a-fA-F]\{4\}$` <br /> |
| `deviceId` _string_ | DeviceId is the 4-digit hexadecimal device ID<br />default for AWS: `cd01` (NVMe SSD) |  | Pattern: `^[0-9a-fA-F]\{4\}$` <br /> |


#### PVCConfig







_Appears in:_
- [WekaClientSpec](#wekaclientspec)
- [WekaClusterSpec](#wekaclusterspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ |  |  |  |
| `path` _string_ |  |  |  |


#### PodConfiguration







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `topologySpreadConstraints` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | controls the distribution of weka containers across the failure domains |  | Schemaless: \{\} <br /> |
| `roleTopologySpreadConstraints` _[RoleTopologySpreadConstraints](#roletopologyspreadconstraints)_ | takes precedence over the `topologySpreadConstraints` |  |  |
| `affinity` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | advanced scheduling constraints |  | Schemaless: \{\} <br /> |
| `roleAffinity` _[RoleAffinity](#roleaffinity)_ | affinity per container role<br />takes precedence over the `affinity` field |  |  |


#### PodResources







_Appears in:_
- [PodResourcesSpec](#podresourcesspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cpu` _[Quantity](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#quantity-resource-api)_ |  |  |  |
| `memory` _[Quantity](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#quantity-resource-api)_ |  |  |  |


#### PodResourcesSpec







_Appears in:_
- [WekaClientSpec](#wekaclientspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `requests` _[PodResources](#podresources)_ |  |  |  |
| `limits` _[PodResources](#podresources)_ |  |  |  |


#### PolicyPayload







_Appears in:_
- [WekaPolicySpec](#wekapolicyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `signDrivesPayload` _[SignDrivesPayload](#signdrivespayload)_ |  |  |  |
| `schedulingConfigPayload` _[SchedulingConfigPayload](#schedulingconfigpayload)_ |  |  |  |
| `discoverDrivesPayload` _[DiscoverDrivesPayload](#discoverdrivespayload)_ |  |  |  |
| `ensureNICsPayload` _[EnsureNICsPayload](#ensurenicspayload)_ |  |  |  |
| `driverDistPayload` _[DriverDistPayload](#driverdistpayload)_ |  |  |  |
| `remoteTracesSessionPayload` _[RemoteTracesSessionConfig](#remotetracessessionconfig)_ |  |  |  |
| `interval` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ |  | 5m | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(s\|m\|h))+)$` <br />Type: string <br /> |
| `waitForPolicies` _string array_ |  |  |  |


#### PortRange







_Appears in:_
- [WekaClientSpec](#wekaclientspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `basePort` _integer_ |  | 45000 | Maximum: 65535 <br /> |
| `portRange` _integer_ | number of ports to check for availability<br />if 0 - will check all ports from basePort to 65535 | 0 |  |


#### RemoteTracesSessionConfig







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)
- [PolicyPayload](#policypayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cluster` _[ObjectReference](#objectreference)_ |  |  |  |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |
| `hostNetwork` _boolean_ |  |  |  |
| `duration` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | Duration specifies how long the trace session should run.<br />WekaManualOperation: defaults to 1 week if omitted/0. CR auto-deletes after expiration.<br />WekaPolicy: defaults to continuous if omitted/0. Resources cleaned up after expiration.<br />Examples: "30m", "2h", "7d", "168h" |  |  |
| `wekahomeEndpointOverride` _string_ |  |  |  |
| `allowHttpWekahomeEndpoint` _boolean_ |  |  |  |
| `allowInsecureWekahomeEndpoint` _boolean_ |  |  |  |
| `wekahomeCaSecret` _string_ |  |  |  |


#### RoleAffinity







_Appears in:_
- [PodConfiguration](#podconfiguration)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `drive` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `s3` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `nfs` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `smbw` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |


#### RoleAnnotations







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _map[string]string_ | annotations for compute weka containers |  |  |
| `drive` _map[string]string_ | annotations for drive weka containers |  |  |
| `s3` _map[string]string_ | annotations for s3 weka containers |  |  |
| `nfs` _map[string]string_ | annotations for nfs weka containers |  |  |
| `smbw` _map[string]string_ | annotations for smbw weka containers |  |  |
| `dataServices` _map[string]string_ | annotations for data services weka containers |  |  |


#### RoleCoreIds



RoleCoreIds defines CPU core id lists per container role for Manual CPU policy.
Each slice contains the core IDs (as visible to the node OS) that should be
pinned to every container of the corresponding role. If a slice is empty or
omitted, no explicit core pinning will be applied for that role.

_Validation:_
- Type: object

_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _integer array_ |  |  | Optional: \{\} <br /> |
| `drive` _integer array_ |  |  | Optional: \{\} <br /> |
| `s3` _integer array_ |  |  | Optional: \{\} <br /> |
| `nfs` _integer array_ |  |  | Optional: \{\} <br /> |
| `smbw` _integer array_ |  |  | Optional: \{\} <br /> |
| `dataServices` _integer array_ |  |  | Optional: \{\} <br /> |


#### RoleNetworkSelector







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _[Network](#network)_ | network selector for compute weka containers |  |  |
| `drive` _[Network](#network)_ | network selector for drive weka containers |  |  |
| `s3` _[Network](#network)_ | network selector for s3 weka containers |  |  |
| `nfs` _[Network](#network)_ | network selector for nfs weka containers |  |  |
| `smbw` _[Network](#network)_ | network selector for smbw weka containers |  |  |
| `dataServices` _[Network](#network)_ | network selector for data services weka containers |  |  |


#### RoleNodeSelector







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _map[string]string_ | nodeSelector for compute weka containers |  |  |
| `drive` _map[string]string_ | nodeSelector for drive weka containers |  |  |
| `s3` _map[string]string_ | nodeSelector for s3 weka containers |  |  |
| `nfs` _map[string]string_ | nodeSelector for nfs weka containers |  |  |
| `smbw` _map[string]string_ | nodeSelector for smbw weka containers |  |  |
| `dataServices` _map[string]string_ | nodeSelector for data services weka containers |  |  |


#### RoleTopologySpreadConstraints







_Appears in:_
- [PodConfiguration](#podconfiguration)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `compute` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `drive` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `s3` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `nfs` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |
| `smbw` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  | Schemaless: \{\} <br /> |


#### S3Config







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `clusterCreateArgs` _string array_ | No overlap validation, only appended to the cluster create command as-is<br />Useful for settings such as: `--envoy-max-requests 1150 --envoy-max-connections 1300 --envoy-max-pending-requests 1450`<br />Not propagated to already created cluster, and direct weka control should be used for that |  |  |




#### SchedulingConfigPayload







_Appears in:_
- [PolicyPayload](#policypayload)



#### SignDrivesPayload







_Appears in:_
- [ManualOperatorPayload](#manualoperatorpayload)
- [PolicyPayload](#policypayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `type` _string_ |  |  | Enum: [aws-all gcp-all device-identifiers device-paths all-not-root] <br /> |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |
| `devicePaths` _string array_ |  |  |  |
| `pciDevices` _[PCIDevices](#pcidevices)_ | PCI vendor and device IDs of the drives to sign.<br />To get the values for VendorId and DeviceId:<br />1. Run the following command to list all PCI devices on your system:<br />   ```bash<br />   lspci -nn<br />   ```<br />2. Find the relevant PCI device in the output, which will display both the<br />   vendor and device IDs in square brackets in the format [vendorId:deviceId].<br />   For example:<br />   ```<br />   00:1f.0 Non-Volatile memory controller [0108]: Amazon.com, Inc. NVMe SSD Controller [1d0f:cd01]<br />   ``` |  |  |
| `options` _[SignOptions](#signoptions)_ |  |  |  |
| `shared` _boolean_ | Shared enables shared drive signing for proxy mode (defaults to false).<br />When enabled:<br />- Drives are signed for proxy using 'weka-sign-drive sign proxy' command<br />- Drives are signed with a proxy system GUID<br />- Results are stored in weka.io/shared-drives annotation (instead of weka.io/weka-drives)<br />- Physical UUIDs, serial IDs, and capacities are captured<br />- Enables multi-tenant drive sharing via SSD proxy |  |  |


#### SignOptions







_Appears in:_
- [SignDrivesPayload](#signdrivespayload)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowEraseWekaPartitions` _boolean_ |  |  |  |
| `allowEraseNonWekaPartitions` _boolean_ |  |  |  |
| `allowNonEmptyDevice` _boolean_ |  |  |  |
| `skipTrimFormat` _boolean_ |  |  |  |


#### SmbwConfig







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `clusterName` _string_ | ClusterName is the SMB-W cluster name, defaults to "default" |  |  |
| `domainName` _string_ | DomainName is the domain name for SMB-W, required for SMB-W cluster creation |  |  |
| `domainJoinSecret` _string_ |  |  |  |
| `userName` _string_ |  |  |  |
| `ipRanges` _string array_ | IpRanges specifies floating IP ranges for SMB-W high availability |  |  |
| `symlink` _boolean_ | Creation-time configuration flags<br />Symlink enables symlink support for SMB-W shares |  |  |
| `domainNetbiosName` _string_ | DomainNetbiosName is the NetBIOS name for the domain |  |  |
| `idmapBackend` _string_ | IdmapBackend specifies the identity mapping backend (e.g., "ad", "rfc2307") |  |  |
| `defaultDomainMappingFromId` _integer_ | DefaultDomainMappingFromId is the start of the UID/GID range for default domain mapping |  |  |
| `defaultDomainMappingToId` _integer_ | DefaultDomainMappingToId is the end of the UID/GID range for default domain mapping |  |  |
| `joinedDomainMappingFromId` _integer_ | JoinedDomainMappingFromId is the start of the UID/GID range for joined domain mapping |  |  |
| `joinedDomainMappingToId` _integer_ | JoinedDomainMappingToId is the end of the UID/GID range for joined domain mapping |  |  |
| `encryption` _string_ | Encryption specifies the encryption level for SMB connections |  | Enum: [enabled disabled desired required] <br /> |
| `scaleOutMode` _string_ | ScaleOutMode specifies the scale-out mode for SMB-W clustering |  | Enum: [none full partial] <br /> |
| `smbConfExtra` _string_ | SmbConfExtra contains additional smb.conf configuration |  |  |
| `ipPools` _string array_ | IpPools specifies IP pools for SMB-W service assignment |  |  |


#### SplunkExportConfig



SplunkExportConfig defines Splunk-specific export configuration



_Appears in:_
- [TelemetryExport](#telemetryexport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `authTokenSecretRef` _string_ | AuthTokenSecretRef references a secret containing the Splunk HEC authentication token.<br />Format: "secretName.keyName" where secretName is the name of the secret in the same namespace<br />and keyName is the key within the secret's data that contains the token. |  |  |
| `endpoint` _string_ | Endpoint is the Splunk HEC endpoint URL (maps to --target in weka CLI) |  |  |
| `caCertSecretRef` _string_ | CACertSecretRef optionally references a secret containing a user-provided CA certificate PEM file.<br />Format: "secretName.keyName" where secretName is the name of the secret in the same namespace<br />and keyName is the key within the secret's data that contains the certificate.<br />Maps to --ca-cert in weka CLI. Empty string is treated same as nil (de-configures if was configured).<br />Mutually exclusive with VerifyWithClusterCACert. |  |  |
| `allowUnverifiedCertificate` _boolean_ | AllowUnverifiedCertificate allows accessing without verifying the target certificate.<br />Maps to --allow-unverified-certificate in weka CLI. |  |  |
| `verifyWithClusterCACert` _boolean_ | VerifyWithClusterCACert uses the Weka cluster's internal CA certificate to verify.<br />Maps to --verify-with-cluster-cacert in weka CLI.<br />Mutually exclusive with CACertSecretRef. |  |  |


#### StartIoConditions







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `minNumDrives` _integer_ | minimum number of drives that should be added to the cluster before starting IO |  |  |


#### StatusIops







_Appears in:_
- [IoStats](#iostats)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `read` _[IntMetric](#intmetric)_ |  |  |  |
| `write` _[IntMetric](#intmetric)_ |  |  |  |
| `metadata` _[IntMetric](#intmetric)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `total` _[IntMetric](#intmetric)_ |  |  |  |


#### StatusThroughput







_Appears in:_
- [IoStats](#iostats)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `read` _[IntMetric](#intmetric)_ |  |  |  |
| `write` _[IntMetric](#intmetric)_ |  |  |  |


#### StringMetric

_Underlying type:_ _string_





_Appears in:_
- [ClientPrinterColumns](#clientprintercolumns)
- [ClusterMetrics](#clustermetrics)
- [ClusterPrinterColumns](#clusterprintercolumns)
- [ContainerPrinterColumns](#containerprintercolumns)



#### TelemetryConfig



TelemetryConfig defines the telemetry export configuration for the Weka cluster



_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `exports` _[TelemetryExport](#telemetryexport) array_ | List of telemetry exports to configure |  |  |


#### TelemetryExport



TelemetryExport defines a single telemetry export destination



_Appears in:_
- [TelemetryConfig](#telemetryconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name is the unique identifier for this export |  |  |
| `sources` _string array_ | Sources specifies which telemetry sources to export (e.g., "audit") |  |  |
| `splunk` _[SplunkExportConfig](#splunkexportconfig)_ | Splunk configuration for Splunk HEC export |  |  |


#### TracesConfiguration



TraceConfiguration defines the configuration for the traces, accepts parameters in gigabytes



_Appears in:_
- [WekaClientSpec](#wekaclientspec)
- [WekaClusterSpec](#wekaclusterspec)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `maxCapacityPerIoNode` _integer_ |  | 10 |  |
| `ensureFreeSpace` _integer_ |  | 20 |  |
| `dumperConfigMode` _[DumperConfigMode](#dumperconfigmode)_ |  | auto | Enum: [override partial-override auto cluster] <br /> |


#### TypedPolicyStatus



TypedPolicyStatus holds status fields specific to a policy type



_Appears in:_
- [WekaPolicyStatus](#wekapolicystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `distService` _[DistServiceStatus](#distservicestatus)_ |  |  |  |


#### UpgradePolicy







_Appears in:_
- [WekaClientSpec](#wekaclientspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `type` _[UpgradePolicyType](#upgradepolicytype)_ |  | all-at-once | Enum: [manual all-at-once rolling all-at-once-force] <br /> |


#### UpgradePolicyType

_Underlying type:_ _string_





_Appears in:_
- [UpgradePolicy](#upgradepolicy)
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description |
| --- | --- |
| `manual` |  |
| `rolling` |  |
| `all-at-once` |  |
| `all-at-once-force` |  |


#### VaultConfig







_Appears in:_
- [EncryptionConfig](#encryptionconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `address` _string_ | Vault address, e.g. "https://vault.example.com:8200". |  |  |
| `role` _string_ | Role to authenticate as in Vault. |  |  |
| `authPath` _string_ | Path under auth/ that the weka uses for login. defaults to "kubernetes" | kubernetes |  |
| `transitPath` _string_ | Transit engine mount path, defaults "transit". | transit |  |
| `method` _string_ | Vault Auth method (only “kubernetes” is supported  on operator side.) | kubernetes | Enum: [kubernetes] <br /> |
| `keyName` _string_ | Name of the transit key. defaults to "weka-key" | weka-key |  |


#### VirtualDrive



VirtualDrive represents a virtual drive allocation in drive sharing mode



_Appears in:_
- [ContainerAllocations](#containerallocations)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `virtualUUID` _string_ | VirtualUUID is the virtual drive identifier |  |  |
| `physicalUUID` _string_ | PhysicalUUID is the physical drive UUID obtained from proxy signing |  |  |
| `capacityGiB` _integer_ | CapacityGiB is the allocated capacity in GiB |  |  |
| `serial` _string_ | Serial is the serial number of the physical drive |  |  |
| `type` _string_ | Type is the type of the drive (e.g., TLC, QLC) |  |  |


#### WekaClient



WekaClient is the Schema for the clients API





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `WekaClient` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[WekaClientSpec](#wekaclientspec)_ |  |  |  |
| `status` _[WekaClientStatus](#wekaclientstatus)_ |  |  |  |


#### WekaClientSpec



WekaClientSpec defines the desired state of WekaClient



_Appears in:_
- [WekaClient](#wekaclient)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `image` _string_ | full container image in format of quay.io/weka.io/weka-in-container:VERSION |  | Pattern: `^.+:\d+\.\d+\.\d+.*$` <br /> |
| `imagePullSecret` _string_ |  |  |  |
| `port` _integer_ | if not set (0), weka will find a free port from the portRange |  |  |
| `agentPort` _integer_ | if not set (0), weka will find a free port from the portRange |  |  |
| `portRange` _[PortRange](#portrange)_ | used for dynamic port allocation |  |  |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |
| `wekaSecretRef` _string_ |  |  |  |
| `network` _[Network](#network)_ |  |  |  |
| `driversDistService` _string_ |  |  |  |
| `joinIpPorts` _string array_ |  |  |  |
| `targetCluster` _[ObjectReference](#objectreference)_ |  |  |  |
| `cpuPolicy` _[CpuPolicy](#cpupolicy)_ |  | auto | Enum: [auto shared dedicated dedicated_ht manual] <br /> |
| `cpuRequest` _string_ |  |  |  |
| `coresNum` _integer_ |  |  |  |
| `coreIds` _integer array_ |  |  |  |
| `nonDatapathCoreIds` _integer array_ |  |  |  |
| `tracesConfiguration` _[TracesConfiguration](#tracesconfiguration)_ |  |  |  |
| `tolerations` _string array_ |  |  |  |
| `rawTolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |
| `additionalMemory` _integer_ | memory to add/decrease from "auto-calculated" memory |  |  |
| `resources` _[PodResourcesSpec](#podresourcesspec)_ | experimental: pod resources to be proxied as-is to the pod spec |  |  |
| `hugepages` _integer_ | hugepages, value in megabytes |  |  |
| `hugepagesOffset` _integer_ | value in megabytes to offset |  |  |
| `wekaHomeConfig` _[WekahomeClientConfig](#wekahomeclientconfig)_ | DEPRECATED, kept for compatibility with old API clients, not taking any action, to be removed on new API version |  |  |
| `wekaHome` _[WekahomeClientConfig](#wekahomeclientconfig)_ |  |  |  |
| `upgradePolicy` _[UpgradePolicy](#upgradepolicy)_ |  |  |  |
| `allowHotUpgrade` _boolean_ |  |  |  |
| `overrides` _[WekaClientSpecOverrides](#wekaclientspecoverrides)_ |  |  |  |
| `autoRemoveTimeout` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | sets weka cluster-side timeout, if client is not coming back in specified duration it will be auto removed from cluster config | 24h | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(ns\|us\|µs\|ms\|s\|m\|h))+)$` <br />Type: string <br /> |
| `globalPVC` _[PVCConfig](#pvcconfig)_ |  |  |  |
| `csiConfig` _[ClientCsiConfig](#clientcsiconfig)_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: if set, allows to reuse the same csi resources for multiple clients |  | Type: object <br /> |


#### WekaClientSpecOverrides







_Appears in:_
- [WekaClientSpec](#wekaclientspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `driversBuildId` _string_ | can be used to specify a build_id for a driver in the distributor service, keep empty for auto detection default |  |  |
| `driversLoaderImage` _string_ |  |  |  |
| `machineIdentifierNodeRef` _string_ | used to override machine identifier node reference for client containers |  |  |
| `forceDrain` _boolean_ | unsafe operation, forces drain on the node where the container is running, should not be used unless instructed explicitly by weka personnel, the effect of drain is throwing away all IOs and acknowledging all umounts in unsafe manner |  |  |
| `skipActiveMountsCheck` _boolean_ | option to skip active mounts check before deleting client containers |  |  |
| `umountOnHost` _boolean_ | unsafe operation, runs nsenter in root namespace to umount all wekafs mounts visible on host |  |  |
| `dropAffinityConstraints` _boolean_ | unsafe parameter, disables anti-affinities on client pods, allowing to schedule more than one client pod per node.<br />Running multiple clients for multiple clusters on the same node is not fully supported yet, and this flag should not be used in production. |  |  |
| `wekaContainerName` _string_ | override name used in weka local setup for the container<br />this can be used for integration with external client on the host |  |  |
| `dpdkBaseMemoryMb` _integer_ |  |  |  |


#### WekaClientStatus



WekaClientStatus defines the observed state of WekaClient



_Appears in:_
- [WekaClient](#wekaclient)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ |  |  |  |
| `lastAppliedSpec` _string_ |  |  |  |
| `status` _[WekaClientStatusEnum](#wekaclientstatusenum)_ |  | Init | Enum: [Init Running Upgrading Destroying] <br /> |
| `stats` _[ClientMetrics](#clientmetrics)_ |  |  |  |
| `printer` _[ClientPrinterColumns](#clientprintercolumns)_ |  |  |  |


#### WekaClientStatusEnum

_Underlying type:_ _string_





_Appears in:_
- [WekaClientStatus](#wekaclientstatus)

| Field | Description |
| --- | --- |
| `Init` |  |
| `Running` |  |
| `Upgrading` |  |
| `Destroying` |  |


#### WekaCluster









| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `WekaCluster` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[WekaClusterSpec](#wekaclusterspec)_ |  |  |  |
| `status` _[WekaClusterStatus](#wekaclusterstatus)_ |  |  |  |


#### WekaClusterSpec



WekaClusterSpec defines the desired state of WekaCluster



_Appears in:_
- [WekaCluster](#wekacluster)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `template` _string_ | A template/strategy of how to build a cluster, right now only "dynamic" supported, explicitly specifying config of a cluster | dynamic |  |
| `image` _string_ | full container image name in format of quay.io/weka.io/weka-in-container:VERSION |  | Pattern: `^.+:\d+\.\d+\.\d+.*$` <br /> |
| `imagePullSecret` _string_ | image pull secret to use for pulling the image |  |  |
| `driversDistService` _string_ | endpoint for distribution service, global https://drivers.weka.io or in-k8s-cluster "https://weka-drivers-dist.namespace.svc.cluster.local:60001" |  |  |
| `nodeSelector` _object (keys:string, values:string)_ | node selector for the weka containers |  |  |
| `roleNodeSelector` _[RoleNodeSelector](#rolenodeselector)_ | node selector for the weka containers per role, overrides global nodeSelector |  |  |
| `roleAnnotations` _[RoleAnnotations](#roleannotations)_ | annotations for the weka containers per role |  |  |
| `roleNetworkSelector` _[RoleNetworkSelector](#rolenetworkselector)_ | network selector for the weka containers per role, overrides global network |  |  |
| `failureDomain` _[FailureDomain](#failuredomain)_ | failure domain configuration for weka containers |  |  |
| `podConfig` _[PodConfiguration](#podconfiguration)_ | advanced pod affinities configuration |  |  |
| `cpuPolicy` _[CpuPolicy](#cpupolicy)_ | cpu policy to use for scheduling cores for weka, unless instructed by weka team, keep default of auto<br />manual and shared are same, with shared being deprecated<br />when manual is used - no exclusive cores will be allocaated on k8s/cgroup level, assuming good alignment of cores usage across different applications, like weka and slurm<br />there is no need to specify siblings in this list, but on the side of other applications like slurm, both weka core and its siblings should be excluded from used cpu set | auto | Enum: [auto shared dedicated dedicated_ht manual] <br /> |
| `tracesConfiguration` _[TracesConfiguration](#tracesconfiguration)_ | traces capacities configuration for weka containers |  |  |
| `tolerations` _string array_ | simplified tolerations, checked only by key existence, expanding to NoExecute\|NoSchedule tolerations |  |  |
| `rawTolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ | tolerations in standard k8s format |  |  |
| `wekaHome` _[WekaHomeConfig](#wekahomeconfig)_ | weka home configuration |  |  |
| `ipv6` _boolean_ | use ipv6 for weka cluster networking configuration |  |  |
| `additionalMemory` _[AdditionalMemory](#additionalmemory)_ | additional memory to allocate for weka containers |  |  |
| `ports` _[ClusterPorts](#clusterports)_ | port allocation for weka containers, if not set, free range will be auto selected. Currently allocated ports can be seen in wekacluster.status.ports |  |  |
| `operatorSecretRef` _string_ | reference to the secret containing the weka system credentials used by operator, used in flow of migration |  |  |
| `expandEndpoints` _string array_ | endpoint of existing weka cluster, containers created for this k8s-driver cluster will join existing weka cluster, used in flow of migration |  |  |
| `dynamicTemplate` _[WekaClusterTemplate](#wekaclustertemplate)_ | weka cluster topology configuration |  |  |
| `network` _[Network](#network)_ | weka cluster network configuration |  |  |
| `hotSpare` _integer_ | A hot spare is reserved capacity designed to handle data rebuilds while maintaining the system's net capacity, even in the event of failure domains being lost<br />See: https://docs.weka.io/weka-system-overview/ssd-capacity-management#hot-spare | 0 |  |
| `redundancyLevel` _integer_ | storage capacity dedicated to system protection (2/4). https://docs.weka.io/weka-system-overview/ssd-capacity-management#protection-level |  |  |
| `stripeWidth` _integer_ | stripe width is the number of blocks within a common protection set, ranging from 3 to 16 https://docs.weka.io/weka-system-overview/ssd-capacity-management#stripe-width |  |  |
| `leadershipRaftSize` _integer_ | size of raft for leadership, defaults to 5, 5/9 are supported |  |  |
| `bucketRaftSize` _integer_ | size of raft for buckets, defaults to 5, 5/9 are supported |  |  |
| `startIoConditions` _[StartIoConditions](#startioconditions)_ | conditions that must be met before starting IO |  |  |
| `gracefulDestroyDuration` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | During this period the cluster will not be destroyed (protection from accidental deletion)<br />Note: due to discrepancies in validation vs parsing, we use a Pattern instead of `Format=duration`. See<br />https://bugzilla.redhat.com/show_bug.cgi?id=2050332<br />https://github.com/kubernetes/apimachinery/issues/131<br />https://github.com/kubernetes/apiextensions-apiserver/issues/56 | 24h | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(ns\|us\|µs\|ms\|s\|m\|h))+)$` <br />Type: string <br /> |
| `overrides` _[WekaClusterSpecOverrides](#wekaclusterspecoverrides)_ |  |  |  |
| `csiConfig` _[CsiConfig](#csiconfig)_ |  |  |  |
| `globalPVC` _[PVCConfig](#pvcconfig)_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |
| `roleCoreIds` _[RoleCoreIds](#rolecoreids)_ | RoleCoreIds defines a list of CPU core IDs (as seen by the host) that should<br />be assigned to containers of the specific role when CpuPolicy is set to<br />"manual". If the slice for the given role is empty, core ids will not be<br />set for that role, and the manual policy will fail validation on pod start.<br />NOTE: The semantics are the same as for NodeSelector/Annotations structures –<br />a single list per role which will be copied to every container of that role.<br />Users are responsible to provide a set that makes sense for their topology.<br />Example:<br />  roleCoreIds:<br />    compute: [0,1,2,3]<br />    drive:   [4,5,6,7]<br />will result in every compute container getting coreIds [0,1,2,3] and every<br />drive container getting [4,5,6,7]. |  | Type: object <br /> |
| `roleNonDatapathCoreIds` _[RoleCoreIds](#rolecoreids)_ | RoleNonDatapathCoreIds defines CPU core IDs (as seen by the host) to pin<br />management/aux (non-IONode) processes to, per container role. Applicable<br />when CpuPolicy is "manual" or "shared".<br />When set, weka pins management processes to these cores instead of deriving them automatically. |  | Type: object <br /> |
| `encryption` _[EncryptionConfig](#encryptionconfig)_ |  |  |  |
| `nfs` _[NfsConfig](#nfsconfig)_ |  |  |  |
| `s3` _[S3Config](#s3config)_ |  |  |  |
| `smbw` _[SmbwConfig](#smbwconfig)_ |  |  |  |
| `telemetry` _[TelemetryConfig](#telemetryconfig)_ | Telemetry configuration for exporting audit logs and other telemetry data |  |  |
| `catalog` _[CatalogConfig](#catalogconfig)_ | Catalog configuration for data catalog service |  |  |


#### WekaClusterSpecOverrides







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowS3ClusterDestroy` _boolean_ |  |  |  |
| `allowSmbwClusterDestroy` _boolean_ |  |  |  |
| `disregardRedundancy` _boolean_ | disregard redundancy constraints, useful for testing, should not be used in production as misaligns failure domains |  |  |
| `driversBuildId` _string_ | can be used to specify a build_id for a driver in the distributor service, keep empty for auto detection default |  |  |
| `driversLoaderImage` _string_ | image to be used for loading drivers, do not use unless explicitly instructed by Weka team |  |  |
| `forceAio` _boolean_ | force weka to use drives in aio mode and not direct nvme (impacts performance, but might serve as a fallback in case of incompatible device) |  |  |
| `postFormClusterScript` _string_ | script to run post cluster create (i.e before starting io) |  |  |
| `upgradeForceReplace` _boolean_ | unsafe operation, skips graceful stop of weka container for a quick replacement to a new image, should not be used unless instructed explicitly by weka personnel |  |  |
| `upgradeForceReplaceDrives` _boolean_ | unsafe operation, skips graceful stop of drive weka container for a quick replacement to a new image, should not be used unless instructed explicitly by weka personnel |  |  |
| `upgradeAllAtOnce` _boolean_ | unsafe operation, should not be used unless instructed explicitly by weka personnel |  |  |
| `upgradePaused` _boolean_ | Pause upgrade |  |  |
| `upgradePausePreCompute` _boolean_ | Prevent from moving into compute phase |  |  |
| `podTerminationDeactivationTimeout` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | Timeout duration for deactivating pods that are terminating longer than this duration.<br />When nil (default), the default timeout of 5 minutes is used.<br />When set to 0, deactivation of terminating pods is disabled.<br />Otherwise, the specified duration is used. |  | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(ns\|us\|µs\|ms\|s\|m\|h))+)$` <br />Type: string <br />Optional: \{\} <br /> |
| `paused` _boolean_ | Pause the cluster - all containers will be stopped forcefully.<br />nil (not set): no propagation, allows direct container-level state manipulation.<br />true: pause all containers.<br />false: actively unpause containers that are in paused state. |  |  |
| `cancelDeletion` _boolean_ | Cancel deletion of the cluster if it is in graceful destroy period, a disaster recovery mechanism |  |  |
| `dpdkBaseMemoryMb` _[DpdkBaseMemoryMbOverride](#dpdkbasememorymboverride)_ |  |  |  |
| `machineIdentifierNodeRef` _string_ | used to override machine identifier node reference for backend containers (drive, compute, etc.) |  |  |


#### WekaClusterStatus



WekaClusterStatus defines the observed state of WekaCluster



_Appears in:_
- [WekaCluster](#wekacluster)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `status` _[WekaClusterStatusEnum](#wekaclusterstatusenum)_ |  |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ |  |  |  |
| `clusterID` _string_ |  |  |  |
| `traceId` _string_ |  |  |  |
| `spanId` _string_ |  |  |  |
| `lastAppliedImage` _string_ |  |  |  |
| `lastAppliedSpec` _string_ |  |  |  |
| `lastAppliedPodConfigHash` _string_ |  |  |  |
| `ports` _[ClusterPorts](#clusterports)_ |  |  |  |
| `stats` _[ClusterMetrics](#clustermetrics)_ |  |  |  |
| `printer` _[ClusterPrinterColumns](#clusterprintercolumns)_ |  |  |  |
| `timestamps` _object (keys:string, values:[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta))_ |  |  |  |


#### WekaClusterStatusEnum

_Underlying type:_ _string_





_Appears in:_
- [WekaClusterStatus](#wekaclusterstatus)

| Field | Description |
| --- | --- |
| `Init` |  |
| `Ready` |  |
| `WaitForDrives` |  |
| `StartingIO` |  |
| `Paused` |  |
| `GracePeriod` |  |
| `Destroying` |  |
| `Deallocating` |  |


#### WekaClusterTemplate







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `computeContainers` _integer_ |  |  | Minimum: 0 <br /> |
| `driveContainers` _integer_ |  |  | Minimum: 0 <br /> |
| `s3Containers` _integer_ |  |  | Minimum: 0 <br /> |
| `computeCores` _integer_ |  |  | Minimum: 0 <br /> |
| `driveCores` _integer_ |  |  | Minimum: 0 <br /> |
| `s3Cores` _integer_ |  |  | Minimum: 0 <br /> |
| `numDrives` _integer_ |  |  | Minimum: 0 <br /> |
| `computeExtraCores` _integer_ |  |  | Minimum: 0 <br /> |
| `driveExtraCores` _integer_ |  |  | Minimum: 0 <br /> |
| `s3ExtraCores` _integer_ |  |  | Minimum: 0 <br /> |
| `driveHugepages` _integer_ |  |  | Minimum: 0 <br /> |
| `driveHugepagesOffset` _integer_ |  |  | Minimum: 0 <br /> |
| `computeHugepages` _integer_ |  |  | Minimum: 0 <br /> |
| `computeHugepagesOffset` _integer_ |  |  | Minimum: 0 <br /> |
| `s3FrontendHugepages` _integer_ |  |  | Minimum: 0 <br /> |
| `s3FrontendHugepagesOffset` _integer_ |  |  | Minimum: 0 <br /> |
| `envoyCores` _integer_ |  |  | Minimum: 0 <br /> |
| `nfsContainers` _integer_ |  |  | Minimum: 0 <br /> |
| `nfsCores` _integer_ |  |  | Minimum: 0 <br /> |
| `nfsExtraCores` _integer_ |  |  | Minimum: 0 <br /> |
| `nfsFrontendHugepages` _integer_ |  |  | Minimum: 0 <br /> |
| `nfsFrontendHugepagesOffset` _integer_ |  |  | Minimum: 0 <br /> |
| `smbwContainers` _integer_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: number of SMB-W containers (3-8) |  | Minimum: 0 <br /> |
| `smbwCores` _integer_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: number of SMB-W cores per container |  | Minimum: 0 <br /> |
| `smbwExtraCores` _integer_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: number of SMB-W extra cores per container |  | Minimum: 0 <br /> |
| `smbwFrontendHugepages` _integer_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: hugepage allocation for SMB-W frontend |  | Minimum: 0 <br /> |
| `smbwFrontendHugepagesOffset` _integer_ | EXPERIMENTAL, ALPHA STATE, should not be used in production: hugepage offset for SMB-W frontend |  | Minimum: 0 <br /> |
| `driveCapacity` _integer_ | DriveCapacity is the capacity in GiB to allocate per single virtual drive.<br />NumDrives multiplied by DriveCapacity gives the total capacity requested by each drive container.<br />This value determines how much capacity each container receives from shared drives. |  | Minimum: 0 <br /> |
| `containerCapacity` _integer_ | ContainerCapacity specifies the total capacity (in GiB) requested by each container when using shared drives via SSD proxy.<br />This value takes precedence over DriveCapacity when both are set. It allows more flexible capacity allocation. |  | Minimum: 0 <br /> |
| `driveTypesRatio` _[DriveTypesRatio](#drivetypesratio)_ | DriveTypesRatio specifies the desired ratio of drive types (TLC vs QLC) when allocating drives for the cluster. |  |  |
| `dataServicesContainers` _integer_ |  |  | Minimum: 0 <br /> |
| `dataServicesCores` _integer_ |  |  | Minimum: 0 <br /> |
| `dataServicesExtraCores` _integer_ |  |  |  |
| `dataServicesHugepages` _integer_ |  |  | Minimum: 0 <br /> |
| `dataServicesHugepagesOffset` _integer_ |  |  | Minimum: 0 <br /> |
| `dataServicesFeCores` _integer_ |  |  | Minimum: 0 <br /> |


#### WekaContainer









| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `WekaContainer` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[WekaContainerSpec](#wekacontainerspec)_ |  |  |  |
| `status` _[WekaContainerStatus](#wekacontainerstatus)_ |  |  |  |


#### WekaContainerBuilder







_Appears in:_
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `weka_version` _string_ |  |  |  |


#### WekaContainerMetrics







_Appears in:_
- [WekaContainerStatus](#wekacontainerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `processes` _[EntityStatefulNum](#entitystatefulnum)_ |  |  |  |
| `cpuUtilization` _[FloatMetric](#floatmetric)_ |  |  |  |
| `drives` _[DriveMetrics](#drivemetrics)_ |  |  |  |
| `activeMounts` _[IntMetric](#intmetric)_ |  |  |  |
| `lastUpdate` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ |  |  |  |




#### WekaContainerSpec







_Appears in:_
- [WekaContainer](#wekacontainer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nodeAffinity` _[NodeName](#nodename)_ | name of the node where the container should run on |  |  |
| `failureDomain` _[FailureDomain](#failuredomain)_ | failure domain configuration |  |  |
| `topologySpreadConstraints` _[TopologySpreadConstraint](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#topologyspreadconstraint-v1-core) array_ | controls the distribution of weka containers across the failure domains |  |  |
| `affinity` _[Affinity](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#affinity-v1-core)_ | advanced scheduling constraints |  |  |
| `nodeSelector` _object (keys:string, values:string)_ |  |  |  |
| `port` _integer_ |  |  |  |
| `exposePorts` _integer array_ | deprecated, use ExposedPorts instead |  |  |
| `exposedPorts` _[ContainerPort](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#containerport-v1-core) array_ | ports to be exposed on the container, proxied to pod |  |  |
| `agentPort` _integer_ |  |  |  |
| `portRange` _[PortRange](#portrange)_ |  |  |  |
| `image` _string_ |  |  |  |
| `podConfigHash` _string_ | a hash that represents the config state of a WekaContainer, will recreate the pod if stale |  |  |
| `imagePullSecret` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `mode` _string_ |  |  | Enum: [drive compute client dist drivers-dist drivers-loader drivers-builder discovery s3 adhoc-op-with-container adhoc-op envoy nfs smbw telemetry ssdproxy data-services data-services-fe] <br /> |
| `numCores` _integer_ |  |  |  |
| `extraCores` _integer_ |  |  |  |
| `coreIds` _integer array_ |  |  |  |
| `nonDatapathCoreIds` _integer array_ | NonDatapathCoreIds pins management/aux (non-IONode) processes to specific CPUs. |  |  |
| `cpuPolicy` _[CpuPolicy](#cpupolicy)_ |  | auto | Enum: [auto shared dedicated dedicated_ht manual] <br /> |
| `network` _[Network](#network)_ |  |  |  |
| `hugepages` _integer_ |  |  |  |
| `hugepagesOffset` _integer_ |  |  |  |
| `hugepagesSize` _string_ |  |  |  |
| `numDrives` _integer_ |  |  |  |
| `driversDistService` _string_ |  |  |  |
| `driversLoaderImage` _string_ |  |  |  |
| `driversBuildId` _string_ |  |  |  |
| `builder` _[WekaContainerBuilder](#wekacontainerbuilder)_ |  |  |  |
| `wekaSecretRef` _[EnvVarSource](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#envvarsource-v1-core)_ |  |  |  |
| `joinIpPorts` _string array_ |  |  |  |
| `tracesConfiguration` _[TracesConfiguration](#tracesconfiguration)_ |  |  |  |
| `tolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `nodeInfoConfigMap` _string_ |  |  |  |
| `ipv6` _boolean_ |  |  |  |
| `additionalMemory` _integer_ |  |  |  |
| `group` _string_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |
| `additionalSecrets` _object (keys:string, values:string)_ |  |  |  |
| `instructions` _[Instructions](#instructions)_ |  |  |  |
| `dropAffinityConstraints` _boolean_ |  |  |  |
| `uploadResultsTo` _string_ |  |  |  |
| `upgradePolicyType` _[UpgradePolicyType](#upgradepolicytype)_ |  | manual | Enum: [manual all-at-once rolling all-at-once-force] <br /> |
| `state` _[ContainerState](#containerstate)_ |  | active | Enum: [active paused destroying deleting] <br /> |
| `allowHotUpgrade` _boolean_ |  |  |  |
| `driveCapacity` _integer_ | DriveCapacity specifies the capacity (in GiB) per virtual drive, indicates this container uses shared drives via SSD proxy.<br />When enabled, the container will:<br />- Use virtual UUIDs instead of device paths for drives<br />- Allocate capacity from shared drives rather than exclusive drives<br />- Require an SSD proxy container to be running on the same node<br />This value is copied from the cluster's DriveSharing.DriveCapacity configuration.<br />Used to calculate total capacity request: NumDrives * DriveCapacity |  |  |
| `containerCapacity` _integer_ | ContainerCapacity specifies the total capacity (in GiB) requested by this container when using shared drives via SSD proxy.<br />This value takes precedence over DriveCapacity when both are set. It allows more flexible capacity allocation. |  |  |
| `driveTypesRatio` _[DriveTypesRatio](#drivetypesratio)_ | DriveTypesRatio specifies the desired ratio of drive types (TLC vs QLC) when allocating drives for the container. |  |  |
| `autoRemoveTimeout` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | sets weka cluster-side timeout, if client is not coming back in specified duration it will be auto removed from cluster config | 0s | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(ns\|us\|µs\|ms\|s\|m\|h))+)$` <br />Type: string <br /> |
| `overrides` _[WekaContainerSpecOverrides](#wekacontainerspecoverrides)_ |  |  |  |
| `hostPID` _boolean_ |  |  |  |
| `resources` _[PodResourcesSpec](#podresourcesspec)_ | resources to be proxied as-is to the pod spec |  |  |
| `pvc` _[PVCConfig](#pvcconfig)_ |  |  |  |
| `dpdkBaseMemoryMb` _integer_ |  |  |  |
| `dataServicesConfig` _[DataServicesConfig](#dataservicesconfig)_ |  |  |  |


#### WekaContainerSpecOverrides







_Appears in:_
- [WekaContainerSpec](#wekacontainerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `skipDeactivate` _boolean_ | skips deactivation of container, this is unsafe operation that should be used only when this container will never be back into cluster |  |  |
| `skipDrivesForceResign` _boolean_ | skips resign of drives, if we did not resign drives on removal of drive container we will not be able to reuse them, and manual operation with force resign will be required |  |  |
| `skipVirtualDrivesRemoval` _boolean_ | skips removal of virtual drives from ssdproxy - unsafe operation that can lead to virtual drives leftovers<br />should not be used unless instructed explicitly by weka personnel |  |  |
| `skipCleanupPersistentDir` _boolean_ | skips cleanup of persistent directory, if this operation was omit local data of container will remain in persistent location(/opt/k8s-weka on vanilla OS/k8s distributions) |  |  |
| `upgradeForceReplace` _boolean_ | unsafe operation, skips graceful stop of weka container for a quick replacement to a new image, should not be used unless instructed explicitly by weka personnel |  |  |
| `upgradePreventEviction` _boolean_ |  |  |  |
| `podDeleteForceReplace` _boolean_ |  |  |  |
| `machineIdentifierNodeRef` _string_ |  |  |  |
| `preRunScript` _string_ | script to be executed post initial persistency(if needed) configuration, before running actual workload |  |  |
| `forceDrain` _boolean_ | unsafe operation, forces drain on the node where the container is running, should not be used unless instructed explicitly by weka personnel, the effect of drain is throwing away all IOs and acknowledging all umounts in unsafe manner |  |  |
| `skipActiveMountsCheck` _boolean_ | option to skip active mounts check before deleting client containers |  |  |
| `umountOnHost` _boolean_ | unsafe operation, runs nsenter in root namespace to umount all wekafs mounts visible on host |  |  |
| `debugSleepOnTerminate` _integer_ | DebugSleepOnTerminate specifies the number of seconds to sleep on container abnormal exit for debugging purposes |  |  |
| `migrateOutFromPvc` _boolean_ | MigrateOutFromPvc specifies that the container should be migrated out from PVC into local storage, this will be done prior to starting pod |  |  |


#### WekaContainerStatus







_Appears in:_
- [WekaContainer](#wekacontainer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `status` _[ContainerStatus](#containerstatus)_ |  | Init |  |
| `internalStatus` _string_ |  |  |  |
| `managementIP` _string_ |  |  |  |
| `managementIPs` _string array_ |  |  |  |
| `containerID` _integer_ |  |  |  |
| `clusterID` _string_ |  |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ |  |  |  |
| `lastAppliedImage` _string_ |  |  |  |
| `lastAppliedSpec` _string_ |  |  |  |
| `lastAppliedPodConfigHash` _string_ |  |  |  |
| `nodeAffinity` _[NodeName](#nodename)_ |  |  |  |
| `result` _string_ |  |  |  |
| `allocations` _[ContainerAllocations](#containerallocations)_ |  |  |  |
| `addedDrives` _[Drive](#drive) array_ |  |  |  |
| `stats` _[WekaContainerMetrics](#wekacontainermetrics)_ |  |  |  |
| `printer` _[ContainerPrinterColumns](#containerprintercolumns)_ |  |  |  |
| `timestamps` _object (keys:string, values:[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta))_ |  |  |  |
| `notToleratedOnReschedule` _boolean_ |  |  |  |


#### WekaHomeConfig







_Appears in:_
- [WekaClusterSpec](#wekaclusterspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `endpoint` _string_ |  |  |  |
| `allowInsecure` _boolean_ |  |  |  |
| `cacertSecret` _string_ |  |  |  |
| `enableStats` _boolean_ |  |  |  |


#### WekaManualOperation



WekaManualOperation is the Schema for the wekamanualoperations API





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `WekaManualOperation` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[WekaManualOperationSpec](#wekamanualoperationspec)_ |  |  |  |
| `status` _[WekaManualOperationStatus](#wekamanualoperationstatus)_ |  |  |  |


#### WekaManualOperationSpec



WekaManualOperationSpec defines the desired state of WekaManualOperation



_Appears in:_
- [WekaManualOperation](#wekamanualoperation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ |  |  | Enum: [sign-drives discover-drives force-resign-drives block-drives unblock-drives ensure-nics remote-traces-session] <br /> |
| `payload` _[ManualOperatorPayload](#manualoperatorpayload)_ |  |  |  |
| `image` _string_ |  |  |  |
| `imagePullSecret` _string_ |  |  |  |
| `tolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |
| `deletionDelay` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#duration-v1-meta)_ | DeletionDelay specifies how long to wait after completion before deleting the resource.<br />Defaults to 5m if not specified. | 5m | Pattern: `^(0\|([0-9]+(\.[0-9]+)?(s\|m\|h))+)$` <br />Type: string <br /> |


#### WekaManualOperationStatus



WekaManualOperationStatus defines the observed state of WekaManualOperation



_Appears in:_
- [WekaManualOperation](#wekamanualoperation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `result` _string_ |  |  |  |
| `status` _string_ |  |  |  |
| `completedAt` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ |  |  |  |




#### WekaPolicy



WekaPolicy is the Schema for the wekapolicies API





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `weka.weka.io/v1alpha1` | | |
| `kind` _string_ | `WekaPolicy` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[WekaPolicySpec](#wekapolicyspec)_ |  |  |  |
| `status` _[WekaPolicyStatus](#wekapolicystatus)_ |  |  |  |


#### WekaPolicySpec



WekaPolicySpec defines the desired state of WekaPolicy



_Appears in:_
- [WekaPolicy](#wekapolicy)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `type` _string_ |  |  |  |
| `payload` _[PolicyPayload](#policypayload)_ |  |  |  |
| `image` _string_ |  |  |  |
| `imagePullSecret` _string_ |  |  |  |
| `tolerations` _[Toleration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#toleration-v1-core) array_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |


#### WekaPolicyStatus



WekaPolicyStatus defines the observed state of WekaPolicy



_Appears in:_
- [WekaPolicy](#wekapolicy)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `status` _string_ |  |  |  |
| `result` _string_ |  |  |  |
| `lastRunTime` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ |  |  |  |
| `progress` _string_ |  |  |  |
| `typedStatus` _[TypedPolicyStatus](#typedpolicystatus)_ |  |  |  |


#### WekahomeClientConfig







_Appears in:_
- [WekaClientSpec](#wekaclientspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cacertSecret` _string_ |  |  |  |


