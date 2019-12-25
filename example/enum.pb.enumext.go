package common

// -------------------------------------------------------------------
// State
// -------------------------------------------------------------------
func (e State) IsValid() bool {
	_, ok := State_name[int32(e)]
	return ok
}

var gql_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "felix_text",
	2: "STATE_INACTIVE",
	3: "STATE_DELETED",
	4: "STATE_DENY",
}
var gql_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"felix_text":        1,
	"STATE_INACTIVE":    2,
	"STATE_DELETED":     3,
	"STATE_DENY":        4,
}

func (e State) ExtValue() (string, error) {
	if !e.IsValid() {
		return "", fmt.Errorf("invalid State '%s'", e)
	}
	return gql_State_name[int32(e)], nil
}
