package user

// func GetAuthUserHandler(c way.Context) error {
// 	uuid := "7ccbce2f-3654-4497-8f62-7e11b89e98ce"
// 	id := models.UUIDStringToBytes(uuid)
// 	ctx := c.Request().Context()
// 	u, err := user.GetAuthUser(db.Auth, ctx, id)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, "invalid token")
// 	}

// 	// Return user details
// 	return c.JSON(http.StatusOK, u)
// }

// func GetUserByProvider(conn *database.Conn, ctx context.Context, provider *models.ProviderInfo) (*AuthUser, error) {
// 	id, err := conn.FindUserIdByProvider(ctx, *provider)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if id == nil {
// 		return nil, nil
// 	}

// 	return GetUserById(conn, ctx, id)
// }

// func GetUserByEmail(conn *database.Conn, ctx context.Context, email string) (*AuthUser, error) {
// 	id, err := conn.FindUserIdByEmail(ctx, email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if id == nil {
// 		return nil, nil
// 	}

// 	return GetUserById(conn, ctx, id)
// }

// func GetUserById(conn *database.Conn, ctx context.Context, id []byte) (*AuthUser, error) {
// 	user, err := conn.FindUserById(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ConvertUser(user), nil
// }

// func CreateUser(conn *database.Conn, ctx context.Context, user *models.User) error {
// 	return conn.CreateUser(ctx, user)
// }

// func AddOAuthProvider(conn *database.Conn, ctx context.Context, provider *models.OAuthProvider) error {
// 	return conn.CreateOAuthProvider(ctx, provider)
// }

// func GetAuthUser(conn *database.Conn, ctx context.Context, id []byte) (*AuthUser, error) {
// 	// get user
// 	user, err := conn.FindUserById(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// get user pictures
// 	pictures, err := conn.FindUserPicturesByUserId(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// get user emails
// 	emails, err := conn.FindUserEmailsByUserId(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// get user oauth providers
// 	providers, err := conn.FindUserOAuthProvidersByUserId(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// convert user
// 	authUser := ConvertUser(user)

// 	// convert pictures
// 	authUser.Pictures = ConvertUserPictures(pictures)

// 	// convert emails
// 	authUser.Email = ConvertUserEmails(emails)

// 	// convert oauth providers
// 	authUser.OAuthProviders = ConvertOAuthProviders(providers)

// 	return authUser, nil

// }
