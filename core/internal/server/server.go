package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/luckystrike561/vizimind/core/internal/datastore"
	"github.com/luckystrike561/vizimind/core/pkg/regiondo"
	v1 "github.com/luckystrike561/vizimind/core/proto/v1"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type ServerHTTPConfig struct {
	Addr string `koanf:"addr"`
}

type ServerGRPCConfig struct {
	Addr string `koanf:"addr"`
}

type Config struct {
	HTTP   *ServerHTTPConfig `koanf:"http"`
	GRPC   *ServerGRPCConfig `koanf:"grpc"`
	Public bool              `koanf:"public"`
}

type Server struct {
	v1.UnimplementedCoreServiceServer

	cfg          *Config
	grpc         *grpc.Server
	gwMux        *runtime.ServeMux
	http         *http.Server
	listener     net.Listener
	clientConn   *grpc.ClientConn
	datastoreSvc datastore.Datastore
	regiondoSvc  regiondo.Client
}

func New(cfg *Config, datastoreSvc datastore.Datastore, regiondoSvc regiondo.Client) *Server {
	srv := &Server{
		cfg: cfg,
		gwMux: runtime.NewServeMux([]runtime.ServeMuxOption{
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			}),
		}...),
		grpc: grpc.NewServer([]grpc.ServerOption{
			grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
				grpcrecovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
				grpcrecovery.UnaryServerInterceptor(),
			)),
		}...),
		regiondoSvc:  regiondoSvc,
		datastoreSvc: datastoreSvc,
	}

	srv.http = &http.Server{
		Addr:    cfg.HTTP.Addr,
		Handler: srv.gwMux,
	}

	return srv
}

func (s *Server) Init() error {
	var err error
	s.listener, err = net.Listen("tcp", s.cfg.GRPC.Addr)
	grpc_health_v1.RegisterHealthServer(s.grpc, health.NewServer())
	reflection.Register(s.grpc)

	if err != nil {
		return fmt.Errorf("couldn't init server err: %w", err)
	}

	s.http.Handler = cors.AllowAll().Handler(s.http.Handler)

	s.clientConn, err = grpc.DialContext(
		context.Background(),
		s.cfg.GRPC.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	v1.RegisterCoreServiceServer(
		s.grpc, s,
	)

	return v1.RegisterCoreServiceHandler(
		context.Background(), s.gwMux, s.clientConn,
	)
}

func (s *Server) Run() error {
	log.Info().Msgf("starting grpc server on %s", s.cfg.GRPC.Addr)

	go s.grpc.Serve(s.listener)

	log.Info().Msgf("starting http server on %s", s.cfg.HTTP.Addr)

	if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
