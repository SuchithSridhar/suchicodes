package components

const NavbarJsonFile = "./web/components/navbar.json"
const FooterJsonFile = "./web/components/footer.json"

type NavbarJSON struct {
	Brand string `json:"brand"`
	Links []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"links"`
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
