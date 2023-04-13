package model

type SyncMode string

const (
	SyncModeNormal = "normal"
	SyncModeGood   = "good"
	SyncModeMerge  = "merge"
)

type GetSwaggerByUrlResp struct {
	CommonResp
	Data any
}

type UpdateSwaggerSyncConfigReq struct {
	ProjectID   int      `json:"project_id"`
	IsSyncOpen  bool     `json:"is_sync_open"`
	UID         int      `json:"uid"`
	ID          int      `json:"id"`
	SyncMode    SyncMode `json:"sync_mode"`
	SyncJSONURL string   `json:"sync_json_url"`
	SyncCron    string   `json:"sync_cron"`
}
