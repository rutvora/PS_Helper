package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func checkErrors(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//func generatePostRequestData() url.Values {
//	loginRequest := url.Values{}
//	loginRequest.Set("__EVENTTARGET", "")
//	loginRequest.Set("__EVENTARGUMENT", "")
//	loginRequest.Set("__VIEWSTATE", "/wEPDwULLTE1NjMxNjMxNzFkZCo3T3kAnddTDFryr26qaofiTp5p")
//	loginRequest.Set("__VIEWSTATEGENERATOR", "C2EE9ABB")
//	loginRequest.Set("__EVENTVALIDATION", "/wEdAAYNcEy/uvEwBm4by+oKLWkjSvD5Cbpu3w0ab2H9f5rbFEPTPkdPWl+8YN2NtDCtxifN+DvxnwFeFeJ9MIBWR693w+qCzNvQHKCQwl8+YzOKE62xJNKuHibH70Ul6qoa4F8sDaR1uxEyo1xbP9xcXI4vvNcYtQ==")
//	loginRequest.Set("TxtEmail", "id")
//	loginRequest.Set("txtPass", "pass")
//	loginRequest.Set("Button1", "Login")
//	loginRequest.Set("txtEmailId", "")
//	return loginRequest
//}

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,gu;q=0.8,hi;q=0.7")
	//req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "20")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", "_ga=GA1.3.236302189.1557160384; _fbp=fb.2.1557160390387.730603858; ASP.NET_SessionId=sdsu3el5cdfybni3kdqimjbm")
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

func findInMapArray(mapArray []map[string]interface{}, key string, value interface{}) map[string]interface{} {
	for _, object := range mapArray {
		if val, ok := object[key]; ok && val == value {
			return object
		}
	}
	return nil
}

func writeCSV(stationList, problemBank []map[string]interface{}) {
	//Open file for csv
	csvFile, err := os.Create("StationDetails.csv")
	checkErrors(err)
	csvWriter := csv.NewWriter(csvFile)

	csvData := make([]string, 8)
	//Head
	csvData[0] = "Station ID"
	csvData[1] = "Company Name"
	csvData[2] = "Location"
	csvData[3] = "Industry Domain"
	csvData[4] = "Preferred Branches"
	csvData[5] = "Stipend (UG)"
	csvData[6] = "Stipend (PG)"
	csvData[7] = "Have Accommodation?"
	err = csvWriter.Write(csvData)
	checkErrors(err)

	for _, object := range stationList {
		problemBankCounterpart := findInMapArray(problemBank, "StationId", object["StationId"])
		if problemBankCounterpart != nil {
			csvData[0] = fmt.Sprintf("%v", object["StationId"])
			csvData[1] = fmt.Sprintf("%v", problemBankCounterpart["CompanyName"])
			csvData[2] = fmt.Sprintf("%v", problemBankCounterpart["City"])
			csvData[3] = fmt.Sprintf("%v", problemBankCounterpart["IndustryDomain"])
			csvData[4] = fmt.Sprintf("%v", problemBankCounterpart["Tags"])
			csvData[5] = fmt.Sprintf("%v", problemBankCounterpart["stipend"])
			csvData[6] = fmt.Sprintf("%v", problemBankCounterpart["stipendforpg"])
			csvData[7] = "No"

		} else {

			temp := strings.Split(fmt.Sprintf("%v", object["Companyname"]), "-")
			companyDomain := strings.TrimSpace(temp[0])
			companyName := strings.TrimSpace(temp[1])

			csvData[0] = fmt.Sprintf("%v", object["StationId"])
			csvData[1] = companyName
			csvData[2] = fmt.Sprintf("%v", object["City"])
			csvData[3] = companyDomain
			csvData[4] = "Unavailable"
			csvData[5] = "Unavailable"
			csvData[6] = "Unavailable"
			csvData[7] = "No"
		}
		err = csvWriter.Write(csvData)
		checkErrors(err)
	}
}

func getData(url string, data string) []map[string]interface{} {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	checkErrors(err)
	setHeaders(req)

	client := &http.Client{}
	checkErrors(err)
	resp, err := client.Do(req)
	checkErrors(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalln("HTTP Response code: " + string(resp.StatusCode))
	}
	return decodeJSON(&resp.Body)
}

func getUpdateJSON() []byte {
	csvFile, err := os.Open("StationDetails.csv")
	checkErrors(err)
	csvReader := csv.NewReader(csvFile)

	var finalPostData = make(map[string]interface{}, 3)
	finalPostData["jsonvalue"] = ""
	finalPostData["contistation"] = "0"

	var preferenceListArr []map[string]string

	record, err := csvReader.Read() //First Row has titles

	//Make JSONArray for sending
	var i int64 = 1
	for {
		var StationEntry = make(map[string]string, 4)
		record, err = csvReader.Read()
		if err != io.EOF {
			checkErrors(err)
		} else {
			break
		}
		StationEntry["isActive"] = "1"
		StationEntry["PreferenceNo"] = strconv.FormatInt(i, 10)
		i += 1
		StationEntry["StationId"] = record[0]
		var accommodation string
		if string(record[7][0]) == "Y" || string(record[7][0]) == "y" {
			accommodation = "true"
		} else {
			accommodation = "false"
		}
		StationEntry["Accomodation"] = accommodation

		preferenceListArr = append(preferenceListArr, StationEntry)

	}

	updateJSON, err := json.Marshal(preferenceListArr)
	encoder := json.NewEncoder(os.Stderr)
	encoder.SetEscapeHTML(true)
	encoder.Encode()
	checkErrors(err)
	return updateJSON
	//fmt.Println(preferenceListArr)
}

func main() {
	//Create CSV
	//stationList := getData("http://psd.bits-pilani.ac.in/Student/StudentStationPreference.aspx/getinfoStation", "{CompanyId: \"0\" }")
	//problemBank := getData("http://psd.bits-pilani.ac.in/Student/ViewActiveStationProblemBankData.aspx/getPBdetail", "{batchid: \"undefined\" }")
	//writeCSV(stationList, problemBank)

	updateList()

	//Update pref list on website
}
