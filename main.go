package main

import (
	"flag"
	"net/http"
	"os"

	conf "./config"
	lib "./lib"
	log "./logging"
	services "./services"
	"./transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initDB() lib.DbConnection {
	var dbConn *lib.DbConnection

	dbConn, err := lib.New(conf.Param.Query)
	if err != nil {
		log.Errorf("Unable to initialize database %v", err)
		os.Exit(1)
	}
	dbConn.Db, err = dbConn.Open()
	if err != nil {
		log.Errorf("Unable to open database %v", err)
		os.Exit(1)
	}
	return *dbConn
}

func initHandlers(dbConn lib.DbConnection) {

	var svcJWT services.JWTServices
	svcJWT = services.JWTService{}

	var svcCorporateCategoryList services.CorporateCategoryListServices
	svcCorporateCategoryList = services.CorporateCategoryListService{}

	var svcCorporateRequiredInfo services.CorporateRequiredInfoServices
	svcCorporateRequiredInfo = services.CorporateRequiredInfoService{}

	var svcCorporateCreateTicket services.CorporateCreateTicketServices
	svcCorporateCreateTicket = services.CorporateCreateTicketService{}

	JWTHandler := httptransport.NewServer(
		transport.JWTEndpoint(svcJWT),
		transport.JWTDecodeRequest,
		transport.JWTEncodeResponse,
	)

	CorporateCategoryListHandler := httptransport.NewServer(
		transport.CorporateCategoryListEndpoint(svcCorporateCategoryList, dbConn),
		transport.CorporateCategoryListDecodeRequest,
		transport.CorporateCategoryListEncodeResponse,
	)

	CorporateCreateTicketHandler := httptransport.NewServer(
		transport.CorporateCreateTicketEndpoint(svcCorporateCreateTicket, dbConn),
		transport.CorporateCreateTicketDecodeRequest,
		transport.CorporateCreateTicketEncodeResponse,
	)

	CorporateRequiredInfoHandler := httptransport.NewServer(
		transport.CorporateRequiredInfoEndpoint(svcCorporateRequiredInfo, dbConn),
		transport.CorporateRequiredInfoDecodeRequest,
		transport.CorporateRequiredInfoEncodeResponse,
	)

	http.Handle("/api/corporate/getToken", JWTHandler)
	http.Handle("/api/corporate/categoryList", CorporateCategoryListHandler)
	http.Handle("/api/corporate/requiredInfo", CorporateRequiredInfoHandler)
	http.Handle("/api/corporate/createTicket", CorporateCreateTicketHandler)

}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	configFile := flag.String("conf", "config/conf.yml", "main configuration file")
	flag.Parse()

	log.Logf("Reads configuration from %s", *configFile)
	conf.LoadConfigFromFile(configFile)

	//log.Init(conf.Config.Log.Level, common.Config.Log.FileName)

	// initiate Service Database connection
	dbConn := initDB()

	// Register and Initiate Listener
	initHandlers(dbConn)
	var err error
	err = http.ListenAndServe(conf.Param.ListenPort, nil)

	if err != nil {
		log.Errorf("Unable to start the server %v", err)
		os.Exit(1)
	}

}
