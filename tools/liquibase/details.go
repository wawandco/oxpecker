package liquibase

import (
	"errors"
	"fmt"
	"regexp"
)

type connectionDetails struct {
	URL      string
	username string
	password string
}

func (lb Command) buildRunArgs() ([]string, error) {
	deets, err := lb.connectionDetails()
	if err != nil {
		return []string{}, err
	}

	runArgs := []string{
		"--driver=org.postgresql.Driver",
		"--url=" + deets.URL,
		"--logLevel=debug",
		"--username=" + deets.username,
		"--password=" + deets.password,
	}
	return runArgs, nil
}

func (lb *Command) connectionDetails() (connectionDetails, error) {
	conn := lb.connections[lb.connectionName]
	if conn == nil {
		return connectionDetails{}, errors.New("connection not found")
	}

	r := regexp.MustCompile(`postgres:\/\/(?P<username>.*):(?P<password>.*)@(?P<host>.*):(?P<port>.*)\/(?P<database>.*)\?(?P<extras>.*)`)
	match := r.FindStringSubmatch(conn.URL())
	if match == nil {
		return connectionDetails{}, fmt.Errorf("could not convert `%v` url into Liquibase", lb.connectionName)
	}

	return connectionDetails{
		URL:      fmt.Sprintf("jdbc:postgresql://%v:%v/%v?%v", match[3], match[4], match[5], match[6]),
		username: match[1],
		password: match[2],
	}, nil
}
