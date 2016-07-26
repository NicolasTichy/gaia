package gaia

import "github.com/aporeto-inc/elemental"

const (
	MapNodeAttributeNameID     elemental.AttributeSpecificationNameKey = "mapnode/ID"
	MapNodeAttributeNameLabels elemental.AttributeSpecificationNameKey = "mapnode/labels"
	MapNodeAttributeNameName   elemental.AttributeSpecificationNameKey = "mapnode/name"
	MapNodeAttributeNameType   elemental.AttributeSpecificationNameKey = "mapnode/type"
)

// MapNodeTypeValue represents the possible values for attribute "type".
type MapNodeTypeValue string

const (
	MapNodeTypeContainer MapNodeTypeValue = "Container"
	MapNodeTypeVolume    MapNodeTypeValue = "Volume"
)

// MapNodeIdentity represents the Identity of the object
var MapNodeIdentity = elemental.Identity{
	Name:     "mapnode",
	Category: "mapnodes",
}

// MapNodesList represents a list of MapNodes
type MapNodesList []*MapNode

// MapNode represents the model of a mapnode
type MapNode struct {
	ID     string           `json:"ID,omitempty" cql:"-"`
	Labels []string         `json:"labels,omitempty" cql:"-"`
	Name   string           `json:"name,omitempty" cql:"-"`
	Type   MapNodeTypeValue `json:"type,omitempty" cql:"-"`
}

// NewMapNode returns a new *MapNode
func NewMapNode() *MapNode {

	return &MapNode{}
}

// Identity returns the Identity of the object.
func (o *MapNode) Identity() elemental.Identity {

	return MapNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapNode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapNode) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *MapNode) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Container", "Volume"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o MapNode) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return MapNodeAttributesMap[name]
}

var MapNodeAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	MapNodeAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	MapNodeAttributeNameLabels: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "labels",
		Required:       true,
		SubType:        "tag_list",
		Type:           "external",
	},
	MapNodeAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	MapNodeAttributeNameType: elemental.AttributeSpecification{
		AllowedChoices: []string{"Container", "Volume"},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}
