// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type LabelColor struct {
	LabelColorID uuid.UUID `json:"label_color_id"`
	ColorHex     string    `json:"color_hex"`
	Position     float64   `json:"position"`
	Name         string    `json:"name"`
}

type Notification struct {
	NotificationID       uuid.UUID `json:"notification_id"`
	NotificationObjectID uuid.UUID `json:"notification_object_id"`
	NotifierID           uuid.UUID `json:"notifier_id"`
	Read                 bool      `json:"read"`
}

type NotificationObject struct {
	NotificationObjectID uuid.UUID `json:"notification_object_id"`
	EntityID             uuid.UUID `json:"entity_id"`
	ActionType           int32     `json:"action_type"`
	ActorID              uuid.UUID `json:"actor_id"`
	EntityType           int32     `json:"entity_type"`
	CreatedOn            time.Time `json:"created_on"`
}

type Organization struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
}

type PersonalProject struct {
	PersonalProjectID uuid.UUID `json:"personal_project_id"`
	ProjectID         uuid.UUID `json:"project_id"`
	UserID            uuid.UUID `json:"user_id"`
}

type Project struct {
	ProjectID uuid.UUID `json:"project_id"`
	TeamID    uuid.UUID `json:"team_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type ProjectLabel struct {
	ProjectLabelID uuid.UUID      `json:"project_label_id"`
	ProjectID      uuid.UUID      `json:"project_id"`
	LabelColorID   uuid.UUID      `json:"label_color_id"`
	CreatedDate    time.Time      `json:"created_date"`
	Name           sql.NullString `json:"name"`
}

type ProjectMember struct {
	ProjectMemberID uuid.UUID `json:"project_member_id"`
	ProjectID       uuid.UUID `json:"project_id"`
	UserID          uuid.UUID `json:"user_id"`
	AddedAt         time.Time `json:"added_at"`
	RoleCode        string    `json:"role_code"`
}

type RefreshToken struct {
	TokenID   uuid.UUID `json:"token_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Role struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SystemOption struct {
	OptionID uuid.UUID      `json:"option_id"`
	Key      string         `json:"key"`
	Value    sql.NullString `json:"value"`
}

type Task struct {
	TaskID      uuid.UUID      `json:"task_id"`
	TaskGroupID uuid.UUID      `json:"task_group_id"`
	CreatedAt   time.Time      `json:"created_at"`
	Name        string         `json:"name"`
	Position    float64        `json:"position"`
	Description sql.NullString `json:"description"`
	DueDate     sql.NullTime   `json:"due_date"`
	Complete    bool           `json:"complete"`
	CompletedAt sql.NullTime   `json:"completed_at"`
}

type TaskAssigned struct {
	TaskAssignedID uuid.UUID `json:"task_assigned_id"`
	TaskID         uuid.UUID `json:"task_id"`
	UserID         uuid.UUID `json:"user_id"`
	AssignedDate   time.Time `json:"assigned_date"`
}

type TaskChecklist struct {
	TaskChecklistID uuid.UUID `json:"task_checklist_id"`
	TaskID          uuid.UUID `json:"task_id"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

type TaskChecklistItem struct {
	TaskChecklistItemID uuid.UUID    `json:"task_checklist_item_id"`
	TaskChecklistID     uuid.UUID    `json:"task_checklist_id"`
	CreatedAt           time.Time    `json:"created_at"`
	Complete            bool         `json:"complete"`
	Name                string       `json:"name"`
	Position            float64      `json:"position"`
	DueDate             sql.NullTime `json:"due_date"`
}

type TaskGroup struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	ProjectID   uuid.UUID `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

type TaskLabel struct {
	TaskLabelID    uuid.UUID `json:"task_label_id"`
	TaskID         uuid.UUID `json:"task_id"`
	ProjectLabelID uuid.UUID `json:"project_label_id"`
	AssignedDate   time.Time `json:"assigned_date"`
}

type Team struct {
	TeamID         uuid.UUID `json:"team_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

type TeamMember struct {
	TeamMemberID uuid.UUID `json:"team_member_id"`
	TeamID       uuid.UUID `json:"team_id"`
	UserID       uuid.UUID `json:"user_id"`
	Addeddate    time.Time `json:"addeddate"`
	RoleCode     string    `json:"role_code"`
}

type UserAccount struct {
	UserID           uuid.UUID      `json:"user_id"`
	CreatedAt        time.Time      `json:"created_at"`
	Email            string         `json:"email"`
	Username         string         `json:"username"`
	PasswordHash     string         `json:"password_hash"`
	ProfileBgColor   string         `json:"profile_bg_color"`
	FullName         string         `json:"full_name"`
	Initials         string         `json:"initials"`
	ProfileAvatarUrl sql.NullString `json:"profile_avatar_url"`
	RoleCode         string         `json:"role_code"`
	Bio              string         `json:"bio"`
}
