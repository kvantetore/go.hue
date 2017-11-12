package hue

import (
	"encoding/json"
)

type GroupAction struct {
	On			bool		`json:"on"`
	Bri			int			`json:"bri"`
	Hue			int			`json:"hue"`
	Sat			int			`json:"sat"`
	Effect		string		`json:"effect"`
	Xy			[]float32	`json:"xy"`
	Ct			int			`json:"ct"`
	Alert		string		`json:"alert"`
	ColorMode	string		`json:"colormode"`
}

type Group struct {
	Id			string
	Name        string      		`json:"name"`
	Lights      []string    		`json:"lights"`
	Type        string      		`json:"type"`
	Action		GroupAction			`json:"action"`
}

// GetAllGroups retrieves groups defined on the bridge pr
// https://www.developers.meethue.com/documentation/groups-api#21_get_all_groups
func (bridge *Bridge) GetAllGroups() ([]*Group, error) {
	response, err := bridge.get("/groups")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resultMap := make(map[string]*Group)
	err = json.NewDecoder(response.Body).Decode(&resultMap)
	
	result := []*Group {}
	for id, group := range resultMap {
		group.Id = id
		result = append(result, group)
	}

	return result, err
}


// GetAllRooms returns all groups of type Room defined on the bridge
func (bridge *Bridge) GetAllRooms() ([]*Group, error) {
	groups, err := bridge.GetAllGroups()
	if err != nil {
		return nil, err
	}

	var rooms []*Group
	for _, group := range groups {
		if group.Type ==  "Room" {
			rooms = append(rooms, group)
		}
	}

	return rooms, nil
}

