# Model
model:
  rest_name: remoteprocessor
  resource_name: remoteprocessors
  entity_name: RemoteProcessor
  package: rufus
  group: policy/hooks
  description: Hook to integrate a Segment service.
  aliases:
  - hks
  - hk

# Attributes
attributes:
  v1:
  - name: claims
    description: Represents the claims of the currently managed object.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - '@auth:realm=certificate'
    - '@auth:commonname=john'

  - name: input
    description: Represents data received from the service.
    type: external
    exposed: true
    subtype: json.RawMessage
    required: true
    example_value: |-
      {
        "name": "hello",
        "description": "hello",
      }

  - name: mode
    description: Defines the hook's type.
    type: enum
    exposed: true
    allowed_choices:
    - Post
    - Pre
    example_value: Pre

  - name: namespace
    description: Represents the current namespace.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: operation
    description: Defines the operation that is currently handled by the service.
    type: external
    exposed: true
    subtype: elemental.Operation
    required: true
    example_value: create

  - name: output
    description: Returns `OutputData` filled with the processor information.
    type: external
    exposed: true
    subtype: _elemental_identifiable
    read_only: true
    autogenerated: true

  - name: requestID
    description: Gives the ID of the request coming from the main server.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: targetIdentity
    description: Represents the identity name of the managed object.
    type: string
    exposed: true
    required: true
    example_value: processingunit
