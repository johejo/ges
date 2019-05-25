package model

import (
	"bytes"
	"encoding/json"
)

type Message struct {
	ID    string                 `json:"id"`
	Title map[string]interface{} `json:"title"`
	Text  map[string]interface{} `json:"text"`
}

func (m *Message) ToJSON() (*MessageJSON, error) {
	mj := &MessageJSON{
		ID: m.ID,
	}

	var err error
	mj.Title, err = json.Marshal(m.Title)
	if err != nil {
		return nil, err
	}
	mj.Text, err = json.Marshal(m.Text)
	if err != nil {
		return nil, err
	}

	return mj, nil
}

type MessageJSON struct {
	ID    string          `json:"id" db:"id"`
	Title json.RawMessage `json:"title" db:"title"`
	Text  json.RawMessage `json:"text" db:"text"`
}

func (mj *MessageJSON) ToMessage() (*Message, error) {
	m := &Message{
		ID: mj.ID,
	}

	if err := json.NewDecoder(bytes.NewReader(mj.Title)).Decode(&m.Title); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(bytes.NewReader(mj.Text)).Decode(&m.Text); err != nil {
		return nil, err
	}

	return m, nil
}
