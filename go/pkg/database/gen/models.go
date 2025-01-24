// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gen

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type ApisAuthType string

const (
	ApisAuthTypeKey ApisAuthType = "key"
	ApisAuthTypeJwt ApisAuthType = "jwt"
)

func (e *ApisAuthType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ApisAuthType(s)
	case string:
		*e = ApisAuthType(s)
	default:
		return fmt.Errorf("unsupported scan type for ApisAuthType: %T", src)
	}
	return nil
}

type NullApisAuthType struct {
	ApisAuthType ApisAuthType
	Valid        bool // Valid is true if ApisAuthType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullApisAuthType) Scan(value interface{}) error {
	if value == nil {
		ns.ApisAuthType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ApisAuthType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullApisAuthType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ApisAuthType), nil
}

type EventsEvent string

const (
	EventsEventVerificationsusagerecord EventsEvent = "verifications.usage.record"
	EventsEventXx                       EventsEvent = "xx"
)

func (e *EventsEvent) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EventsEvent(s)
	case string:
		*e = EventsEvent(s)
	default:
		return fmt.Errorf("unsupported scan type for EventsEvent: %T", src)
	}
	return nil
}

type NullEventsEvent struct {
	EventsEvent EventsEvent
	Valid       bool // Valid is true if EventsEvent is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEventsEvent) Scan(value interface{}) error {
	if value == nil {
		ns.EventsEvent, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EventsEvent.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEventsEvent) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EventsEvent), nil
}

type EventsState string

const (
	EventsStateCreated   EventsState = "created"
	EventsStateRetrying  EventsState = "retrying"
	EventsStateDelivered EventsState = "delivered"
	EventsStateFailed    EventsState = "failed"
)

func (e *EventsState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EventsState(s)
	case string:
		*e = EventsState(s)
	default:
		return fmt.Errorf("unsupported scan type for EventsState: %T", src)
	}
	return nil
}

type NullEventsState struct {
	EventsState EventsState
	Valid       bool // Valid is true if EventsState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEventsState) Scan(value interface{}) error {
	if value == nil {
		ns.EventsState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EventsState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEventsState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EventsState), nil
}

type RatelimitOverridesSharding string

const (
	RatelimitOverridesShardingEdge RatelimitOverridesSharding = "edge"
)

func (e *RatelimitOverridesSharding) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RatelimitOverridesSharding(s)
	case string:
		*e = RatelimitOverridesSharding(s)
	default:
		return fmt.Errorf("unsupported scan type for RatelimitOverridesSharding: %T", src)
	}
	return nil
}

type NullRatelimitOverridesSharding struct {
	RatelimitOverridesSharding RatelimitOverridesSharding
	Valid                      bool // Valid is true if RatelimitOverridesSharding is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRatelimitOverridesSharding) Scan(value interface{}) error {
	if value == nil {
		ns.RatelimitOverridesSharding, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RatelimitOverridesSharding.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRatelimitOverridesSharding) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RatelimitOverridesSharding), nil
}

type VercelBindingsEnvironment string

const (
	VercelBindingsEnvironmentDevelopment VercelBindingsEnvironment = "development"
	VercelBindingsEnvironmentPreview     VercelBindingsEnvironment = "preview"
	VercelBindingsEnvironmentProduction  VercelBindingsEnvironment = "production"
)

func (e *VercelBindingsEnvironment) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VercelBindingsEnvironment(s)
	case string:
		*e = VercelBindingsEnvironment(s)
	default:
		return fmt.Errorf("unsupported scan type for VercelBindingsEnvironment: %T", src)
	}
	return nil
}

type NullVercelBindingsEnvironment struct {
	VercelBindingsEnvironment VercelBindingsEnvironment
	Valid                     bool // Valid is true if VercelBindingsEnvironment is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVercelBindingsEnvironment) Scan(value interface{}) error {
	if value == nil {
		ns.VercelBindingsEnvironment, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VercelBindingsEnvironment.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVercelBindingsEnvironment) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VercelBindingsEnvironment), nil
}

