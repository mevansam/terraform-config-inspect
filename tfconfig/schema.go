package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
)

var rootSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "terraform",
			LabelNames: nil,
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

var terraformBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "required_version",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "required_providers",
		},
		{
			Type: "backend",
			LabelNames: []string{"s3"},
		},
		{
			Type: "backend",
			LabelNames: []string{"azurerm"},
		},
		{
			Type: "backend",
			LabelNames: []string{"gcs"},
		},
		{
			Type: "backend",
			LabelNames: []string{"oss"},
		},
		{
			Type: "backend",
			LabelNames: []string{"cos"},
		},
		{
			Type: "backend",
			LabelNames: []string{"pg"},
		},
		{
			Type: "backend",
			LabelNames: []string{"consul"},
		},
		{
			Type: "backend",
			LabelNames: []string{"kubernetes"},
		},
		{
			Type: "backend",
			LabelNames: []string{"http"},
		},
	},
}

var providerConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "version",
		},
		{
			Name: "alias",
		},
	},
}

var variableSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "type",
		},
		{
			Name: "description",
		},
		{
			Name: "default",
		},
		{
			Name: "sensitive",
		},
	},
}

var outputSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "sensitive",
		},
	},
}

var moduleCallSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "source",
		},
		{
			Name: "version",
		},
		{
			Name: "providers",
		},
	},
}

var resourceSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "provider",
		},
	},
}
