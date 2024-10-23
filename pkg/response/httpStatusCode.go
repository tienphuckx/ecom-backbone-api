package response

// Constants for error codes
const (
	ErrCode_SUCCESS               = 2000
	ErrCode_INTERNAL_SERVER_ERROR = 5000
	ErrCode_INVALID_REQUEST       = 4000
	ErrCode_UNAUTHORIZED          = 4010
	ErrCode_FORBIDDEN             = 4030
	ErrCode_NOT_FOUND             = 4040
	ErrCode_CONFLICT              = 4090
	ErrCode_TOO_MANY_REQUESTS     = 4290
	ErrCode_SERVICE_UNAVAILABLE   = 5030
	ErrCode_INVALID_CREDENTIALS   = 6000
	ErrCode_UNKNOWN_ERROR         = 9999
	ErrCode_INVALID_TOKEN         = 7000
	ErrCode_INVALID_REFRESH_TOKEN = 7001
	ErrCode_INVALID_CLIENT_ID     = 7002
	ErrCode_EXPIRED_TOKEN         = 7003
	ErrCode_TOKEN_REJECTED        = 7004
)

// Map for error code messages
var msg = map[int]string{
	ErrCode_SUCCESS:               "Success",
	ErrCode_INTERNAL_SERVER_ERROR: "Internal server error",
	ErrCode_INVALID_REQUEST:       "Invalid request",
	ErrCode_UNAUTHORIZED:          "Unauthorized",
	ErrCode_FORBIDDEN:             "Forbidden",
	ErrCode_NOT_FOUND:             "Not found",
	ErrCode_CONFLICT:              "Conflict",
	ErrCode_TOO_MANY_REQUESTS:     "Too many requests",
	ErrCode_SERVICE_UNAVAILABLE:   "Service unavailable",
	ErrCode_INVALID_CREDENTIALS:   "Invalid credentials",
	ErrCode_UNKNOWN_ERROR:         "Unknown error",
	ErrCode_INVALID_TOKEN:         "Invalid token",
	ErrCode_INVALID_REFRESH_TOKEN: "Invalid refresh token",
	ErrCode_INVALID_CLIENT_ID:     "Invalid client ID",
	ErrCode_EXPIRED_TOKEN:         "Expired token",
	ErrCode_TOKEN_REJECTED:        "Token rejected",
}
