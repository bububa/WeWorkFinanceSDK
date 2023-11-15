# 企业微信会话内容存档SDK golang版 

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/WeWorkFinanceSDK.svg)](https://pkg.go.dev/github.com/bububa/WeWorkFinanceSDK)
[![Go](https://github.com/bububa/WeWorkFinanceSDK/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/WeWorkFinanceSDK/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/WeWorkFinanceSDK/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/WeWorkFinanceSDK/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/WeWorkFinanceSDK.svg)](https://github.com/bububa/WeWorkFinanceSDK)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/WeWorkFinanceSDK)](https://goreportcard.com/report/github.com/bububa/WeWorkFinanceSDK)
[![GitHub license](https://img.shields.io/github/license/bububa/WeWorkFinanceSDK.svg)](https://github.com/bububa/WeWorkFinanceSDK/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/WeWorkFinanceSDK.svg)](https://GitHub.com/bububa/WeWorkFinanceSDK/releases/)

用于拉取企业聊天记录与媒体消息，该SDK是对官方linux c库的封装

## Usage
```golang
package main

import (
    "log"
    "fmt"
    "bytes"

    sdk "github.com/bububa/WeWorkFinanceSDK"
)

func main() {
    corpId := "企业ID"
    corpSecret := "secret"
    clt, err := sdk.NewClient(corpId, corpSecret)
    if err != nil {
        log.Fatalln(err)
    }
    defer clt.Free()
    var (
        seq uint64 
        limit uint64 = 1000
        proxy = ""
        passwd = ""
        timeout = 300
    )
    for {
        chatsData, err := clt.GetChatData(seq, limit, proxy, passwd, timeout)
        if err != nil {
            log.Fatalln(err)
        }
        if len(chatsData) == 0 {
            break
        }
        // iterate chatsData
        for _, data := range chatsData {
            // update seq with chatdata.Seq for next loop
            seq = data.Seq
            // decrypt chat data
            msg, err := clt.DecryptData(data.EncryptRandomKey, data.EncryptChatMsg)
            if err != nil {
                log.Fatalln(err)
                continue
            }
            fmt.Printf("msg: %+v\n", msg)
            if msg.MessageType() == sdk.IMG_MSG {
                var (
                    w = new(bytes.Buffer) 
                    sdkField = msg.(sdk.ImageMessage).SdkFieldId
                )
                err := clt.DownloadMedia(w, sdkField, proxy, passwd, timeout)
                if err != nil {
                    log.Fatalln(err)
                    continue
                }
                log.Printf("download mediadata, len:%d bytes\n", w.Len())
            }
        }
    }
}
```

## Reference
[企业微信获取会话内容文档](https://work.weixin.qq.com/api/doc/90000/90135/91774)

