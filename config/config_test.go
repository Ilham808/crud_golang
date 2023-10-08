package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	config := InitConfig()

	assert.NotNil(t, config, "Config should not be nil")
	assert.Equal(t, "root", config.DB_Username, "DB_Username should be 'root'")
	assert.Equal(t, "mysql", config.DB_Password, "DB_Password should be 'mysql'")
	assert.Equal(t, "3306", config.DB_Port, "DB_Port should be '3306'")
	assert.Equal(t, "127.0.0.1", config.DB_Host, "DB_Host should be '127.0.0.1'")
	assert.Equal(t, "crud_go", config.DB_Name, "DB_Name should be 'crud_go'")
}

func TestLoadConfigTest(t *testing.T) {
	config := loadConfigTest()

	assert.NotNil(t, config, "Config should not be nil")
	assert.Equal(t, "root", config.DB_Username, "DB_Username should be 'root'")
	assert.Equal(t, "mysql", config.DB_Password, "DB_Password should be 'mysql'")
	assert.Equal(t, "3306", config.DB_Port, "DB_Port should be '3306'")
	assert.Equal(t, "127.0.0.1", config.DB_Host, "DB_Host should be '127.0.0.1'")
	assert.Equal(t, "crud_go", config.DB_Name, "DB_Name should be 'crud_go'")
}
