package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/datacontext"
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
	user.Roleid = create.Roleid
	user.Configid = create.Configid
	user.Countrycode = create.Countrycode
	user.CurrencyCode = create.Currencycode
	user.Currencysymbol = create.Currencysymbol
	user.Dialcode = create.Dialcode
	var userid int64
	var err error
	// user.Status = "Active"
	id := user.CheckUser()
	if id.ID != 0 {
		if id.Email == create.Email {
			return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Email Already Exists",
				UserInfo: &model.UserData{}}, nil
		} else if id.Mobile == create.Mobile {
			return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Contactno Already Exists",
				UserInfo: &model.UserData{}}, nil
		}
	} else {
		if create.Password != "" {

			userid, err = user.Create()
			if err != nil {
				if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'authname'", user.Email) {
					print("true")
					return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Email Already Exists",
						UserInfo: &model.UserData{}}, nil
				} else if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'contactno'", user.Mobile) {
					return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Contactno Already Exists",
						UserInfo: &model.UserData{}}, nil
				} else {
					return nil, err
				}

			}
		} else {
			userid, err = user.Createwithoutpassword()
			if err != nil {
				if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'authname'", user.Email) {
					print("true")
					return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Email Already Exists",
						UserInfo: &model.UserData{}}, nil
				} else if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'contactno'", user.Mobile) {
					return &model.UserCreatedData{Status: false, Code: http.StatusConflict, Message: "Contactno Already Exists",
						UserInfo: &model.UserData{}}, nil
				} else {
					return nil, err
				}

			}
		}
	}

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
			UserID:         user.ID,
			Firstname:      user.FirstName,
			Lastname:       user.LastName,
			Email:          user.Email,
			Mobile:         user.Mobile,
			Token:          "",
			CreatedDate:    user.CreatedDate,
			Status:         user.Status,
			Currencysymbol: user.Currencysymbol,
			Currencycode:   user.CurrencyCode,
			Countrycode:    user.Countrycode,
			Dialcode:       user.Dialcode,
		}}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.LoginData, error) {
	var user users.User
	var tenantlist []*model.Tenantdata
	var t []users.Tenant
	var loc []users.Location
	var loclist []*model.Location
	user.FirstName = input.Username
	user.Password = input.Password
	user.Devicetype = input.Devicetype
	if input.Password != "" {
		correct := user.Authenticate()
		if !correct {
			// 1
			return &model.LoginData{Status: false, Code: http.StatusBadRequest, Message: "Incorrect Username or password", UserInfo: nil,
				Tenantinfo: tenantlist}, nil
		}
	} else {
		stat := user.Checkauthname()
		if stat == false {
			return &model.LoginData{Status: false, Code: http.StatusBadRequest, Message: "Incorrect Username", UserInfo: nil, Tenantinfo: tenantlist}, nil
		}
	}
	var token string
	var err error
	if user.ID != 0 && user.Configid != 0 {
		token, err = accesstoken.GenerateToken(user.ID, user.Configid)
		if err != nil {
			return nil, err
		}
	} else {
		return &model.LoginData{Status: false, Code: http.StatusBadGateway, Message: "update user now"}, nil
	}

	user.LoginResponse(int64(user.ID))
	print("refid", user.Referenceid)

	user.InsertToken(token)
	if user.Referenceid != 0 {
		print("not 0")
		if input.Tenanttoken != "" || input.Devicetype != "" {
			status := users.Updatetenant(input.Tenanttoken, input.Devicetype, user.Referenceid)
			print("tentokenupdate=", status)
		}

		t = users.Tenantresponse(user.Referenceid)
		if len(t) != 0 {
			for _, k := range t {
				tenantlist = append(tenantlist, &model.Tenantdata{Subscriptionid: k.Subscriptionid,
					Packageid: k.Packageid, Packagename: k.Packagename, Moduleid: k.Moduleid, Validitydate: k.Validiydate,
					Validity: k.Validity, Categoryid: k.Categoryid, Subcategoryid: k.Subcategoryid,
					Taxpercent: k.Taxpercent, Status: k.Status, Featureid: k.Featureid, Taxamount: k.Taxamount, Totalamount: k.Totalamount, Subscriptionaccid: k.Subscriptionaccid, Subscriptionmethodid: k.Subscriptionmethodid, Paymentstatus: k.Paymentstatus, Modulename: k.Modulename, Iconurl: k.Iconurl, Logourl: k.Logourl})
			}
		}
		loc = users.Locationresponse(user.Referenceid)
		for _, n := range loc {
			loclist = append(loclist, &model.Location{Locationid: n.LocationId, Tenantid: n.Tenantid, Locationname: n.Locationname,
				Email: n.Email, Contactno: n.Contactno, Address: n.Address, City: n.City, State: n.State, Postcode: n.Postcode,
				Suburb: n.Suburb, Latitude: n.Latitude, Longitude: n.Longitude, Opentime: n.Opentime, Closetime: n.Closetime})
		}

	} else {
		print("not 0")
	}

	return &model.LoginData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		UserInfo: &model.UserData1{
			UserID:             user.ID,
			Tenantid:           &user.Referenceid,
			Devicetype:         user.Devicetype,
			Firstname:          user.FirstName,
			Lastname:           user.LastName,
			Email:              user.Email,
			Mobile:             user.Mobile,
			Token:              token,
			Usercurrencysymbol: user.Usercurrencysymbol,
			Usercurrencycode:   user.UsercurrencyCode,
			Usercountrycode:    user.Usercountrycode,
			CreatedDate:        user.CreatedDate,
			Status:             user.Status,
			Roleid:             &user.Roleid,
			Configid:           &user.Configid,
			Tenantname:         &user.Tenantname,
			Tenantimageurl:     &user.Tenantimage,
			Profileimage:       user.Profileimage,
			Tenantaccid:        user.Tenantaccid,
			Countrycode:        user.Countrycode,
			Currencyid:         user.Currencyid,
			Currencycode:       user.CurrencyCode,
			Currencysymbol:     user.Currencysymbol,
			Dialcode:           user.Dialcode,
			Tenantstatus:       user.Tenantstatus,
		}, Tenantinfo: tenantlist, Locationinfo: loclist,
	}, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, input model.Reset) (string, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
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
	_, _, err := accesstoken.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := accesstoken.GenerateToken(0, 0)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Updateuser(ctx context.Context, input *model.Userupdateinput) (*model.Updateddata, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		return nil, &gqlerror.Error{

			Path:    graphql.GetPath(ctx),
			Message: "no userfound",

			Extensions: map[string]interface{}{
				"code": "400",
			},
		}
	}
	print(id.From)
	var d users.User
	d.FirstName = input.Firstname
	d.LastName = input.Lastname
	d.Email = input.Email
	d.Profileimage = input.Profileimage
	d.Mobile = input.Contactno
	d.ID = input.Userid
	d.Dialcode = input.Dialcode
	d.Configid = id.Configid
	print("updateconfig", d.Configid)
	res := d.CheckUserforupdate()
	if res.ID != 0 {
		if res.Email == input.Email {
			return &model.Updateddata{Status: false, Code: http.StatusConflict, Message: "Email Already Exists"}, nil
		} else if res.Mobile == input.Contactno {
			return &model.Updateddata{Status: false, Code: http.StatusConflict, Message: "Contactno Already Exists"}, nil
		}
	} else {
		stat, err := d.Updateappuser()
		if err != nil {
			if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'authname'", d.Email) {
				print("true")
				return &model.Updateddata{Status: false, Code: http.StatusConflict, Message: "Email Already Exists"}, nil
			} else if err.Error() == fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'contactno'", d.Mobile) {
				return &model.Updateddata{Status: false, Code: http.StatusConflict, Message: "Contactno Already Exists"}, nil
			} else {
				return nil, err
			}

		}
		if stat == true {
			st, er1 := d.Updateuserprofile()
			if er1 != nil {
				return nil, er1
			}
			print("userprofileupdate", st)
		}
	}

	return &model.Updateddata{Status: true, Code: http.StatusCreated, Message: "Profile Update Successfully"}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.GetUser, error) {
	var userResult []*model.GetUser
	var userGetAll []users.User

	userGetAll = users.GetAllUsers()
	for _, user := range userGetAll {
		userResult = append(userResult, &model.GetUser{Userid: user.ID, Firstname: user.FirstName, Lastname: user.LastName, Mobile: user.Mobile, Email: user.Email, Profileimage: user.Profileimage,
			Dialcode: user.Dialcode, Currencysymbol: user.Currencysymbol, Currencycode: user.CurrencyCode, Countrycode: user.Countrycode, Created: user.CreatedDate, Status: user.Status})
	}
	return userResult, nil
}

func (r *queryResolver) Getuser(ctx context.Context) (*model.LoginData, error) {
	id, usererr := datacontext.ForAuthContext(ctx)
	if usererr != nil {
		print("errr", usererr)
		return nil, &gqlerror.Error{

			Path:    graphql.GetPath(ctx),
			Message: "no userfound",

			Extensions: map[string]interface{}{
				"code": "400",
			},
		}
	}

	return &model.LoginData{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Success",
		UserInfo: &model.UserData1{
			UserID:         id.ID,
			Firstname:      id.FirstName,
			Lastname:       id.LastName,
			Email:          id.Email,
			Mobile:         id.Mobile,
			CreatedDate:    id.CreatedDate,
			Status:         id.Status,
			Profileimage:   id.Profileimage,
			Roleid:         &id.Roleid,
			Configid:       &id.Configid,
			Countrycode:    id.Countrycode,
			Currencycode:   id.CurrencyCode,
			Currencysymbol: id.Currencysymbol,
			Dialcode:       id.Dialcode,
		}}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
