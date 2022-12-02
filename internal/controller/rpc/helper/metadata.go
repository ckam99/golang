package helper

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent     string
	ClientIP      string
	Authorization string
}

func GetMetaData(ctx context.Context) *Metadata {
	md, ok := metadata.FromIncomingContext(ctx)
	mtd := &Metadata{}
	if ok {
		/*
		 --- http---
		 grpcgateway-user-agent:
		 grpcgateway-authorization:
		 -- common ---
		 x-forwarded-host:
		 x-forwarded-for:
		 authorization:
		 -- grpc ---
		 user-agent:
		 :authority:
		*/
		if agents := md.Get("grpcgateway-user-agent"); len(agents) > 0 {
			mtd.UserAgent = agents[0]
		}
		if forwarded := md.Get("x-forwarded-for"); len(forwarded) > 0 {
			mtd.ClientIP = forwarded[0]
		}
		if auth := md.Get("authorization"); len(auth) > 0 {
			mtd.Authorization = auth[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		mtd.ClientIP = p.Addr.String()
	}
	return mtd
}
