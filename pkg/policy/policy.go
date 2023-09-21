package policy

import (
	"encoding/json"
	"strings"
)

type Policy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid       string                 `json:"Sid"`
	Effect    string                 `json:"Effect"`
	Action    interface{}            `json:"Action"`
	Resource  interface{}            `json:"Resource"`
	Condition map[string]interface{} `json:"Condition"`
}

func ParsePolicyDocument(d []byte) (Policy, error) {

	var p Policy

	if err := json.Unmarshal(d, &p); err != nil {
		return p, err
	}

	return p, nil

}

func CheckIfContainsAction(a string, s Statement) bool {

	// Asserting the action into string
	i, ok := s.Action.(string)

	if ok {
		// If assertion to string is successful then compare it.
		if strings.Compare(i, a) == 0 {
			return true
		}

	} else {
		// If assertion is failed then the action is []interface{}
		// Asserting the action into []interface{}
		al, ok := s.Action.([]interface{})

		if ok {
			// Iterating through the slice of interface and asserting it into the string
			for _, k := range al {
				c := k.(string)

				if strings.Compare(c, a) == 0 {
					return true
				}
			}

		}

	}

	return false
}

func CheckIfContainsResource(r string, s Statement) bool {

	i, ok := s.Resource.(string)

	if ok {
		if strings.Compare(r, i) == 0 {
			return true
		}
	} else {
		rl, ok := s.Resource.([]interface{})

		if ok {
			// Iterating through the slice of interface and asserting it into the string
			for _, k := range rl {
				c := k.(string)

				if strings.Compare(c, r) == 0 {
					return true
				}
			}

		}
	}

	return false
}
