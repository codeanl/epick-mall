// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package pmsclient

import (
	"context"

	"epick-mall/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BrandAddReq                        = pms.BrandAddReq
	BrandAddResp                       = pms.BrandAddResp
	BrandDeleteReq                     = pms.BrandDeleteReq
	BrandDeleteResp                    = pms.BrandDeleteResp
	BrandListByIdsReq                  = pms.BrandListByIdsReq
	BrandListData                      = pms.BrandListData
	BrandListReq                       = pms.BrandListReq
	BrandListResp                      = pms.BrandListResp
	BrandUpdateReq                     = pms.BrandUpdateReq
	BrandUpdateResp                    = pms.BrandUpdateResp
	ProductAddReq                      = pms.ProductAddReq
	ProductAddResp                     = pms.ProductAddResp
	ProductAttributeAddReq             = pms.ProductAttributeAddReq
	ProductAttributeAddResp            = pms.ProductAttributeAddResp
	ProductAttributeCategoryAddReq     = pms.ProductAttributeCategoryAddReq
	ProductAttributeCategoryAddResp    = pms.ProductAttributeCategoryAddResp
	ProductAttributeCategoryDeleteReq  = pms.ProductAttributeCategoryDeleteReq
	ProductAttributeCategoryDeleteResp = pms.ProductAttributeCategoryDeleteResp
	ProductAttributeCategoryListData   = pms.ProductAttributeCategoryListData
	ProductAttributeCategoryListReq    = pms.ProductAttributeCategoryListReq
	ProductAttributeCategoryListResp   = pms.ProductAttributeCategoryListResp
	ProductAttributeCategoryUpdateReq  = pms.ProductAttributeCategoryUpdateReq
	ProductAttributeCategoryUpdateResp = pms.ProductAttributeCategoryUpdateResp
	ProductAttributeDeleteReq          = pms.ProductAttributeDeleteReq
	ProductAttributeDeleteResp         = pms.ProductAttributeDeleteResp
	ProductAttributeListData           = pms.ProductAttributeListData
	ProductAttributeListReq            = pms.ProductAttributeListReq
	ProductAttributeListResp           = pms.ProductAttributeListResp
	ProductAttributeUpdateReq          = pms.ProductAttributeUpdateReq
	ProductAttributeUpdateResp         = pms.ProductAttributeUpdateResp
	ProductCategoryAddReq              = pms.ProductCategoryAddReq
	ProductCategoryAddResp             = pms.ProductCategoryAddResp
	ProductCategoryDeleteReq           = pms.ProductCategoryDeleteReq
	ProductCategoryDeleteResp          = pms.ProductCategoryDeleteResp
	ProductCategoryFirstListReq        = pms.ProductCategoryFirstListReq
	ProductCategoryListData            = pms.ProductCategoryListData
	ProductCategoryListReq             = pms.ProductCategoryListReq
	ProductCategoryListResp            = pms.ProductCategoryListResp
	ProductCategorySecondListReq       = pms.ProductCategorySecondListReq
	ProductCategoryUpdateReq           = pms.ProductCategoryUpdateReq
	ProductCategoryUpdateResp          = pms.ProductCategoryUpdateResp
	ProductDeleteReq                   = pms.ProductDeleteReq
	ProductDeleteResp                  = pms.ProductDeleteResp
	ProductListData                    = pms.ProductListData
	ProductListReq                     = pms.ProductListReq
	ProductListResp                    = pms.ProductListResp
	ProductUpdateReq                   = pms.ProductUpdateReq
	ProductUpdateResp                  = pms.ProductUpdateResp

	Pms interface {
		ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error)
		ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error)
		// rpc ProductListByIds(ProductByIdsReq) returns(ProductListResp);
		ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error)
		ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error)
		BrandAdd(ctx context.Context, in *BrandAddReq, opts ...grpc.CallOption) (*BrandAddResp, error)
		BrandList(ctx context.Context, in *BrandListReq, opts ...grpc.CallOption) (*BrandListResp, error)
		BrandListByIds(ctx context.Context, in *BrandListByIdsReq, opts ...grpc.CallOption) (*BrandListResp, error)
		BrandUpdate(ctx context.Context, in *BrandUpdateReq, opts ...grpc.CallOption) (*BrandUpdateResp, error)
		BrandDelete(ctx context.Context, in *BrandDeleteReq, opts ...grpc.CallOption) (*BrandDeleteResp, error)
		ProductCategoryAdd(ctx context.Context, in *ProductCategoryAddReq, opts ...grpc.CallOption) (*ProductCategoryAddResp, error)
		ProductCategoryList(ctx context.Context, in *ProductCategoryListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error)
		ProductCategoryFirstList(ctx context.Context, in *ProductCategoryFirstListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error)
		ProductCategorySecondList(ctx context.Context, in *ProductCategorySecondListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error)
		ProductCategoryUpdate(ctx context.Context, in *ProductCategoryUpdateReq, opts ...grpc.CallOption) (*ProductCategoryUpdateResp, error)
		ProductCategoryDelete(ctx context.Context, in *ProductCategoryDeleteReq, opts ...grpc.CallOption) (*ProductCategoryDeleteResp, error)
		AttributeAdd(ctx context.Context, in *ProductAttributeAddReq, opts ...grpc.CallOption) (*ProductAttributeAddResp, error)
		AttributeList(ctx context.Context, in *ProductAttributeListReq, opts ...grpc.CallOption) (*ProductAttributeListResp, error)
		AttributeUpdate(ctx context.Context, in *ProductAttributeUpdateReq, opts ...grpc.CallOption) (*ProductAttributeUpdateResp, error)
		AttributeDelete(ctx context.Context, in *ProductAttributeDeleteReq, opts ...grpc.CallOption) (*ProductAttributeDeleteResp, error)
		AttributeCategoryAdd(ctx context.Context, in *ProductAttributeCategoryAddReq, opts ...grpc.CallOption) (*ProductAttributeCategoryAddResp, error)
		AttributeCategoryList(ctx context.Context, in *ProductAttributeCategoryListReq, opts ...grpc.CallOption) (*ProductAttributeCategoryListResp, error)
		AttributeCategoryUpdate(ctx context.Context, in *ProductAttributeCategoryUpdateReq, opts ...grpc.CallOption) (*ProductAttributeCategoryUpdateResp, error)
		AttributeCategoryDelete(ctx context.Context, in *ProductAttributeCategoryDeleteReq, opts ...grpc.CallOption) (*ProductAttributeCategoryDeleteResp, error)
	}

	defaultPms struct {
		cli zrpc.Client
	}
)

