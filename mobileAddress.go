package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type MobileAddress struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Provice  string `json:"provice"`
	City     string `json:"city"`
	Isp      string `json:"isp"`
	AreaCode string `json:"areacode"`
	ZipCode  string `json:"zipcode"`
}

var mobileAddressMap = make(map[string]*MobileAddress)

func NewMobileAddress(code string, areacode string, province string, city string, isp string) *MobileAddress {
	return &MobileAddress{
		Code:     code,
		AreaCode: areacode,
		Provice:  province,
		City:     city,
		Isp:      isp,
	}
}

func QueryMobile(code string) *MobileAddress {
	if len(code) <= 10 {
		return nil
	}
	var part string

	if strings.HasPrefix(code, "0") {
		// FIXME 根据区号查询只有归属地是对的。什么运营商都是错的。。。
		part = code[1:4]
	} else if strings.HasPrefix(code, "1") {
		part = code[0:7]
	}
	log.Printf("part: '%s'\n", part)
	if part == "" {
		return nil
	}
	return mobileAddressMap[part]
}

func LoadMobileAddress() {
	log.Println("LoadMobileAddress")
	fileName := "data/mobile.data"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when open file %s: %s", fileName, err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		a := strings.Fields(line)
		var address *MobileAddress
		if len(a) == 5 {
			address = NewMobileAddress(a[0], a[1], a[2], a[3], a[4])
		} else if len(a) == 4 {
			address = NewMobileAddress(a[0], a[1], a[2], a[2], a[3])
		} else {
			continue
		}
		mobileAddressMap[address.Code] = address
		mobileAddressMap[address.AreaCode] = address
	}

	log.Printf("%d\n", len(mobileAddressMap))
}
