package servers

import (
	"awesomeProject"
	utils "awesomeProject/internal/util"
	"awesomeProject/network"
	pb "awesomeProject/protoFiles/files"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
)

type Server struct {
	pb.UnimplementedApiServer
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterApiServer(grpcServer, &Server{})
	log.Println("server is running on port 8080...")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func CreatePayoutCard(url string, sum int, card string, token string) (network.PayoutResponse, string) {
	CreateClientOrderID := uuid.NewString()
	PayerID := fmt.Sprintf("Invoice_%d", utils.RandRanged(10000, 99999))

	createPayout := network.CreateRequestPayout{
		ClientOrderID: CreateClientOrderID,
		Sum:           sum,
		Ttl:           3600,
		Message:       "PayoutP2P",
		Type:          "card",
		WalletID:      1,
		WebhookUrl:    "https://google.com/example_webhook",
		CardNumber:    card,
		PayerInfo: network.PayerInfo{
			PayerID: PayerID,
		},
	}

	invoice, err := json.Marshal(createPayout)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invoice))

	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var requestClass network.PayoutResponse
	err = json.Unmarshal(body, &requestClass)

	if err != nil {
		fmt.Println("Ошибка, обратитесь в поддержку к нам +79932703447:", err)
	}
	return requestClass, ""
}

func CreateRequestInvoiceP2PH2H(url string, sum int, token string) (network.P2PInvoiceResponse, int) {

	CreateClientOrderID := uuid.NewString()
	PayerID := fmt.Sprintf("Invoice_%d", utils.RandRanged(10000, 99999))

	createInvoice := network.CreateInvoiceP2P{
		ClientOrderID: CreateClientOrderID,
		ClientIP:      utils.CreateIP(),
		PayerID:       PayerID,
		Sum:           sum,
		WalletID:      1,
		Bank:          "sberbank",
		Message:       "invoiceP2P",
		WebhookUrl:    "https://google.com/example_webhook",
		ExpireAt:      3600,
	}

	invoice, err := json.Marshal(createInvoice)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invoice))

	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	var requestClass network.P2PInvoiceResponse

	if err != nil {
		log.Fatal(err)
		return requestClass, resp.StatusCode
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		json.Unmarshal(
			body,
			&requestClass)
		return requestClass, resp.StatusCode
	}

	err = json.Unmarshal(body, &requestClass)

	if err != nil {
		return requestClass, 504
	}

	return requestClass, 200
}
func CreateRequestInvoiceFPSH2H(url string, sum int, token string) (network.FPSInvoiceResponse, int) {

	CreateClientOrderID := uuid.NewString()
	PayerID := fmt.Sprintf("Invoice_%d", utils.RandRanged(10000, 99999))

	createInvoice := network.CreateInvoiceFPS{
		ClientOrderID: CreateClientOrderID,
		ClientIP:      utils.CreateIP(),
		PayerID:       PayerID,
		Amount:        sum,
		CurrencyID:    1,
		Comment:       "InvoiceFPS",
		WebhookURL:    "https://google.com/example_webhook",
		ExpireAt:      3600,
	}

	invoice, err := json.Marshal(createInvoice)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invoice))

	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	var requestClass network.FPSInvoiceResponse

	if err != nil {
		log.Fatal(err)
		return requestClass, resp.StatusCode
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		json.Unmarshal(
			body,
			&requestClass)
		return requestClass, resp.StatusCode
	}

	err = json.Unmarshal(body, &requestClass)

	if err != nil {
		return requestClass, 504
	}

	return requestClass, 200
}
func logRequest(handler http.Handler, config *awesomeProject.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/" {
			{ //ввод
				var requestClass network.RequestPostmanInvoice
				body, _ := io.ReadAll(r.Body)
				json.Unmarshal(body, &requestClass)

				//if err := requestClass.Validate(); err != nil {
				//	createErrorRequest := network.FinalResponseError{Error: "ошибка"}
				//	request, _ := json.Marshal(createErrorRequest)
				//	w.WriteHeader(400)
				//	_, _ = w.Write(request)
				//	handler.ServeHTTP(w, r)
				//}

				switch requestClass.Type {
				case "P2P":

					req, err := CreateRequestInvoiceP2PH2H(network.PaymentP2PURL, requestClass.Sum, config.Token)
					if err != 200 {
						createErrorRequest := network.FinalResponseError{Error: "Что-то пошло не так, попробуйте позже или обратитесь в поддержку"}
						request, _ := json.Marshal(createErrorRequest)
						w.WriteHeader(err)
						_, _ = w.Write(request)
						handler.ServeHTTP(w, r)
						break
					}

					createRequest := network.FinalResponseInvoice{
						Uuid:           req.Uuid,
						Card:           req.Card,
						Sum:            req.Sum,
						CardHolderName: req.CardHolderName,
						BankName:       req.BankName,
					}

					request, _ := json.Marshal(createRequest)
					_, _ = w.Write(request)
					handler.ServeHTTP(w, r)
					break

				case "FPS":
					req, err := CreateRequestInvoiceFPSH2H(network.PaymentFPSURL, requestClass.Sum, config.Token)
					if err != 200 {
						createErrorRequest := network.FinalResponseError{Error: "Что-то пошло не так, попробуйте позже или обратитесь в поддержку"}
						request, _ := json.Marshal(createErrorRequest)
						w.WriteHeader(err)
						_, _ = w.Write(request)
						handler.ServeHTTP(w, r)
						break
					}

					createRequest := network.FinalResponseInvoice{
						Uuid:           req.ExternalID,
						Card:           req.PhoneNumber,
						Sum:            req.Amount,
						CardHolderName: req.CardHolderName,
						BankName:       req.BankName,
					}

					request, _ := json.Marshal(createRequest)
					_, _ = w.Write(request)
					handler.ServeHTTP(w, r)
					break

				default:
					createErrorRequest := network.FinalResponseError{Error: "Нет такого метода"}
					request, _ := json.Marshal(createErrorRequest)
					w.WriteHeader(400)
					_, _ = w.Write(request)
					handler.ServeHTTP(w, r)

				}
			}
		} else if r.RequestURI == "/payout" { //вывод
			var requestClass network.RequestPostmanPayout
			body, _ := io.ReadAll(r.Body)
			json.Unmarshal(body, &requestClass)

			//if err := requestClass.Validate(); err != nil {
			//	createErrorRequest := network.FinalResponseError{Error: "ошибка"}
			//	request, _ := json.Marshal(createErrorRequest)
			//	_, _ = w.Write(request)
			//	handler.ServeHTTP(w, r)
			//}

			req, err := CreatePayoutCard(network.PayoutURL, requestClass.Sum, requestClass.Card, config.Token)
			if err != "" {
				createErrorRequest := network.FinalResponseError{Error: err}
				request, _ := json.Marshal(createErrorRequest)
				_, _ = w.Write(request)
				handler.ServeHTTP(w, r)
			}

			createRequest := network.FinalResponsePayout{
				Id:      req.Id,
				Uuid:    req.Uuid,
				Success: req.Success,
			}

			request, _ := json.Marshal(createRequest)
			_, _ = w.Write(request)
			handler.ServeHTTP(w, r)
		}
	})
}
