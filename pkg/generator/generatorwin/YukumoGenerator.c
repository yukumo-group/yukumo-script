#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "AquesTalk2.h"

// file_load loads the phont file
void* file_load(const char* file, int* psize)
{
    FILE* fp;
    char* data;
    struct _stat st;
    *psize = 0;
    // Get file size
    if (_stat(file, &st) != 0)
        return NULL;
    // If size = 0
    if ((size_t)st.st_size == 0){
        return NULL;
    }
    // Allocate mameory
    data = (char*)malloc(st.st_size);
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
    char* phont_path,
    char* text,
    char* result_path,
    int speed
){
    int size_phont, size_wav;
    void *p_phont = file_load(phont_path, &size_phont);
    if (p_phont == NULL) {
        // File loading error
        return -1;
    }
    unsigned char* wav = AquesTalk2_Synthe(text, speed, &size_wav, p_phont);
    if (wav == NULL) {
        fprintf(stderr, "AquesTalk2 error code: %d\n", size_wav);
        free(p_phont);
        // wav generating error
        return -2;
    }
    free(p_phont);
    FILE* fp = fopen(result_path, "wb");
    if (fp == NULL){
        AquesTalk2_FreeWave(wav);
        // open error
        return -3;
    }
    size_t write_result = fwrite(wav, 1, size_wav, fp);
    if(write_result < size_wav){
        // write incomplete
        return -4;
    }
    fclose(fp);
    AquesTalk2_FreeWave(wav);
    return 0;
}
