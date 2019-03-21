package img

import (
	"angenalZZZ/go-program/api-svr/cors"
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

/**
Build and Run the Demo: nginx config
server {
        listen 80;
        server_name captcha.mojotv.cn;
        charset utf-8;

        location / {
            try_files /_not_exists_ @backend;
        }
        location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           pro=xy_set_header Host $http_host;
           proxy_pass http://127.0.0.1:8008;
        }
        access_log  /home/wwwlogs/captcha.mojotv.cn.log;
}
*/

// json request body
type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// create http handler
func CaptchaGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodGet, http.MethodPost}) {
		return
	}

	//output format
	outputJson := r.URL.Query().Get("dataType") == "json"

	//parse request parameters
	var postParameters ConfigJsonBody
	id := r.URL.Query().Get("id")
	if id == "" {
		id = r.URL.Query().Get("lastCode")
	}
	if id == "" && r.Method == http.MethodPost {
		defer r.Body.Close()
		if e := json.NewDecoder(r.Body).Decode(&postParameters); e != nil {
			FError(&w, id, e, outputJson)
			return
		}
	} else {
		postParameters = ConfigJsonBody{
			Id:          id,
			CaptchaType: getCaptchaType(r), //get query captchaType
			//VerifyValue: "",
			ConfigAudio: base64Captcha.ConfigAudio{CaptchaLen: 4, Language: "zh"},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height: 40,
				Width:  120,
				//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合
				Mode:               base64Captcha.CaptchaModeArithmetic,
				ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
				ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
				IsUseSimpleFont:    true,
				IsShowHollowLine:   false,
				IsShowNoiseDot:     false,
				IsShowNoiseText:    false,
				IsShowSlimeLine:    false,
				IsShowSineLine:     false,
				CaptchaLen:         6,
			},
			ConfigDigit: base64Captcha.ConfigDigit{
				Height:     35,
				Width:      70,
				CaptchaLen: 4,
				MaxSkew:    0.8,
				DotCount:   60,
			},
		}
	}

	//create base64 encoding captcha
	var config interface{} = postParameters.ConfigDigit
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	}
	captchaId, instance := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(instance)

	//or you can just write the captcha content to the httpResponseWriter.
	//before you put the captchaId into the response COOKIE.
	//instance.WriteTo(w)

	//set response
	FOk(&w, captchaId, base64blob, outputJson)
}

// verify http handler
func CaptchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodPost}) {
		return
	}

	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var postParameters ConfigJsonBody
	body := map[string]interface{}{"code": 1} // response error
	err := decoder.Decode(&postParameters)
	if err == nil {
		id, verifyValue := postParameters.Id, postParameters.VerifyValue
		if id == "" {
			id = r.URL.Query().Get("id")
			if id == "" {
				id = r.URL.Query().Get("lastCode")
			}
		}
		if id != "" || verifyValue != "" {
			//verify the captcha
			verifyResult := base64Captcha.VerifyCaptcha(id, verifyValue)

			//set response
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			if verifyResult {
				body = map[string]interface{}{"code": 0} // response ok
			}
		}
	}
	json.NewEncoder(w).Encode(body)
}

// response ok
func FOk(response *http.ResponseWriter, id string, data string, outputJson bool) {
	w := *response
	if outputJson == true {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		body := map[string]interface{}{"code": 0, "data": data, "captchaId": id, "msg": "success"}
		json.NewEncoder(w).Encode(body)
	} else {
		fmt.Fprint(w, data)
	}
}

// response error
func FError(response *http.ResponseWriter, id string, err error, outputJson bool) {
	w := *response
	//set json response
	w.WriteHeader(http.StatusAccepted)
	if outputJson == true {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		body := map[string]interface{}{"code": 1, "data": "", "captchaId": id, "msg": fmt.Sprintf("%v", err)}
		json.NewEncoder(w).Encode(body)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

// get query captchaType
func getCaptchaType(r *http.Request) (captchaType string) {
	captchaType = r.URL.Query().Get("captchaType")
	ok := false
	switch captchaType {
	case "audio":
	case "character":
	case "digit":
		ok = true
	}
	if ok == false {
		captchaType = "digit"
	}
	return
}
