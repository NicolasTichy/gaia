{
    "attributes": [
        {
            "allowed_chars": "^[^\\*\\=]*$",
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Name is the name of the namespace.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": true,
            "identifier": null,
            "index": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "ns"
        ],
        "create": false,
        "delete": true,
        "description": "A Namespace represents the core organizational unit of the system. All objects always exists in a single namespace. A Namespace can also have child namespaces. They can be used to split the system into organizations, business units, applications, services or any combination you like.",
        "entity_name": "Namespace",
        "extends": [
            "@base",
            "@described",
            "@identifiable-nopk-stored",
            "@metadatable"
        ],
        "get": true,
        "package": "Authorization",
        "resource_name": "namespaces",
        "rest_name": "namespace",
        "root": null,
        "update": true
    }
}