package auth

import (
	"context"
	"database/sql"
  "fmt"
  "strings"
)

type repo struct {
	*sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{
		DB: db,
	}
}

func(r *repository) Find(ctx context.Context,  user *User) error{
args := make([]interface{},0)
f := make([]string,0)
  if user.ID > 0 {
    f = append(f, fmt.Sprintf("id=$%d", len(args)+1))
   args = append(args, user.ID)
  }

  if user.Email != ""{
    f = append(f, fmt.Sprintf("email=$%d", len(args)+1))
   args = append(args, user.Email)
  }

  if user.Phone != nil {
    f = append(f, fmt.Sprintf("phone=$%d", len(args)+1))
   args = append(args, user.Phone)
  }

  if user.Password != nil {
    f = append(f, fmt.Sprintf("password=$%d", len(args)+1))
   args = append(args, user.Password)
  }

  if len(f) == 0{
    return fmt.Errorf("No parameter set")
  }
  q := fmt.Sprintf(`select * from users where %s limit 1`, strings.Join(f, " and ") )
  
  var user User
 if err:=r.QueryRowContext(ctx, q, args...).Scan(
     &user.ID,
     &user.FullName,
     &user.Email,
     &user.Phone,
     &User.Role,
     &user.IsActive,
     &user.EmailConfirmedAt,
     &user.PhoneConfirmedAt,
     &user.PasswordChangedAt,
     &user.CreatedAt,
     &user.UpatedAt,
    );err!= nil{
     return User{}, err
    }
    return user, nil
}

(r *repository) Create(ctx context.Context, user *User) error{
 q := `insert into users(
  full_name,email,phone,role,password
) values($1,$2,$3,$4,$5) returning 
 id,created_at, updated_at
  `
  if err := r.QueryRowContext(ctx, q, user.FullName, user.Email,user.Phone,user.Role,user.Password).Scan(
    &user.ID,
    &user.CreatedAt,
    &user.UpatedAt,
  );err != nil {
    return err
  }
  return nil
}

(r *repository) Update(ctx context.Context, user *User) error {
  panic("not implemented")
}

(r *repository) Delete(ctx context.Context, id int64, soft bool) error {
  panic("not implemented")
}
