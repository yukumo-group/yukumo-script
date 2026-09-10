#include "errmessage.h"
#include <stdlib.h>
#include <string.h>

ErrorMessage new_error_message(int code, char* information) {
	ErrorMessage msg;
	msg.code = code;
	msg.information = information;
	return msg;
}