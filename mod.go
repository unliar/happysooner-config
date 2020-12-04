package happyConfig

// MySQL Config Model
type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database,omitempty"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database,omitempty"`
}

type SMTP struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type NSQ struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type WechatMP struct {
	AppID      string `json:"appId"`
	TemplateID string `json:"templateId"`
}
