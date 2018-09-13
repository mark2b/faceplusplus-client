package faceplusplus_client

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"strings"
)

func (self *fppClient) DetectBody(filename string, returnAttributes []string) (response DetectBodyReponse, e error) {
	body := &bytes.Buffer{}
	returnAttributesField := strings.Join(returnAttributes, ",")
	if contentType, err := self.newRequestBodyWithParametersAndInputFile(map[string]string{"return_attributes": returnAttributesField}, filename, body); err == nil {
		if request, err := http.NewRequest("POST", DETECT_BODY_URL, body); err == nil {
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

func (self *fppClient) DetectBodyWithImageData(imageData []byte) (response DetectBodyReponse, e error) {
	body := &bytes.Buffer{}
	imageDataEncoded := base64.StdEncoding.EncodeToString(imageData)
	if contentType, err := self.newRequestBodyWithParameters(map[string]string{"image_base64": imageDataEncoded}, body); err == nil {
		if request, err := http.NewRequest("POST", DETECT_BODY_URL, body); err == nil {
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

type DetectBodyReponse struct {
	TimeUsed    int    `json:"time_used",omitempty`
	ImageId     string `json:"image_id",omitempty`
	HumanBodies []struct {
		Attributes struct {
			Gender struct {
				Confidence float32 `json:"confidence",omitempty`
				Value      string  `json:"value",omitempty`
			} `json:"gender",omitempty`
			UpperBodyCloth struct {
				UpperBodyClothColor    string `json:"upper_body_cloth_color",omitempty`
				UpperBodyClothColorRGB string `json:"upper_body_cloth_color_rgb",omitempty`
			} `json:"upper_body_cloth",omitempty`
			LowerBodyCloth struct {
				LowerBodyClothColor    string `json:"lower_body_cloth_color",omitempty`
				LowerBodyClothColorRGB string `json:"lower_body_cloth_color_rgb",omitempty`
			} `json:"lower_body_cloth",omitempty`
		} `json:"attributes",omitempty`
		HumanBodyRectangle struct {
			Top    int `json:"top",omitempty`
			Left   int `json:"left",omitempty`
			Width  int `json:"width",omitempty`
			Height int `json:"height",omitempty`
		} `json:"humanbody_rectangle",omitempty`
		Confidence float32 `json:"confidence",omitempty`
	} `json:"humanbodies",omitempty`
}
