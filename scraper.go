package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type preferenceList struct {
	Sno         int
	City        string
	Companyname string
	StationId   int
	CompanyId   int
}

func checkErrors(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func generatePostRequestData() url.Values {
	loginRequest := url.Values{}
	loginRequest.Set("__EVENTTARGET", "")
	loginRequest.Set("__EVENTARGUMENT", "")
	loginRequest.Set("__VIEWSTATE", "/wEPDwULLTE1NjMxNjMxNzFkZCo3T3kAnddTDFryr26qaofiTp5p")
	loginRequest.Set("__VIEWSTATEGENERATOR", "C2EE9ABB")
	loginRequest.Set("__EVENTVALIDATION", "/wEdAAYNcEy/uvEwBm4by+oKLWkjSvD5Cbpu3w0ab2H9f5rbFEPTPkdPWl+8YN2NtDCtxifN+DvxnwFeFeJ9MIBWR693w+qCzNvQHKCQwl8+YzOKE62xJNKuHibH70Ul6qoa4F8sDaR1uxEyo1xbP9xcXI4vvNcYtQ==")
	loginRequest.Set("TxtEmail", "id")
	loginRequest.Set("txtPass", "pass")
	loginRequest.Set("Button1", "Login")
	loginRequest.Set("txtEmailId", "")
	return loginRequest
}

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,gu;q=0.8,hi;q=0.7")
	//req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "20")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", "ASP.NET_SessionId=sdsu3el5cdfybni3kdqimjbm")
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "psd.bits-pilani.ac.in")
	req.Header.Set("Origin", "http://psd.bits-pilani.ac.in")
	req.Header.Set("Referer", "ttp://psd.bits-pilani.ac.in/Student/ViewActiveStationProblemBankData.aspx")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Mobile Safari/537.36)")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

}

func decodeJSON(response *io.ReadCloser) []map[string]interface{} {
	decoder := json.NewDecoder(*response)
	var array map[string]string
	err := decoder.Decode(&array)
	checkErrors(err)
	decoder = json.NewDecoder(strings.NewReader(array["d"]))
	var dataArray []map[string]interface{}
	err = decoder.Decode(&dataArray)
	checkErrors(err)

	return dataArray
}

func writeCSV(dataArray []map[string]interface{}) {
	for _, object := range dataArray {
		for key, value := range object {
			fmt.Println(key + ": " + fmt.Sprintf("%v", value))
		}
		fmt.Println()
	}
}

func main() {
	//loginRequest := generatePostRequestData()
	//req, err := http.NewRequest("POST", "http://psd.bits-pilani.ac.in/Student/ViewActiveStationProblemBankData.aspx/getPBPOPUP", strings.NewReader("{StationId: \"3514\" }"))
	req, err := http.NewRequest("POST", "http://psd.bits-pilani.ac.in/Student/StudentStationPreference.aspx/getinfoStation", strings.NewReader("{CompanyId: \"0\" }"))
	setHeaders(req)

	client := &http.Client{}
	checkErrors(err)
	resp, err := client.Do(req)
	checkErrors(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalln("HTTP Response code: " + string(resp.StatusCode))
	}
	dataArray := decodeJSON(&resp.Body)
	writeCSV(dataArray)

	//bytearr, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bytearr))
	//responseBytes, err := ioutil.ReadAll(resp.Body)
	//checkErrors(err)
	//fmt.Println(string(responseBytes))
}
