package tensorflow_service_apis

import (
	"github.com/Golang-Tools/feasthelper/feast/serving"
	"github.com/Golang-Tools/grpcsdk"
)

var DefaultOnlineStoreSDK = grpcsdk.New(serving.NewOnlineStoreClient, &serving.OnlineStore_ServiceDesc)
var DefaultServingServiceSDK = grpcsdk.New(serving.NewServingServiceClient, &serving.ServingService_ServiceDesc)
var DefaultTransformationServiceSDK = grpcsdk.New(serving.NewTransformationServiceClient, &serving.TransformationService_ServiceDesc)
