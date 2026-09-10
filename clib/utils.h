#include <stdlib.h>
#include <string.h>
#ifndef utils
#define utils
// StringList is a NULL-terminated list of C strings allocated by this library.
typedef struct StringList {
	char** items;
	int count;
} StringList;
#endif