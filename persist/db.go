package persist

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/norhe/transit-benchmark/utils"
)

// Config : Describes the database to persist records in
type Config struct {
	Addr     string
	Username string
	Password string
	DbName   string
}

var db *sql.DB

func initDb(cfg *Config) {
	if nil == db {
		d, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.Username, cfg.Password, cfg.Addr, cfg.DbName))
		utils.FailOnError(err, "Failed to connect to database")
		db = d
	}

}

// CreateTables : Init the db for saving results
func CreateTables(conf *Config) {
	initDb(conf)
	_, err := db.Exec("USE " + conf.DbName)
	utils.FailOnError(err, "Failed to connect to DB")

	createTestTable :=
		"CREATE TABLE IF NOT EXISTS `test_run`(" +
			"`test_id` INT(11) NOT NULL AUTO_INCREMENT, " +
			"`start_date` VARCHAR(256) NOT NULL," +
			"PRIMARY KEY (test_id) " +
			") engine=InnoDB;"

	log.Println("Creating test_run table (if not exist)")

	_, err = db.Exec(createTestTable)
	utils.FailOnError(err, "Failed to create test_run table")

	// WorkUnits
	createWorkUnitTable :=
		"CREATE TABLE IF NOT EXISTS `workunit`(" +
			"`workunit_id` INT(11) NOT NULL AUTO_INCREMENT, " +
			"`test_id` INT(11) NOT NULL, " +
			"`start_time` DATETIME NOT NULL," +
			"`end_time` DATETIME NOT NULL," +
			"`run_millis` INT(11) NOT NULL, " +
			"`operation_type_id` INT(11) NOT NULL, " +
			"`exception` varchar(256) NULL, " +
			"`payload` TEXT(32768) NULL, " +
			"`payload_size` INT(11) NOT NULL, " +
			"`output` VARCHAR(32768) NULL, " +
			"PRIMARY KEY (workunit_id) " +
			") engine=InnoDB;"

	log.Println("Creating workunit table (if not exist)")

	_, err = db.Exec(createWorkUnitTable)
	utils.FailOnError(err, "Failed to create workunit table")

	// Operations
	createOperationTypeTable :=
		"CREATE TABLE IF NOT EXISTS `op_type`(" +
			"`operation_type_id` INT(11) NOT NULL, " +
			"`operation_type` varchar(256) NOT NULL, " +
			"PRIMARY KEY (operation_type_id) " +
			") engine=InnoDB;"

	log.Println("Creating op_type table (if not exist)")

	_, err = db.Exec(createOperationTypeTable)
	utils.FailOnError(err, "Failed to create operation_type table")

	err = populateOpTable(db)
	utils.FailOnError(err, "Failed to populate operation table")
}

func populateOpTable(db *sql.DB) error {
	statement := "INSERT INTO `op_type` (operation_type_id, operation_type)" +
		"VALUES (0, \"Encrypt\"), (1, \"Decrypt\"), (2, \"Rewrap\")," +
		"(3, \"GenerateDataKey\"), (4, \"GenerateRandomBytes\")," +
		"(5, \"HashData\"), (6, \"GenerateHMAC\")," +
		"(7, \"SignData\"), (8, \"VerifySignedData\");"
	_, err := db.Exec(statement)
	// we expect duplicate entry error here on subsequent starts
	if err != nil && !strings.Contains(err.Error(), "Error 1062") {
		log.Print(err)
		return err
	}
	return nil
}
