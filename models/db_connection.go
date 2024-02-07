package models

import (
	"bytes"
)

type DbConnection struct {
	host     string
	port     string
	user     string
	dbName   string
	password string
	sslMode  string
}

func (connection *DbConnection) New() *DbConnection {
	connection = &DbConnection{
		host:     "your-host-name",
		port:     "your-port-name",
		user:     "your-user-name",
		dbName:   "your-dbname-name",
		password: "your-password",
		sslMode:  "your-sslmode",
	}
	return connection
}

func (connection DbConnection) GetConnection() string {
	var buffer bytes.Buffer
	buffer.WriteString("host=")
	buffer.WriteString(connection.host)
	buffer.WriteString(" port=")
	buffer.WriteString(connection.port)
	buffer.WriteString(" user=")
	buffer.WriteString(connection.user)
	buffer.WriteString(" dbname=")
	buffer.WriteString(connection.dbName)
	buffer.WriteString(" password=")
	buffer.WriteString(connection.password)
	buffer.WriteString(" sslmode=")
	buffer.WriteString(connection.sslMode)
	return buffer.String()
}
