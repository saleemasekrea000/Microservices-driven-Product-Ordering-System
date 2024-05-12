package product

import (
	"context"
	"product_order/internal/application/core/domains"
	proto_gen "product_order/internal_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductAdapter struct {
	client proto_gen.ProductServiceClient
}

func New(product_url string) (*ProductAdapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(product_url, opts...)
	if err != nil {
		return nil, err
	}
	// defer conn.Close()
	cleint := proto_gen.NewProductServiceClient(conn)
	return &ProductAdapter{client: cleint}, nil
}

func (p *ProductAdapter) GetProduct(ctx context.Context, id int32) (domains.Product, error) {
	res, err := p.client.GetProduct(ctx, &proto_gen.GetProductRequest{Id: id})
	if err != nil {
		return domains.Product{}, err
	}
	product := domains.Product{
		ID:        res.Product.Id,
		Name:      res.Product.Name,
		Price:     res.Product.Price,
		Image:     res.Product.Image,
		CreatedAt: res.Product.CreatedAt.AsTime(),
		UpdatedAt: res.Product.UpdatedAt.AsTime(),
	}
	return product, nil
}
