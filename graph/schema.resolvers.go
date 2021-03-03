package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/graph/generated"
	"github.com/engajerest/auth/graph/model"
	"github.com/engajerest/auth/utils/accesstoken"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateUser(ctx context.Context, create model.NewUser) (*model.UserCreatedData, error) {
	var user users.User
	user.FirstName = create.Firstname
	user.LastName = create.Lastname
	user.Password = create.Password
	user.Email = create.Email
	user.Mobile = create.Mobile

	user.Roleid = *create.Roleid

	// user.Status = "Active"

	userid := user.Create()
	user.ID = int(userid)
	user.InsertUserintoProfile()
	user.GetByUserId(userid)
	// token, err := jwt.GenerateToken(user.Password)
	// if err != nil {
	// 	return "", err
	// }
	return &model.UserCreatedData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		UserInfo: &model.UserData{
			UserID:      user.ID,
			Firstname:   user.FirstName,
			Lastname:    user.LastName,
			Email:       user.Email,
			Mobile:      user.Mobile,
			Token:       "",
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.LoginData, error) {
	var user users.User
	user.FirstName = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return &model.LoginData{Status: false, Code: http.StatusBadRequest, Message: "Incorrect Username or password", UserInfo: nil}, nil
	}

	token, err := accesstoken.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	user.LoginResponse(int64(user.ID))

	user.InsertToken(token)
	if user.Referenceid!=0{
		
		if input.Tenanttoken!=""{
			status:=users.Updatetenant(input.Tenanttoken,user.Referenceid)
			print("tentokenupdate=",status)
		}
	}
	

	return &model.LoginData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		UserInfo: &model.UserData{
			UserID:      user.ID,
			Tenantid:    &user.Referenceid,
			Locationid:  &user.LocationId,
			Moduleid:    &user.LocationId,
			Packageid:   &user.Packageid,
			Modulename:  &user.Modulename,
			Tenantname:  &user.Tenantname,
			Firstname:   user.FirstName,
			Lastname:    user.LastName,
			Email:       user.Email,
			Mobile:      user.Mobile,
			Token:       token,
			Opentime:    &user.Opentime,
			Closetime:   &user.Closetime,
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}}, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, input model.Reset) (string, error) {
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return "nil", &gqlerror.Error{

			Path:    graphql.GetPath(ctx),
			Message: "no userfound",

			Extensions: map[string]interface{}{
				"code": "400",
			},
		}
	}
	var user users.User
	// user.Username = input.Username
	user.Password = input.Password
	user.ID = id.ID
	print(user.ID)
	reset := user.ResetPassword()
	if reset != false {
		return "success", nil
	}
	return "failure", nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	_, err := accesstoken.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := accesstoken.GenerateToken(0)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.GetUser, error) {
	var userResult []*model.GetUser
	var userGetAll []users.User

	userGetAll = users.GetAllUsers()
	for _, user := range userGetAll {
		userResult = append(userResult, &model.GetUser{UserID: user.ID, Firstname: user.FirstName, Lastname: user.LastName, Mobile: user.Mobile, Email: user.Email, Created: user.CreatedDate, Status: user.Status})
	}
	return userResult, nil
}

func (r *queryResolver) Getuser(ctx context.Context) (*model.LoginData, error) {
	id, usererr := controller.ForContext(ctx)
	if usererr != nil {
		return nil, &gqlerror.Error{

			Path:    graphql.GetPath(ctx),
			Message: "no userfound",

			Extensions: map[string]interface{}{
				"code": "400",
			},
		}
	}
	// print(id.ID)

	return &model.LoginData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		UserInfo: &model.UserData{
			UserID:      id.ID,
			Firstname:   id.FirstName,
			Lastname:    id.LastName,
			Email:       id.Email,
			Mobile:      id.Mobile,
			CreatedDate: id.CreatedDate,
			Status:      id.Status,
		}}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
