package auth


type Role int64

const (
   UNDEFINED Role = iota
   USER
   ADMIN 
   SUPERUSER
)

func(r Role) String() string {
  switch r {
    case USER:
     return "user"
    case ADMIN:
     return "admin"
    case SUPERUSER:
      return "superuser"
    default:
      return "undefined"
  }
}
