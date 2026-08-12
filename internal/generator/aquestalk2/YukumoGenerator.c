/*
yukumo-script  Copyright (C) 2026  1Vewton
This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
This is free software, and you are welcome to redistribute it
under certain conditions; type `show c' for details.
*/

#include "YukumoGenerator.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>

#ifdef _WIN32
#include <windows.h>
typedef unsigned char *(__stdcall *AquesTalk2_Synthe_Utf8_t)(const char *koe, int iSpeed, int *pSize, void *phontDat);
typedef void(__stdcall *AquesTalk2_FreeWave_t)(unsigned char *wav);
typedef HMODULE lib_handle_t;
#else
#include <dlfcn.h>
typedef unsigned char *(*AquesTalk2_Synthe_Utf8_t)(const char *koe, int iSpeed, int *pSize, void *phontDat);
typedef void (*AquesTalk2_FreeWave_t)(unsigned char *wav);
typedef void *lib_handle_t;
#endif

static lib_handle_t aquestalk2_module = NULL;
static AquesTalk2_Synthe_Utf8_t p_AquesTalk2_Synthe_Utf8 = NULL;
static AquesTalk2_FreeWave_t p_AquesTalk2_FreeWave = NULL;
static char loaded_lib_path[1024] = {0};

static void *sym_lookup(lib_handle_t handle, const char *name)
{
#ifdef _WIN32
    return (void *)GetProcAddress(handle, name);
#else
    return dlsym(handle, name);
#endif
}

static void lib_close(lib_handle_t handle)
{
#ifdef _WIN32
    FreeLibrary(handle);
#else
    dlclose(handle);
#endif
}

static lib_handle_t lib_open(const char *path)
{
#ifdef _WIN32
    return LoadLibraryA(path);
#else
    return dlopen(path, RTLD_NOW);
#endif
}

// load_aquestalk2 loads the AquesTalk2 shared library from lib_path and resolves common exports.
static int load_aquestalk2(const char *lib_path)
{
    if (aquestalk2_module != NULL)
    {
        if (strcmp(loaded_lib_path, lib_path) == 0)
        {
            return 0;
        }
        lib_close(aquestalk2_module);
        aquestalk2_module = NULL;
        p_AquesTalk2_Synthe_Utf8 = NULL;
        p_AquesTalk2_FreeWave = NULL;
        loaded_lib_path[0] = '\0';
    }

    aquestalk2_module = lib_open(lib_path);
    if (aquestalk2_module == NULL)
    {
        fprintf(stderr, "failed to load AquesTalk2 library: %s\n", lib_path);
        return -1;
    }

    p_AquesTalk2_Synthe_Utf8 = (AquesTalk2_Synthe_Utf8_t)sym_lookup(aquestalk2_module, "AquesTalk2_Synthe_Utf8");
    p_AquesTalk2_FreeWave = (AquesTalk2_FreeWave_t)sym_lookup(aquestalk2_module, "AquesTalk2_FreeWave");
    if (p_AquesTalk2_Synthe_Utf8 == NULL || p_AquesTalk2_FreeWave == NULL)
    {
        fprintf(stderr, "failed to resolve AquesTalk2_Synthe_Utf8 / AquesTalk2_FreeWave\n");
        lib_close(aquestalk2_module);
        aquestalk2_module = NULL;
        return -1;
    }

    strncpy(loaded_lib_path, lib_path, sizeof(loaded_lib_path) - 1);
    loaded_lib_path[sizeof(loaded_lib_path) - 1] = '\0';
    return 0;
}

// file_load loads the phont file into memory.
static void *file_load(const char *file, int *psize)
{
    FILE *fp;
    char *data;
    struct stat st;
    *psize = 0;
    if (stat(file, &st) != 0)
    {
        return NULL;
    }
    if ((size_t)st.st_size == 0)
    {
        return NULL;
    }
    data = (char *)malloc((size_t)st.st_size);
    if (data == NULL)
    {
        fprintf(stderr, "can not alloc memory(file_load)\n");
        return NULL;
    }
    fp = fopen(file, "rb");
    if (fp == NULL)
    {
        free(data);
        perror(file);
        return NULL;
    }
    size_t readLen = fread(data, 1, (size_t)st.st_size, fp);
    if (readLen < (size_t)st.st_size)
    {
        fprintf(stderr, "can not read data (file_load)\n");
        free(data);
        fclose(fp);
        return NULL;
    }
    fclose(fp);
    *psize = (int)st.st_size;
    return data;
}

// generate_wav generates the wav file using AquesTalk2_Synthe_Utf8.
int generate_wav(
    char *lib_path,
    char *phont_path,
    char *text,
    char *result_path,
    int speed)
{
    if (lib_path == NULL || load_aquestalk2(lib_path) != 0)
    {
        return -1;
    }

    int size_phont, size_wav;
    void *p_phont = file_load(phont_path, &size_phont);
    if (p_phont == NULL)
    {
        return -1;
    }
    unsigned char *wav = p_AquesTalk2_Synthe_Utf8(text, speed, &size_wav, p_phont);
    if (wav == NULL)
    {
        fprintf(stderr, "AquesTalk2 error code: %d\n", size_wav);
        free(p_phont);
        return -2;
    }
    free(p_phont);
    FILE *fp = fopen(result_path, "wb");
    if (fp == NULL)
    {
        p_AquesTalk2_FreeWave(wav);
        return -3;
    }
    size_t write_result = fwrite(wav, 1, (size_t)size_wav, fp);
    if (write_result < (size_t)size_wav)
    {
        fclose(fp);
        p_AquesTalk2_FreeWave(wav);
        return -4;
    }
    fclose(fp);
    p_AquesTalk2_FreeWave(wav);
    return 0;
}
