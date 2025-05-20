package kgoutil

const (
    ERRNO_SUCCESS                   =  0
    ERRNO_FAIL                      = -1
    ERRNO_FAIL_LOGIN                = -2
    ERRNO_FAIL_LOGIN_PASS           = -3
    ERRNO_FAIL_LOGIN_WX             = -4
    ERRNO_FAIL_INVALID_REQUEST      = -5
)

var _messages =  map[int]string{
    ERRNO_SUCCESS:                  "Success.",
    ERRNO_FAIL:                     "Failed.",
    ERRNO_FAIL_LOGIN:               "Failed to login.",
    ERRNO_FAIL_LOGIN_PASS:          "Failed to login: username or password is not correct.",
    ERRNO_FAIL_LOGIN_WX:            "Failed to login by wx.",
    ERRNO_FAIL_INVALID_REQUEST:     "Failed to request: invalid request.",
}

func Message(errno int) string {
    message, exists := _messages[errno]
    if exists {
        return message
    } else {
        return "unknown error"
    }
}

