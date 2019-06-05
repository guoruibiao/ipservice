package utils

import (
	"testing"
	"github.com/guoruibiao/ipservice/constants"
	"fmt"
	"log"
	"github.com/ipipdotnet/ipdb-go"
)

func TestGetOfficialData(t *testing.T) {
	e, err := GetOfficialData("223.104.148.108")
	if err != nil {
		t.Error(err)
	}else{
		t.Log(e)
	}
	e, err = GetOfficialData("22304.148.108")
	if err != nil {
		t.Log(err)
	}else{
		t.Log(e)
	}
}

