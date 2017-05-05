package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// AWSAccountIdentity represents the Identity of the object
var AWSAccountIdentity = elemental.Identity{
	Name:     "awsaccount",
	Category: "awsaccounts",
}

// AWSAccountsList represents a list of AWSAccounts
type AWSAccountsList []*AWSAccount

// ContentIdentity returns the identity of the objects in the list.
func (o AWSAccountsList) ContentIdentity() elemental.Identity {

	return AWSAccountIdentity
}

// List converts the object to and elemental.IdentifiablesList.
func (o AWSAccountsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAccountsList) DefaultOrder() []string {

	return []string{}
}

// AWSAccount represents the model of a awsaccount
type AWSAccount struct {
	// ID of the object.
	ID string `json:"ID" bson:"_id"`

	// AccessKeyID contains the aws access key ID. This is used to retrieve your account id, and it is not stored.
	AccessKeyID string `json:"accessKeyID" bson:"-"`

	// AccessToken contains your aws access token. It is used to retrieve your account id, and it not stored.
	AccessToken string `json:"accessToken" bson:"-"`

	// accountID contains your verified accound id.
	AccountID string `json:"accountID" bson:"accountid"`

	// createdAt represents the creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// ParentID contains the parent Vince account ID.
	ParentID string `json:"parentID" bson:"parentid"`

	// ParentName contains the name of the Vince parent Account.
	ParentName string `json:"parentName" bson:"parentname"`

	// Region contains your the region where your AWS account is located
	Region string `json:"region" bson:"region"`

	// secretAccessKey contains the secret key. It is used to retrieve your account id, and it is not stored.
	SecretAccessKey string `json:"secretAccessKey" bson:"-"`

	// UpdateTime represents the last update date of the objct.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAWSAccount returns a new *AWSAccount
func NewAWSAccount() *AWSAccount {

	return &AWSAccount{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *AWSAccount) Identity() elemental.Identity {

	return AWSAccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AWSAccount) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AWSAccount) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *AWSAccount) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *AWSAccount) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AWSAccount) Doc() string {
	return `Allows to bind an AWS account to your Aporeto account to allow auto registration of enforcers running on EC2 `
}

func (o *AWSAccount) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *AWSAccount) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("accessKeyID", o.AccessKeyID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("accessKeyID", o.AccessKeyID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("region", o.Region); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("region", o.Region); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("secretAccessKey", o.SecretAccessKey); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("secretAccessKey", o.SecretAccessKey); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*AWSAccount) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return AWSAccountAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSAccount) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSAccountAttributesMap
}

// AWSAccountAttributesMap represents the map of attribute for AWSAccount.
var AWSAccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AccessKeyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `AccessKeyID contains the aws access key ID. This is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessKeyID",
		Required:       true,
		Type:           "string",
	},
	"AccessToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `AccessToken contains your aws access token. It is used to retrieve your account id, and it not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "accessToken",
		Type:           "string",
	},
	"AccountID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `accountID contains your verified accound id.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "accountID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `createdAt represents the creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID contains the parent Vince account ID.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentName contains the name of the Vince parent Account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentName",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Region": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Region contains your the region where your AWS account is located`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "region",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SecretAccessKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `secretAccessKey contains the secret key. It is used to retrieve your account id, and it is not stored.`,
		Exposed:        true,
		Format:         "free",
		Name:           "secretAccessKey",
		Required:       true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}