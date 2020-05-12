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
		host:     "database-1.cemgdq2nipfl.us-east-2.rds.amazonaws.com",
		port:     "5432",
		user:     "postgres",
		dbName:   "postgres",
		password: "f5S3*8eY*GlQLQmzAgIw",
		sslMode:  "require",
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