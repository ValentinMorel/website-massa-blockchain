package rpc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type GetDataEntriesObject struct {
	Address string `json:"address"`
	Key     []byte `json:"key"`
}

type GetDatastoreEntriesResponse struct {
	Result []struct {
		CandidateValue []byte `json:"candidate_value"`
		FinalValue     []byte `json:"final_value"`
	} `json:"result"`
}

func (r *Client) Get(address string, key string, prefix string) (*GetDatastoreEntriesResponse, error) {
	resp, err := r.getDatastoreEntries(address, key, prefix)
	if err != nil {
		return nil, err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	var responseStruct GetDatastoreEntriesResponse
	json.Unmarshal(response, &responseStruct)
	return &responseStruct, nil
}

func (r *Client) ResolveDns(address string, key string, prefix string) (*GetDatastoreEntriesResponse, error) {
	resp, err := r.getDatastoreEntries(address, key, prefix)
	if err != nil {
		return nil, err
	}
	response, _ := ioutil.ReadAll(resp.Body)
	var responseStruct GetDatastoreEntriesResponse
	json.Unmarshal(response, &responseStruct)
	return &responseStruct, nil
}

func (r *Client) getDatastoreEntries(address string, key string, prefix string) (*http.Response, error) {

	keySequence := []byte(prefix + key)

	request := `{"jsonrpc": "2.0", "id": "0", "method": "get_datastore_entries", "params": [[{"address":"` + address + `", "key": [`
	for i := 0; i <= len(keySequence)-1; i++ {
		if i != len(keySequence)-1 {
			request = request + strconv.Itoa(int(keySequence[i])) + ","
		} else {
			request = request + strconv.Itoa(int(keySequence[i]))
		}
	}
	request = request + `]}]]}`
	resp, err := http.Post(r.Url, "application/json", strings.NewReader(request))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
