// seed.go
package main

import (
	"fmt"
	"log"
	"alephcode-backend/config"
	"alephcode-backend/models"
)

func main() {
	config.ConnectDB()

	// Define initial mission
	mission := models.Mission{
		Title:       "Fix the Greeting Bug",
		Description: `The robot says the wrong greeting. Can you fix it so it says "Hello, World!"?`,
		BlocklyXML: `<xml xmlns="https://developers.google.com/blockly/xml">
  <block type="text_print" x="20" y="20">
    <value name="TEXT">
      <shadow type="text">
        <field name="TEXT">Hi Mom!</field>
      </shadow>
    </value>
  </block>
</xml>`,
	}

	// Check if it exists
	var existing models.Mission
	err := config.DB.Where("title = ?", mission.Title).First(&existing).Error
	if err == nil {
		fmt.Println("Mission already exists, skipping seed.")
		return
	}

	// Insert mission
	if err := config.DB.Create(&mission).Error; err != nil {
		log.Fatalf("Failed to seed mission: %v", err)
	}

	fmt.Println("Seeded mission successfully.")
}
