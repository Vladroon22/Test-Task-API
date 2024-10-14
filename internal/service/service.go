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

type DataToChange struct {
	Song         string `json:"songtochange"`
	DataToChange string `json:"datatochange"`
}

type Service struct {
	Data
	DataToChange
}

func NewService() *Service {
	return &Service{}
}
