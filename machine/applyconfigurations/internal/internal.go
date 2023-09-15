// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.openshift.api.machine.v1.AWSFailureDomain
  map:
    fields:
    - name: placement
      type:
        namedType: com.github.openshift.api.machine.v1.AWSFailureDomainPlacement
      default: {}
    - name: subnet
      type:
        namedType: com.github.openshift.api.machine.v1.AWSResourceReference
- name: com.github.openshift.api.machine.v1.AWSFailureDomainPlacement
  map:
    fields:
    - name: availabilityZone
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1.AWSResourceFilter
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.openshift.api.machine.v1.AWSResourceReference
  map:
    fields:
    - name: arn
      type:
        scalar: string
    - name: filters
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.AWSResourceFilter
          elementRelationship: atomic
    - name: id
      type:
        scalar: string
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: arn
        discriminatorValue: ARN
      - fieldName: filters
        discriminatorValue: Filters
      - fieldName: id
        discriminatorValue: ID
- name: com.github.openshift.api.machine.v1.AzureFailureDomain
  map:
    fields:
    - name: subnet
      type:
        scalar: string
    - name: zone
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1.ControlPlaneMachineSetSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.machine.v1.ControlPlaneMachineSetStatus
      default: {}
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSetSpec
  map:
    fields:
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
      default: {}
    - name: state
      type:
        scalar: string
      default: Inactive
    - name: strategy
      type:
        namedType: com.github.openshift.api.machine.v1.ControlPlaneMachineSetStrategy
      default: {}
    - name: template
      type:
        namedType: com.github.openshift.api.machine.v1.ControlPlaneMachineSetTemplate
      default: {}
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSetStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: unavailableReplicas
      type:
        scalar: numeric
    - name: updatedReplicas
      type:
        scalar: numeric
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSetStrategy
  map:
    fields:
    - name: type
      type:
        scalar: string
      default: RollingUpdate
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSetTemplate
  map:
    fields:
    - name: machineType
      type:
        scalar: string
    - name: machines_v1beta1_machine_openshift_io
      type:
        namedType: com.github.openshift.api.machine.v1.OpenShiftMachineV1Beta1MachineTemplate
    unions:
    - discriminator: machineType
      fields:
      - fieldName: machines_v1beta1_machine_openshift_io
        discriminatorValue: OpenShiftMachineV1Beta1Machine
- name: com.github.openshift.api.machine.v1.ControlPlaneMachineSetTemplateObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: labels
      type:
        map:
          elementType:
            scalar: string
- name: com.github.openshift.api.machine.v1.FailureDomains
  map:
    fields:
    - name: aws
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.AWSFailureDomain
          elementRelationship: atomic
    - name: azure
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.AzureFailureDomain
          elementRelationship: atomic
    - name: gcp
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.GCPFailureDomain
          elementRelationship: atomic
    - name: openstack
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.OpenStackFailureDomain
          elementRelationship: atomic
    - name: platform
      type:
        scalar: string
      default: ""
    - name: vsphere
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1.VSphereFailureDomain
          elementRelationship: atomic
    unions:
    - discriminator: platform
      fields:
      - fieldName: aws
        discriminatorValue: AWS
      - fieldName: azure
        discriminatorValue: Azure
      - fieldName: gcp
        discriminatorValue: GCP
      - fieldName: openstack
        discriminatorValue: OpenStack
      - fieldName: vsphere
        discriminatorValue: VSphere
- name: com.github.openshift.api.machine.v1.GCPFailureDomain
  map:
    fields:
    - name: zone
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1.OpenShiftMachineV1Beta1MachineTemplate
  map:
    fields:
    - name: failureDomains
      type:
        namedType: com.github.openshift.api.machine.v1.FailureDomains
    - name: metadata
      type:
        namedType: com.github.openshift.api.machine.v1.ControlPlaneMachineSetTemplateObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineSpec
      default: {}
- name: com.github.openshift.api.machine.v1.OpenStackFailureDomain
  map:
    fields:
    - name: availabilityZone
      type:
        scalar: string
    - name: rootVolume
      type:
        namedType: com.github.openshift.api.machine.v1.RootVolume
