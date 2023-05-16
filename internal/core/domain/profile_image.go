package domain

type ProfileImage struct {
	ID       uint64 `json:"id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FilePath string `json:"file_path,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
}
