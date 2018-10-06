package translator

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

type IResponseHandler interface {
	HandleResponse(body io.ReadCloser) ([]TranslationResult, error)
}

func NewResponseHandler() IResponseHandler {
	return &ResponseHandler{}
}

type ResponseHandler struct {
}

func (r *ResponseHandler) HandleResponse(body io.ReadCloser) ([]TranslationResult, error) {

	bodyBytes, err := ioutil.ReadAll(body)
	body.Close()

	traslationProposals, err := r.parseJsonToProposalList(bodyBytes)
	if err != nil {
		return nil, err
	}

	return r.createTranslationResultFromProposal(traslationProposals), nil
}

func (self *ResponseHandler) parseJsonToProposalList(data []byte) ([]string, error) {

	type anyArray []interface{}

	var res []interface{} = make([]interface{}, 0)

	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		firstArray, ok := res[0].([]interface{})

		if ok && len(firstArray) > 0 {
			secondArray, ok := firstArray[0].([]interface{})

			if ok && len(secondArray) > 0 {

				if translation, ok := secondArray[0].(string); ok {
					return []string{translation}, nil
				}
			}
		}
	}
	return nil, errors.New("Unable to parse response")
}

func (self *ResponseHandler) createTranslationResultFromProposal(proposals []string) (results []TranslationResult) {
	for _, proposal := range proposals {
		results = append(results, TranslationResult{
			word: proposal,
			desc: "",
		})
	}
	return
}
