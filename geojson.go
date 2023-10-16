package backendgcf

import "encoding/json"

func GCHandlerFunc(Mongostring, dbname, colname string) []byte {
	koneksi := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllGeoData(koneksi, colname)

	jsonhei, _ := json.Marshal(datageo)
	return jsonhei
}
