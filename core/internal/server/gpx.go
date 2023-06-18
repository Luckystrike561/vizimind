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

func (s *Server) CreateGPX(ctx context.Context, req *v1.CreateGPXRequest) (*v1.GPX, error) {
	gpx, err := s.postgresSvc.UpsertGPX(ctx, &model.GPX{
		ID:   req.GetSku(),
		Data: req.GetGpx(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't create gpx")

		return nil, status.Error(codes.Internal, "couldn't create gpx")
	}

	return &v1.GPX{
		Sku: gpx.ID,
		Gpx: gpx.Data,
	}, nil
}

func (s *Server) UpdateGPX(ctx context.Context, req *v1.UpdateGPXRequest) (*v1.GPX, error) {
	gpx, err := s.postgresSvc.UpsertGPX(ctx, &model.GPX{
		ID:   req.GetSku(),
		Data: req.GetGpx(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't update gpx")

		return nil, status.Error(codes.Internal, "couldn't update gpx")
	}

	return &v1.GPX{
		Sku: gpx.ID,
		Gpx: gpx.Data,
	}, nil
}

func (s *Server) DeleteGPX(ctx context.Context, req *v1.DeleteGPXRequest) (*empty.Empty, error) {
	if err := s.postgresSvc.DeleteGPX(ctx, req.GetSku()); err != nil {
		log.Error().
			Err(err).
			Msg("couldn't delete gpx")

		return nil, status.Error(codes.Internal, "couldn't delete gpx")
	}

	return &empty.Empty{}, nil
}

func (s *Server) ListGPX(ctx context.Context, req *v1.ListGPXRequest) (*v1.ListGPXResponse, error) {
	gpxs, err := s.postgresSvc.ListGPX(ctx, req.GetOffset(), req.GetLimit())
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't list gpx")

		return nil, status.Error(codes.Internal, "couldn't list gpx")
	}

	resp := &v1.ListGPXResponse{
		Items: make([]*v1.GPX, 0, len(gpxs)),
		Total: int32(len(gpxs)),
	}
	for i, gpx := range gpxs {
		resp.Items[i] = &v1.GPX{
			Sku: gpx.ID,
			Gpx: gpx.Data,
		}
	}

	return resp, nil
}

func (s *Server) GetGPX(ctx context.Context, req *v1.GetGPXRequest) (*v1.GPX, error) {
	gpx, err := s.postgresSvc.GetGPX(ctx, req.GetSku())
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't get gpx")

		return nil, status.Error(codes.Internal, "couldn't get gpx")
	}

	return &v1.GPX{
		Sku: gpx.ID,
		Gpx: gpx.Data,
	}, nil
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