type VercelBindingsResourceType string

const (
	VercelBindingsResourceTypeRootKey VercelBindingsResourceType = "rootKey"
	VercelBindingsResourceTypeApiId   VercelBindingsResourceType = "apiId"
)

func (e *VercelBindingsResourceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VercelBindingsResourceType(s)
	case string:
		*e = VercelBindingsResourceType(s)
	default:
		return fmt.Errorf("unsupported scan type for VercelBindingsResourceType: %T", src)
	}
	return nil
}

type NullVercelBindingsResourceType struct {
	VercelBindingsResourceType VercelBindingsResourceType
	Valid                      bool // Valid is true if VercelBindingsResourceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVercelBindingsResourceType) Scan(value interface{}) error {
	if value == nil {
		ns.VercelBindingsResourceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VercelBindingsResourceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVercelBindingsResourceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VercelBindingsResourceType), nil
}

type WorkspacesPlan string

const (
	WorkspacesPlanFree       WorkspacesPlan = "free"
	WorkspacesPlanPro        WorkspacesPlan = "pro"
	WorkspacesPlanEnterprise WorkspacesPlan = "enterprise"
)

func (e *WorkspacesPlan) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkspacesPlan(s)
	case string:
		*e = WorkspacesPlan(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkspacesPlan: %T", src)
	}
	return nil
}

type NullWorkspacesPlan struct {
	WorkspacesPlan WorkspacesPlan
	Valid          bool // Valid is true if WorkspacesPlan is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkspacesPlan) Scan(value interface{}) error {
	if value == nil {
		ns.WorkspacesPlan, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkspacesPlan.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkspacesPlan) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkspacesPlan), nil
}

type WorkspacesPlanDowngradeRequest string

const (
	WorkspacesPlanDowngradeRequestFree WorkspacesPlanDowngradeRequest = "free"
)

func (e *WorkspacesPlanDowngradeRequest) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkspacesPlanDowngradeRequest(s)
	case string:
		*e = WorkspacesPlanDowngradeRequest(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkspacesPlanDowngradeRequest: %T", src)
	}
	return nil
}

type NullWorkspacesPlanDowngradeRequest struct {
	WorkspacesPlanDowngradeRequest WorkspacesPlanDowngradeRequest
	Valid                          bool // Valid is true if WorkspacesPlanDowngradeRequest is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkspacesPlanDowngradeRequest) Scan(value interface{}) error {
	if value == nil {
		ns.WorkspacesPlanDowngradeRequest, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkspacesPlanDowngradeRequest.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkspacesPlanDowngradeRequest) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkspacesPlanDowngradeRequest), nil
}

type Api struct {
	ID               string           `db:"id"`
	Name             string           `db:"name"`
	WorkspaceID      string           `db:"workspace_id"`
	IpWhitelist      sql.NullString   `db:"ip_whitelist"`
	AuthType         NullApisAuthType `db:"auth_type"`
	KeyAuthID        sql.NullString   `db:"key_auth_id"`
	CreatedAt        sql.NullTime     `db:"created_at"`
	DeletedAt        sql.NullTime     `db:"deleted_at"`
	DeleteProtection sql.NullBool     `db:"delete_protection"`
}

type AuditLog struct {
	ID          string          `db:"id"`
	WorkspaceID string          `db:"workspace_id"`
	BucketID    string          `db:"bucket_id"`
	Event       string          `db:"event"`
	Time        int64           `db:"time"`
	Display     string          `db:"display"`
	RemoteIp    sql.NullString  `db:"remote_ip"`
	UserAgent   sql.NullString  `db:"user_agent"`
	ActorType   string          `db:"actor_type"`
	ActorID     string          `db:"actor_id"`
	ActorName   sql.NullString  `db:"actor_name"`
	ActorMeta   json.RawMessage `db:"actor_meta"`
	CreatedAt   int64           `db:"created_at"`
	UpdatedAt   sql.NullInt64   `db:"updated_at"`
}

