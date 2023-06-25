package server

import (
	"context"

	v1 "github.com/luckystrike561/vizimind/core/proto/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetActivity(
	ctx context.Context,
	req *v1.GetActivityRequest,
) (*v1.GetActivityResponse, error) {
	order, err := s.regiondoSvc.GetOrder(ctx, req.GetOrderId(), req.GetLanguage())
	if err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't get order")

		return nil, status.Error(codes.Internal, "couldn't get order")
	}

	// Get product ID
	product, err := s.regiondoSvc.GetProduct(ctx, order.GetProductID(), req.GetLanguage())
	if err != nil {
		log.Error().
			Err(err).
			Msg("Couldn't get product")

		return nil, status.Error(codes.Internal, "couldn't get product")
	}

	return &v1.GetActivityResponse{
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
	}, nil
}
