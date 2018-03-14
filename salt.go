package go_salt

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

type SaltTokenJson struct {
	Return []struct {
		Perms  []string `json:"perms"`
		Start  float64 `json:"start"`
		Token  string `json:"token"`
		Expire float64 `json:"expire"`
		User   string `json:"user"`
		Eauth  string `json:"eauth"`
	} `json:"return"`
}


type SaltHostJson struct {
	Return []struct {
		Code02MorepayCn struct {
			Kernelrelease string `json:"kernelrelease"`
			Selinux       struct {
				Enforced string `json:"enforced"`
				Enabled  bool   `json:"enabled"`
			} `json:"selinux"`
			Serialnumber     string `json:"serialnumber"`
			MemTotal         int    `json:"mem_total"`
			Saltversioninfo  []int  `json:"saltversioninfo"`
			Host             string `json:"host"`
			ID               string `json:"id"`
			Osrelease        string `json:"osrelease"`
			NumCpus          int    `json:"num_cpus"`
			HwaddrInterfaces struct {
				Eth0 string `json:"eth0"`
			} `json:"hwaddr_interfaces"`
			IP4Interfaces struct {
				Eth0 []string `json:"eth0"`
			} `json:"ip4_interfaces"`
			Init          string   `json:"init"`
			LsbDistribID  string   `json:"lsb_distrib_id"`
			FqdnIP4       []string `json:"fqdn_ip4"`
			Saltversion   string   `json:"saltversion"`
			ServerID      int      `json:"server_id"`
			Oscodename    string   `json:"oscodename"`
			Osfinger      string   `json:"osfinger"`
			Virtual       string   `json:"virtual"`
			Productname   string   `json:"productname"`
			OsreleaseInfo []int    `json:"osrelease_info"`
			Os            string   `json:"os"`
		} `json:"code02.morepay.cn"`
	} `json:"return"`
}

type Salt struct {
	Config map[string]string //配置文件
	Info   map[string]string
}

func (S *Salt)GET_TOKEN() (string, error) {
	SALT_LOGIN_URL := "http://"+S.Config["salt_addr"] +":"+S.Config["salt_prot"]+ "/login"
	client := &http.Client{}
	post := fmt.Sprintf(`{"eauth": "%s", "username": "%s", "password": "%s"}`, "pam", S.Config["salt_user"], S.Config["salt_passwd"])

	req, err := http.NewRequest("POST", SALT_LOGIN_URL, bytes.NewBuffer([]byte(post)))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Auth-Token", "")
	req.Header.Set("Content-Type", "application/json")
	resp, client_err := client.Do(req)
	if client_err != nil {
		return "", client_err
	}
	defer resp.Body.Close()
	body, ioutil_err := ioutil.ReadAll(resp.Body)
	if ioutil_err != nil {
		return "", ioutil_err
	}
	var JsonRes SaltTokenJson
	json.Unmarshal([]byte(string(body)), &JsonRes)
	return JsonRes.Return[0].Token, nil
}

func (S *Salt)CMD_SALT(post_data string) (string, error) {
	//SALT_CONGIF := SALT_CONFIG()
	token, get_token_err := S.GET_TOKEN()
	if get_token_err != nil {
		return "", get_token_err
	} else {
		client := &http.Client{}
		req1, http_err := http.NewRequest("POST", "http://"+S.Config["salt_addr"] +":"+S.Config["salt_prot"], bytes.NewBuffer([]byte(post_data)))
		if http_err != nil {
			return "", http_err
		}
		req1.Header.Set("X-Auth-Token", token)
		req1.Header.Set("Content-Type", "application/json")
		resp1, client_err := client.Do(req1)
		if client_err != nil {
			return "", client_err
		}
		defer resp1.Body.Close()
		body1, ioutil_err := ioutil.ReadAll(resp1.Body)
		if ioutil_err != nil {
			return "", ioutil_err
		}

		return string(body1), nil
	}
}