type AuditLogBucket struct {
	ID               string        `db:"id"`
	WorkspaceID      string        `db:"workspace_id"`
	Name             string        `db:"name"`
	RetentionDays    sql.NullInt32 `db:"retention_days"`
	CreatedAt        int64         `db:"created_at"`
	UpdatedAt        sql.NullInt64 `db:"updated_at"`
	DeleteProtection sql.NullBool  `db:"delete_protection"`
}

type AuditLogTarget struct {
	WorkspaceID string          `db:"workspace_id"`
	BucketID    string          `db:"bucket_id"`
	AuditLogID  string          `db:"audit_log_id"`
	DisplayName string          `db:"display_name"`
	Type        string          `db:"type"`
	ID          string          `db:"id"`
	Name        sql.NullString  `db:"name"`
	Meta        json.RawMessage `db:"meta"`
	CreatedAt   int64           `db:"created_at"`
	UpdatedAt   sql.NullInt64   `db:"updated_at"`
}

type EncryptedKey struct {
	WorkspaceID     string `db:"workspace_id"`
	KeyID           string `db:"key_id"`
	Encrypted       string `db:"encrypted"`
	EncryptionKeyID string `db:"encryption_key_id"`
}

type Event struct {
	ID          string          `db:"id"`
	WorkspaceID string          `db:"workspace_id"`
	WebhookID   string          `db:"webhook_id"`
	Event       EventsEvent     `db:"event"`
	Time        int64           `db:"time"`
	Payload     json.RawMessage `db:"payload"`
	State       EventsState     `db:"state"`
	CreatedAt   int64           `db:"created_at"`
	UpdatedAt   sql.NullInt64   `db:"updated_at"`
}

type Gateway struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	WorkspaceID string       `db:"workspace_id"`
	Origin      string       `db:"origin"`
	CreatedAt   sql.NullTime `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type GatewayHeaderRewrite struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	SecretID    string       `db:"secret_id"`
	WorkspaceID string       `db:"workspace_id"`
	GatewaysID  string       `db:"gateways_id"`
	CreatedAt   sql.NullTime `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type Identity struct {
	ID          string          `db:"id"`
	ExternalID  string          `db:"external_id"`
	WorkspaceID string          `db:"workspace_id"`
	Environment string          `db:"environment"`
	CreatedAt   int64           `db:"created_at"`
	UpdatedAt   sql.NullInt64   `db:"updated_at"`
	Meta        json.RawMessage `db:"meta"`
}

type Key struct {
	ID                string         `db:"id"`
	KeyAuthID         string         `db:"key_auth_id"`
	Hash              string         `db:"hash"`
	Start             string         `db:"start"`
	WorkspaceID       string         `db:"workspace_id"`
	ForWorkspaceID    sql.NullString `db:"for_workspace_id"`
	Name              sql.NullString `db:"name"`
	OwnerID           sql.NullString `db:"owner_id"`
	IdentityID        sql.NullString `db:"identity_id"`
	Meta              sql.NullString `db:"meta"`
	CreatedAt         time.Time      `db:"created_at"`
	Expires           sql.NullTime   `db:"expires"`
	CreatedAtM        int64          `db:"created_at_m"`
	UpdatedAtM        sql.NullInt64  `db:"updated_at_m"`
	DeletedAtM        sql.NullInt64  `db:"deleted_at_m"`
	DeletedAt         sql.NullTime   `db:"deleted_at"`
	RefillDay         sql.NullInt16  `db:"refill_day"`
	RefillAmount      sql.NullInt32  `db:"refill_amount"`
	LastRefillAt      sql.NullTime   `db:"last_refill_at"`
	Enabled           bool           `db:"enabled"`
	RemainingRequests sql.NullInt32  `db:"remaining_requests"`
	RatelimitAsync    sql.NullBool   `db:"ratelimit_async"`
	RatelimitLimit    sql.NullInt32  `db:"ratelimit_limit"`
	RatelimitDuration sql.NullInt64  `db:"ratelimit_duration"`
	Environment       sql.NullString `db:"environment"`
}

