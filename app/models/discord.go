package models

// DiscordUser stores all information from a successful request to discord's API.
type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
}

// DiscordToken just stores the AccessToken.
type DiscordToken struct {
	AccessToken string
}
