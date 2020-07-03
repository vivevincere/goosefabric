package goosefabric

import(
	"log"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

)

type SdkObjects struct{
	sdk *fabsdk.FabricSDK
	client *channel.Client
}

func NewKey(id string, timestamp string)[][]byte{
return [][]byte{[]byte(id),[]byte(timestamp)}
}


func NewPayload(id string, timestamp string, goosePacket string) [][]byte{
return [][]byte{[]byte(id),[]byte(timestamp), []byte(goosePacket)}

}

func Init(configFile string, channelID string, user string, org string) SdkObjects{
	var toReturn SdkObjects
	var err error
	toReturn.sdk, err = fabsdk.New(config.FromFile(configFile))
	if err != nil {
		log.Fatal("failed to create sdk: %v", err)
	}
	clientContext := toReturn.sdk.ChannelContext(channelID,fabsdk.WithUser(user), fabsdk.WithOrg(org))
	toReturn.client, err = channel.New(clientContext)
	if err != nil {
		log.Fatal("Failed to create new channel: %v", err)
	}
	return toReturn
}

func (thesdk SdkObjects) Set( chaincodeID string, id string, timestamp string, goosePacket string) error{
	defArgs := NewPayload(id, timestamp, goosePacket)
	_, err := thesdk.client.Execute(channel.Request{ChaincodeID: chaincodeID, Fcn: "LogEvent", Args: defArgs})
	if err != nil{
		return err	
	}
	return  nil
}

func (thesdk SdkObjects) Get(chaincodeID string, id string, timestamp string) (string , error){
	defArgs := NewKey(id, timestamp)
	response, err := thesdk.client.Query(channel.Request{ChaincodeID: chaincodeID, Fcn: "QueryEvent", Args: defArgs})
	if err != nil{
		return "", err
	}
	k := string(response.Payload)
	return k, nil
}