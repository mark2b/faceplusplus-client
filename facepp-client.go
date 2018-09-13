package faceplusplus_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func (self *fppClient) executeRequest(request *http.Request, contentType string, response interface{}) (e error) {
	request.Header.Set("Content-Type", contentType)
	if httpResponse, err := self.httpClient.Do(request); err == nil {
		if httpResponse.StatusCode == 200 {
			if responsePayload, err := ioutil.ReadAll(httpResponse.Body); err == nil {
				if err := json.Unmarshal(responsePayload, &response); err == nil {
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			if responsePayload, err := ioutil.ReadAll(httpResponse.Body); err == nil {
				e = errors.New(fmt.Sprintf("%s (%s)", httpResponse.Status, string(responsePayload)))
			} else {
				e = errors.New(httpResponse.Status)
			}
		}
	} else {
		e = err
	}
	return
}

func (self *fppClient) newMultipartWriterWithParameters(parameters map[string]string, writer io.Writer) (mpw *multipart.Writer, e error) {
	mpw = multipart.NewWriter(writer)
	parameters["api_key"] = self.ApiKey
	parameters["api_secret"] = self.ApiSecret
	for k, v := range parameters {
		if err := mpw.WriteField(k, v); err != nil {
			e = err
			break
		}
	}
	return
}

func (self *fppClient) newRequestBodyWithInputFile(filename string, bodyWriter io.Writer) (contentType string, e error) {
	return self.newRequestBodyWithParametersAndInputFile(map[string]string{}, filename, bodyWriter)
}

func (self *fppClient) newRequestBodyWithParametersAndInputFile(parameters map[string]string, filename string, bodyWriter io.Writer) (contentType string, e error) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if writer, err := self.newMultipartWriterWithParameters(parameters, bodyWriter); err == nil {
			if part, err := writer.CreateFormFile("image_file", "image.jpg"); err == nil {
				if _, err = io.Copy(part, file); err == nil {
					if err := writer.Close(); err == nil {
						contentType = writer.FormDataContentType()
					} else {
						e = err
					}
				} else {
					e = err
				}
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

func (self *fppClient) newRequestBodyWithParameters(parameters map[string]string, bodyWriter io.Writer) (contentType string, e error) {
	if writer, err := self.newMultipartWriterWithParameters(parameters, bodyWriter); err == nil {
		if err := writer.Close(); err == nil {
			contentType = writer.FormDataContentType()
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

//func (self *fppClient) newRequestBodyWithParametersAndInputImageData(parameters map[string]string, imageData []byte, bodyWriter io.Writer) (contentType string, e error) {
//	imageDataEncoded := base64.StdEncoding.EncodeToString(imageData)
//	if writer, err := self.newMultipartWriterWithParameters(parameters, bodyWriter); err == nil {
//		if part, err := writer.CreateFormFile("image_file", "image.jpg"); err == nil {
//			if _, err = io.Copy(part, file); err == nil {
//				if err := writer.Close(); err == nil {
//					contentType = writer.FormDataContentType()
//				} else {
//					e = err
//				}
//			} else {
//				e = err
//			}
//		} else {
//			e = err
//		}
//	} else {
//		e = err
//	}
//	return
//}

type fppClient struct {
	ApiKey     string
	ApiSecret  string
	httpClient *http.Client
}

func NewClient(apiKey string, apiSecret string) (client *fppClient) {
	client = new(fppClient)
	client.ApiKey = apiKey
	client.ApiSecret = apiSecret
	client.httpClient = &http.Client{}
	return
}

const (
	DETECT_BODY_URL     = "https://api-us.faceplusplus.com/humanbodypp/beta/detect"
	DETECT_SKELETON_URL = "https://api-us.faceplusplus.com/humanbodypp/v1/skeleton"
)
