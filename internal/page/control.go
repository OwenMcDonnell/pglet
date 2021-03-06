package page

import (
	"encoding/json"

	"github.com/pglet/pglet/internal/utils"
)

const (
	Row     string = "row"
	Col            = "col"
	Label          = "label"
	Textbox        = "textbox"
	Link           = "link"
	Button         = "button"
)

var (
	systemAttrs = []string{
		"id",
		"to",
		"from",
		"at",
	}
)

// Control is an element of a page.
type Control map[string]interface{}

// NewControl initializes a new control object.
func NewControl(controlType string, parentID string, id string) *Control {
	ctl := Control{}
	ctl["t"] = controlType
	ctl["p"] = parentID
	ctl["i"] = id
	ctl["c"] = make([]string, 0, 0)
	return &ctl
}

// NewControlFromJSON initializes a new control instance from JSON.
func NewControlFromJSON(jsonCtrl string) (*Control, error) {
	ctrl := &Control{}
	err := json.Unmarshal([]byte(jsonCtrl), ctrl)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (ctl *Control) GetAttr(name string) interface{} {
	return (*ctl)[name]
}

func (ctl *Control) SetAttr(name string, value interface{}) {
	(*ctl)[name] = value
}

func (ctl *Control) AppendAttr(name string, value string) {
	(*ctl)[name] = (*ctl)[name].(string) + value
}

// ID returns control's ID.
func (ctl *Control) ID() string {
	return (*ctl)["i"].(string)
}

func (ctl *Control) At() int {
	at, ok := (*ctl)["at"].(int)
	if ok {
		return at
	}
	return -1
}

// ParentID returns the ID of parent control.
func (ctl *Control) ParentID() string {
	return (*ctl)["p"].(string)
}

// AddChildID appends the child to the parent control.
func (ctl *Control) AddChildID(childID string) {
	childIds, _ := (*ctl)["c"].([]string)
	(*ctl)["c"] = append(childIds, childID)
}

func (ctl *Control) InsertChildID(childID string, at int) {
	childIds, _ := (*ctl)["c"].([]string)
	(*ctl)["c"] = utils.InsertString(childIds, childID, at)
}

func (ctl *Control) RemoveChild(childID string) {
	childIds, _ := (*ctl)["c"].([]string)
	(*ctl)["c"] = utils.RemoveString(childIds, childID)
}

func (ctl *Control) RemoveChildren() {
	(*ctl)["c"] = make([]string, 0, 0)
}

func (ctl *Control) GetChildrenIds() []string {
	ids, _ := (*ctl)["c"].([]string)
	return ids
}

func IsSystemAttr(attr string) bool {
	return utils.ContainsString(systemAttrs, attr)
}
