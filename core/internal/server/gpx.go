package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/luckystrike561/vizimind/core/internal/model"
	v1 "github.com/luckystrike561/vizimind/core/proto/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGPX(ctx context.Context, req *v1.CreateGPXRequest) (*empty.Empty, error) {
	if err := s.postgresSvc.CreateGPX(ctx, &model.GPX{
		ID:   req.GetSku(),
		Data: req.GetGpx(),
	}); err != nil {
		log.Error().
			Err(err).
			Msg("couldn't create gpx")

		return nil, status.Error(codes.Internal, "couldn't create gpx")
	}

	return &empty.Empty{}, nil
}

func (s *Server) DownloadGPX(ctx context.Context, req *v1.DownloadGPXRequest) (*v1.DownloadGPXResponse, error) {
	gpx, err := s.postgresSvc.GetGPX(ctx, req.GetSku())
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't get gpx")

		return nil, status.Error(codes.Internal, "couldn't get gpx")
	}

	if gpx == nil {
		return nil, status.Error(codes.NotFound, "gpx not found")
	}

	return &v1.DownloadGPXResponse{
		Gpx: gpx.Data,
	}, nil
}
