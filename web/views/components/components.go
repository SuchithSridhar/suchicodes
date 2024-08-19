package components

const NavbarJsonFile = "./web/views/components/navbar.json"
const FooterJsonFile = "./web/views/components/footer.json"

type NavbarJSON struct {
	Brand  string `json:"brand"`
	Links  []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"links"`
	Settings struct {
		ThemeToggler []string `json:"theme-toggler"`
	} `json:"settings"`
}

type FooterJSON struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Sections []struct {
		SectionTitle string `json:"section_title"`
		Links        []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"links"`
	} `json:"sections"`
}
