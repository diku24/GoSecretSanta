package repository

import (
	"SecretSanta/entity"
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
)

//Global Variables
var (
	username       = os.Getenv("USERNAME")
	password       = os.Getenv("PASSWORD")
	bucketName     = os.Getenv("BUCKETNAME")
	connectionPath = os.Getenv("CONNECTION_PATH")
)

//The New Repository Struct to implemet the Interface
type CouchbaseRepository struct{}

//The New repository function to initilize the func
func NewCouchbaseRepository() PersonRepository {

	return &CouchbaseRepository{}
}

func (*CouchbaseRepository) SavePerson(personWish *entity.PersonWish) (*entity.PersonWish, error) {

	//The Initialize the Connection
	cluster, err := gocb.Connect(connectionPath, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		SecurityConfig: gocb.SecurityConfig{
			//Do not set this to true in Production, use it only for testing !
			TLSSkipVerify: true,
		},
	})

	if err != nil {
		logrus.Errorln("Error in Connecting to the Database Server: ", err)
	}

	//Configuration connection for couchbase Bucket.
	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)

	if err != nil {
		logrus.Errorln("Error in Bucket Connection", err)
	}

	//Insert Query execuion
	query := fmt.Sprintf("INSERT INTO '%s' (KEY, VALUE) values ($1, {'name' : $2, 'santa': $3, 'wishes' : $4})", bucketName)
	rows, err := cluster.Query(query, &gocb.QueryOptions{PositionalParameters: []interface{}{personWish.Name, personWish.Name, personWish.Santa, personWish.Wishes}})

	if err != nil {
		logrus.Errorln(err)
	}

	//Printing each found rows
	for rows.Next() {
		var result interface{}
		err := rows.Row(&result)

		if err != nil {
			logrus.Errorln("Error in Rows of Data")
		}
	}

	if err := rows.Err(); err != nil {
		logrus.Errorln(err)
	}

	return personWish, nil
}

func (*CouchbaseRepository) GetAllWishes() ([]entity.PersonWish, error) {
	//Couchbase Logging
	//gocb.SetLogger(gocb.VerboseStdioLogger())

	var data []entity.PersonWish

	//Initialize the Connection
	cluster, err := gocb.Connect(connectionPath, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		SecurityConfig: gocb.SecurityConfig{
			TLSSkipVerify: true,
		},
	})

	if err != nil {
		logrus.Errorln("error in connection")
	}

	//Bucket connection
	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(5*time.Second, nil)

	if err != nil {
		logrus.Errorln("Bucket connnection error", err)
	}

	//Querying SELECT Statement
	query := fmt.Sprintf("SELECT * FROM '%s' ", bucketName)
	rows, err := cluster.Query(query, nil)

	if err != nil {
		logrus.Errorln("Query execution error", err)
	}

	//Traversing on each found rows
	for rows.Next() {
		var results entity.PersonWish
		err := rows.Row(&results)
		if err != nil {
			logrus.Errorln("Error in reading the rows", err)
		}

		data = append(data, results)
	}

	//Checking rows for errors
	if err := rows.Err(); err != nil {
		logrus.Errorln(err)
	}
	return data, nil
}

func (*CouchbaseRepository) AllocateSanta() error {
	//Couchbase Logging
	//gocb.SetLogger(gocb.VerboseStdioLogger())

	//Initialize the Connection
	cluster, err := gocb.Connect(connectionPath, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		SecurityConfig: gocb.SecurityConfig{
			TLSSkipVerify: true,
		},
	})

	if err != nil {
		logrus.Errorln("error in connection")
	}

	//Bucket connection
	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(5*time.Second, nil)

	if err != nil {
		logrus.Errorln("Bucket connnection error", err)
	}

	//Querying UPDATE Statement
	query := fmt.Sprintf("UPDATE '%s' s SET s.santa = s.name ", bucketName)
	rows, err := cluster.Query(query, nil)

	if err != nil {
		logrus.Errorln("Query execution error", err)
	}

	//Checking rows for errors
	if err := rows.Err(); err != nil {
		logrus.Errorln(err)
	}
	return err
}
