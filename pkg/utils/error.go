package utils

// アプリケーションで扱うエラーを列挙する
type ErrorCode int

// GraphqlError Lambdaの返却で扱うエラー型
type ErrorFmt struct {
	Name    string `json:"errType"`
	Message string `json:"errMessage"`
}

const (
	ERR_INTERNAL ErrorCode = iota
	ERR_INVALID_EP
	ERR_INVALID_ID
	ERR_INVALID_S3_CONTENT
	ERR_INVALID_FILETYPE
	ERR_INVALID_SORYBY_TIPE
	ERR_INVALID_TS
	ERR_DENIED_PROJECT
	ERR_DENIED_UT
	ERR_FAILED_GET_ORG
	ERR_FAILED_GET_FOSTOKEN
	ERR_FAILED_GET_PP
	ERR_FAILED_GET_PRJ
	ERR_FAILED_GET_FLIGHT
	ERR_FAILED_GET_POI
	ERR_REQUIRED_PROJECT_ID
	ERR_REQUIRED_FLIGHT_ID
	ERR_REQUIRED_PLAN_ID
	ERR_REQUIRED_ID
	ERR_REQUIRED_IDS
	ERR_REQUIRED_GEOMETRY
	ERR_REQUIRED_CONFIG
	ERR_REQUIRED_FILE_IDS
	ERR_REQUIRED_FILE_SRC_ID
	ERR_REQUIRED_NAME
	ERR_REQUIRED_TYPE
	ERR_REQUIRED_POINTS
	ERR_REQUIRED_COMMENT
	ERR_REQUIRED_TIMESTAMP
	ERR_REQUIRED_CRITICALITY
	ERR_REQUIRED_SRC
	ERR_REQUIRED_COLOR
	ERR_EXCEEDED_STORAGE
)

// アプリケーションで扱うエラータイプを列挙する
var ErrorCodeMap = map[ErrorCode]string{
	ERR_INTERNAL:             "InternalError",
	ERR_INVALID_EP:           "InvalidEndpoint",
	ERR_INVALID_ID:           "InvalidId",
	ERR_INVALID_S3_CONTENT:   "InvalidS3Content",
	ERR_INVALID_FILETYPE:     "InvalidFileType",
	ERR_INVALID_SORYBY_TIPE:  "InvalidSoryByType",
	ERR_INVALID_TS:           "InvalidTimeStamp",
	ERR_DENIED_PROJECT:       "DeniedProject",
	ERR_DENIED_UT:            "DeniedUt",
	ERR_FAILED_GET_ORG:       "FailedGetOrganization",
	ERR_FAILED_GET_FOSTOKEN:  "FailedGetFosToken",
	ERR_FAILED_GET_PP:        "FailedGetPaymentPlans",
	ERR_FAILED_GET_PRJ:       "FailedGetProject",
	ERR_FAILED_GET_FLIGHT:    "FailedGetFlight",
	ERR_FAILED_GET_POI:       "FailedGetPoi",
	ERR_REQUIRED_PROJECT_ID:  "RequiredProjectId",
	ERR_REQUIRED_FLIGHT_ID:   "RequiredFlightId",
	ERR_REQUIRED_PLAN_ID:     "RequiredPlanId",
	ERR_REQUIRED_ID:          "RequiredId",
	ERR_REQUIRED_IDS:         "RequiredIds",
	ERR_REQUIRED_GEOMETRY:    "RequiredGeometry",
	ERR_REQUIRED_CONFIG:      "RequiredConfig",
	ERR_REQUIRED_FILE_IDS:    "RequiredFileIds",
	ERR_REQUIRED_FILE_SRC_ID: "RequiredFileIdOrSrc",
	ERR_REQUIRED_NAME:        "RequiredName",
	ERR_REQUIRED_TYPE:        "RequiredType",
	ERR_REQUIRED_COMMENT:     "RequiredComment",
	ERR_REQUIRED_POINTS:      "RequiredPoints",
	ERR_REQUIRED_TIMESTAMP:   "RequiredTimestamp",
	ERR_REQUIRED_CRITICALITY: "RequiredCriticality",
	ERR_REQUIRED_SRC:         "RequiredSrc",
	ERR_REQUIRED_COLOR:       "RequiredColor",
	ERR_EXCEEDED_STORAGE:     "ExceededStoregeUsage",
}

