package handler

import (
	"github.com/ngonghi/VyLog/common"
	"encoding/json"
	"net/http"
	"net/url"
	"errors"
)

const (
	SLACK_COLOR_DANGER = "danger"
	SLACK_COLOR_WARNING = "warning"
	SLACK_COLOR_GOOD = "good"
	SLACK_COLOR_DEFAULT = "#e3e4e6"
)

type SlackHandler struct {
	AbstractHandler

	channel string
	hookUrl string
	username string
	preText string
}

func GetSlackHandler(channel string,hookUrl string,username string, preText string) *SlackHandler {
	h := &SlackHandler{
		channel: channel,
		hookUrl: hookUrl,
		username: username,
		preText: preText,
	}

	return h
}

func (handler *SlackHandler) getColor(level int) string {
	switch level {
	case common.INFO:
		return SLACK_COLOR_GOOD
	case common.WARN:
		return SLACK_COLOR_WARNING
	case common.ERROR, common.FATAL:
		return SLACK_COLOR_DANGER
	default:
		return SLACK_COLOR_DEFAULT
	}
}

func (handler *SlackHandler) Handle(message *common.Message) error {
	if handler.IsHandling(message) {
		err := handler.Post(message)

		return err
	}

	return nil
}

func (handler * SlackHandler) Post(msg *common.Message) error {

	msgContent,err := handler.createSlackMsgContent(msg)

	if err != nil {
		return err
	}

	resp, err := http.PostForm(handler.hookUrl, url.Values{"payload": {msgContent}})

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Send Slack Message Fail")
	}

	return nil
}

func (handler * SlackHandler) createSlackMsgContent(msg *common.Message) (string,error) {

	var fields []interface{}
	fields = append(fields,map[string]string{
		"short" : "false",
	})

	var attachments []interface{}
	attachment_content := make(map[string]interface{})
	attachment_content["pretext"] = handler.preText
	attachment_content["color"] = handler.getColor(msg.Level)
	attachment_content["text"] = handler.GetFormatter().Format(msg)
	attachment_content["fields"] = fields
	attachments = append(attachments,attachment_content)

	slackMsgContent := make(map[string]interface{})
	slackMsgContent["channel"] = handler.channel
	slackMsgContent["username"] = handler.username
	slackMsgContent["attachments"] = attachments

	b,err := json.Marshal(slackMsgContent)

	return string(b),err
}