package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LDAPProviderConnSecurityProtocolValue represents the possible values for attribute "connSecurityProtocol".
type LDAPProviderConnSecurityProtocolValue string

const (
	// LDAPProviderConnSecurityProtocolInbandTLS represents the value InbandTLS.
	LDAPProviderConnSecurityProtocolInbandTLS LDAPProviderConnSecurityProtocolValue = "InbandTLS"

	// LDAPProviderConnSecurityProtocolTLS represents the value TLS.
	LDAPProviderConnSecurityProtocolTLS LDAPProviderConnSecurityProtocolValue = "TLS"
)

// LDAPProviderIdentity represents the Identity of the object.
var LDAPProviderIdentity = elemental.Identity{
	Name:     "ldapprovider",
	Category: "ldapproviders",
	Package:  "cactuar",
	Private:  false,
}

// LDAPProvidersList represents a list of LDAPProviders
type LDAPProvidersList []*LDAPProvider

// Identity returns the identity of the objects in the list.
func (o LDAPProvidersList) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Copy returns a pointer to a copy the LDAPProvidersList.
func (o LDAPProvidersList) Copy() elemental.Identifiables {

	copy := append(LDAPProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the LDAPProvidersList.
func (o LDAPProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(LDAPProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*LDAPProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o LDAPProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o LDAPProvidersList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the LDAPProvidersList converted to SparseLDAPProvidersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o LDAPProvidersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseLDAPProvidersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseLDAPProvider)
	}

	return out
}

// Version returns the version of the content.
func (o LDAPProvidersList) Version() int {

	return 1
}

// LDAPProvider represents the model of a ldapprovider
type LDAPProvider struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Contains the fully qualified domain name (FQDN) or IP address of the private
	// LDAP server.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Contains the base distinguished name (DN) to use for LDAP queries. Example:
	// `dc=example,dc=com`.
	BaseDN string `json:"baseDN" msgpack:"baseDN" bson:"basedn" mapstructure:"baseDN,omitempty"`

	// Contains the DN to use to bind to the LDAP server. Example:
	// `cn=admin,dc=example,dc=com`.
	BindDN string `json:"bindDN" msgpack:"bindDN" bson:"binddn" mapstructure:"bindDN,omitempty"`

	// Contains the password to be used with the `bindDN` to authenticate to the LDAP
	// server.
	BindPassword string `json:"bindPassword" msgpack:"bindPassword" bson:"bindpassword" mapstructure:"bindPassword,omitempty"`

	// The filter to use to locate the relevant user accounts. For Windows-based
	// systems, the value may
	// be `sAMAccountName={USERNAME}`. For Linux and other systems, the value may be
	// `uid={USERNAME}`.
	BindSearchFilter string `json:"bindSearchFilter" msgpack:"bindSearchFilter" bson:"bindsearchfilter" mapstructure:"bindSearchFilter,omitempty"`

	// Can be left empty if the LDAP server's certificate is signed by a public,
	// trusted certificate
	// authority. Otherwise, include the public key of the certificate authority that
	// signed the
	// LDAP server's certificate.
	CertificateAuthority string `json:"certificateAuthority" msgpack:"certificateAuthority" bson:"certificateauthority" mapstructure:"certificateAuthority,omitempty"`

	// Specifies the connection type for the LDAP provider. `TLS` or `InbandTLS`
	// (default).
	ConnSecurityProtocol LDAPProviderConnSecurityProtocolValue `json:"connSecurityProtocol" msgpack:"connSecurityProtocol" bson:"connsecurityprotocol" mapstructure:"connSecurityProtocol,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// If set, this will be the default LDAP provider. There can be only one default
	// provider in your account. When logging in with LDAP, if no provider name is
	// given, the default will be used.
	Default bool `json:"default" msgpack:"default" bson:"default" mapstructure:"default,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// A list of keys that must not be imported into Aporeto authorization system.
	IgnoredKeys []string `json:"ignoredKeys" msgpack:"ignoredKeys" bson:"ignoredkeys" mapstructure:"ignoredKeys,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// The key to be used to populate the subject of the Midguard token. If you want to
	// use the user as a subject, for Windows-based systems you may use
	// `sAMAccountName`.
	// For Linux and other systems, you may wish to use `uid` (default). You can also
	// use
	// any alternate key.
	SubjectKey string `json:"subjectKey" msgpack:"subjectKey" bson:"subjectkey" mapstructure:"subjectKey,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLDAPProvider returns a new *LDAPProvider
func NewLDAPProvider() *LDAPProvider {

	return &LDAPProvider{
		ModelVersion:         1,
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		BindSearchFilter:     "uid={USERNAME}",
		ConnSecurityProtocol: LDAPProviderConnSecurityProtocolInbandTLS,
		IgnoredKeys:          []string{},
		MigrationsLog:        map[string]string{},
		NormalizedTags:       []string{},
		SubjectKey:           "uid",
	}
}

// Identity returns the Identity of the object.
func (o *LDAPProvider) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LDAPProvider) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LDAPProvider) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *LDAPProvider) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *LDAPProvider) BleveType() string {

	return "ldapprovider"
}

