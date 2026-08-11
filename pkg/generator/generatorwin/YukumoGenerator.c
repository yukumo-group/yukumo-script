#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <windows.h>
#include "AquesTalk2.h"

typedef unsigned char *(__stdcall *AquesTalk2_Synthe_t)(const char *koe, int iSpeed, int *pSize, void *phontDat);
typedef void(__stdcall *AquesTalk2_FreeWave_t)(unsigned char *wav);

static HMODULE aquestalk2_module = NULL;
static AquesTalk2_Synthe_t p_AquesTalk2_Synthe = NULL;
static AquesTalk2_FreeWave_t p_AquesTalk2_FreeWave = NULL;

// load_aquestalk2 loads AquesTalk2.dll from the known third_party path
static int load_aquestalk2(void)
{
    if (aquestalk2_module != NULL)
    {
        return 0;
    }

    // Prefer path relative to the process working directory (repo root when developing)
    const char *candidates[] = {
        "third_party/aquestalk2/win64/AquesTalk2.dll",
        "third_party\\aquestalk2\\win64\\AquesTalk2.dll",
        "AquesTalk2.dll",
        NULL,
    };

    for (int i = 0; candidates[i] != NULL; i++)
    {
        aquestalk2_module = LoadLibraryA(candidates[i]);
        if (aquestalk2_module != NULL)
        {
            break;
        }
    }

    if (aquestalk2_module == NULL)
    {
        fprintf(stderr, "failed to LoadLibrary AquesTalk2.dll (expected under third_party/aquestalk2/win64)\n");
        return -1;
    }

    p_AquesTalk2_Synthe = (AquesTalk2_Synthe_t)GetProcAddress(aquestalk2_module, "AquesTalk2_Synthe");
    p_AquesTalk2_FreeWave = (AquesTalk2_FreeWave_t)GetProcAddress(aquestalk2_module, "AquesTalk2_FreeWave");
    if (p_AquesTalk2_Synthe == NULL || p_AquesTalk2_FreeWave == NULL)
    {
        fprintf(stderr, "failed to GetProcAddress AquesTalk2 exports\n");
        FreeLibrary(aquestalk2_module);
        aquestalk2_module = NULL;
        return -1;
    }
    return 0;
}

// file_load loads the phont file
void *file_load(const char *file, int *psize)
{
    FILE *fp;
    char *data;
    struct _stat st;
    *psize = 0;
    // Get file size
    if (_stat(file, &st) != 0)
        return NULL;
    // If size = 0
    if ((size_t)st.st_size == 0)
    {
        return NULL;
    }
    // Allocate mameory
    data = (char *)malloc(st.st_size);
    if (data == NULL)
    {
        fprintf(stderr, "can not alloc memory(file_load)\n");
        return NULL;
    }
    // Read file
    fp = fopen(file, "rb");
    if (fp == NULL)
    {
        free(data);
        perror(file);
        return NULL;
    }
    // Read all data
    size_t readLen = fread(data, 1, st.st_size, fp);
    if (readLen < (size_t)st.st_size)
    {
        fprintf(stderr, "can not read data (file_load)\n");
        free(data);
        fclose(fp);
        return NULL;
    }
    fclose(fp);
    *psize = st.st_size;
    return data;
}

// generate_wav generates the wav file
int generate_wav(
    char *phont_path,
    char *text,
    char *result_path,
    int speed)
{
    if (load_aquestalk2() != 0)
    {
        return -1;
    }

    int size_phont, size_wav;
    void *p_phont = file_load(phont_path, &size_phont);
    if (p_phont == NULL)
    {
        // File loading error
        return -1;
    }
    unsigned char *wav = p_AquesTalk2_Synthe(text, speed, &size_wav, p_phont);
    if (wav == NULL)
    {
        fprintf(stderr, "AquesTalk2 error code: %d\n", size_wav);
        free(p_phont);
        // wav generating error
        return -2;
    }
    free(p_phont);
    FILE *fp = fopen(result_path, "wb");
    if (fp == NULL)
    {
        p_AquesTalk2_FreeWave(wav);
        // open error
        return -3;
    }
    size_t write_result = fwrite(wav, 1, size_wav, fp);
    if (write_result < size_wav)
    {
        // write incomplete
        return -4;
    }
    fclose(fp);
    p_AquesTalk2_FreeWave(wav);
    return 0;
}
