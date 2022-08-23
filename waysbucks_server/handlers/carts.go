package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	cartdto "waysbucks/dto/carts"
	dto "waysbucks/dto/result"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var path_file_cart = "http://localhost:5000/uploads/"

type handlersCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlersCart {
	return &handlersCart{CartRepository}
}

func (h *handlersCart) FindCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	carts, err := h.CartRepository.FindCarts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: carts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlersCart) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlersCart) CreateCart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// get data user token
// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userId := int(userInfo["id"].(float64))

// 	// get image filename
// 	dataContex := r.Context().Value("dataFile")
// 	filename := dataContex.(string)

// 	var categoriesId []int
// 	for _, r := range r.FormValue("categoryId") {
// 		if int(r-'0') >= 0 {
// 			categoriesId = append(categoriesId, int(r-'0'))
// 		}
//     }

// 	price, _ := strconv.Atoi(r.FormValue("price"))
// 	qty, _ := strconv.Atoi(r.FormValue("qty"))

// 	request := cartdto.CreateCart{
// 		Name: 		r.FormValue("name"),
// 		Desc:		r.FormValue("desc"),
// 		Price:  	price,
// 		Qty:		qty,
// 		UserID:     userId,
// 		CategoryID:	categoriesId,
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Get all category data by id []
// 	category, _ := h.CartRepository.FindCategoriesById(categoriesId)

// 	product := models.Product{
// 		Name:   request.Name,
// 		Desc:   request.Desc,
// 		Price:  request.Price,
// 		Image:  filename,
// 		Qty:    request.Qty,
// 		UserID: userId,
// 		Category:	category,
// 	}

// 	product, err = h.ProductRepository.CreateProduct(product)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	product, _ = h.CartRepository.GetCart(product.ID)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "success", Data: product}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlersCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	Transaction_ID := int(userInfo["time"].(float64))

	request := new(cartdto.CreateCart)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println(request.Product_ID)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var toppingsId []int
	// for _, r := range r.FormValue("topping_Id") {
	// 	if int(r-'0') >= 0 {
	// 		toppingsId = append(toppingsId, int(r-'0'))
	// 	}
	// }

	requestForm := models.Cart{
		Product_ID:    request.Product_ID,
		TransactionID: Transaction_ID,
		SubTotal:      request.SubTotal,
		ToppingID:     request.Topping_ID,
	}

	validatee := validator.New()
	errr := validatee.Struct(requestForm)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ := h.CartRepository.FindToppingsID(request.Topping_ID)

	cart := models.Cart{
		Product_ID:    request.Product_ID,
		SubTotal:      request.SubTotal,
		Topping:       topping,
		TransactionID: Transaction_ID,
	}
	fmt.Println(cart)

	data, err := h.CartRepository.CreateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlersCart) UpdateCart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(cartdto.UpdateCart)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	cart, err := h.CartRepository.GetCart(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	// len > 0
// 	if request.ProductID != 0 {
// 		cart.Product_ID = request.ProductID
// 	}

// 	data, err := h.CartRepository.UpdateCart(cart)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlersCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	cart, err := h.CartRepository.GetCart(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	data, err := h.CartRepository.DeleteCart(cart)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlersCart) FindCartsByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	idTrans := int(userInfo["time"].(float64))

// 	cart, err := h.CartRepository.FindCartsTransaction(idTrans)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	for i, p := range cart {
// 		cart[i].Product.Image = path_file_cart + p.Product.Image
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: cart}
// 	json.NewEncoder(w).Encode(response)
// }

// func convertResponseCart(u models.Cart) models.Cart {
// 	return models.Cart{
// 		ID:       u.ID,
// 		SubTotal: u.SubTotal,
// 		Product:  u.Product,
// 		Topping:  u.Topping,
// 		// Transaction: u.Transaction,
// 	}
// }

func (h *handlersCart) FindCartsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	Transaction_ID := int(userInfo["time"].(float64))

	cart, err := h.CartRepository.FindCartsTransaction(Transaction_ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range cart {
		cart[i].Product.Image = path_file_cart + p.Product.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func convertResponseCart(u models.Cart) models.Cart {
	return models.Cart{
		ID:       u.ID,
		SubTotal: u.SubTotal,
		Product:  u.Product,
		Topping:  u.Topping,
		// Transaction: u.Transaction,
	}
}
