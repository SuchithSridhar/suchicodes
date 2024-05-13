package components

const NavbarJsonFile = "./web/components/navbar.json"
const FooterJsonFile = "./web/components/footer.json"

type NavbarJSON struct {
	Brand    string `json:"brand"`
	Link1    string `json:"link_1"`
	Link1URL string `json:"link_1_url"`
	Link2    string `json:"link_2"`
	Link2URL string `json:"link_2_url"`
	Link3    string `json:"link_3"`
	Link3URL string `json:"link_3_url"`
	Link4    string `json:"link_4"`
	Link4URL string `json:"link_4_url"`
}

type FooterJSON struct {
	Title                  string `json:"title"`
	Subtitle               string `json:"subtitle"`
	MainSection            string `json:"main_section"`
	MainLink1Title         string `json:"main_link_1_title"`
	MainLink1URL           string `json:"main_link_1_url"`
	MainLink2Title         string `json:"main_link_2_title"`
	MainLink2URL           string `json:"main_link_2_url"`
	MainLink3Title         string `json:"main_link_3_title"`
	MainLink3URL           string `json:"main_link_3_url"`
	MainLink4Title         string `json:"main_link_4_title"`
	MainLink4URL           string `json:"main_link_4_url"`
	MainLink5Title         string `json:"main_link_5_title"`
	MainLink5URL           string `json:"main_link_5_url"`
	ProfessionalSection    string `json:"professional_section"`
	ProfessionalLink1Title string `json:"professional_link_1_title"`
	ProfessionalLink1URL   string `json:"professional_link_1_url"`
	ProfessionalLink2Title string `json:"professional_link_2_title"`
	ProfessionalLink2URL   string `json:"professional_link_2_url"`
	ProfessionalLink3Title string `json:"professional_link_3_title"`
	ProfessionalLink3URL   string `json:"professional_link_3_url"`
	AcademicSection        string `json:"academic_section"`
	AcademicLink1Title     string `json:"academic_link_1_title"`
	AcademicLink1URL       string `json:"academic_link_1_url"`
	AcademicLink2Title     string `json:"academic_link_2_title"`
	AcademicLink2URL       string `json:"academic_link_2_url"`
	AcademicLink3Title     string `json:"academic_link_3_title"`
	AcademicLink3URL       string `json:"academic_link_3_url"`
	AcademicLink4Title     string `json:"academic_link_4_title"`
	AcademicLink4URL       string `json:"academic_link_4_url"`
	DevelopersSection      string `json:"developers_section"`
	DevelopersLink1Title   string `json:"developers_link_1_title"`
	DevelopersLink1URL     string `json:"developers_link_1_url"`
	DevelopersLink2Title   string `json:"developers_link_2_title"`
	DevelopersLink2URL     string `json:"developers_link_2_url"`
	DevelopersLink3Title   string `json:"developers_link_3_title"`
	DevelopersLink3URL     string `json:"developers_link_3_url"`
}
