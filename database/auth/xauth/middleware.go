package xauth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"opentaxi/database/auth/tokens"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaim struct {
	USERID string `json:"id"`
	ROL    string `json:"rol"`
	ROLID  int    `json:"rolid"`
	jwt.StandardClaims
}

type AuthData struct {
	Clains *JwtCustomClaim
	// Usuario *model.UsuarioLogin
	TOKEN string `json:"token"`
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

func GenerateToken(ctx context.Context, userID string, rol string, rolid int) (string, error) {
	tokenduration := os.Getenv("TOKEN_DURATION_MIN")
	duration := 10
	id, err := strconv.Atoi(tokenduration)
	if err == nil {
		duration = id
	}
	return jwtGenerate(ctx, userID, rol, rolid, duration)
}

func jwtGenerate(ctx context.Context, userID string, rol string, rolid int, minutos int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		USERID: userID,
		ROL:    rol,
		ROLID:  rolid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(minutos)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}

type authString string

func AuthMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("dataname")
			dataname, er := decrypt(auth)
			if er != nil {
				dataname = "x"
			}
			token := tokens.GetToken(db, dataname)
			validate, err := JwtValidate(context.Background(), token.Token)
			if err != nil && validate == nil {
				next.ServeHTTP(w, r)
				return
			}
			customClaim, _ := validate.Claims.(*JwtCustomClaim)

			data := AuthData{}
			data.Clains = customClaim
			data.TOKEN = token.Token
			ctx := context.WithValue(r.Context(), authString("auth"), &data)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}

func CtxValue(ctx context.Context, db *sql.DB, metodo string) (*AuthData, error) {
	clains, _ := ctx.Value(authString("auth")).(*AuthData)
	if clains == nil {
		return nil, errors.New("debes iniciar session y usar un rol")
	}
	validate, err := JwtValidate(context.Background(), clains.TOKEN)
	if err != nil || !validate.Valid {
		txt := err.Error()
		if strings.HasPrefix(txt, "token is expired by") {
			txt = strings.Replace(txt, "token is expired by", "Su sessión expiró hace", 1)
			return nil, errors.New(txt)
		} else {
			return nil, errors.New(txt)
		}
	}

	err = CheckPermiso(db, metodo, clains.Clains.ROLID, clains.Clains.ROL)
	if err != nil {
		return nil, err
	}
	return clains, nil
}