- name: com.github.openshift.api.machine.v1.RootVolume
  map:
    fields:
    - name: availabilityZone
      type:
        scalar: string
    - name: volumeType
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1.VSphereFailureDomain
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1beta1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: severity
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1beta1.LastOperation
  map:
    fields:
    - name: description
      type:
        scalar: string
    - name: lastUpdated
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: state
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: com.github.openshift.api.machine.v1beta1.LifecycleHook
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: owner
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.machine.v1beta1.LifecycleHooks
  map:
    fields:
    - name: preDrain
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1beta1.LifecycleHook
          elementRelationship: associative
          keys:
          - name
    - name: preTerminate
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1beta1.LifecycleHook
          elementRelationship: associative
          keys:
          - name
- name: com.github.openshift.api.machine.v1beta1.Machine
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineStatus
      default: {}
- name: com.github.openshift.api.machine.v1beta1.MachineHealthCheck
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineHealthCheckSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineHealthCheckStatus
      default: {}
- name: com.github.openshift.api.machine.v1beta1.MachineHealthCheckSpec
  map:
    fields:
    - name: maxUnhealthy
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: nodeStartupTimeout
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
    - name: remediationTemplate
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
      default: {}
    - name: unhealthyConditions
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1beta1.UnhealthyCondition
          elementRelationship: atomic
- name: com.github.openshift.api.machine.v1beta1.MachineHealthCheckStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1beta1.Condition
          elementRelationship: atomic
    - name: currentHealthy
      type:
        scalar: numeric
    - name: expectedMachines
      type:
        scalar: numeric
    - name: remediationsAllowed
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.machine.v1beta1.MachineSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineSetSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineSetStatus
      default: {}
- name: com.github.openshift.api.machine.v1beta1.MachineSetSpec
  map:
    fields:
    - name: deletePolicy
      type:
        scalar: string
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
      default: {}
    - name: template
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineTemplateSpec
      default: {}
- name: com.github.openshift.api.machine.v1beta1.MachineSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: errorMessage
      type:
        scalar: string
    - name: errorReason
      type:
        scalar: string
    - name: fullyLabeledReplicas
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
- name: com.github.openshift.api.machine.v1beta1.MachineSpec
  map:
    fields:
    - name: lifecycleHooks
      type:
        namedType: com.github.openshift.api.machine.v1beta1.LifecycleHooks
      default: {}
    - name: metadata
      type:
        namedType: com.github.openshift.api.machine.v1beta1.ObjectMeta
      default: {}
    - name: providerID
      type:
        scalar: string
    - name: providerSpec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.ProviderSpec
      default: {}
    - name: taints
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Taint
          elementRelationship: atomic
- name: com.github.openshift.api.machine.v1beta1.MachineStatus
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeAddress
          elementRelationship: atomic
    - name: conditions
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.machine.v1beta1.Condition
          elementRelationship: atomic
    - name: errorMessage
      type:
        scalar: string
    - name: errorReason
      type:
        scalar: string
    - name: lastOperation
      type:
        namedType: com.github.openshift.api.machine.v1beta1.LastOperation
    - name: lastUpdated
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: nodeRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: phase
      type:
        scalar: string
    - name: providerStatus
      type:
        namedType: __untyped_atomic_
- name: com.github.openshift.api.machine.v1beta1.MachineTemplateSpec
  map:
    fields:
    - name: metadata
      type:
        namedType: com.github.openshift.api.machine.v1beta1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.machine.v1beta1.MachineSpec
      default: {}
- name: com.github.openshift.api.machine.v1beta1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: generateName
      type:
        scalar: string
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
- name: com.github.openshift.api.machine.v1beta1.ProviderSpec
  map:
    fields:
    - name: value
      type:
        namedType: __untyped_atomic_
- name: com.github.openshift.api.machine.v1beta1.UnhealthyCondition
  map:
    fields:
    - name: status
      type:
        scalar: string
      default: ""
    - name: timeout
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
      default: 0
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.NodeAddress
  map:
    fields:
    - name: address
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldPath
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resourceVersion
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.core.v1.Taint
  map:
    fields:
    - name: effect
      type:
        scalar: string
      default: ""
    - name: key
      type:
        scalar: string
      default: ""
    - name: timeAdded
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: value
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
      default: ""
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
  scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: io.k8s.apimachinery.pkg.runtime.RawExtension
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.util.intstr.IntOrString
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
