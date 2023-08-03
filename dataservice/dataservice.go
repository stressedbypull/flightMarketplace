package dataservice

import "tui/flightmarketplace/database"

type DataServiceInterface interface {
	FlightDataservice
	//UserDataservice
	//OrderDataservice
	DB() database.PostgresInterface
}

type DataGatewayService struct {
	Db database.PostgresInterface
	//redisConn *redis.Client
}

// DB get the internal DB manager instance.
func (dgs *DataGatewayService) DB() database.PostgresInterface {
	return dgs.Db
}
