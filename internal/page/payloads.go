package page

import (
	"github.com/pglet/pglet/internal/model"
	"github.com/pglet/pglet/internal/page/command"
)

type RegisterHostClientRequestPayload struct {
	PageName string `json:"pageName"`
	IsApp    bool   `json:"isApp"`
}

type RegisterHostClientResponsePayload struct {
	SessionID string `json:"sessionID"`
	PageName  string `json:"pageName"`
	Error     string `json:"error"`
}

type RegisterWebClientRequestPayload struct {
	PageName string `json:"pageName"`
	IsApp    bool   `json:"isApp"`
}

type RegisterWebClientResponsePayload struct {
	Session *model.Session `json:"session"`
	Error   string         `json:"error"`
}

type SessionCreatedPayload struct {
	PageName  string `json:"pageName"`
	SessionID string `json:"sessionID"`
}

type PageCommandRequestPayload struct {
	PageName  string          `json:"pageName"`
	SessionID string          `json:"sessionID"`
	Command   command.Command `json:"command"`
}

type PageCommandResponsePayload struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type PageEventPayload struct {
	PageName    string `json:"pageName"`
	SessionID   string `json:"sessionID"`
	EventTarget string `json:"eventTarget"`
	EventName   string `json:"eventName"`
	EventData   string `json:"eventData"`
}

type AddPageControlsPayload struct {
	Controls []*model.Control `json:"controls"`
}

type UpdateControlPropsPayload struct {
	Props []map[string]interface{} `json:"props"`
}

type AppendControlPropsPayload struct {
	Props []map[string]string `json:"props"`
}

type RemoveControlPayload struct {
	IDs []string `json:"ids"`
}

type CleanControlPayload struct {
	IDs []string `json:"ids"`
}