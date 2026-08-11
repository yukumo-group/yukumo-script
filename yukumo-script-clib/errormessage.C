// ErrorMessage manages the result of golang function
struct ErrorMessage{
    int code;
    char* information;
};

// NewErrorMessage creates new error message
struct ErrorMessage NewErrorMessage(int code, char* information) {
    struct ErrorMessage newErrorMessage;
    newErrorMessage.code = code;
    newErrorMessage.information = information;
    return newErrorMessage;
}
