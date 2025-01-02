package controllers

// Base struct for common fields across all payloads
type BasePayload struct {
	MessageID  string   `json:"messageID"`
	AgentID    string   `json:"agentID"`
	CampaignID string   `json:"campaignID,omitempty"` // Optional field
	Contacts   []string `json:"contacts"`
	DataSMS    *DataSMS `json:"data_sms,omitempty"` // Optional for APIs with SMS data
}

// Payload Structs for Each API
type SuggestedDailerPayload struct {
	BasePayload
	Data RichCardData `json:"data"`
}

type SimpleTextMessagePayload struct {
	BasePayload
	Data TextData `json:"data"`
}

type SimpleImgMessagePayload struct {
	BasePayload
	Data ImgData `json:"data"`
}

type SuggestedReplyPayload struct {
	BasePayload
	Data RichCardData `json:"data"`
}

type SuggestedCalenderPayload struct {
	BasePayload
	Data RichCardData `json:"data"`
}

type SuggestedLocationPayload struct {
	BasePayload
	Data RichCardData `json:"data"`
}

// Payload-Specific Data Chains
type TextData struct {
	Content TextContent `json:"content"`
}

type ImgData struct {
	Content ImgContent `json:"content"`
}

type RichCardData struct {
	Content RichCardContent `json:"content"`
}

// Content Structs
type TextContent struct {
	PlainText string `json:"plainText"`
}

type ImgContent struct {
	RichCardDetails RichCardDetails `json:"richCardDetails"`
}

type RichCardContent struct {
	RichCardDetails RichCardDetails `json:"richCardDetails"`
}

// Rich Card Details
type RichCardDetails struct {
	Standalone Standalone `json:"standalone"`
}

// Standalone
type Standalone struct {
	CardOrientation string      `json:"cardOrientation"`
	Content         CardContent `json:"content"`
}

// Card Content
type CardContent struct {
	CardTitle       string       `json:"cardTitle"`
	CardDescription string       `json:"cardDescription"`
	CardMedia       CardMedia    `json:"cardMedia"`
	Suggestions     []Suggestion `json:"suggestions"`
}

// Card Media
type CardMedia struct {
	MediaHeight string      `json:"mediaHeight"`
	ContentInfo ContentInfo `json:"contentInfo"`
}

// Content Info
type ContentInfo struct {
	FileURL string `json:"fileUrl"`
}

// Suggestion
type Suggestion struct {
	Action *Action `json:"action,omitempty"`
	Reply  *Reply  `json:"reply,omitempty"`
}

// Action
type Action struct {
	PlainText           string               `json:"plainText,omitempty"`
	PostBack            *PostBack            `json:"postBack,omitempty"`
	OpenUrl             *OpenUrl             `json:"openUrl,omitempty"`
	CreateCalendarEvent *CreateCalendarEvent `json:"createCalendarEvent,omitempty"`
	ShowLocation        *ShowLocation        `json:"showLocation,omitempty"`
	DialerAction        *DialerAction        `json:"dialerAction,omitempty"`
}

// Dailer
type DialerAction struct {
	PhoneNumber string `json:"phoneNumber"`
}

// PostBack
type PostBack struct {
	Data string `json:"data"`
}

// OpenUrl
type OpenUrl struct {
	URL string `json:"url"`
}

// Reply
type Reply struct {
	PlainText string   `json:"plainText"`
	PostBack  PostBack `json:"postBack"`
}

// Create Calendar Event
type CreateCalendarEvent struct {
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Show Location
type ShowLocation struct {
	Coordinates Coordinates `json:"coordinates"`
	Label       string      `json:"label"`
}

// Coordinates
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// SMS Data (Optional for APIs)
type DataSMS struct {
	SenderID       string `json:"sender_id"`
	DomainID       string `json:"domain_id"`
	SMSType        string `json:"sms_type"`
	SMSContentType string `json:"sms_content_type"`
	DLTEntityID    string `json:"dlt_entity_id"`
	Body           string `json:"body"`
	DLTTemplateID  string `json:"dlt_template_id"`
}
