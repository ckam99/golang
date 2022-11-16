package auth

import (
	"context"
	"fmt"
	"main/pkg/clients/postgresql"
	"strings"
)

type repository struct {
	postgresql.Client
}

func NewRepository(c postgresql.Client) Repository {
	return &repository{
		Client: c,
	}
}

func (r *repository) Find(ctx context.Context, user *User) error {
	args := make([]interface{}, 0)
	f := make([]string, 0)
	if user.ID > 0 {
		f = append(f, fmt.Sprintf("id=$%d", len(args)+1))
		args = append(args, user.ID)
	}

	if user.Email != "" {
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

	if len(f) == 0 {
		return fmt.Errorf("no parameter set")
	}
	q := fmt.Sprintf(`select * from users where %s limit 1`, strings.Join(f, " and "))

	if err := r.QueryRow(ctx, q, args...).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.IsActive,
		&user.EmailConfirmedAt,
		&user.PhoneConfirmedAt,
		&user.PasswordChangedAt,
		&user.CreatedAt,
		&user.UpatedAt,
	); err != nil {
		return postgresql.Error(err)
	}
	return nil
}

func (r *repository) Create(ctx context.Context, user *User) error {
	q := `insert into users(
  full_name,email,phone,role,password
) values($1,$2,$3,$4,$5) returning 
 id,created_at, updated_at
  `
	if err := r.QueryRow(ctx, q, user.FullName, user.Email, user.Phone, user.Role, user.Password).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, user *User) error {
	q := `update  users set
  coalesce($1,full_name),
  coalesce($2,email),
  coalesce($3,phone),
  coalesce($4,role),
  coalesce($5,password),
  returning updated_at where id = $6`
	if err := r.QueryRow(ctx, q,
		user.FullName, user.Email, user.Phone, user.Role, user.Password, user.ID).
		Scan(&user.UpatedAt); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int64, soft bool) error {
	q := "delete from users where id = $1"
	if soft {
		q = "update users set deleted_at=now() where id = $1"
	}
	if _, err := r.Exec(ctx, q, id); err != nil {
		return err
	}
	return nil
}
