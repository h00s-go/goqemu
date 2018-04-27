package qemu

// Guest represent one guest data
type Guest struct {
	Password string `json:"password"`
	Monitor  struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"monitor"`
	Params map[string]interface{} `json:"params"`
}
