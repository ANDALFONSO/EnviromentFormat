package service

import (
	"bytes"
	"fmt"
	"format_enviroment/entity"
	"log"
	"strings"
)

type IFormatService interface {
	Format(text string) entity.Response
}

type FormatService struct {
	response entity.Response
}

func NewPS() IFormatService {
	return &FormatService{}
}

func (s *FormatService) Format(text string) entity.Response {
	chains := strings.Split(text, "\n")
	for i, c := range chains {
		log.Println(i, c)
	}
	envs := []string{chains[2], chains[3], chains[4], chains[9], chains[10], chains[11], chains[12]}
	mapEnv := getEnv(envs)
	kvsName := strings.Replace(strings.Split(chains[1], ":")[1], " ", "", 1)
	dsName := strings.Replace(strings.Split(chains[8], ":")[1], " ", "", 1)
	s.response.NameServiceKvs = kvsName
	s.response.NameServiceDs = dsName
	s.response.Vscode = mapEnv
	s.response.Golang = mapToString(mapEnv)
	return s.response
}

func getEnv(envs []string) map[string]string {
	m := make(map[string]string)
	for _, env := range envs {
		chains := strings.Split(env, " ")
		//log.Printf("key:%v, value:%v", chains[1], chains[2])
		chains[1] = strings.Replace(chains[1], ":", "", 1)
		m[chains[1]] = chains[2]
	}
	m["SCOPE"] = "local"
	m["GO_ENVIRONMENT"] = "production"
	m["CONF_DIR"] = "${workspaceRoot}/conf"
	m["configFileName"] = "${workspaceRoot}/pkg/config/application.properties"
	m["checksumEnabled"] = "false"
	m["NEW_RELIC_LICENSE_KEY"] = "1111111111111111111111111111111111111111"
	m["NEW_RELIC_APP_NAME"] = "wdm-api"
	m["CACHE_TESTWDM_NODES_ENDPOINT"] = "localhost:11211"
	return m
}

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%v=%v;", key, value)
	}
	return b.String()
}
