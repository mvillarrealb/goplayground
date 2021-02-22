package geocoder

import "context"

type server struct {
}

func (s *server) GetLocation(ctx context.Context, input *GeoCodeRequest) (*GeocodeResponse, error) {
	return &GeocodeResponse{
		LatLng: input.location,
		Addres: "Algun lugar lejano",
	}, nil
}
