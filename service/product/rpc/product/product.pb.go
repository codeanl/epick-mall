// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: product.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId                    int64   `protobuf:"varint,2,opt,name=BrandId,proto3" json:"BrandId,omitempty"`
	ProductCategoryId          int64   `protobuf:"varint,3,opt,name=ProductCategoryId,proto3" json:"ProductCategoryId,omitempty"`
	FreightTemplateId          int64   `protobuf:"varint,4,opt,name=FreightTemplateId,proto3" json:"FreightTemplateId,omitempty"`
	ProductAttributeCategoryId int64   `protobuf:"varint,5,opt,name=ProductAttributeCategoryId,proto3" json:"ProductAttributeCategoryId,omitempty"`
	Name                       string  `protobuf:"bytes,6,opt,name=Name,proto3" json:"Name,omitempty"`
	Pic                        string  `protobuf:"bytes,7,opt,name=Pic,proto3" json:"Pic,omitempty"`
	ProductSn                  string  `protobuf:"bytes,8,opt,name=ProductSn,proto3" json:"ProductSn,omitempty"`
	DeleteStatus               string  `protobuf:"bytes,9,opt,name=DeleteStatus,proto3" json:"DeleteStatus,omitempty"`
	NewStatus                  string  `protobuf:"bytes,11,opt,name=NewStatus,proto3" json:"NewStatus,omitempty"`
	RecommendStatus            string  `protobuf:"bytes,12,opt,name=RecommendStatus,proto3" json:"RecommendStatus,omitempty"`
	Sort                       string  `protobuf:"bytes,14,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Sale                       string  `protobuf:"bytes,15,opt,name=Sale,proto3" json:"Sale,omitempty"`
	Price                      float32 `protobuf:"fixed32,16,opt,name=Price,proto3" json:"Price,omitempty"`
	PromotionPrice             float32 `protobuf:"fixed32,17,opt,name=PromotionPrice,proto3" json:"PromotionPrice,omitempty"`
	GiftGrowth                 int64   `protobuf:"varint,18,opt,name=GiftGrowth,proto3" json:"GiftGrowth,omitempty"`
	GiftPoint                  int64   `protobuf:"varint,19,opt,name=GiftPoint,proto3" json:"GiftPoint,omitempty"`
	UsePointLimit              int64   `protobuf:"varint,20,opt,name=UsePointLimit,proto3" json:"UsePointLimit,omitempty"`
	SubTitle                   string  `protobuf:"bytes,21,opt,name=SubTitle,proto3" json:"SubTitle,omitempty"`
	Description                string  `protobuf:"bytes,22,opt,name=Description,proto3" json:"Description,omitempty"`
	OriginalPrice              float32 `protobuf:"fixed32,23,opt,name=OriginalPrice,proto3" json:"OriginalPrice,omitempty"`
	Stock                      int64   `protobuf:"varint,24,opt,name=Stock,proto3" json:"Stock,omitempty"`
	LowStock                   int64   `protobuf:"varint,25,opt,name=LowStock,proto3" json:"LowStock,omitempty"`
	Unit                       string  `protobuf:"bytes,26,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Weight                     float32 `protobuf:"fixed32,27,opt,name=Weight,proto3" json:"Weight,omitempty"`
	PreviewStatus              int64   `protobuf:"varint,28,opt,name=PreviewStatus,proto3" json:"PreviewStatus,omitempty"`
	ServiceIds                 string  `protobuf:"bytes,29,opt,name=ServiceIds,proto3" json:"ServiceIds,omitempty"`
	Keywords                   string  `protobuf:"bytes,30,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	Note                       string  `protobuf:"bytes,31,opt,name=Note,proto3" json:"Note,omitempty"`
	AlbumPics                  string  `protobuf:"bytes,32,opt,name=AlbumPics,proto3" json:"AlbumPics,omitempty"`
	DetailTitle                string  `protobuf:"bytes,33,opt,name=DetailTitle,proto3" json:"DetailTitle,omitempty"`
	DetailDesc                 string  `protobuf:"bytes,34,opt,name=DetailDesc,proto3" json:"DetailDesc,omitempty"`
	DetailHtml                 string  `protobuf:"bytes,35,opt,name=DetailHtml,proto3" json:"DetailHtml,omitempty"`
	DetailMobileHtml           string  `protobuf:"bytes,36,opt,name=DetailMobileHtml,proto3" json:"DetailMobileHtml,omitempty"`
	PromotionStartTime         string  `protobuf:"bytes,37,opt,name=PromotionStartTime,proto3" json:"PromotionStartTime,omitempty"`
	PromotionEndTime           string  `protobuf:"bytes,38,opt,name=PromotionEndTime,proto3" json:"PromotionEndTime,omitempty"`
	PromotionPerLimit          int64   `protobuf:"varint,39,opt,name=PromotionPerLimit,proto3" json:"PromotionPerLimit,omitempty"`
	PromotionType              int64   `protobuf:"varint,40,opt,name=PromotionType,proto3" json:"PromotionType,omitempty"`
	BrandName                  string  `protobuf:"bytes,41,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	ProductCategoryName        string  `protobuf:"bytes,42,opt,name=ProductCategoryName,proto3" json:"ProductCategoryName,omitempty"`
}

