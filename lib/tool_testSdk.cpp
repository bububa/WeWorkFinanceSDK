#include "WeWorkFinanceSdk_C.h"
#include <dlfcn.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string>
using std::string;

typedef WeWorkFinanceSdk_t* newsdk_t();
typedef int Init_t(WeWorkFinanceSdk_t*, const char*, const char*);

typedef int GetChatData_t(WeWorkFinanceSdk_t*, unsigned long long, unsigned int, const char*, const char*, int, Slice_t*);
typedef Slice_t* NewSlice_t();
typedef void FreeSlice_t(Slice_t*);

typedef int GetMediaData_t(WeWorkFinanceSdk_t*, const char*, const char*, const char*, const char*, int, MediaData_t*);
typedef int DecryptData_t(const char*, const char*, Slice_t*);
typedef MediaData_t* NewMediaData_t();
typedef void FreeMediaData_t(MediaData_t*);

int main(int argc, char* argv[])
{
    int ret = 0;
    if (argc < 2) {
        printf("./sdktools 1(chatmsg) 2(mediadata) 3(decryptdata)\n");
        printf("./sdktools 1 seq limit proxy passwd timeout\n");
        printf("./sdktools 2 fileid proxy passwd timeout savefile\n");
        printf("./sdktools 3 encrypt_key encrypt_msg\n");
        return -1;
    }

    // new sdk api
    void* so_handle = dlopen("./libWeWorkFinanceSdk_C.so", RTLD_LAZY);
    if (!so_handle) {
        printf("load sdk so fail:%s\n", dlerror());
        return -1;
    }
    newsdk_t* newsdk_fn = (newsdk_t*)dlsym(so_handle, "NewSdk");
    WeWorkFinanceSdk_t* sdk = newsdk_fn();

    // init api
    Init_t* init_fn = (Init_t*)dlsym(so_handle, "Init");
    ret = init_fn(sdk, "wwd08c8e7c775ab44d", "zJ6k0naVVQ--gt9PUSSEvs03zW_nlDVmjLCTOTAfrew");
    if (ret != 0) {
        printf("init sdk err ret:%d\n", ret);
        return -1;
    }

    int type = strtoul(argv[1], NULL, 10);
    if (type == 1) {
        uint64_t iSeq = strtoul(argv[2], NULL, 10);
        uint64_t iLimit = strtoul(argv[3], NULL, 10);
        uint64_t timeout = strtoul(argv[6], NULL, 10);
        
        NewSlice_t* newslice_fn = (NewSlice_t*)dlsym(so_handle, "NewSlice");
        FreeSlice_t* freeslice_fn = (FreeSlice_t*)dlsym(so_handle, "FreeSlice");

        Slice_t* chatDatas = newslice_fn();
        // getchatdata api
        GetChatData_t* getchatdata_fn = (GetChatData_t*)dlsym(so_handle, "GetChatData");
        ret = getchatdata_fn(sdk, iSeq, iLimit, argv[4], argv[5], timeout, chatDatas);
        if (ret != 0) {
            printf("GetChatData err ret:%d\n", ret);
            return -1;
        }
        printf("GetChatData len:%d data:%s\n", chatDatas->len, chatDatas->buf);
        freeslice_fn(chatDatas);
    } else if (type == 2) {
        std::string index;
        uint64_t timeout = strtoul(argv[5], NULL, 10);
        int isfinish = 0;

        GetMediaData_t* getmediadata_fn = (GetMediaData_t*)dlsym(so_handle, "GetMediaData");
        NewMediaData_t* newmediadata_fn = (NewMediaData_t*)dlsym(so_handle, "NewMediaData");
        FreeMediaData_t* freemediadata_fn = (FreeMediaData_t*)dlsym(so_handle, "FreeMediaData");

        while (isfinish == 0) {
            printf("index:%s\n", index.c_str());
            MediaData_t* mediaData = newmediadata_fn();
            // getmediadata api
            ret = getmediadata_fn(sdk, index.c_str(), argv[2], argv[3], argv[4], timeout, mediaData);
            if (ret != 0) {
                printf("GetMediaData err ret:%d\n", ret);
                return -1;
            }
            printf("content size:%d isfin:%d outindex:%s\n", mediaData->data_len, mediaData->is_finish, mediaData->outindexbuf);

            char file[200];
            snprintf(file, sizeof(file), "%s", argv[6]);
            FILE* fp = fopen(file, "ab+");
            printf("filename:%s \n", file);
            if (NULL == fp) {
                printf("open file err\n");
                return -1;
            }

            fwrite(mediaData->data, mediaData->data_len, 1, fp);
            fclose(fp);

            index.assign(string(mediaData->outindexbuf));
            isfinish = mediaData->is_finish;
            freemediadata_fn(mediaData);
        }
    } else if (type == 3) {
        NewSlice_t* newslice_fn = (NewSlice_t*)dlsym(so_handle, "NewSlice");
        FreeSlice_t* freeslice_fn = (FreeSlice_t*)dlsym(so_handle, "FreeSlice");

        Slice_t* Msgs = newslice_fn();
        // decryptdata api
        DecryptData_t* decryptdata_fn = (DecryptData_t*)dlsym(so_handle, "DecryptData");
        ret = decryptdata_fn(argv[2], argv[3], Msgs);
        printf("chatdata :%s ret :%d\n", Msgs->buf, ret);

        freeslice_fn(Msgs);
    }

    return ret;
}
