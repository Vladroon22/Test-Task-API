package service

type Data struct {
	Group_name   string `json:"group"`
	Song_title   string `json:"song"`
	Genre        string `json:"genre"`
	Album        string `json:"album"`
	Release_date string `json:"release_data"`
	Lyrics       string `json:"lyrics"`
	Link         string `json:"link"`
}

type Service struct {
	Data
}

func NewService() *Service {
	return &Service{}
}
