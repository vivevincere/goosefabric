package goosefabric

import(
	"log"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

)

type SdkObjects struct{
	sdk fabsdk.FabricSDK
	client channel.Client
}

func NewKey(id string, timestamp int64)[][]byte{
return [][]byte{[]byte(id),[]byte(timestamp)}
}


func NewPayload(id string, timestamp int64, goosePacket string){
return [][]byte{[]byte(id),[]byte(timestamp), []byte(goosePacket)}

}

func Init(configFile string, channelID string, user string, org string) SdkObjects{
	var SdkObjects toReturn
	toReturn.sdk, err := fabsdk.New(config.FromFile(configFile))
	if err != nil {
		log.Fatal("failed to create sdk: %v", err)
	}
	clientContext := sdk.ChannelContext(channelID,fabsdk.WithUser(user), fabsdk.WithOrg(org))
	toReturn.client, err := channel.New(clientContext)
	if err != nil {
		log.Fatal("Failed to create new channel: %v", err)
	}
	return toReturn
}

func (thesdk SdkObjects) Set( chaincodeID string, id string, timestamp int64, goosePacket string) error{
	defArgs := NewPayload(id, timestamp, goosePacket)
	response, err := thesdk.client.Execute(channel.Request{ChaincodeID: chaincodeID, Fcn: "LogEvent", Args: defArgs})
	if err != nil{
		return error	
	}
	return  nil
}

func (thesdk SdkObjects) Get(chaincodeID string, id string, timestamp int64) (string , error){
	defArgs := NewKey(id, timestamp)
	response, err := thesdk.client.Query(channel.Request{ChaincodeID: chaincodeID, Fcn: "QueryEvent", Args: defArgs})
	if err != nil{
		return nil, error	
	}
	return string(response.Payload), nil
}