# golang_test_app

**golang_test_app** work in deamon mode.

* all errors send to **stderr**


**ARGs LIST:**\
--help - Show help\
--version - Show version\
--api-port string - API Port (default "8080")\
--log-level string - Set log level (trace,debug,info,warning,error,fatal). (default "INFO")

**Run examples golang_test_app**:
* golang_test_app

**Examples ENV:**
* export DBAAS_LOG_LEVEL="INFO"
* export DBAAS_API_PORT="8080"
* export DBAAS_DB_URL="XXX"

HOW run  
`http://127.0.0.1:8080/config`
`http://127.0.0.1:8080/getInfo?ip=8.8.8.8`

Be aware! Google Chrome does not run simultaneous sessions with the same URI.   
 `curl "http://127.0.0.1:8080/getInfo?ip=8.8.8.8" & curl "http://127.0.0.1:8080/getInfo?ip=1.1.1.1"`
