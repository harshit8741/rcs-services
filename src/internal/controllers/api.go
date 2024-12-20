package controllers

type Payload struct {
	MessageID  string   `json:"messageID"`
	AgentID    string   `json:"agentID"`
	CampaignID string   `json:"campaignID"`
	Contacts   []string `json:"contacts"`
	Data       Data     `json:"data"`
}

type Data struct {
	Content Content `json:"content"`
}

type Content struct {
	PlainText string `json:"plainText"`
}
