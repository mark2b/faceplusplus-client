package faceplusplus_client

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

func (self *fppClient) DetectSkelethon(filename string) (response DetectSkeletonReponse, e error) {
	body := &bytes.Buffer{}
	if contentType, err := self.newRequestBodyWithInputFile(filename, body); err == nil {
		if request, err := http.NewRequest("POST", DETECT_SKELETON_URL, body); err == nil {
			if err := self.executeRequest(request, contentType, &response); err == nil {
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) DetectSkelethonWithImageData(imageData []byte) (response DetectSkeletonReponse, e error) {
	body := &bytes.Buffer{}
	imageDataEncoded := base64.StdEncoding.EncodeToString(imageData)
	if contentType, err := self.newRequestBodyWithParameters(map[string]string{"image_base64": imageDataEncoded}, body); err == nil {
		if request, err := http.NewRequest("POST", DETECT_SKELETON_URL, body); err == nil {
			if err := self.executeRequest(request, contentType, &response); err == nil {
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

type landmarkPoint struct {
	X     int     `json:"x",omitempty`
	Y     int     `json:"y",omitempty`
	Score float32 `json:"score",omitempty`
}

type DetectSkeletonReponse struct {
	TimeUsed  int    `json:"time_used",omitempty`
	ImageId   string `json:"image_id",omitempty`
	Skeletons []struct {
		BodyRectangle struct {
			Top    int `json:"top",omitempty`
			Left   int `json:"left",omitempty`
			Width  int `json:"width",omitempty`
			Height int `json:"height",omitempty`
		} `json:"body_rectangle",omitempty`
		Landmark struct {
			Head          landmarkPoint `json:"head",omitempty`
			Neck          landmarkPoint `json:"neck",omitempty`
			LeftElbow     landmarkPoint `json:"left_elbow",omitempty`
			LeftButtocks  landmarkPoint `json:"left_buttocks",omitempty`
			LeftShoulder  landmarkPoint `json:"left_shoulder",omitempty`
			LeftKnee      landmarkPoint `json:"left_knee",omitempty`
			LeftHand      landmarkPoint `json:"left_hand",omitempty`
			LeftFoot      landmarkPoint `json:"left_foot",omitempty`
			RightElbow    landmarkPoint `json:"right_elbow",omitempty`
			RightButtocks landmarkPoint `json:"right_buttocks",omitempty`
			RightShoulder landmarkPoint `json:"right_shoulder",omitempty`
			RightKnee     landmarkPoint `json:"right_knee",omitempty`
			RightHand     landmarkPoint `json:"right_hand",omitempty`
			RightFoot     landmarkPoint `json:"right_foot",omitempty`
		} `json:"landmark",omitempty`
	} `json:"skeletons",omitempty`
}
