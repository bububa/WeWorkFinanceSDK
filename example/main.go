package main

import (
	"bytes"
	"fmt"
	"log"

	sdk "github.com/bububa/WeWorkFinanceSDK"
)

func main() {
	// CorpId 企业ID
	corpId := "wwd08c8e7c775ab44d"
	corpSecret := "zJ6k0naVVQ--gt9PUSSEvs03zW_nlDVmjLCTOTAfrew"
	clt, err := sdk.NewClient(corpId, corpSecret)
	if err != nil {
		log.Fatalln(err)
	}
	defer clt.Destroy()
	var (
		seq     uint64
		limit   uint64 = 1000
		proxy          = ""
		passwd         = ""
		timeout        = 300
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
					w        = new(bytes.Buffer)
					sdkField = msg.(sdk.ImageMessage).SdkFileId
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
