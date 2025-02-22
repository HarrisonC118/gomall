// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	"errors"
	product "github.com/PiaoAdmin/gomall/rpc_gen/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListProducts": kitex.NewMethodInfo(
		listProductsHandler,
		newListProductsArgs,
		newListProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetProduct": kitex.NewMethodInfo(
		getProductHandler,
		newGetProductArgs,
		newGetProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchProducts": kitex.NewMethodInfo(
		searchProductsHandler,
		newSearchProductsArgs,
		newSearchProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddProduct": kitex.NewMethodInfo(
		addProductHandler,
		newAddProductArgs,
		newAddProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AssociateProductWithCategory": kitex.NewMethodInfo(
		associateProductWithCategoryHandler,
		newAssociateProductWithCategoryArgs,
		newAssociateProductWithCategoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddCategory": kitex.NewMethodInfo(
		addCategoryHandler,
		newAddCategoryArgs,
		newAddCategoryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productCatalogServiceServiceInfo                = NewServiceInfo()
	productCatalogServiceServiceInfoForClient       = NewServiceInfoForClient()
	productCatalogServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ProductCatalogService"
	handlerType := (*product.ProductCatalogService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "product",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func listProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.ListProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).ListProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListProductsArgs:
		success, err := handler.(product.ProductCatalogService).ListProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListProductsArgs() interface{} {
	return &ListProductsArgs{}
}

func newListProductsResult() interface{} {
	return &ListProductsResult{}
}

type ListProductsArgs struct {
	Req *product.ListProductsReq
}

func (p *ListProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.ListProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.ListProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListProductsArgs_Req_DEFAULT *product.ListProductsReq

func (p *ListProductsArgs) GetReq() *product.ListProductsReq {
	if !p.IsSetReq() {
		return ListProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListProductsResult struct {
	Success *product.ListProductsResp
}

var ListProductsResult_Success_DEFAULT *product.ListProductsResp

func (p *ListProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.ListProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListProductsResult) Unmarshal(in []byte) error {
	msg := new(product.ListProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListProductsResult) GetSuccess() *product.ListProductsResp {
	if !p.IsSetSuccess() {
		return ListProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.ListProductsResp)
}

func (p *ListProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListProductsResult) GetResult() interface{} {
	return p.Success
}

func getProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.GetProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).GetProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetProductArgs:
		success, err := handler.(product.ProductCatalogService).GetProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetProductArgs() interface{} {
	return &GetProductArgs{}
}

func newGetProductResult() interface{} {
	return &GetProductResult{}
}

type GetProductArgs struct {
	Req *product.GetProductReq
}

func (p *GetProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.GetProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetProductArgs) Unmarshal(in []byte) error {
	msg := new(product.GetProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetProductArgs_Req_DEFAULT *product.GetProductReq

func (p *GetProductArgs) GetReq() *product.GetProductReq {
	if !p.IsSetReq() {
		return GetProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetProductResult struct {
	Success *product.GetProductResp
}

var GetProductResult_Success_DEFAULT *product.GetProductResp

func (p *GetProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.GetProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetProductResult) Unmarshal(in []byte) error {
	msg := new(product.GetProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProductResult) GetSuccess() *product.GetProductResp {
	if !p.IsSetSuccess() {
		return GetProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.GetProductResp)
}

func (p *GetProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProductResult) GetResult() interface{} {
	return p.Success
}

func searchProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.SearchProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).SearchProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchProductsArgs:
		success, err := handler.(product.ProductCatalogService).SearchProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchProductsArgs() interface{} {
	return &SearchProductsArgs{}
}

func newSearchProductsResult() interface{} {
	return &SearchProductsResult{}
}

type SearchProductsArgs struct {
	Req *product.SearchProductsReq
}

func (p *SearchProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.SearchProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchProductsArgs_Req_DEFAULT *product.SearchProductsReq

func (p *SearchProductsArgs) GetReq() *product.SearchProductsReq {
	if !p.IsSetReq() {
		return SearchProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchProductsResult struct {
	Success *product.SearchProductsResp
}

var SearchProductsResult_Success_DEFAULT *product.SearchProductsResp

func (p *SearchProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.SearchProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchProductsResult) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchProductsResult) GetSuccess() *product.SearchProductsResp {
	if !p.IsSetSuccess() {
		return SearchProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.SearchProductsResp)
}

func (p *SearchProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchProductsResult) GetResult() interface{} {
	return p.Success
}

func addProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.AddProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).AddProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddProductArgs:
		success, err := handler.(product.ProductCatalogService).AddProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddProductArgs() interface{} {
	return &AddProductArgs{}
}

func newAddProductResult() interface{} {
	return &AddProductResult{}
}

type AddProductArgs struct {
	Req *product.AddProductReq
}

func (p *AddProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.AddProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddProductArgs) Unmarshal(in []byte) error {
	msg := new(product.AddProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddProductArgs_Req_DEFAULT *product.AddProductReq

func (p *AddProductArgs) GetReq() *product.AddProductReq {
	if !p.IsSetReq() {
		return AddProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddProductResult struct {
	Success *product.AddProductResp
}

var AddProductResult_Success_DEFAULT *product.AddProductResp

func (p *AddProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.AddProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddProductResult) Unmarshal(in []byte) error {
	msg := new(product.AddProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddProductResult) GetSuccess() *product.AddProductResp {
	if !p.IsSetSuccess() {
		return AddProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.AddProductResp)
}

func (p *AddProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddProductResult) GetResult() interface{} {
	return p.Success
}

func associateProductWithCategoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.AssociateProductWithCategoryReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).AssociateProductWithCategory(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AssociateProductWithCategoryArgs:
		success, err := handler.(product.ProductCatalogService).AssociateProductWithCategory(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AssociateProductWithCategoryResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAssociateProductWithCategoryArgs() interface{} {
	return &AssociateProductWithCategoryArgs{}
}

func newAssociateProductWithCategoryResult() interface{} {
	return &AssociateProductWithCategoryResult{}
}

type AssociateProductWithCategoryArgs struct {
	Req *product.AssociateProductWithCategoryReq
}

func (p *AssociateProductWithCategoryArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.AssociateProductWithCategoryReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AssociateProductWithCategoryArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AssociateProductWithCategoryArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AssociateProductWithCategoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AssociateProductWithCategoryArgs) Unmarshal(in []byte) error {
	msg := new(product.AssociateProductWithCategoryReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AssociateProductWithCategoryArgs_Req_DEFAULT *product.AssociateProductWithCategoryReq

func (p *AssociateProductWithCategoryArgs) GetReq() *product.AssociateProductWithCategoryReq {
	if !p.IsSetReq() {
		return AssociateProductWithCategoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AssociateProductWithCategoryArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AssociateProductWithCategoryArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AssociateProductWithCategoryResult struct {
	Success *product.AssociateProductWithCategoryResp
}

var AssociateProductWithCategoryResult_Success_DEFAULT *product.AssociateProductWithCategoryResp

func (p *AssociateProductWithCategoryResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.AssociateProductWithCategoryResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AssociateProductWithCategoryResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AssociateProductWithCategoryResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AssociateProductWithCategoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AssociateProductWithCategoryResult) Unmarshal(in []byte) error {
	msg := new(product.AssociateProductWithCategoryResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AssociateProductWithCategoryResult) GetSuccess() *product.AssociateProductWithCategoryResp {
	if !p.IsSetSuccess() {
		return AssociateProductWithCategoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AssociateProductWithCategoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.AssociateProductWithCategoryResp)
}

func (p *AssociateProductWithCategoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AssociateProductWithCategoryResult) GetResult() interface{} {
	return p.Success
}

func addCategoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.AddCategoryReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).AddCategory(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddCategoryArgs:
		success, err := handler.(product.ProductCatalogService).AddCategory(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddCategoryResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddCategoryArgs() interface{} {
	return &AddCategoryArgs{}
}

func newAddCategoryResult() interface{} {
	return &AddCategoryResult{}
}

type AddCategoryArgs struct {
	Req *product.AddCategoryReq
}

func (p *AddCategoryArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.AddCategoryReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddCategoryArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddCategoryArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddCategoryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddCategoryArgs) Unmarshal(in []byte) error {
	msg := new(product.AddCategoryReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddCategoryArgs_Req_DEFAULT *product.AddCategoryReq

func (p *AddCategoryArgs) GetReq() *product.AddCategoryReq {
	if !p.IsSetReq() {
		return AddCategoryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddCategoryArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddCategoryArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddCategoryResult struct {
	Success *product.AddCategoryResp
}

var AddCategoryResult_Success_DEFAULT *product.AddCategoryResp

func (p *AddCategoryResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.AddCategoryResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddCategoryResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddCategoryResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddCategoryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddCategoryResult) Unmarshal(in []byte) error {
	msg := new(product.AddCategoryResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddCategoryResult) GetSuccess() *product.AddCategoryResp {
	if !p.IsSetSuccess() {
		return AddCategoryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddCategoryResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.AddCategoryResp)
}

func (p *AddCategoryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddCategoryResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListProducts(ctx context.Context, Req *product.ListProductsReq) (r *product.ListProductsResp, err error) {
	var _args ListProductsArgs
	_args.Req = Req
	var _result ListProductsResult
	if err = p.c.Call(ctx, "ListProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProduct(ctx context.Context, Req *product.GetProductReq) (r *product.GetProductResp, err error) {
	var _args GetProductArgs
	_args.Req = Req
	var _result GetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchProducts(ctx context.Context, Req *product.SearchProductsReq) (r *product.SearchProductsResp, err error) {
	var _args SearchProductsArgs
	_args.Req = Req
	var _result SearchProductsResult
	if err = p.c.Call(ctx, "SearchProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddProduct(ctx context.Context, Req *product.AddProductReq) (r *product.AddProductResp, err error) {
	var _args AddProductArgs
	_args.Req = Req
	var _result AddProductResult
	if err = p.c.Call(ctx, "AddProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AssociateProductWithCategory(ctx context.Context, Req *product.AssociateProductWithCategoryReq) (r *product.AssociateProductWithCategoryResp, err error) {
	var _args AssociateProductWithCategoryArgs
	_args.Req = Req
	var _result AssociateProductWithCategoryResult
	if err = p.c.Call(ctx, "AssociateProductWithCategory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddCategory(ctx context.Context, Req *product.AddCategoryReq) (r *product.AddCategoryResp, err error) {
	var _args AddCategoryArgs
	_args.Req = Req
	var _result AddCategoryResult
	if err = p.c.Call(ctx, "AddCategory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
