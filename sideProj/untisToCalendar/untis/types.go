package untis

// Lesson represents a single lesson/event from WebUntis timetable
type Lesson struct {
	ID           int    `json:"id"`
	Date         int    `json:"date"`         // YYYYMMDD format (e.g., 20260526)
	StartTime    int    `json:"startTime"`    // HHMM format (e.g., 1055 = 10:55)
	EndTime      int    `json:"endTime"`      // HHMM format (e.g., 1140 = 11:40)
	Code         string `json:"code"`         // "irregular", "cancelled", etc
	ActivityType string `json:"activityType"` // "Unterricht", "Bereitschaft", etc
	Classes      []struct {
		ID int `json:"id"`
	} `json:"kl"`
	Teachers []struct {
		ID    int `json:"id"`
		OrgID int `json:"orgid,omitempty"`
	} `json:"te"`
	Subjects []struct {
		ID int `json:"id"`
	} `json:"su"`
	Rooms []struct {
		ID    int `json:"id"`
		OrgID int `json:"orgid,omitempty"`
	} `json:"ro"`
}

// TimetableResponse is the JSON-RPC response from WebUntis
type TimetableResponse struct {
	JsonRPC string   `json:"jsonrpc"`
	ID      string   `json:"id"`
	Result  []Lesson `json:"result"`
}

// Subject represents a subject entry in WebUntis
type Subject struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	LongName string `json:"longName,omitempty"`
}

// Teacher represents a teacher entry in WebUntis
type Teacher struct {
	ID           int    `json:"id"`
	Name         string `json:"name,omitempty"`
	LongName     string `json:"longName,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
}
