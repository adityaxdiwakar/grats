package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MultiGroupApiResponseObject struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response []GroupChat `json:"response"`
}

type GroupApiResponseObject struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response GroupChat `json:"response"`
}

type GroupChat struct {
	CreatedAt     int    `json:"created_at"`
	CreatorUserID string `json:"creator_user_id"`
	Description   string `json:"description"`
	GroupID       string `json:"group_id"`
	ID            string `json:"id"`
	ImageURL      string `json:"image_url"`
	MaxMembers    int    `json:"max_members"`
	Members       []struct {
		Autokicked bool     `json:"autokicked"`
		ID         string   `json:"id"`
		ImageURL   string   `json:"image_url"`
		Muted      bool     `json:"muted"`
		Name       string   `json:"name"`
		Nickname   string   `json:"nickname"`
		Roles      []string `json:"roles"`
		UserID     string   `json:"user_id"`
	} `json:"members"`
	Messages struct {
		Count                int    `json:"count"`
		LastMessageCreatedAt int    `json:"last_message_created_at"`
		LastMessageID        string `json:"last_message_id"`
		Preview              struct {
			Attachments []interface{} `json:"attachments"`
			ImageURL    string        `json:"image_url"`
			Nickname    string        `json:"nickname"`
			Text        string        `json:"text"`
		} `json:"preview"`
	} `json:"messages"`
	Name           string `json:"name"`
	OfficeMode     bool   `json:"office_mode"`
	PhoneNumber    string `json:"phone_number"`
	ShareQrCodeURL string `json:"share_qr_code_url"`
	ShareURL       string `json:"share_url"`
	Type           string `json:"type"`
	UpdatedAt      int    `json:"updated_at"`
}

func GetGroupInformation(GroupID int) (GroupChat, error) {
	url := fmt.Sprintf("https://api.groupme.com/v3/groups/%d?token=%s", groupID, os.Getenv("ACCESS_TOKEN"))
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		err := fmt.Sprintf("GM-API did not return 200 response, instead returned %d", resp.StatusCode)
		return GroupChat{}, errors.New(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	GroupReturn := GroupApiResponseObject{}
	json.Unmarshal([]byte(body), &GroupReturn)

	return GroupReturn.Response, nil
}

func GetGroupListing() ([]GroupChat, error) {
	GroupChats := []GroupChat{}
	for index := 1; true; index++ {
		url := fmt.Sprintf("https://api.groupme.com/v3/groups?per_page=100&token=%s&page=%d", os.Getenv("ACCESS_TOKEN"), index)
		resp, _ := http.Get(url)
		body, _ := ioutil.ReadAll(resp.Body)

		GroupResp := MultiGroupApiResponseObject{}
		json.Unmarshal([]byte(body), &GroupResp)
		GroupChats = append(GroupChats, GroupResp.Response...)

		if len(GroupResp.Response) < 100 {
			return GroupChats, nil
		}

		data, _ := json.MarshalIndent(GroupResp, "", "    ")
		_ = ioutil.WriteFile("test.json", data, 0644)
	}
	return GroupChat{}, nil
}
