package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/user/gapsi_orders_api/internal/domain"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
)

var validate = validator.New()

type AuthHandler struct {
	authService domain.UserService
}

func NewAuthHandler(authService domain.UserService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "Email" && fieldError.Tag() == "required" {
					common.RespondWithError(w, r, http.StatusNotFound, "el email esta vacio")
					return
				}
			}
		}
		// Fallback for other validation errors
		common.RespondWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	role := domain.RoleUser
	if req.Role == "ADMIN" {
		role = domain.RoleAdmin
	}

	user, err := h.authService.Register(r.Context(), req.Email, req.Password, role)
	if err != nil {
		if err.Error() == "el correo electrónico ya está registrado" {
			common.RespondWithError(w, r, http.StatusBadRequest, err.Error())
			return
		}
		common.RespondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		common.RespondWithError(w, r, http.StatusUnauthorized, "invalid credentials")
		return
	}

	setTokenCookies(w, accessToken, refreshToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token":  accessToken,
		"expires_in":    86400000, // 1 day in ms (example format matching requirements)
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	// Try body first, fallback to cookie
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RefreshToken == "" {
		if cookie, err := r.Cookie("refresh_token"); err == nil {
			req.RefreshToken = cookie.Value
		} else {
			common.RespondWithError(w, r, http.StatusBadRequest, "refresh token required")
			return
		}
	}

	accessToken, newRefreshToken, err := h.authService.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		common.RespondWithError(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	setTokenCookies(w, accessToken, newRefreshToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"access_token":  accessToken,
		"expires_in":    86400000,
		"refresh_token": newRefreshToken,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout exitoso",
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		common.RespondWithError(w, r, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := h.authService.GetMe(r.Context(), userID)
	if err != nil {
		common.RespondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})
}

func setTokenCookies(w http.ResponseWriter, access, refresh string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    access,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