func (x *ProductAddReq) Reset() {
	*x = ProductAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAddReq) ProtoMessage() {}

func (x *ProductAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAddReq.ProtoReflect.Descriptor instead.
func (*ProductAddReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductAddReq) GetBrandId() int64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *ProductAddReq) GetProductCategoryId() int64 {
	if x != nil {
		return x.ProductCategoryId
	}
	return 0
}

func (x *ProductAddReq) GetFreightTemplateId() int64 {
	if x != nil {
		return x.FreightTemplateId
	}
	return 0
}

func (x *ProductAddReq) GetProductAttributeCategoryId() int64 {
	if x != nil {
		return x.ProductAttributeCategoryId
	}
	return 0
}

func (x *ProductAddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductAddReq) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *ProductAddReq) GetProductSn() string {
	if x != nil {
		return x.ProductSn
	}
	return ""
}

func (x *ProductAddReq) GetDeleteStatus() string {
	if x != nil {
		return x.DeleteStatus
	}
	return ""
}

func (x *ProductAddReq) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

func (x *ProductAddReq) GetRecommendStatus() string {
	if x != nil {
		return x.RecommendStatus
	}
	return ""
}

func (x *ProductAddReq) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ProductAddReq) GetSale() string {
	if x != nil {
		return x.Sale
	}
	return ""
}

func (x *ProductAddReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductAddReq) GetPromotionPrice() float32 {
	if x != nil {
		return x.PromotionPrice
	}
	return 0
}

func (x *ProductAddReq) GetGiftGrowth() int64 {
	if x != nil {
		return x.GiftGrowth
	}
	return 0
}

func (x *ProductAddReq) GetGiftPoint() int64 {
	if x != nil {
		return x.GiftPoint
	}
	return 0
}

func (x *ProductAddReq) GetUsePointLimit() int64 {
	if x != nil {
		return x.UsePointLimit
	}
	return 0
}

func (x *ProductAddReq) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *ProductAddReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductAddReq) GetOriginalPrice() float32 {
	if x != nil {
		return x.OriginalPrice
	}
	return 0
}

func (x *ProductAddReq) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductAddReq) GetLowStock() int64 {
	if x != nil {
		return x.LowStock
	}
	return 0
}

func (x *ProductAddReq) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ProductAddReq) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ProductAddReq) GetPreviewStatus() int64 {
	if x != nil {
		return x.PreviewStatus
	}
	return 0
}

func (x *ProductAddReq) GetServiceIds() string {
	if x != nil {
		return x.ServiceIds
	}
	return ""
}

func (x *ProductAddReq) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ProductAddReq) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *ProductAddReq) GetAlbumPics() string {
	if x != nil {
		return x.AlbumPics
	}
	return ""
}

func (x *ProductAddReq) GetDetailTitle() string {
	if x != nil {
		return x.DetailTitle
	}
	return ""
}

func (x *ProductAddReq) GetDetailDesc() string {
	if x != nil {
		return x.DetailDesc
	}
	return ""
}

func (x *ProductAddReq) GetDetailHtml() string {
	if x != nil {
		return x.DetailHtml
	}
	return ""
}

func (x *ProductAddReq) GetDetailMobileHtml() string {
	if x != nil {
		return x.DetailMobileHtml
	}
	return ""
}

func (x *ProductAddReq) GetPromotionStartTime() string {
	if x != nil {
		return x.PromotionStartTime
	}
	return ""
}

func (x *ProductAddReq) GetPromotionEndTime() string {
	if x != nil {
		return x.PromotionEndTime
	}
	return ""
}

func (x *ProductAddReq) GetPromotionPerLimit() int64 {
	if x != nil {
		return x.PromotionPerLimit
	}
	return 0
}

func (x *ProductAddReq) GetPromotionType() int64 {
	if x != nil {
		return x.PromotionType
	}
	return 0
}

func (x *ProductAddReq) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *ProductAddReq) GetProductCategoryName() string {
	if x != nil {
		return x.ProductCategoryName
	}
	return ""
}

type ProductAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductAddResp) Reset() {
	*x = ProductAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAddResp) ProtoMessage() {}

func (x *ProductAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAddResp.ProtoReflect.Descriptor instead.
func (*ProductAddResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductAddResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0xa3, 0x0a, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x46,
	0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x3e, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x50, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x65, 0x77,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x53, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x69, 0x66, 0x74, 0x47, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x47, 0x69, 0x66, 0x74,
	0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x69, 0x66, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x47, 0x69, 0x66, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x55, 0x73, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75,
	0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75,
	0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x77, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x6f, 0x77, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x69, 0x63, 0x73, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x69, 0x63, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44, 0x65, 0x73, 0x63,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x74, 0x6d, 0x6c,
	0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x74,
	0x6d, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x48, 0x74, 0x6d, 0x6c, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x2e,
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x27, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x24,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x32, 0x48, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x3d, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_proto_goTypes = []interface{}{
	(*ProductAddReq)(nil),  // 0: product.ProductAddReq
	(*ProductAddResp)(nil), // 1: product.ProductAddResp
}
var file_product_proto_depIdxs = []int32{
	0, // 0: product.product.ProductAdd:input_type -> product.ProductAddReq
	1, // 1: product.product.ProductAdd:output_type -> product.ProductAddResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAddReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAddResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
