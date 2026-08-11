/*
yukumo-script  Copyright (C) 2026  1Vewton
This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
This is free software, and you are welcome to redistribute it
under certain conditions; type `show c' for details.
*/
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "AquesTalk2.h"
#ifndef generator
#define generator
// generate_wav generates the yukumo wav file
int generate_wav(char* phont_path, char* text, char* result_path, int speed);
#endif