type KeyAuth struct {
	ID                 string         `db:"id"`
	WorkspaceID        string         `db:"workspace_id"`
	CreatedAt          sql.NullTime   `db:"created_at"`
	DeletedAt          sql.NullTime   `db:"deleted_at"`
	CreatedAtM         int64          `db:"created_at_m"`
	UpdatedAtM         sql.NullInt64  `db:"updated_at_m"`
	DeletedAtM         sql.NullInt64  `db:"deleted_at_m"`
	StoreEncryptedKeys bool           `db:"store_encrypted_keys"`
	DefaultPrefix      sql.NullString `db:"default_prefix"`
	DefaultBytes       sql.NullInt32  `db:"default_bytes"`
	SizeApprox         int32          `db:"size_approx"`
	SizeLastUpdatedAt  int64          `db:"size_last_updated_at"`
}

type KeyMigrationError struct {
	ID          string          `db:"id"`
	MigrationID string          `db:"migration_id"`
	CreatedAt   int64           `db:"created_at"`
	WorkspaceID string          `db:"workspace_id"`
	Message     json.RawMessage `db:"message"`
}

type KeysPermission struct {
	TempID       int64        `db:"temp_id"`
	KeyID        string       `db:"key_id"`
	PermissionID string       `db:"permission_id"`
	WorkspaceID  string       `db:"workspace_id"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

type KeysRole struct {
	KeyID       string       `db:"key_id"`
	RoleID      string       `db:"role_id"`
	WorkspaceID string       `db:"workspace_id"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type LlmGateway struct {
	ID          string        `db:"id"`
	Name        string        `db:"name"`
	Subdomain   string        `db:"subdomain"`
	WorkspaceID string        `db:"workspace_id"`
	CreatedAt   int64         `db:"created_at"`
	UpdatedAt   sql.NullInt64 `db:"updated_at"`
}

type Permission struct {
	ID          string         `db:"id"`
	WorkspaceID string         `db:"workspace_id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

type Ratelimit struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	WorkspaceID string         `db:"workspace_id"`
	CreatedAt   int64          `db:"created_at"`
	UpdatedAt   sql.NullInt64  `db:"updated_at"`
	KeyID       sql.NullString `db:"key_id"`
	IdentityID  sql.NullString `db:"identity_id"`
	Limit       int32          `db:"limit"`
	Duration    int64          `db:"duration"`
}

type RatelimitNamespace struct {
	ID          string       `db:"id"`
	WorkspaceID string       `db:"workspace_id"`
	Name        string       `db:"name"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type RatelimitOverride struct {
	ID          string                         `db:"id"`
	WorkspaceID string                         `db:"workspace_id"`
	NamespaceID string                         `db:"namespace_id"`
	Identifier  string                         `db:"identifier"`
	Limit       int32                          `db:"limit"`
	Duration    int32                          `db:"duration"`
	Async       sql.NullBool                   `db:"async"`
	Sharding    NullRatelimitOverridesSharding `db:"sharding"`
	CreatedAt   time.Time                      `db:"created_at"`
	UpdatedAt   sql.NullTime                   `db:"updated_at"`
	DeletedAt   sql.NullTime                   `db:"deleted_at"`
}

type Role struct {
	ID          string         `db:"id"`
	WorkspaceID string         `db:"workspace_id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

type RolesPermission struct {
	RoleID       string       `db:"role_id"`
	PermissionID string       `db:"permission_id"`
	WorkspaceID  string       `db:"workspace_id"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	CreatedAt    time.Time    `db:"created_at"`
}

type Secret struct {
	ID              string         `db:"id"`
	Name            string         `db:"name"`
	WorkspaceID     string         `db:"workspace_id"`
	Encrypted       string         `db:"encrypted"`
	EncryptionKeyID string         `db:"encryption_key_id"`
	CreatedAt       int64          `db:"created_at"`
	UpdatedAt       sql.NullInt64  `db:"updated_at"`
	Comment         sql.NullString `db:"comment"`
}

type UsageReporter struct {
	ID            string        `db:"id"`
	WebhookID     string        `db:"webhook_id"`
	KeySpaceID    string        `db:"key_space_id"`
	WorkspaceID   string        `db:"workspace_id"`
	Interval      int64         `db:"interval"`
	HighWaterMark int64         `db:"high_water_mark"`
	NextExecution int64         `db:"next_execution"`
	CreatedAt     int64         `db:"created_at"`
	UpdatedAt     sql.NullInt64 `db:"updated_at"`
}

type VercelBinding struct {
	ID            string                     `db:"id"`
	IntegrationID string                     `db:"integration_id"`
	WorkspaceID   string                     `db:"workspace_id"`
	ProjectID     string                     `db:"project_id"`
	Environment   VercelBindingsEnvironment  `db:"environment"`
	ResourceID    string                     `db:"resource_id"`
	ResourceType  VercelBindingsResourceType `db:"resource_type"`
	VercelEnvID   string                     `db:"vercel_env_id"`
	CreatedAt     time.Time                  `db:"created_at"`
	UpdatedAt     time.Time                  `db:"updated_at"`
	DeletedAt     sql.NullTime               `db:"deleted_at"`
	LastEditedBy  string                     `db:"last_edited_by"`
}

type VercelIntegration struct {
	ID          string         `db:"id"`
	WorkspaceID string         `db:"workspace_id"`
	TeamID      sql.NullString `db:"team_id"`
	AccessToken string         `db:"access_token"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}

type Webhook struct {
	ID              string        `db:"id"`
	Destination     string        `db:"destination"`
	WorkspaceID     string        `db:"workspace_id"`
	Enabled         bool          `db:"enabled"`
	CreatedAt       int64         `db:"created_at"`
	UpdatedAt       sql.NullInt64 `db:"updated_at"`
	Encrypted       string        `db:"encrypted"`
	EncryptionKeyID string        `db:"encryption_key_id"`
}

type WebhookDeliveryAttempt struct {
	ID             string         `db:"id"`
	WorkspaceID    string         `db:"workspace_id"`
	WebhookID      string         `db:"webhook_id"`
	EventID        string         `db:"event_id"`
	Time           int64          `db:"time"`
	Attempt        int32          `db:"attempt"`
	NextAttemptAt  sql.NullInt64  `db:"next_attempt_at"`
	Success        bool           `db:"success"`
	InternalError  sql.NullString `db:"internal_error"`
	ResponseStatus sql.NullInt32  `db:"response_status"`
	ResponseBody   sql.NullString `db:"response_body"`
}

type Workspace struct {
	ID                   string                             `db:"id"`
	TenantID             string                             `db:"tenant_id"`
	Name                 string                             `db:"name"`
	CreatedAt            sql.NullTime                       `db:"created_at"`
	DeletedAt            sql.NullTime                       `db:"deleted_at"`
	Plan                 NullWorkspacesPlan                 `db:"plan"`
	StripeCustomerID     sql.NullString                     `db:"stripe_customer_id"`
	StripeSubscriptionID sql.NullString                     `db:"stripe_subscription_id"`
	TrialEnds            sql.NullTime                       `db:"trial_ends"`
	BetaFeatures         json.RawMessage                    `db:"beta_features"`
	Features             json.RawMessage                    `db:"features"`
	PlanLockedUntil      sql.NullTime                       `db:"plan_locked_until"`
	PlanDowngradeRequest NullWorkspacesPlanDowngradeRequest `db:"plan_downgrade_request"`
	PlanChanged          sql.NullTime                       `db:"plan_changed"`
	Subscriptions        json.RawMessage                    `db:"subscriptions"`
	Enabled              bool                               `db:"enabled"`
	DeleteProtection     sql.NullBool                       `db:"delete_protection"`
}
