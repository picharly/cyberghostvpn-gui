package resources

type City struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Load     string `json:"load"`
}

type Country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Server struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Load     string `json:"load"`
}
