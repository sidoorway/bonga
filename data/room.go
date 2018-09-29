package data

type Room struct {
	Status       bool   `json:"status"`
	ChatStatus   string `json:"chat_status"`
	ChatURL      string `json:"chat_url"`
	EmbedChatURL string `json:"embed_chat_url"`

	MembersCount int `json:"members_count"`

	UserName    string `json:"username"`
	DisplayName string `json:"display_name"`

	Gender string `json:"gender"`

	Ethnicity string `json:"ethnicity"`
	HomeTown  string `json:"hometown"`

	DisplayAge    string `json:"display_age"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	HairColor     string `json:"hair_color"`
	EyeColor      string `json:"eye_color"`
	BustPenisSize string `json:"bust_penis_size"`
	PubicHair     string `json:"pubic_hair"`

	TurnsOn  string `json:"turns_on"`
	TurnsOff string `json:"turns_off"`

	PrimaryLanguageKey   string `json:"primary_language_key"`
	PrimaryLanguage      string `json:"primary_language"`
	SecondaryLanguageKey string `json:"secondary_language_key"`
	SecondaryLanguage    string `json:"secondary_language"`

	RestrictedCountryIDs string `json:"restricted_country_ids"`
	RestrictedStateIDs   string `json:"restricted_state_ids"`

	SignupDate     string `json:"signup_date"`
	LastUpdateDate string `json:"last_update_date"`
	OnlineTime     int    `json:"online_time"`

	Marker string `json:"marker"`

	IsGeo bool `json:"is_geo"`

	ProfileImages struct {
		ProfileImage             string `json:"profile_image"`
		ThumbnailImageSmall      string `json:"thumbnail_image_small"`
		ThumbnailImageMedium     string `json:"thumbnail_image_medium"`
		ThumbnailImageBig        string `json:"thumbnail_image_big"`
		ThumbnailImageSmallLive  string `json:"thumbnail_image_small_live"`
		ThumbnailImageMediumLive string `json:"thumbnail_image_medium_live"`
		ThumbnailImageBigLive    string `json:"thumbnail_image_big_live"`
	} `json:"profile_images"`
}
