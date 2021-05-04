package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/synthetic_engine_type"
	"github.com/dtcookie/hcl"
)

// SyntheticEngineType Comparison for `SYNTHETIC_ENGINE_TYPE` attributes.
type SyntheticEngineType struct {
	BaseComparison
	Operator synthetic_engine_type.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *synthetic_engine_type.Value   `json:"value,omitempty"` // The value to compare to.
}

func (setc *SyntheticEngineType) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.SyntheticEngineType
}

func (setc *SyntheticEngineType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be SYNTHETIC_ENGINE_TYPE",
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
			Description: "Operator of the comparison. Possible values are  EQUALS and EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to. Possible values are CLASSIC and CUSTOM",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (setc *SyntheticEngineType) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(setc.Unknowns) > 0 {
		data, err := json.Marshal(setc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = setc.Negate
	result["operator"] = setc.Operator
	if setc.Value != nil {
		result["value"] = setc.Value.String()
	}
	return result, nil
}

func (setc *SyntheticEngineType) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), setc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &setc.Unknowns); err != nil {
			return err
		}
		delete(setc.Unknowns, "type")
		delete(setc.Unknowns, "negate")
		delete(setc.Unknowns, "operator")
		delete(setc.Unknowns, "value")
		if len(setc.Unknowns) == 0 {
			setc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		setc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		setc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		setc.Operator = synthetic_engine_type.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		setc.Value = synthetic_engine_type.Value(value.(string)).Ref()
	}
	return nil
}

func (setc *SyntheticEngineType) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(setc.Unknowns) > 0 {
		for k, v := range setc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(setc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(setc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&setc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if setc.Value != nil {
		rawMessage, err := json.Marshal(setc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (setc *SyntheticEngineType) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	setc.Type = setc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &setc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &setc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &setc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		setc.Unknowns = m
	}
	return nil
}