// DefaultOrder returns the list of default ordering fields.
func (o *LDAPProvider) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *LDAPProvider) Doc() string {

	return `Allows you to declare a generic LDAP provider that can be used in exchange
for a Midgard token.`
}

func (o *LDAPProvider) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *LDAPProvider) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *LDAPProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *LDAPProvider) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *LDAPProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *LDAPProvider) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *LDAPProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *LDAPProvider) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *LDAPProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *LDAPProvider) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *LDAPProvider) SetDescription(description string) {

	o.Description = description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *LDAPProvider) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *LDAPProvider) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *LDAPProvider) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *LDAPProvider) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *LDAPProvider) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *LDAPProvider) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *LDAPProvider) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *LDAPProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *LDAPProvider) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *LDAPProvider) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *LDAPProvider) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *LDAPProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *LDAPProvider) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *LDAPProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *LDAPProvider) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *LDAPProvider) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *LDAPProvider) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *LDAPProvider) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *LDAPProvider) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseLDAPProvider{
			ID:                   &o.ID,
			Address:              &o.Address,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			BaseDN:               &o.BaseDN,
			BindDN:               &o.BindDN,
			BindPassword:         &o.BindPassword,
			BindSearchFilter:     &o.BindSearchFilter,
			CertificateAuthority: &o.CertificateAuthority,
			ConnSecurityProtocol: &o.ConnSecurityProtocol,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Default:              &o.Default,
			Description:          &o.Description,
			IgnoredKeys:          &o.IgnoredKeys,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			SubjectKey:           &o.SubjectKey,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseLDAPProvider{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "address":
			sp.Address = &(o.Address)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "baseDN":
			sp.BaseDN = &(o.BaseDN)
		case "bindDN":
			sp.BindDN = &(o.BindDN)
		case "bindPassword":
			sp.BindPassword = &(o.BindPassword)
		case "bindSearchFilter":
			sp.BindSearchFilter = &(o.BindSearchFilter)
		case "certificateAuthority":
			sp.CertificateAuthority = &(o.CertificateAuthority)
		case "connSecurityProtocol":
			sp.ConnSecurityProtocol = &(o.ConnSecurityProtocol)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "default":
			sp.Default = &(o.Default)
		case "description":
			sp.Description = &(o.Description)
		case "ignoredKeys":
			sp.IgnoredKeys = &(o.IgnoredKeys)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subjectKey":
			sp.SubjectKey = &(o.SubjectKey)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseLDAPProvider to the object.
func (o *LDAPProvider) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseLDAPProvider)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Address != nil {
		o.Address = *so.Address
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.BaseDN != nil {
		o.BaseDN = *so.BaseDN
	}
	if so.BindDN != nil {
		o.BindDN = *so.BindDN
	}
	if so.BindPassword != nil {
		o.BindPassword = *so.BindPassword
	}
	if so.BindSearchFilter != nil {
		o.BindSearchFilter = *so.BindSearchFilter
	}
	if so.CertificateAuthority != nil {
		o.CertificateAuthority = *so.CertificateAuthority
	}
	if so.ConnSecurityProtocol != nil {
		o.ConnSecurityProtocol = *so.ConnSecurityProtocol
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Default != nil {
		o.Default = *so.Default
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.IgnoredKeys != nil {
		o.IgnoredKeys = *so.IgnoredKeys
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.SubjectKey != nil {
		o.SubjectKey = *so.SubjectKey
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the LDAPProvider.
func (o *LDAPProvider) DeepCopy() *LDAPProvider {

	if o == nil {
		return nil
	}

	out := &LDAPProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LDAPProvider.
func (o *LDAPProvider) DeepCopyInto(out *LDAPProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LDAPProvider: %s", err))
	}

	*out = *target.(*LDAPProvider)
}

// Validate valides the current information stored into the structure.
func (o *LDAPProvider) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("address", o.Address); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("baseDN", o.BaseDN); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("bindDN", o.BindDN); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("bindPassword", o.BindPassword); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("connSecurityProtocol", string(o.ConnSecurityProtocol), []string{"TLS", "InbandTLS"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
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
func (*LDAPProvider) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LDAPProviderAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LDAPProviderLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LDAPProvider) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LDAPProviderAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LDAPProvider) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "address":
		return o.Address
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "baseDN":
		return o.BaseDN
	case "bindDN":
		return o.BindDN
	case "bindPassword":
		return o.BindPassword
	case "bindSearchFilter":
		return o.BindSearchFilter
	case "certificateAuthority":
		return o.CertificateAuthority
	case "connSecurityProtocol":
		return o.ConnSecurityProtocol
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "default":
		return o.Default
	case "description":
		return o.Description
	case "ignoredKeys":
		return o.IgnoredKeys
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "subjectKey":
		return o.SubjectKey
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// LDAPProviderAttributesMap represents the map of attribute for LDAPProvider.
var LDAPProviderAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Address": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Address",
		Description: `Contains the fully qualified domain name (FQDN) or IP address of the private
LDAP server.`,
		Exposed:    true,
		Filterable: true,
		Name:       "address",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"BaseDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BaseDN",
		Description: `Contains the base distinguished name (DN) to use for LDAP queries. Example:
` + "`" + `dc=example,dc=com` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "baseDN",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"BindDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindDN",
		Description: `Contains the DN to use to bind to the LDAP server. Example:
` + "`" + `cn=admin,dc=example,dc=com` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "bindDN",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"BindPassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindPassword",
		Description: `Contains the password to be used with the ` + "`" + `bindDN` + "`" + ` to authenticate to the LDAP
server.`,
		Exposed:   true,
		Name:      "bindPassword",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
	},
	"BindSearchFilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `The filter to use to locate the relevant user accounts. For Windows-based
systems, the value may
be ` + "`" + `sAMAccountName={USERNAME}` + "`" + `. For Linux and other systems, the value may be
` + "`" + `uid={USERNAME}` + "`" + `.`,
		Exposed:   true,
		Name:      "bindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"CertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description: `Can be left empty if the LDAP server's certificate is signed by a public,
trusted certificate
authority. Otherwise, include the public key of the certificate authority that
signed the
LDAP server's certificate.`,
		Exposed:   true,
		Name:      "certificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"ConnSecurityProtocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "ConnSecurityProtocol",
		DefaultValue:   LDAPProviderConnSecurityProtocolInbandTLS,
		Description: `Specifies the connection type for the LDAP provider. ` + "`" + `TLS` + "`" + ` or ` + "`" + `InbandTLS` + "`" + `
(default).`,
		Exposed:    true,
		Filterable: true,
		Name:       "connSecurityProtocol",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default LDAP provider. There can be only one default
provider in your account. When logging in with LDAP, if no provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"IgnoredKeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoredKeys",
		Description:    `A list of keys that must not be imported into Aporeto authorization system.`,
		Exposed:        true,
		Name:           "ignoredKeys",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"SubjectKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectKey",
		DefaultValue:   "uid",
		Description: `The key to be used to populate the subject of the Midguard token. If you want to
use the user as a subject, for Windows-based systems you may use
` + "`" + `sAMAccountName` + "`" + `.
For Linux and other systems, you may wish to use ` + "`" + `uid` + "`" + ` (default). You can also
use
any alternate key.`,
		Exposed:   true,
		Name:      "subjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// LDAPProviderLowerCaseAttributesMap represents the map of attribute for LDAPProvider.
var LDAPProviderLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"address": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Address",
		Description: `Contains the fully qualified domain name (FQDN) or IP address of the private
LDAP server.`,
		Exposed:    true,
		Filterable: true,
		Name:       "address",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"basedn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BaseDN",
		Description: `Contains the base distinguished name (DN) to use for LDAP queries. Example:
` + "`" + `dc=example,dc=com` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "baseDN",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"binddn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindDN",
		Description: `Contains the DN to use to bind to the LDAP server. Example:
` + "`" + `cn=admin,dc=example,dc=com` + "`" + `.`,
		Exposed:    true,
		Filterable: true,
		Name:       "bindDN",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "string",
	},
	"bindpassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindPassword",
		Description: `Contains the password to be used with the ` + "`" + `bindDN` + "`" + ` to authenticate to the LDAP
server.`,
		Exposed:   true,
		Name:      "bindPassword",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
	},
	"bindsearchfilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "BindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `The filter to use to locate the relevant user accounts. For Windows-based
systems, the value may
be ` + "`" + `sAMAccountName={USERNAME}` + "`" + `. For Linux and other systems, the value may be
` + "`" + `uid={USERNAME}` + "`" + `.`,
		Exposed:   true,
		Name:      "bindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"certificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateAuthority",
		Description: `Can be left empty if the LDAP server's certificate is signed by a public,
trusted certificate
authority. Otherwise, include the public key of the certificate authority that
signed the
LDAP server's certificate.`,
		Exposed:   true,
		Name:      "certificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"connsecurityprotocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "ConnSecurityProtocol",
		DefaultValue:   LDAPProviderConnSecurityProtocolInbandTLS,
		Description: `Specifies the connection type for the LDAP provider. ` + "`" + `TLS` + "`" + ` or ` + "`" + `InbandTLS` + "`" + `
(default).`,
		Exposed:    true,
		Filterable: true,
		Name:       "connSecurityProtocol",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"default": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Default",
		Description: `If set, this will be the default LDAP provider. There can be only one default
provider in your account. When logging in with LDAP, if no provider name is
given, the default will be used.`,
		Exposed: true,
		Name:    "default",
		Stored:  true,
		Type:    "boolean",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ignoredkeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoredKeys",
		Description:    `A list of keys that must not be imported into Aporeto authorization system.`,
		Exposed:        true,
		Name:           "ignoredKeys",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subjectkey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SubjectKey",
		DefaultValue:   "uid",
		Description: `The key to be used to populate the subject of the Midguard token. If you want to
use the user as a subject, for Windows-based systems you may use
` + "`" + `sAMAccountName` + "`" + `.
For Linux and other systems, you may wish to use ` + "`" + `uid` + "`" + ` (default). You can also
use
any alternate key.`,
		Exposed:   true,
		Name:      "subjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseLDAPProvidersList represents a list of SparseLDAPProviders
type SparseLDAPProvidersList []*SparseLDAPProvider

// Identity returns the identity of the objects in the list.
func (o SparseLDAPProvidersList) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Copy returns a pointer to a copy the SparseLDAPProvidersList.
func (o SparseLDAPProvidersList) Copy() elemental.Identifiables {

	copy := append(SparseLDAPProvidersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseLDAPProvidersList.
func (o SparseLDAPProvidersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseLDAPProvidersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseLDAPProvider))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseLDAPProvidersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseLDAPProvidersList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseLDAPProvidersList converted to LDAPProvidersList.
func (o SparseLDAPProvidersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseLDAPProvidersList) Version() int {

	return 1
}

// SparseLDAPProvider represents the sparse version of a ldapprovider.
type SparseLDAPProvider struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Contains the fully qualified domain name (FQDN) or IP address of the private
	// LDAP server.
	Address *string `json:"address,omitempty" msgpack:"address,omitempty" bson:"address,omitempty" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Contains the base distinguished name (DN) to use for LDAP queries. Example:
	// `dc=example,dc=com`.
	BaseDN *string `json:"baseDN,omitempty" msgpack:"baseDN,omitempty" bson:"basedn,omitempty" mapstructure:"baseDN,omitempty"`

	// Contains the DN to use to bind to the LDAP server. Example:
	// `cn=admin,dc=example,dc=com`.
	BindDN *string `json:"bindDN,omitempty" msgpack:"bindDN,omitempty" bson:"binddn,omitempty" mapstructure:"bindDN,omitempty"`

	// Contains the password to be used with the `bindDN` to authenticate to the LDAP
	// server.
	BindPassword *string `json:"bindPassword,omitempty" msgpack:"bindPassword,omitempty" bson:"bindpassword,omitempty" mapstructure:"bindPassword,omitempty"`

	// The filter to use to locate the relevant user accounts. For Windows-based
	// systems, the value may
	// be `sAMAccountName={USERNAME}`. For Linux and other systems, the value may be
	// `uid={USERNAME}`.
	BindSearchFilter *string `json:"bindSearchFilter,omitempty" msgpack:"bindSearchFilter,omitempty" bson:"bindsearchfilter,omitempty" mapstructure:"bindSearchFilter,omitempty"`

	// Can be left empty if the LDAP server's certificate is signed by a public,
	// trusted certificate
	// authority. Otherwise, include the public key of the certificate authority that
	// signed the
	// LDAP server's certificate.
	CertificateAuthority *string `json:"certificateAuthority,omitempty" msgpack:"certificateAuthority,omitempty" bson:"certificateauthority,omitempty" mapstructure:"certificateAuthority,omitempty"`

	// Specifies the connection type for the LDAP provider. `TLS` or `InbandTLS`
	// (default).
	ConnSecurityProtocol *LDAPProviderConnSecurityProtocolValue `json:"connSecurityProtocol,omitempty" msgpack:"connSecurityProtocol,omitempty" bson:"connsecurityprotocol,omitempty" mapstructure:"connSecurityProtocol,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// If set, this will be the default LDAP provider. There can be only one default
	// provider in your account. When logging in with LDAP, if no provider name is
	// given, the default will be used.
	Default *bool `json:"default,omitempty" msgpack:"default,omitempty" bson:"default,omitempty" mapstructure:"default,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// A list of keys that must not be imported into Aporeto authorization system.
	IgnoredKeys *[]string `json:"ignoredKeys,omitempty" msgpack:"ignoredKeys,omitempty" bson:"ignoredkeys,omitempty" mapstructure:"ignoredKeys,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// The key to be used to populate the subject of the Midguard token. If you want to
	// use the user as a subject, for Windows-based systems you may use
	// `sAMAccountName`.
	// For Linux and other systems, you may wish to use `uid` (default). You can also
	// use
	// any alternate key.
	SubjectKey *string `json:"subjectKey,omitempty" msgpack:"subjectKey,omitempty" bson:"subjectkey,omitempty" mapstructure:"subjectKey,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseLDAPProvider returns a new  SparseLDAPProvider.
func NewSparseLDAPProvider() *SparseLDAPProvider {
	return &SparseLDAPProvider{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseLDAPProvider) Identity() elemental.Identity {

	return LDAPProviderIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseLDAPProvider) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseLDAPProvider) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseLDAPProvider) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseLDAPProvider) ToPlain() elemental.PlainIdentifiable {

	out := NewLDAPProvider()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Address != nil {
		out.Address = *o.Address
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.BaseDN != nil {
		out.BaseDN = *o.BaseDN
	}
	if o.BindDN != nil {
		out.BindDN = *o.BindDN
	}
	if o.BindPassword != nil {
		out.BindPassword = *o.BindPassword
	}
	if o.BindSearchFilter != nil {
		out.BindSearchFilter = *o.BindSearchFilter
	}
	if o.CertificateAuthority != nil {
		out.CertificateAuthority = *o.CertificateAuthority
	}
	if o.ConnSecurityProtocol != nil {
		out.ConnSecurityProtocol = *o.ConnSecurityProtocol
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Default != nil {
		out.Default = *o.Default
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.IgnoredKeys != nil {
		out.IgnoredKeys = *o.IgnoredKeys
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.SubjectKey != nil {
		out.SubjectKey = *o.SubjectKey
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseLDAPProvider) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseLDAPProvider) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseLDAPProvider) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseLDAPProvider) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseLDAPProvider) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetDescription(description string) {

	o.Description = &description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseLDAPProvider) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseLDAPProvider) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseLDAPProvider) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseLDAPProvider) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseLDAPProvider) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseLDAPProvider) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseLDAPProvider) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseLDAPProvider) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseLDAPProvider) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseLDAPProvider) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseLDAPProvider.
func (o *SparseLDAPProvider) DeepCopy() *SparseLDAPProvider {

	if o == nil {
		return nil
	}

	out := &SparseLDAPProvider{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseLDAPProvider.
func (o *SparseLDAPProvider) DeepCopyInto(out *SparseLDAPProvider) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseLDAPProvider: %s", err))
	}

	*out = *target.(*SparseLDAPProvider)
}