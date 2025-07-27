package model

import (
	"github.com/go-dev-frame/sponge/pkg/sgorm"
	"time"
)

// BaseModel a base model that includes the ID, CreatedAt, and UpdatedAt fields
// remove deleted_at from sponge/pkg/sgorm/base_model.go
type BaseModel struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type Users struct {
	BaseModel `gorm:"embedded"` // embed id and time

	Email                      string          `gorm:"column:email;type:varchar(255);not null" json:"email"`
	EncryptedPassword          string          `gorm:"column:encrypted_password;type:varchar(255);not null" json:"encryptedPassword"`
	ResetPasswordToken         string          `gorm:"column:reset_password_token;type:varchar(255)" json:"resetPasswordToken"`
	ResetPasswordSentAt        time.Time       `gorm:"column:reset_password_sent_at;type:datetime" json:"resetPasswordSentAt"`
	RememberCreatedAt          time.Time       `gorm:"column:remember_created_at;type:datetime" json:"rememberCreatedAt"`
	SignInCount                int             `gorm:"column:sign_in_count;type:int(11);default:0;not null" json:"signInCount"`
	CurrentSignInAt            time.Time       `gorm:"column:current_sign_in_at;type:datetime" json:"currentSignInAt"`
	LastSignInAt               time.Time       `gorm:"column:last_sign_in_at;type:datetime" json:"lastSignInAt"`
	CurrentSignInIP            string          `gorm:"column:current_sign_in_ip;type:varchar(255)" json:"currentSignInIP"`
	LastSignInIP               string          `gorm:"column:last_sign_in_ip;type:varchar(255)" json:"lastSignInIP"`
	ConfirmationToken          string          `gorm:"column:confirmation_token;type:varchar(255)" json:"confirmationToken"`
	ConfirmedAt                time.Time       `gorm:"column:confirmed_at;type:datetime" json:"confirmedAt"`
	ConfirmationSentAt         time.Time       `gorm:"column:confirmation_sent_at;type:datetime" json:"confirmationSentAt"`
	UnconfirmedEmail           string          `gorm:"column:unconfirmed_email;type:varchar(255)" json:"unconfirmedEmail"`
	FailedAttempts             int             `gorm:"column:failed_attempts;type:int(11);default:0;not null" json:"failedAttempts"`
	UnlockToken                string          `gorm:"column:unlock_token;type:varchar(255)" json:"unlockToken"`
	LockedAt                   time.Time       `gorm:"column:locked_at;type:datetime" json:"lockedAt"`
	InvitationToken            string          `gorm:"column:invitation_token;type:varchar(255)" json:"invitationToken"`
	InvitationCreatedAt        time.Time       `gorm:"column:invitation_created_at;type:datetime" json:"invitationCreatedAt"`
	InvitationSentAt           time.Time       `gorm:"column:invitation_sent_at;type:datetime" json:"invitationSentAt"`
	InvitationAcceptedAt       time.Time       `gorm:"column:invitation_accepted_at;type:datetime" json:"invitationAcceptedAt"`
	InvitationLimit            int             `gorm:"column:invitation_limit;type:int(11)" json:"invitationLimit"`
	InvitedByType              string          `gorm:"column:invited_by_type;type:varchar(255)" json:"invitedByType"`
	InvitedByID                int64           `gorm:"column:invited_by_id;type:bigint(20)" json:"invitedByID"`
	InvitationsCount           int             `gorm:"column:invitations_count;type:int(11);default:0" json:"invitationsCount"`
	PositionTitle              string          `gorm:"column:position_title;type:varchar(255)" json:"positionTitle"`
	ClerkCode                  string          `gorm:"column:clerk_code;type:varchar(255)" json:"clerkCode"`
	ChineseName                string          `gorm:"column:chinese_name;type:varchar(255)" json:"chineseName"`
	DeskPhone                  string          `gorm:"column:desk_phone;type:varchar(255)" json:"deskPhone"`
	JobLevel                   string          `gorm:"column:job_level;type:varchar(255)" json:"jobLevel"`
	WecomID                    string          `gorm:"column:wecom_id;type:varchar(255)" json:"wecomID"`
	PreSsoID                   string          `gorm:"column:pre_sso_id;type:varchar(255)" json:"preSsoID"`
	Mobile                     string          `gorm:"column:mobile;type:varchar(255)" json:"mobile"`
	EntryCompanyDate           time.Time       `gorm:"column:entry_company_date;type:date" json:"entryCompanyDate"`
	Gender                     *sgorm.TinyBool `gorm:"column:gender;type:tinyint(1);default:1" json:"gender"`
	PerPage                    int             `gorm:"column:per_page;type:int(11);default:12" json:"perPage"`
	OpenInNewTab               *sgorm.TinyBool `gorm:"column:open_in_new_tab;type:tinyint(1);default:0" json:"openInNewTab"`
	MajorCode                  string          `gorm:"column:major_code;type:varchar(255)" json:"majorCode"`
	MajorName                  string          `gorm:"column:major_name;type:varchar(255)" json:"majorName"`
	PositionChangedInLastMonth *sgorm.TinyBool `gorm:"column:position_changed_in_last_month;type:tinyint(1);default:0" json:"positionChangedInLastMonth"`
	NewUI                      *sgorm.TinyBool `gorm:"column:new_ui;type:tinyint(1);default:1" json:"newUI"`
	PositionNcPkPost           string          `gorm:"column:position_nc_pk_post;type:varchar(255)" json:"positionNcPkPost"`
	WindowsSid                 string          `gorm:"column:windows_sid;type:varchar(255)" json:"windowsSid"`
}

// UsersColumnNames Whitelist for custom query fields to prevent sql injection attacks
var UsersColumnNames = map[string]bool{
	"id":                             true,
	"created_at":                     true,
	"updated_at":                     true,
	"email":                          true,
	"encrypted_password":             true,
	"reset_password_token":           true,
	"reset_password_sent_at":         true,
	"remember_created_at":            true,
	"sign_in_count":                  true,
	"current_sign_in_at":             true,
	"last_sign_in_at":                true,
	"current_sign_in_ip":             true,
	"last_sign_in_ip":                true,
	"confirmation_token":             true,
	"confirmed_at":                   true,
	"confirmation_sent_at":           true,
	"unconfirmed_email":              true,
	"failed_attempts":                true,
	"unlock_token":                   true,
	"locked_at":                      true,
	"invitation_token":               true,
	"invitation_created_at":          true,
	"invitation_sent_at":             true,
	"invitation_accepted_at":         true,
	"invitation_limit":               true,
	"invited_by_type":                true,
	"invited_by_id":                  true,
	"invitations_count":              true,
	"position_title":                 true,
	"clerk_code":                     true,
	"chinese_name":                   true,
	"desk_phone":                     true,
	"job_level":                      true,
	"wecom_id":                       true,
	"pre_sso_id":                     true,
	"mobile":                         true,
	"entry_company_date":             true,
	"gender":                         true,
	"per_page":                       true,
	"open_in_new_tab":                true,
	"major_code":                     true,
	"major_name":                     true,
	"position_changed_in_last_month": true,
	"new_ui":                         true,
	"position_nc_pk_post":            true,
	"windows_sid":                    true,
}
