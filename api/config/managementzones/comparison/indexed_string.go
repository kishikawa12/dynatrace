package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/managementzones/comparison/indexed_string"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// IndexedString Comparison for `INDEXED_STRING` attributes.
type IndexedString struct {
	BaseComparison
	Operator indexed_string.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *string                 `json:"value,omitempty"` // The value to compare to.
}

func (isc *IndexedString) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.IndexedString
}

func (isc *IndexedString) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be INDEXED_STRING",
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
			Description: "Either EQUALS or EXISTS. You can reverse it by setting **negate** to `true`",
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

func (isc *IndexedString) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(isc.Unknowns) > 0 {
		data, err := json.Marshal(isc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = isc.Negate
	result["operator"] = isc.Operator
	if isc.Value != nil {
		result["value"] = *isc.Value
	}
	return result, nil
}

func (isc *IndexedString) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), isc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &isc.Unknowns); err != nil {
			return err
		}
		delete(isc.Unknowns, "type")
		delete(isc.Unknowns, "negate")
		delete(isc.Unknowns, "operator")
		delete(isc.Unknowns, "value")
		if len(isc.Unknowns) == 0 {
			isc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		isc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		isc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		isc.Operator = indexed_string.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		isc.Value = opt.NewString(value.(string))
	}
	return nil
}

func (isc *IndexedString) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(isc.Unknowns) > 0 {
		for k, v := range isc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(isc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(isc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&isc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if isc.Value != nil {
		rawMessage, err := json.Marshal(isc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (isc *IndexedString) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	isc.Type = isc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &isc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &isc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &isc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		isc.Unknowns = m
	}
	return nil
}
