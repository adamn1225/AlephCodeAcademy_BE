package models

import "gorm.io/gorm"

type Mission struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    BlocklyXML  string `json:"blockly_xml"`
}
