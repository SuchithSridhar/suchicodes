package home

const IndexJsonFile = "./web/views/home/index.json"

type IndexJSON struct {
	Title   string `json:"title"`
	Landing struct {
		Title string `json:"title"`
		About string `json:"about"`
		Image struct {
			Src string `json:"src"`
			Alt string `json:"alt"`
		}
		Buttons []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		} `json:"buttons"`
	} `json:"landing"`
	VisitSection struct {
		Title  string `json:"title"`
		Topics []struct {
			Title    string `json:"title"`
			Subtitle string `json:"subtitle"`
			Buttons  []struct {
				Text string `json:"text"`
				URL  string `json:"url"`
			} `json:"buttons"`
		} `json:"topics"`
	} `json:"visit_section"`
	AboutSection struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"about_section"`
	WorkSection struct {
		Title       string `json:"title"`
		Experiences []struct {
			Position    string `json:"position"`
			Location    string `json:"location"`
			Start       string `json:"start"`
			End         string `json:"end"`
			Description string `json:"description"`
		} `json:"experiences"`
	} `json:"work_section"`
}
