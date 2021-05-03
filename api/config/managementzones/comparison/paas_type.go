package comparison

import (
	"encoding/json"

	paastype "github.com/dtcookie/dynatrace/api/config/managementzones/comparison/paas_type"
	"github.com/dtcookie/hcl"
)

// PaasType Comparison for `PAAS_TYPE` attributes.
type PaasType struct {
	BaseComparison
	Operator paastype.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *paastype.Value   `json:"value,omitempty"` // The value to compare to.
}

func (ptc *PaasType) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.PaasType
}

func (ptc *PaasType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be PAAS_TYPE",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns the **begins with** into **does not begin with**",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison. Possible values are EQUALS and EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to. Possible values are AWS_ECS_EC2, AWS_ECS_FARGATE, AWS_LAMBDA, AZURE_FUNCTIONS, AZURE_WEBSITES, CLOUD_FOUNDRY, GOOGLE_APP_ENGINE, HEROKU, KUBERNETES and OPENSHIFT.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (ptc *PaasType) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ptc.Unknowns) > 0 {
		data, err := json.Marshal(ptc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = ptc.Negate
	result["operator"] = ptc.Operator
	if ptc.Value != nil {
		result["value"] = ptc.Value.String()
	}
	return result, nil
}

func (ptc *PaasType) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ptc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ptc.Unknowns); err != nil {
			return err
		}
		delete(ptc.Unknowns, "type")
		delete(ptc.Unknowns, "negate")
		delete(ptc.Unknowns, "operator")
		delete(ptc.Unknowns, "value")
		if len(ptc.Unknowns) == 0 {
			ptc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		ptc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		ptc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		ptc.Operator = paastype.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		ptc.Value = paastype.Value(value.(string)).Ref()
	}
	return nil
}

func (ptc *PaasType) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ptc.Unknowns) > 0 {
		for k, v := range ptc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(ptc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ComparisonBasicTypes.PaasType)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&ptc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if ptc.Value != nil {
		rawMessage, err := json.Marshal(ptc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (ptc *PaasType) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	ptc.Type = ptc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &ptc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &ptc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &ptc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		ptc.Unknowns = m
	}
	return nil
}
