package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pglet/pglet/internal/cache"
	"github.com/pglet/pglet/internal/model"
	"github.com/pglet/pglet/internal/utils"
)

const (
	pageNextIDKey           = "page_next_id"                  // Inc integer with the next page ID
	pageKey                 = "page:%s"                       // page JSON data
	pageHostClientsKey      = "page_host_clients:%d"          // a Set with client IDs
	pageSessionsKey         = "page_sessions:%d"              // a Set with session IDs
	sessionKey              = "session:%d:%s"                 // session JSON data
	sessionNextControlIDKey = "session_next_control_id:%d:%s" // Inc integer with the next control ID for a given session
	sessionControlsKey      = "session_controls:%d:%s"        // session controls, value is JSON data
	sessionHostClientsKey   = "session_host_clients:%d:%s"    // a Set with client IDs
	sessionWebClientsKey    = "session_web_clients:%d:%s"     // a Set with client IDs
)

//
// Pages
// ==============================

func GetPage(pageName string) *model.Page {
	j := cache.GetString(fmt.Sprintf(pageKey, pageName))
	if j == "" {
		return nil
	}
	p := new(model.Page)
	json.Unmarshal([]byte(j), p)
	return p
}

func AddPage(page *model.Page) {

	// TODO - check if the page exists

	pageID := cache.Inc(pageNextIDKey, 1)
	page.ID = pageID
	cache.SetString(fmt.Sprintf(pageKey, page.Name), utils.ToJSON(page), 0)
}

//
// Page Host Clients
// ==============================

func GetPageHostClients(page *model.Page) []string {
	return cache.SetGet(fmt.Sprintf(pageHostClientsKey, page.ID))
}

func AddPageHostClient(page *model.Page, clientID string) {
	cache.SetAdd(fmt.Sprintf(pageHostClientsKey, page.ID), clientID)
}

func RemovePageHostClient(page *model.Page, clientID string) {
	cache.SetRemove(fmt.Sprintf(pageHostClientsKey, page.ID), clientID)
}

//
// Sessions
// ==============================

func GetSession(page *model.Page, sessionID string) *model.Session {

	j := cache.GetString(fmt.Sprintf(sessionKey, page.ID, sessionID))
	if j == "" {
		return nil
	}
	session := new(model.Session)
	json.Unmarshal([]byte(j), &session)
	session.Page = page
	return session
}

func AddSession(session *model.Session) {
	cache.SetString(fmt.Sprintf(sessionKey, session.Page.ID, session.ID), utils.ToJSON(session), 0)
	cache.SetAdd(fmt.Sprintf(pageSessionsKey, session.Page.ID), session.ID)
}

func DeleteSession(session *model.Session) {
	cache.SetRemove(fmt.Sprintf(pageSessionsKey, session.Page.ID), session.ID)
	cache.Remove(fmt.Sprintf(sessionKey, session.Page.ID, session.ID))
	cache.Remove(fmt.Sprintf(sessionNextControlIDKey, session.Page.ID, session.ID))
	cache.Remove(fmt.Sprintf(sessionControlsKey, session.Page.ID, session.ID))
}

//
// Controls
// ==============================

func GetSessionNextControlID(session *model.Session) int {
	return cache.Inc(fmt.Sprintf(sessionNextControlIDKey, session.Page.ID, session.ID), 1)
}

func GetSessionControl(session *model.Session, ctrlID string) *model.Control {
	cj := cache.HashGet(fmt.Sprintf(sessionControlsKey, session.Page.ID, session.ID), ctrlID)
	if cj == "" {
		return nil
	}
	ctrl, err := model.NewControlFromJSON(cj)
	if err != nil {
		log.Fatal(err)
	}
	return ctrl
}

func GetAllSessionControls(session *model.Session) map[string]*model.Control {
	fields := cache.HashGetAll(fmt.Sprintf(sessionControlsKey, session.Page.ID, session.ID))
	controls := make(map[string]*model.Control, len(fields))
	for k, v := range fields {
		ctrl, _ := model.NewControlFromJSON(v)
		controls[k] = ctrl
	}
	return controls
}

func SetSessionControl(session *model.Session, ctrl *model.Control) {
	cj := utils.ToJSON(ctrl)
	cache.HashSet(fmt.Sprintf(sessionControlsKey, session.Page.ID, session.ID), ctrl.ID(), cj)
}

func DeleteSessionControl(session *model.Session, ctrlID string) {
	cache.HashRemove(fmt.Sprintf(sessionControlsKey, session.Page.ID, session.ID), ctrlID)
}

//
// Session Host Clients
// ==============================

func GetSessionHostClients(session *model.Session) []string {
	return cache.SetGet(fmt.Sprintf(sessionHostClientsKey, session.Page.ID, session.ID))
}

func AddSessionHostClient(session *model.Session, clientID string) {
	cache.SetAdd(fmt.Sprintf(sessionHostClientsKey, session.Page.ID, session.ID), clientID)
}

func RemoveSessionHostClient(session *model.Session, clientID string) {
	cache.SetRemove(fmt.Sprintf(sessionHostClientsKey, session.Page.ID, session.ID), clientID)
}

//
// Session Web Clients
// ==============================

func GetSessionWebClients(session *model.Session) []string {
	return cache.SetGet(fmt.Sprintf(sessionWebClientsKey, session.Page.ID, session.ID))
}

func AddSessionWebClient(session *model.Session, clientID string) {
	cache.SetAdd(fmt.Sprintf(sessionWebClientsKey, session.Page.ID, session.ID), clientID)
}

func RemoveSessionWebClient(session *model.Session, clientID string) {
	cache.SetRemove(fmt.Sprintf(sessionWebClientsKey, session.Page.ID, session.ID), clientID)
}
