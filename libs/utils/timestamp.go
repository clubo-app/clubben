package utils

import "google.golang.org/protobuf/types/known/timestamppb"

func TimestamppIsZero(tsp *timestamppb.Timestamp) bool {
	return tsp.AsTime().Unix() == 0 || tsp.AsTime().IsZero()
}
