package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/makzzz1986/s3-zookeeper-uploader/cmd"
	log "github.com/sirupsen/logrus"
)

var (
	zkHost        string
	AwsRegionName = "eu-west-1"
)

func init() {
	zkHost = os.Getenv("ZK_HOST")
	if zkHost == "" {
		zkHost = "127.0.0.1"
	}

	log.SetOutput(os.Stdout)
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" || strings.ToLower(logLevel) == "info" {
		log.SetLevel(log.InfoLevel) // default is INFO
	} else if strings.ToLower(logLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if strings.ToLower(logLevel) == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if strings.ToLower(logLevel) == "error" {
		log.SetLevel(log.ErrorLevel)
	} else if strings.ToLower(logLevel) == "critical" || strings.ToLower(logLevel) == "fatal" {
		log.SetLevel(log.FatalLevel)
	}
}

func main() {
	fmt.Println("The app has been started")

	// conn, err := cmd.ZkConnection(zkHost)
	// if err != nil {
	// 	panic(err)
	// }

	// cmd.List(conn, "/")
	// data, err := cmd.Get(conn, "/security.json")
	// // data, err := cmd.Get(conn, "/aliases.json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data[:]))

	// data, _ = cmd.Get(conn, "/autoscaling")
	// fmt.Println(string(data[:]))

	// treePath := "/overseer_elect"
	// tree, err := cmd.Tree(conn, treePath)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("\nPrinting file tree of %s\n", treePath)
	// 	for _, file := range tree {
	// 		fmt.Println(file)
	// 	}
	// }

	// data, err := os.ReadFile("notes/testfile.txt") // just pass the file name
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	cmd.GetHash(data)
	// }

	// uploaded, err := cmd.Upload(conn, "/tmp", []byte{})
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(uploaded)
	// }

	// uploaded, err := cmd.Upload(conn, "/tmp/more/testfile.txt", data)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(uploaded)
	// }

	// hash, _ := cmd.Hash(conn, "/tmp/more/testfile.txt")
	// fmt.Println(hash)
	client, err := cmd.S3Connection(AwsRegionName)
	if err != nil {
		panic(err)
	}
	result, err := cmd.GetS3ListObjects(client, "solr-updater-2", "TAG2/")
	if err != nil {
		panic(err)
	} else {
		log.Infoln(result)
	}
}
