package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/user/gapsi_orders_api/internal/domain"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
)

type AuthHandler struct {
	authService domain.UserService
}

func NewAuthHandler(authService domain.UserService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithCustomError(w, r, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}

	if err := req.Validate(); err != nil {
		common.RespondWithCustomError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	role := domain.RoleUser
	if req.Role == "ADMIN" {
		role = domain.RoleAdmin
	}

	user, err := h.authService.Register(r.Context(), req.Email, req.Password, role)
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.RespondWithCustomError(w, r, http.StatusBadRequest, "cuerpo de la solicitud inválido")
		return
	}

	if err := req.Validate(); err != nil {
		common.RespondWithCustomError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	setTokenCookies(w, accessToken, refreshToken)

	common.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"expires_in":    86400000, // 1 day in ms
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	// Try body first
	json.NewDecoder(r.Body).Decode(&req)

	// Fallback to cookie if body is empty
	if req.RefreshToken == "" {
		if cookie, err := r.Cookie("refresh_token"); err == nil {
			req.RefreshToken = cookie.Value
		}
	}

	if err := req.Validate(); err != nil {
		common.RespondWithCustomError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, newRefreshToken, err := h.authService.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	setTokenCookies(w, accessToken, newRefreshToken)

	common.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
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

	common.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Logout exitoso",
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		common.RespondWithCustomError(w, r, http.StatusUnauthorized, "no autorizado")
		return
	}

	user, err := h.authService.GetMe(r.Context(), userID)
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	common.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
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
