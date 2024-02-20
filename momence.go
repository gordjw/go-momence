package momence

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Momence struct {
	hostId  string
	token   string
	baseUrl string
}

type MomenceEvent struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	DateTime        string `json:"dateTime"`
	Duration        int    `json:"duration"`
	OriginalTeacher string `json:"originalTeacher"`
}

type MomenceTeacher struct {
	Id           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profileImage"`
	IsDeleted    bool   `json:"isDeleted"`
}

var (
	teachersEndpoint = "/Teachers"
	eventsEndpoint   = "/Events"
)

func NewMomence(hostId, token, baseUrl string) Momence {
	return Momence{
		hostId:  hostId,
		token:   token,
		baseUrl: baseUrl,
	}
}

func (m *Momence) buildApiUrl(endpoint string) string {
	return fmt.Sprintf("%s%s?hostId=%s&token=%s",
		m.baseUrl,
		endpoint,
		m.hostId,
		m.token,
	)
}

func (m *Momence) GetTeachers() ([]MomenceTeacher, error) {
	url := m.buildApiUrl(teachersEndpoint)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving data from %s: %v\n", url, err.Error())
		return []MomenceTeacher{}, err
	}

	var teacherList []MomenceTeacher
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&teacherList)
	if err != nil {
		fmt.Printf("Error decoding class list: %v\n", err.Error())
		return []MomenceTeacher{}, err
	}

	return teacherList, nil
}

func (m *Momence) GetEvents() ([]MomenceEvent, error) {
	url := m.buildApiUrl(eventsEndpoint)
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving data from %s: %v\n", url, err.Error())
		return []MomenceEvent{}, err
	}

	var eventList []MomenceEvent
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&eventList)
	if err != nil {
		fmt.Printf("MomenceEvent decoding class list: %v\n", err.Error())
		return []MomenceEvent{}, err
	}

	return eventList, nil
}
