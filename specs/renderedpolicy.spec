# Model
model:
  rest_name: renderedpolicy
  resource_name: renderedpolicies
  entity_name: RenderedPolicy
  package: squall
  group: core/policy
  description: Retrieve the aggregated policies applied to a particular processing
    unit.
  aliases:
  - rpol
  - rpols

# Attributes
attributes:
  v1:
  - name: certificate
    description: |-
      The certificate associated with this processing unit. It will identify the
      processing unit to any internal or external services.
    type: string
    exposed: true
    read_only: true
    deprecated: true
    omit_empty: true

  - name: datapathType
    description: |-
      The datapath type that this processing unit must implement according to
      the rendered policy:
      - `Default`: This policy is not making a decision for the datapath.
      - `Aporeto`: The enforcer is managing and handling the datapath.
      - `EnvoyAuthorizer`: The enforcer is serving Envoy-compatible gRPC APIs
      that for example can be used by an Envoy proxy to use the Microsegmentation PKI
      and implement Microsegmentation network policies. NOTE: The enforcer is not
      owning the datapath in this case. It is merely providing an authorizer API.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Default
    - Aporeto
    - EnvoyAuthorizer
    autogenerated: true
    deprecated: true
    omit_empty: true

  - name: defaultPUIncomingTrafficAction
    description: Describes the default for incoming traffic.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Reject

  - name: defaultPUOutgoingTrafficAction
    description: Describes the default for outgoing traffic.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Reject

  - name: dependendServices
    description: The list of services that this processing unit depends on.
    type: refList
    exposed: true
    subtype: service
    extensions:
      noInit: true

  - name: egressPolicies
    description: Lists all the egress policies attached to processing unit.
    type: external
    exposed: true
    subtype: _rendered_policy
    read_only: true
    autogenerated: true
    deprecated: true
    omit_empty: true

  - name: exposedServices
    description: The list of services that this processing unit is implementing.
    type: refList
    exposed: true
    subtype: service
    extensions:
      noInit: true

  - name: hashedTags
    description: Contains the list of tags that matched the policies and their hashes.
    type: external
    exposed: true
    subtype: map[string]string
    read_only: true
    autogenerated: true

  - name: ingressPolicies
    description: Lists all the ingress policies attached to the processing unit.
    type: external
    exposed: true
    subtype: _rendered_policy
    read_only: true
    autogenerated: true
    deprecated: true
    omit_empty: true

  - name: matchingTags
    description: Contains the list of tags that matched the policies.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true
    deprecated: true
    omit_empty: true

  - name: processingUnit
    description: |-
      Can be set during a `POST` operation to render a policy on a processing unit
      that has not been created yet.
    type: ref
    exposed: true
    subtype: processingunit
    creation_only: true
    deprecated: true
    example_value: |-
      {
        "name": "pu",
        "type": "Docker",
        "normalizedTags": [
          "a=a",
          "b=b"
        ]
      }
    omit_empty: true
    extensions:
      noInit: true
      refMode: pointer

  - name: processingUnitID
    description: Identifier of the processing unit.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    deprecated: true
    omit_empty: true

  - name: processingUnitTags
    description: |-
      Can be set during a `POST` operation to render a policy on a processing unit
      tags.
    type: list
    exposed: true
    subtype: string
    creation_only: true
    omit_empty: true

  - name: ruleSetPolicies
    description: Lists all the rule set policies attached to processing unit.
    type: refList
    exposed: true
    subtype: policyrule
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: scopes
    description: |-
      The set of scopes granted to this processing unit that has to be
      present in HTTP requests.
    type: list
    exposed: true
    subtype: string
    stored: true
