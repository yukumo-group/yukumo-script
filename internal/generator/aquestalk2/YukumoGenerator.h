/*
yukumo-script  Copyright (C) 2026  1Vewton
This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
This is free software, and you are welcome to redistribute it
under certain conditions; type `show c' for details.
*/

#ifndef YUKUMO_GENERATOR_H
#define YUKUMO_GENERATOR_H

// generate_wav loads lib_path at runtime, synthesizes UTF-8 koe via AquesTalk2, and writes a WAV.
int generate_wav(char *lib_path, char *phont_path, char *text, char *result_path, int speed);

#endif