func NewPms(cli zrpc.Client) Pms {
	return &defaultPms{
		cli: cli,
	}
}

func (m *defaultPms) ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductAdd(ctx, in, opts...)
}

func (m *defaultPms) ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

// rpc ProductListByIds(ProductByIdsReq) returns(ProductListResp);
func (m *defaultPms) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductUpdate(ctx, in, opts...)
}

func (m *defaultPms) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductDelete(ctx, in, opts...)
}

func (m *defaultPms) BrandAdd(ctx context.Context, in *BrandAddReq, opts ...grpc.CallOption) (*BrandAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.BrandAdd(ctx, in, opts...)
}

func (m *defaultPms) BrandList(ctx context.Context, in *BrandListReq, opts ...grpc.CallOption) (*BrandListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.BrandList(ctx, in, opts...)
}

func (m *defaultPms) BrandListByIds(ctx context.Context, in *BrandListByIdsReq, opts ...grpc.CallOption) (*BrandListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.BrandListByIds(ctx, in, opts...)
}

func (m *defaultPms) BrandUpdate(ctx context.Context, in *BrandUpdateReq, opts ...grpc.CallOption) (*BrandUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.BrandUpdate(ctx, in, opts...)
}

func (m *defaultPms) BrandDelete(ctx context.Context, in *BrandDeleteReq, opts ...grpc.CallOption) (*BrandDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.BrandDelete(ctx, in, opts...)
}

func (m *defaultPms) ProductCategoryAdd(ctx context.Context, in *ProductCategoryAddReq, opts ...grpc.CallOption) (*ProductCategoryAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategoryAdd(ctx, in, opts...)
}

func (m *defaultPms) ProductCategoryList(ctx context.Context, in *ProductCategoryListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategoryList(ctx, in, opts...)
}

func (m *defaultPms) ProductCategoryFirstList(ctx context.Context, in *ProductCategoryFirstListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategoryFirstList(ctx, in, opts...)
}

func (m *defaultPms) ProductCategorySecondList(ctx context.Context, in *ProductCategorySecondListReq, opts ...grpc.CallOption) (*ProductCategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategorySecondList(ctx, in, opts...)
}

func (m *defaultPms) ProductCategoryUpdate(ctx context.Context, in *ProductCategoryUpdateReq, opts ...grpc.CallOption) (*ProductCategoryUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategoryUpdate(ctx, in, opts...)
}

func (m *defaultPms) ProductCategoryDelete(ctx context.Context, in *ProductCategoryDeleteReq, opts ...grpc.CallOption) (*ProductCategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductCategoryDelete(ctx, in, opts...)
}

func (m *defaultPms) AttributeAdd(ctx context.Context, in *ProductAttributeAddReq, opts ...grpc.CallOption) (*ProductAttributeAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeAdd(ctx, in, opts...)
}

func (m *defaultPms) AttributeList(ctx context.Context, in *ProductAttributeListReq, opts ...grpc.CallOption) (*ProductAttributeListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeList(ctx, in, opts...)
}

func (m *defaultPms) AttributeUpdate(ctx context.Context, in *ProductAttributeUpdateReq, opts ...grpc.CallOption) (*ProductAttributeUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeUpdate(ctx, in, opts...)
}

func (m *defaultPms) AttributeDelete(ctx context.Context, in *ProductAttributeDeleteReq, opts ...grpc.CallOption) (*ProductAttributeDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeDelete(ctx, in, opts...)
}

func (m *defaultPms) AttributeCategoryAdd(ctx context.Context, in *ProductAttributeCategoryAddReq, opts ...grpc.CallOption) (*ProductAttributeCategoryAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryAdd(ctx, in, opts...)
}

func (m *defaultPms) AttributeCategoryList(ctx context.Context, in *ProductAttributeCategoryListReq, opts ...grpc.CallOption) (*ProductAttributeCategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryList(ctx, in, opts...)
}

func (m *defaultPms) AttributeCategoryUpdate(ctx context.Context, in *ProductAttributeCategoryUpdateReq, opts ...grpc.CallOption) (*ProductAttributeCategoryUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryUpdate(ctx, in, opts...)
}

func (m *defaultPms) AttributeCategoryDelete(ctx context.Context, in *ProductAttributeCategoryDeleteReq, opts ...grpc.CallOption) (*ProductAttributeCategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryDelete(ctx, in, opts...)
}
