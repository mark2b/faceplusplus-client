package faceplusplus_client

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

func (self *fppClient) DetectSkelethon(filename string) (response DetectHumanBodySkeletonReponse, e error) {
	body := &bytes.Buffer{}
	if contentType, err := self.newRequestBodyWithInputFile(filename, body); err == nil {
		if rsp, err := self.detectHumanBodySkeleton(body, contentType); err == nil {
			response = rsp
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) DetectSkeletonWithImageData(imageData []byte) (response DetectHumanBodySkeletonReponse, e error) {
	body := &bytes.Buffer{}
	imageDataEncoded := base64.StdEncoding.EncodeToString(imageData)
	if contentType, err := self.newRequestBodyWithParameters(map[string]string{"image_base64": imageDataEncoded}, body); err == nil {
		if rsp, err := self.detectHumanBodySkeleton(body, contentType); err == nil {
			response = rsp
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) detectHumanBodySkeleton(body *bytes.Buffer, contentType string) (response DetectHumanBodySkeletonReponse, e error) {
	if request, err := http.NewRequest("POST", DETECT_HUMAN_BODY_SKELETON_URL, body); err == nil {
		if err := self.executeRequest(request, contentType, &response); err == nil {
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

type HumanBodySkeletonLandmarkPoint struct {
	X     int     `json:"x",omitempty`
	Y     int     `json:"y",omitempty`
	Score float32 `json:"score",omitempty`
}

type HumanBodySkeleton struct {
	BodyRectangle struct {
		Top    int `json:"top",omitempty`
		Left   int `json:"left",omitempty`
		Width  int `json:"width",omitempty`
		Height int `json:"height",omitempty`
	} `json:"body_rectangle",omitempty`
	Landmark HumanBodySkeletonLandmark `json:"landmark",omitempty`
}

type HumanBodySkeletonLandmark struct {
	Head          HumanBodySkeletonLandmarkPoint `json:"head",omitempty`
	Neck          HumanBodySkeletonLandmarkPoint `json:"neck",omitempty`
	LeftElbow     HumanBodySkeletonLandmarkPoint `json:"left_elbow",omitempty`
	LeftButtocks  HumanBodySkeletonLandmarkPoint `json:"left_buttocks",omitempty`
	LeftShoulder  HumanBodySkeletonLandmarkPoint `json:"left_shoulder",omitempty`
	LeftKnee      HumanBodySkeletonLandmarkPoint `json:"left_knee",omitempty`
	LeftHand      HumanBodySkeletonLandmarkPoint `json:"left_hand",omitempty`
	LeftFoot      HumanBodySkeletonLandmarkPoint `json:"left_foot",omitempty`
	RightElbow    HumanBodySkeletonLandmarkPoint `json:"right_elbow",omitempty`
	RightButtocks HumanBodySkeletonLandmarkPoint `json:"right_buttocks",omitempty`
	RightShoulder HumanBodySkeletonLandmarkPoint `json:"right_shoulder",omitempty`
	RightKnee     HumanBodySkeletonLandmarkPoint `json:"right_knee",omitempty`
	RightHand     HumanBodySkeletonLandmarkPoint `json:"right_hand",omitempty`
	RightFoot     HumanBodySkeletonLandmarkPoint `json:"right_foot",omitempty`
}

type DetectHumanBodySkeletonReponse struct {
	TimeUsed  int                 `json:"time_used",omitempty`
	ImageId   string              `json:"image_id",omitempty`
	Skeletons []HumanBodySkeleton `json:"skeletons",omitempty`
}
