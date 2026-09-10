#include <stdlib.h>
#include <string.h>
#ifndef calculator
#define calculator
// ErrorMessage manages the result of a golang function for C callers.
typedef struct ErrorMessage {
	int code;
	char* information;
} ErrorMessage;
// new_error_message creates new error message
ErrorMessage new_error_message(int code, char* information);
#endif