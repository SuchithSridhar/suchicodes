package home

const ContactJsonFile = "./web/views/home/contact.json"

type ContactJSON struct {
	Title string `json:"title"`
	Form  struct {
		FormTitle          string `json:"form_title"`
		SubjectPlaceholder string `json:"subject_placeholder"`
		MessagePlaceholder string `json:"message_placeholder"`
		ContactInfoNote    string `json:"contact_info_note"`
		HumanTestLabel     string `json:"human_test_label"`
		HumanTestText      string `json:"human_test_text"`
		SubmitButtonText   string `json:"submit_button_text"`
	} `json:"form"`
	ContactButtons []struct {
		Platform string `json:"platform"`
		URL      string `json:"url"`
		Icon     string `json:"icon"`
		AltText  string `json:"alt_text"`
	} `json:"contact_buttons"`
}
