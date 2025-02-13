// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestayOrder "looklook/app/order/cmd/api/internal/handler/homestayOrder"
	"looklook/app/order/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/homestayOrder/createHomestayOrder",
				Handler: homestayOrder.CreateHomestayOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayOrder/userHomestayOrderList",
				Handler: homestayOrder.UserHomestayOrderListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/homestayOrder/userHomestayOrderDetail",
				Handler: homestayOrder.UserHomestayOrderDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/order/v1"),
	)
}
