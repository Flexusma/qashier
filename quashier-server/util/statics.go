package util

import "gorm.io/gorm"

const Version string = "0.0.1"
const Dev bool = true

const DBURL = ""

// Holds the Db pointer on runtime to prevent circle imports
var DBI *gorm.DB
