package thaiaddressdatabase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type SubDistrict struct {
	id              int
	code            int
	name_in_thai    string
	name_in_english string
	latitude        float64
	longitude       float64
	district_id     int
	zip_code        int
}
type District struct {
	id              int
	code            int
	name_in_thai    string
	name_in_english string
	province_id     int
}

type Province struct {
	id              int
	code            int
	name_in_thai    string
	name_in_english string
}

type Address struct {
	SubDistricts []SubDistrict
	Districts    []District
	Provinces    []Province
}

func LoadJSONData(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read data: %v", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}

func Init() (*Address, error) {
	var address Address
	if err := LoadJSONData("data/province.json", &address.Provinces); err != nil {
		return nil, err
	}
	if err := LoadJSONData("data/district.json", &address.Districts); err != nil {
		return nil, err
	}
	if err := LoadJSONData("data/sub_district.json", &address.SubDistricts); err != nil {
		return nil, err
	}
	return &address, nil
}

func (a *Address) GetProvince(provinceId string, lang string) string {

	for _, p := range a.Provinces {
		if fmt.Sprint(p.id) == provinceId {
			if lang == "th" {
				return p.name_in_thai
			} else {
				return p.name_in_english
			}
		}
	}

	return ""
}

func (a *Address) GetDistrict(districtId string, lang string) string {

	for _, d := range a.Districts {
		if fmt.Sprint(d.id) == districtId {
			if lang == "th" {
				return d.name_in_thai
			} else {
				return d.name_in_english
			}
		}
	}

	return ""
}

func (a *Address) GetSubDistrict(subDistrictId string, lang string) string {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.id) == subDistrictId {
			if lang == "th" {
				return s.name_in_thai
			} else {
				return s.name_in_english
			}
		}
	}

	return ""
}

func (a *Address) GetSubDistrictByZipCode(zipCode string, lang string) string {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.zip_code) == zipCode {
			if lang == "th" {
				return s.name_in_thai
			} else {
				return s.name_in_english
			}
		}
	}

	return ""
}

func (a *Address) GetSubDistrictByDistrictId(districtId string, lang string) []SubDistrict {

	var subDistricts []SubDistrict

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.district_id) == districtId {
			subDistricts = append(subDistricts, s)
		}
	}

	return subDistricts
}

func (a *Address) GetDistrictByProvinceId(provinceId string, lang string) []District {

	var districts []District

	for _, d := range a.Districts {
		if fmt.Sprint(d.province_id) == provinceId {
			districts = append(districts, d)
		}
	}

	return districts
}

func (a *Address) GetProvinceByDistrictName(districtName string, lang string) string {

	for _, d := range a.Districts {
		if d.name_in_thai == districtName {
			return a.GetProvince(fmt.Sprint(d.province_id), lang)
		}
	}

	return ""
}

func (a *Address) GetProvinceBySubDistrictName(subDistrictName string, lang string) string {

	for _, s := range a.SubDistricts {
		if s.name_in_thai == subDistrictName {
			return a.GetProvinceByDistrictName(a.GetDistrict(fmt.Sprint(s.district_id), lang), lang)
		}
	}

	return ""
}

func (a *Address) GetProvinceByZipCode(zipCode string, lang string) string {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.zip_code) == zipCode {
			return a.GetProvinceByDistrictName(a.GetDistrict(fmt.Sprint(s.district_id), lang), lang)
		}
	}

	return ""
}

func (a *Address) GetDistrictBySubDistrictName(subDistrictName string, lang string) string {

	for _, s := range a.SubDistricts {
		if s.name_in_thai == subDistrictName {
			return a.GetDistrict(fmt.Sprint(s.district_id), lang)
		}
	}

	return ""
}

func (a *Address) GetDistrictByZipCode(zipCode string, lang string) string {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.zip_code) == zipCode {
			return a.GetDistrict(fmt.Sprint(s.district_id), lang)
		}
	}

	return ""
}

func (a *Address) GetAddressByLatLong(latitude string, longitude string, lang string) (string, string, string) {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.latitude) == latitude && fmt.Sprint(s.longitude) == longitude {
			return a.GetProvinceByDistrictName(a.GetDistrict(fmt.Sprint(s.district_id), lang), lang), a.GetDistrict(fmt.Sprint(s.district_id), lang), a.GetSubDistrict(fmt.Sprint(s.id), lang)
		}
	}

	return "", "", ""
}

func (a *Address) GetAddressByZipCode(zipCode string, lang string) (string, string, string) {

	for _, s := range a.SubDistricts {
		if fmt.Sprint(s.zip_code) == zipCode {
			return a.GetProvinceByDistrictName(a.GetDistrict(fmt.Sprint(s.district_id), lang), lang), a.GetDistrict(fmt.Sprint(s.district_id), lang), a.GetSubDistrict(fmt.Sprint(s.id), lang)
		}
	}

	return "", "", ""
}

func (a *Address) Search(query string) []string {
	results := []string{}
	lowerQuery := strings.ToLower(query)

	// Search in Provinces
	for _, p := range a.Provinces {
		if strings.Contains(strings.ToLower(p.name_in_thai), lowerQuery) || strings.Contains(strings.ToLower(p.name_in_english), lowerQuery) {
			results = append(results, fmt.Sprintf("Province: %s / %s", p.name_in_thai, p.name_in_english))
		}
	}

	// Search in Districts
	for _, d := range a.Districts {
		if strings.Contains(strings.ToLower(d.name_in_thai), lowerQuery) || strings.Contains(strings.ToLower(d.name_in_english), lowerQuery) {
			results = append(results, fmt.Sprintf("District: %s / %s", d.name_in_thai, d.name_in_english))
		}
	}

	// Search in SubDistricts
	for _, s := range a.SubDistricts {
		if strings.Contains(strings.ToLower(s.name_in_thai), lowerQuery) || strings.Contains(strings.ToLower(s.name_in_english), lowerQuery) {
			results = append(results, fmt.Sprintf("SubDistrict: %s / %s", s.name_in_thai, s.name_in_english))
		}
	}

	return results
}
