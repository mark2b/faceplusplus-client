package faceplusplus_client

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

func (self *fppClient) DetectHumanBodySegment(filename string) (response DetectHumanBodySegmentReponse, e error) {
	body := &bytes.Buffer{}
	if contentType, err := self.newRequestBodyWithInputFile(filename, body); err == nil {
		if rsp, err := self.detectHumanBodySegment(body, contentType); err == nil {
			response = rsp
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) DetectHumanBodySegmentWithImageData(imageData []byte) (response DetectHumanBodySegmentReponse, e error) {
	body := &bytes.Buffer{}
	imageDataEncoded := base64.StdEncoding.EncodeToString(imageData)
	if contentType, err := self.newRequestBodyWithParameters(map[string]string{"image_base64": imageDataEncoded}, body); err == nil {
		if rsp, err := self.detectHumanBodySegment(body, contentType); err == nil {
			response = rsp
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) detectHumanBodySegment(body *bytes.Buffer, contentType string) (response DetectHumanBodySegmentReponse, e error) {
	if request, err := http.NewRequest("POST", DETECT_HUMAN_BODY_SEGMENT_URL, body); err == nil {
		if err := self.executeRequest(request, contentType, &response); err == nil {
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

type DetectHumanBodySegmentReponse struct {
	TimeUsed     int    `json:"time_used",omitempty`
	ImageId      string `json:"image_id",omitempty`
	Result       string `json:"result",omitempty`
	BodyImage    string `json:"body_image",omitempty`
	ErrorMessage string `json:"error_message",omitempty`
}
