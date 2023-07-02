package server

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/luckystrike561/vizimind/core/internal/model"
	v1 "github.com/luckystrike561/vizimind/core/proto/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetActivityByOrderID(
	ctx context.Context,
	req *v1.GetActivityByOrderIDRequest,
) (*v1.GetActivityByOrderIDResponse, error) {
	order, err := s.regiondoSvc.GetOrder(ctx, req.GetOrderId(), req.GetLanguage())
	if err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't get order")

		return nil, status.Error(codes.Internal, "couldn't get order")
	}

	product, err := s.regiondoSvc.GetProduct(ctx, order.GetProductID(), req.GetLanguage())
	if err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't get product")

		return nil, status.Error(codes.Internal, "couldn't get product")
	}

	activity, err := s.datastoreSvc.GetActivity(ctx, strings.ToLower(product.Data.Sku))
	if err != nil {
		log.Error().
			Err(err).
			Str("sku", product.Data.Sku).
			Msg("Couldn't get activity")

		return nil, status.Error(codes.Internal, "couldn't get activity")
	}

	return &v1.GetActivityByOrderIDResponse{
		OrderId:          req.GetOrderId(),
		ProductId:        order.GetProductID(),
		Name:             product.Data.Name,
		Sku:              product.Data.Sku,
		ShortDescription: product.Data.ShortDescription,
		GeoLat:           product.Data.GeoLat,
		GeoLon:           product.Data.GeoLon,
		Distance:         product.Data.Distance,
		LocationAddress:  product.Data.LocationAddress,
		City:             product.Data.City,
		Zipcode:          product.Data.Zipcode,
		CityId:           product.Data.CityID,
		RegionId:         product.Data.RegionID,
		PoiIds:           product.Data.PoiIDS,
		CountryId:        product.Data.CountryID,
		Thumbnail:        product.Data.Thumbnail,
		Image:            product.Data.Image,
		ImageLabel:       product.Data.ImageLabel,
		RatingSummary:    product.Data.RatingSummary,
		ReviewsCount:     product.Data.ReviewsCount,
		Description:      product.Data.Description,
		Date:             order.GetEventDate(),
		PurchasedAt:      order.PurchasedAt,
		Resources:        order.GetResources(),
		Details: &v1.Activity{
			Sku:          activity.ID,
			Gpx:          activity.GPX,
			ProductId:    activity.ProductID,
			Names:        activity.Names,
			Descriptions: activity.Descriptions,
			Image:        activity.Image,
			Transports:   activity.Transports,
			Supplier: &v1.Supplier{
				Name:    activity.Supplier.Name,
				Email:   activity.Supplier.Email,
				Phone:   activity.Supplier.Phone,
				Address: activity.Supplier.Address,
				City:    activity.Supplier.City,
				Zipcode: activity.Supplier.Zipcode,
				Country: activity.Supplier.Country,
			},
			ExtraMeetingInfo: activity.ExtraMeetingInfo,
		},
	}, nil
}

func (s *Server) CreateActivity(ctx context.Context, req *v1.CreateActivityRequest) (*v1.Activity, error) {
	if s.cfg.Public == true {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	activity, err := s.datastoreSvc.UpsertActivity(ctx, &model.Activity{
		ID:           strings.ToLower(req.GetSku()),
		GPX:          req.GetGpx(),
		ProductID:    req.GetProductId(),
		Names:        req.GetNames(),
		Descriptions: req.GetDescriptions(),
		Image:        req.GetImage(),
		Transports:   req.GetTransports(),
		Supplier: model.Supplier{
			Name:    req.GetSupplier().GetName(),
			Email:   req.GetSupplier().GetEmail(),
			Phone:   req.GetSupplier().GetPhone(),
			Address: req.GetSupplier().GetAddress(),
			City:    req.GetSupplier().GetCity(),
			Zipcode: req.GetSupplier().GetZipcode(),
			Country: req.GetSupplier().GetCountry(),
		},
		ExtraMeetingInfo: req.GetExtraMeetingInfo(),
	})
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't create activity")

		return nil, status.Error(codes.Internal, "couldn't create activity")
	}

	return &v1.Activity{
		Sku:          activity.ID,
		Gpx:          activity.GPX,
		ProductId:    activity.ProductID,
		Names:        activity.Names,
		Descriptions: activity.Descriptions,
		Image:        activity.Image,
		Transports:   activity.Transports,
		Supplier: &v1.Supplier{
			Name:    activity.Supplier.Name,
			Email:   activity.Supplier.Email,
			Phone:   activity.Supplier.Phone,
			Address: activity.Supplier.Address,
			City:    activity.Supplier.City,
			Zipcode: activity.Supplier.Zipcode,
			Country: activity.Supplier.Country,
		},
		ExtraMeetingInfo: activity.ExtraMeetingInfo,
	}, nil
}

