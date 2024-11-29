package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type config struct {
	Database string  `xml:"database"`
	Username string  `xml:"username"`
	Password string  `xml:"password"`
	Options  Options `xml:"options"`
}
type Options struct {
	AutoBackup     bool `xml:"auto_backup"`
	MaxConnections int  `xml:"max_connections"`
}

func createAndWriteXMLFile(str config) {
	file, err := os.OpenFile("task/xml/config.xml", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error happened on creating file:", err)
	}
	defer file.Close()

	x, err := xml.Marshal(str)
	if err != nil {
		fmt.Println("Error occurred while marshalling", err)
	}

	_, err = file.Write(x)
	if err != nil {
		fmt.Println("Error occurred while writing", err)
	}
}

func Task6() {
	fmt.Println("Task-6   ********************")
	Config := config{Database: "mydb", Username: "admin", Password: "secret", Options: Options{AutoBackup: true, MaxConnections: 100}}
	createAndWriteXMLFile(Config)
}

//<config><database>mydb</database><username>admin</username><password>secret</password><options><auto_backup>true</auto_backup><max_connections>100</max_connections></options></config>