// アプリケーションで扱うエラータイプごとのエラーメッセージを列挙する
var ErrorCodeMessageMap = map[ErrorCode]string{
	ERR_INTERNAL:             "Internal server error! Contact system admins",
	ERR_INVALID_EP:           "No endpoint",
	ERR_INVALID_ID:           "Your ID would be mistaken",
	ERR_INVALID_S3_CONTENT:   "S3 content can't be found",
	ERR_INVALID_FILETYPE:     "Your fileType would be mistaken",
	ERR_INVALID_SORYBY_TIPE:  "Your soryByType would be mistaken",
	ERR_INVALID_TS:           "stamp.txt would be wrong",
	ERR_DENIED_PROJECT:       "Denied to access a project",
	ERR_DENIED_UT:            "Denied to access UT data",
	ERR_FAILED_GET_ORG:       "Failed to get organization info",
	ERR_FAILED_GET_FOSTOKEN:  "Failed to get FOS IdToken",
	ERR_FAILED_GET_PP:        "Failed to get paymentPlans",
	ERR_FAILED_GET_PRJ:       "Failed to get project info",
	ERR_FAILED_GET_FLIGHT:    "Failed to get flight info",
	ERR_FAILED_GET_POI:       "Failed to get poi info",
	ERR_REQUIRED_PROJECT_ID:  "Field 'project_id' is requried",
	ERR_REQUIRED_FLIGHT_ID:   "Field 'flight_id' is requried",
	ERR_REQUIRED_PLAN_ID:     "Field 'plan_id' is requried",
	ERR_REQUIRED_ID:          "Field 'id' is requried",
	ERR_REQUIRED_IDS:         "Field 'ids' is requried",
	ERR_REQUIRED_GEOMETRY:    "Field 'geometry' is requried",
	ERR_REQUIRED_CONFIG:      "Field 'config' is requried",
	ERR_REQUIRED_FILE_IDS:    "Field 'fileIds' is requried",
	ERR_REQUIRED_FILE_SRC_ID: "Field 'file.id' or 'src' is requried",
	ERR_REQUIRED_NAME:        "Field 'name' is requried",
	ERR_REQUIRED_TYPE:        "Field 'type' is requried",
	ERR_REQUIRED_COMMENT:     "Field 'comment' is requried",
	ERR_REQUIRED_POINTS:      "Field 'points' is requried",
	ERR_REQUIRED_TIMESTAMP:   "Field 'timestamp' is requried",
	ERR_REQUIRED_CRITICALITY: "Field 'criticality' is requried",
	ERR_REQUIRED_SRC:         "Field 'src' is requried",
	ERR_REQUIRED_COLOR:       "Field 'color' is requried",
	ERR_EXCEEDED_STORAGE:     "Storage usage has exceeded the maximum allowed",
}

// String 該当するエラータイプを文字列で返却する
func (e ErrorCode) Type() string {
	if s, ok := ErrorCodeMap[e]; ok {
		return s
	}
	panic("Invalid ErrorCode")
}

// Error 該当するエラーメッセージを文字列で返却する
func (e ErrorCode) Error() string {
	if s, ok := ErrorCodeMessageMap[e]; ok {
		return s
	}
	panic("Invalid ErrorCode")
}

// errTypeからErrorを生成する
func NewErr(e ErrorCode) *ErrorFmt {
	return &ErrorFmt{
		Name:    e.Type(),
		Message: e.Error(),
	}
}

// Error エラーメッセージを返却する
func (e *ErrorFmt) Error() string {
	return e.Message
}