func (s *Server) DeleteActivity(ctx context.Context, req *v1.DeleteActivityRequest) (*empty.Empty, error) {
	if s.cfg.Public == true {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	if err := s.datastoreSvc.DeleteActivity(ctx, strings.ToLower(req.GetSku())); err != nil {
		log.Error().
			Err(err).
			Msg("couldn't delete activity")

		return nil, status.Error(codes.Internal, "couldn't delete activity")
	}

	return &empty.Empty{}, nil
}

func (s *Server) ListActivity(ctx context.Context, req *v1.ListActivityRequest) (*v1.ListActivityResponse, error) {
	if s.cfg.Public == true {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	offset := req.GetOffset()
	limit := req.GetLimit()

	if offset <= 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = 10
	}

	activities, err := s.datastoreSvc.ListActivity(ctx, offset, limit)
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't list activities")

		return nil, status.Error(codes.Internal, "couldn't list activities")
	}

	resp := &v1.ListActivityResponse{
		Items: make([]*v1.Activity, len(activities)),
		Total: int32(len(activities)),
	}
	for i, activity := range activities {
		resp.Items[i] = &v1.Activity{
			Sku:          activity.ID,
			Gpx:          activity.GPX,
			ProductId:    activity.ProductID,
			Names:        activity.Names,
			Descriptions: activity.Descriptions,
			Image:        activity.Image,
			Transports:   activity.Transports,
			Supplier: &v1.Supplier{
				Name:    activity.Supplier.Name,
				Email:   activity.Supplier.Email,
				Phone:   activity.Supplier.Phone,
				Address: activity.Supplier.Address,
				City:    activity.Supplier.City,
				Zipcode: activity.Supplier.Zipcode,
				Country: activity.Supplier.Country,
			},
			ExtraMeetingInfo: activity.ExtraMeetingInfo,
		}
	}

	return resp, nil
}

func (s *Server) GetActivity(ctx context.Context, req *v1.GetActivityRequest) (*v1.Activity, error) {
	if s.cfg.Public == true {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	activity, err := s.datastoreSvc.GetActivity(ctx, strings.ToLower(req.GetSku()))
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't get activity")

		return nil, status.Error(codes.Internal, "couldn't get activity")
	}

	return &v1.Activity{
		Sku:          activity.ID,
		Gpx:          activity.GPX,
		ProductId:    activity.ProductID,
		Names:        activity.Names,
		Descriptions: activity.Descriptions,
		Image:        activity.Image,
		Transports:   activity.Transports,
		Supplier: &v1.Supplier{
			Name:    activity.Supplier.Name,
			Email:   activity.Supplier.Email,
			Phone:   activity.Supplier.Phone,
			Address: activity.Supplier.Address,
			City:    activity.Supplier.City,
			Zipcode: activity.Supplier.Zipcode,
			Country: activity.Supplier.Country,
		},
		ExtraMeetingInfo: activity.ExtraMeetingInfo,
	}, nil
}

func (s *Server) DownloadGPX(ctx context.Context, req *v1.DownloadGPXRequest) (*v1.DownloadGPXResponse, error) {
	activity, err := s.datastoreSvc.GetActivity(ctx, strings.ToLower(req.GetSku()))
	if err != nil {
		log.Error().
			Err(err).
			Msg("couldn't get activity")

		return nil, status.Error(codes.Internal, "couldn't get activity")
	}

	if activity == nil {
		return nil, status.Error(codes.NotFound, "activity not found")
	}

	return &v1.DownloadGPXResponse{
		Gpx: activity.GPX,
	}, nil
}
