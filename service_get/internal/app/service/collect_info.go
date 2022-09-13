package service

import (
	"context"
	ci "github.com/xamust/petbot/service_get/api"
	"github.com/xamust/petbot/service_get/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	collectService     *CollectService
	collectDataClubs   map[string]*ci.ClubsName
	collectDataShedule map[string]*ci.Shedule
}

func (s *server) GetClubs(ctx context.Context, club *ci.Club) (*ci.ClubsName, error) {

	options := options.Find()
	filter := bson.M{}
	results := make(map[string]string)

	cur, err := s.collectService.store.Collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while search data of club names table %v", err)
	}

	for cur.Next(context.TODO()) {
		var elem model.ClubData
		if err = cur.Decode(&elem); err != nil {
			return nil, status.Errorf(codes.Internal, "error while search data of club names table", err)
		}
		results[elem.Name] = elem.Url
	}

	if err = cur.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "error with cursor", err)
	}

	if err = cur.Close(context.TODO()); err != nil {
		return nil, status.Errorf(codes.Internal, "error with close cursor", err)
	}

	//if err = s.collectService.store.Disconnect(); err != nil {
	//	return nil, status.Errorf(codes.Internal, "error while database disconnect", err)
	//}
	return &ci.ClubsName{ClubsName: results}, status.New(codes.OK, "").Err()
}

func (s *server) GetShedule(ctx context.Context, url *ci.ClubUrl) (*ci.Shedule, error) {
	//TODO implement me
	panic("implement me")
}
