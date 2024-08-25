package home

const AboutJsonFile = "./web/views/home/about.json"

type AboutJSON struct {
	Title   string `json:"title"`
	Content []string `json:"content"`
}
