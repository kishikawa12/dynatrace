package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/managementzones/comparison/indexed_name"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// IndexedName Comparison for `INDEXED_NAME` attributes.
type IndexedName struct {
	BaseComparison
	Operator indexed_name.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *string               `json:"value,omitempty"` // The value to compare to.
}

func (inc *IndexedName) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.IndexedName
}

func (inc *IndexedName) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be INDEXED_NAME",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns EQUALS into DOES NOT EQUAL",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Either EQUALS, CONTAINS or EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (inc *IndexedName) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(inc.Unknowns) > 0 {
		data, err := json.Marshal(inc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = inc.Negate
	result["operator"] = inc.Operator
	if inc.Value != nil {
		result["value"] = *inc.Value
	}
	return result, nil
}

func (inc *IndexedName) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), inc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &inc.Unknowns); err != nil {
			return err
		}
		delete(inc.Unknowns, "type")
		delete(inc.Unknowns, "negate")
		delete(inc.Unknowns, "operator")
		delete(inc.Unknowns, "value")
		if len(inc.Unknowns) == 0 {
			inc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		inc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		inc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		inc.Operator = indexed_name.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		inc.Value = opt.NewString(value.(string))
	}
	return nil
}

func (inc *IndexedName) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(inc.Unknowns) > 0 {
		for k, v := range inc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(inc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(inc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&inc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if inc.Value != nil {
		rawMessage, err := json.Marshal(inc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (inc *IndexedName) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	inc.Type = inc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &inc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &inc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &inc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		inc.Unknowns = m
	}
	return nil
}
