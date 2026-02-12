variable "database_url" {
  type = string
  default = "postgresql://postgres:postgres@localhost:5432/vet_shifter"
}

env "local" {
  src = "file://sqlc/schemas.sql"
  dev = "docker://postgres/18/dev"
  url = var.database_url
}