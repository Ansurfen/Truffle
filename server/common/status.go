package common

/*
	Service err
*/

var (
	SUCCESS = newResult(SUCCESS_CODE, "ok")
	ERROR   = newResult(ERROR_CODE, "")

	// json err

	JSON_MARSHAL_ERR   = newResult(JSON_MARSHAL_ERR_CODE, "").WithMsg("system error, please retrying")
	JSON_UNMARSHAL_ERR = newResult(JSON_UNMARSHAL_ERR_CODE, "").WithMsg("system error, please retrying")
)

const (
	DB_ERR_CODE = 1000

	SUCCESS_CODE             = 200
	ERROR_BUT_NO_ISSUCE_CODE = 201
	ERROR_CODE               = 500

	JSON_MARSHAL_ERR_CODE   = 600
	JSON_UNMARSHAL_ERR_CODE = 601
)
